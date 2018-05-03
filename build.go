//
// https://github.com/txthinking/pac
//
package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strings"
	"text/template"

	"github.com/codegangsta/cli"
	"github.com/txthinking/ant"
)

var whiteList string
var whiteCIDR string
var blackList string
var proxy string

func main() {
	app := cli.NewApp()
	app.Name = "PAC"
	app.Version = "20180503"
	app.Usage = "PAC file generator"
	app.Author = "Cloud"
	app.Email = "cloud@txthinking.com"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "whitelist",
			Value:       "./white.list",
			Usage:       "White list file.",
			Destination: &whiteList,
		},
		cli.StringFlag{
			Name:        "whitecidr",
			Value:       "./china_cidr.list",
			Usage:       "White CIDR file",
			Destination: &whiteCIDR,
		},
		cli.StringFlag{
			Name:        "blacklist",
			Value:       "./black.list",
			Usage:       "Black list file",
			Destination: &blackList,
		},
		cli.StringFlag{
			Name:        "proxy",
			Value:       "SOCKS5 127.0.0.1:1080; SOCKS 127.0.0.1:1080; DIRECT",
			Usage:       "Proxy",
			Destination: &proxy,
		},
	}
	app.Action = func(c *cli.Context) error {
		return build()
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func build() error {
	t := template.New("pac")
	t, err := t.Parse(js)
	if err != nil {
		return err
	}

	data, err := ioutil.ReadFile(whiteList)
	if err != nil {
		return err
	}
	wl := makeList(data)
	data, err = ioutil.ReadFile(whiteCIDR)
	if err != nil {
		return err
	}
	wc := makeWhiteCIDR(data)
	f, err := os.OpenFile("./white.pac", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	if err := t.Execute(f, map[string]interface{}{
		"proxy": proxy,
		"wl":    wl,
		"wc":    wc,
	}); err != nil {
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}

	data, err = ioutil.ReadFile(blackList)
	if err != nil {
		return err
	}
	bl := makeList(data)
	f, err = os.OpenFile("./black.pac", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	if err := t.Execute(f, map[string]interface{}{
		"proxy": proxy,
		"bl":    bl,
	}); err != nil {
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}

	f, err = os.OpenFile("./all.pac", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	if err := t.Execute(f, map[string]interface{}{
		"proxy": proxy,
	}); err != nil {
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}

	return nil
}

func makeList(data []byte) map[string]map[string][]string {
	w := make(map[string]map[string][]string)
	w["_"] = map[string][]string{"_": []string{}}
	bb := bytes.Split(bytes.TrimSpace(data), []byte("\n"))
	for _, d := range bb {
		if net.ParseIP(string(d)) != nil {
			w["_"]["_"] = append(w["_"]["_"], string(d))
			continue
		}

		rd := reverseAsCopy(d)
		i := bytes.IndexByte(rd, '.')
		if i == 0 || i == len(rd)-1 {
			// invalid
			continue
		}
		if i == -1 {
			// cn/local/...
			w["_"]["_"] = append(w["_"]["_"], string(d))
			continue
		}

		suffix := string(rd[:i])
		index := string(rd[i+1 : i+2])
		_, ok := w[suffix]
		if !ok {
			w[suffix] = make(map[string][]string)
		}
		_, ok = w[suffix][index]
		if !ok {
			w[suffix][index] = make([]string, 0)
		}
		w[suffix][index] = append(w[suffix][index], string(d))
	}
	return w
}

func makeWhiteCIDR(data []byte) []map[string]int64 {
	wc := make([]map[string]int64, 0)
	ss := strings.Split(strings.TrimSpace(bytes.NewBuffer(data).String()), "\n")
	for _, s := range ss {
		c, err := ant.CIDR(s)
		if err != nil {
			continue
		}
		first, err := ant.IP2Decimal(c.First)
		if err != nil {
			continue
		}
		last, err := ant.IP2Decimal(c.Last)
		if err != nil {
			continue
		}
		m := make(map[string]int64)
		m["first"] = first
		m["last"] = last
		wc = append(wc, m)
	}
	return wc
}

func reverseAsCopy(s []byte) []byte {
	a := make([]byte, len(s))
	copy(a, s)
	i := 0
	j := len(a) - 1
	for i < j {
		x := a[i]
		a[i] = a[j]
		a[j] = x
		i++
		j--
	}
	return a
}

const js = `
//
// https://github.com/txthinking/pac
//
var proxy="{{.proxy}}";

{{if .wl}}
var wl = {
    {{range $k, $v := .wl}}
    "{{$k}}" : {
        {{range $kk, $vv := $v}}
        "{{$kk}}": [
            {{range $vv}}
            "{{.}}",{{end}}
        ],{{end}}
    },{{end}}
};{{end}}

{{if .wc}}
var wc = [
    {{range .wc}}
    [{{.first}},{{.last}}],{{end}}
];{{end}}

{{if .bl}}
var bl = {
    {{range $k, $v := .bl}}
    "{{$k}}" : {
        {{range $kk, $vv := $v}}
        "{{$kk}}": [
            {{range $vv}}
            "{{.}}",{{end}}
        ],{{end}}
    },{{end}}
};{{end}}

function ip2decimal(ip) {
    var d = ip.split('.');
    return ((((((+d[0])*256)+(+d[1]))*256)+(+d[2]))*256)+(+d[3]);
}

function FindProxyForURL(url, host){
    // internal
    if(/\d+\.\d+\.\d+\.\d+/.test(host)){
        if (isInNet(dnsResolve(host), "10.0.0.0", "255.0.0.0") ||
                isInNet(dnsResolve(host), "172.16.0.0",  "255.240.0.0") ||
                isInNet(dnsResolve(host), "192.168.0.0", "255.255.0.0") ||
                isInNet(dnsResolve(host), "127.0.0.0", "255.255.255.0")){
            return "DIRECT";
        }
        {{if .wc}}
        var d = ip2decimal(host);
        var l = wc.length;
        var min = 0;
        var max = l;
        for(;;){
            if (min+1 > max) {
                break;
            }
            var mid = Math.floor(min+(max-min)/2);
            if(d >= wc[mid][0] && d<=wc[mid][1]){
                return "DIRECT";
            }else if(d < wc[mid][0]){
                max = mid;
            }else{
                min = mid+1;
            }
        }{{end}}
    }

    // plain
    if (isPlainHostName(host)){
        return "DIRECT";
    }

    {{if .wl}}

    var l = wl["_"]["_"].length;
    for(var i=0;i<l;i++){
        if(dnsDomainIs(host, wl["_"]["_"][i])){
            return "DIRECT";
        }
    }

    var rd = host.split("").reverse().join("")
    var i = rd.indexOf(".")
    if (i == 0 || i == rd.length-1){
        return "DIRECT";
    }
    if (i == -1){
        return "DIRECT";
    }
    var suffix = rd.substring(0, i)
    var index = rd.substring(i+1, i+2)
    if (!wl.hasOwnProperty(suffix)){
        return proxy;
    }
    if (!wl[suffix].hasOwnProperty(index)){
        return proxy;
    }
    var l = wl[suffix][index].length;
    for(var i=0;i<l;i++){
        if(dnsDomainIs(host, wl[suffix][index][i])){
            return "DIRECT";
        }
    }
    return proxy;

    {{else if .bl}}

    var l = bl["_"]["_"].length;
    for(var i=0;i<l;i++){
        if(dnsDomainIs(host, bl["_"]["_"][i])){
            return proxy;
        }
    }

    var rd = host.split("").reverse().join("")
    var i = rd.indexOf(".")
    if (i == 0 || i == rd.length-1){
        return "DIRECT";
    }
    if (i == -1){
        return "DIRECT";
    }
    var suffix = rd.substring(0, i)
    var index = rd.substring(i+1, i+2)
    if (!bl.hasOwnProperty(suffix)){
        return "DIRECT";
    }
    if (!bl[suffix].hasOwnProperty(index)){
        return "DIRECT";
    }
    var l = bl[suffix][index].length;
    for(var i=0;i<l;i++){
        if(dnsDomainIs(host, bl[suffix][index][i])){
            return proxy;
        }
    }
    return "DIRECT";

    {{else}}
    return proxy;
    {{end}}
}
`
