//
// https://github.com/txthinking/pac
//
package main

import (
	"bytes"
	"io/ioutil"
	"log"
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
	wl := strings.Split(strings.TrimSpace(bytes.NewBuffer(data).String()), "\n")
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
	bl := strings.Split(strings.TrimSpace(bytes.NewBuffer(data).String()), "\n")
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

	return nil
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

const js = `
//
// https://github.com/txthinking/pac
//

var proxy="{{.proxy}}";

{{if .wl}}
var wl = {
	{{range .wl}}
	"{{.}}": 1,
	{{end}}
};
{{end}}

{{if .wc}}
var wc = [
    {{range .wc}}
    [{{.first}},{{.last}}],
	{{end}}
];
{{end}}

{{if .bl}}
var bl = {
	{{range .bl}}
	"{{.}}": 1,
	{{end}}
};
{{end}}

function ip2decimal(ip) {
    var d = ip.split('.');
    return ((((((+d[0])*256)+(+d[1]))*256)+(+d[2]))*256)+(+d[3]);
}

function FindProxyForURL(url, host){
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
        }
		{{end}}
    }

    if (isPlainHostName(host)){
        return "DIRECT";
    }

    {{if .wl}}
    var a = host.split(".");
    for(var i=a.length-1; i>=0; i--){
        if (wl.hasOwnProperty(a.slice(i).join("."))){
            return "DIRECT";
        }
    }
    return proxy;
	{{end}}

    {{if .bl}}
    var a = host.split(".");
    for(var i=a.length-1; i>=0; i--){
        if (bl.hasOwnProperty(a.slice(i).join("."))){
			return proxy;
        }
    }
	return "DIRECT";
    {{end}}
}
`
