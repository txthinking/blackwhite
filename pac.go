//
// https://github.com/txthinking/pac
//
package main

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"
	"time"

	"github.com/codegangsta/cli"
	"github.com/txthinking/ant"
)

var mode string
var domainURL string
var cidrURL string
var proxy string
var listen string

func main() {
	app := cli.NewApp()
	app.Name = "PAC"
	app.Version = "20180918"
	app.Usage = "PAC file generator"
	app.Author = "Cloud"
	app.Email = "cloud@txthinking.com"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "mode, m",
			Usage:       "white/black/global [required]",
			Destination: &mode,
		},
		cli.StringFlag{
			Name:        "domainURL, d",
			Usage:       "domains url, http(s):// or file:// [required when mode is not global]",
			Destination: &domainURL,
		},
		cli.StringFlag{
			Name:        "cidrURL, c",
			Usage:       "CIDR url, http(s):// or file:// [optional]",
			Destination: &cidrURL,
		},
		cli.StringFlag{
			Name:        "proxy, p",
			Usage:       "Proxy, like: SOCKS5 127.0.0.1:1080; SOCKS 127.0.0.1:1080; DIRECT [required]",
			Destination: &proxy,
		},
		cli.StringFlag{
			Name:        "listen, l",
			Usage:       "HTTP server address, like: 127.0.0.1:1980 [optional]",
			Destination: &listen,
		},
	}
	app.Action = func(c *cli.Context) error {
		if mode != "global" && mode != "white" && mode != "black" {
			return errors.New("Invalid mode")
		}
		if mode != "global" && domainURL == "" {
			return errors.New("domainURL required")
		}
		if proxy == "" {
			return errors.New("proxy required")
		}
		return run()
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	r, err := PAC(proxy, mode, domainURL, cidrURL)
	if err != nil {
		return err
	}
	if listen == "" {
		if _, err := io.Copy(os.Stdout, r); err != nil {
			return err
		}
		return nil
	}

	b, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/x-ns-proxy-autoconfig")
		w.Write(b)
	})
	return http.ListenAndServe(listen, nil)
}

func PAC(proxy, mode, domainURL, cidrURL string) (io.Reader, error) {
	t := template.New("pac")
	t, err := t.Parse(js)
	if err != nil {
		return nil, err
	}
	b := &bytes.Buffer{}

	if mode == "global" {
		if err := t.Execute(b, map[string]interface{}{
			"mode":  "global",
			"proxy": proxy,
		}); err != nil {
			return nil, err
		}
		return b, nil
	}

	var ds []string
	var cs []map[string]int64
	if domainURL != "" {
		data, err := readData(domainURL)
		if err != nil {
			return nil, err
		}
		ds = makeDomains(data)
	}
	if cidrURL != "" {
		data, err := readData(cidrURL)
		if err != nil {
			return nil, err
		}
		cs = makeCIDRs(data)
	}

	if err := t.Execute(b, map[string]interface{}{
		"proxy":   proxy,
		"mode":    mode,
		"domains": ds,
		"cidrs":   cs,
	}); err != nil {
		return nil, err
	}
	return b, nil
}

func readData(url string) ([]byte, error) {
	if strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://") {
		c := &http.Client{
			Timeout: 9 * time.Second,
		}
		r, err := c.Get(url)
		if err != nil {
			return nil, err
		}
		defer r.Body.Close()
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return nil, err
		}
		return data, nil
	}
	if strings.HasPrefix(url, "file://") {
		data, err := ioutil.ReadFile(url[7:])
		if err != nil {
			return nil, err
		}
		return data, nil
	}
	return nil, errors.New("Unsupport URL")
}

func makeDomains(data []byte) []string {
	data = bytes.TrimSpace(data)
	data = bytes.Replace(data, []byte{0x20}, []byte{}, -1)
	data = bytes.Replace(data, []byte{0x0d, 0x0a}, []byte{0x0a}, -1)
	ds := strings.Split(string(data), "\n")
	return ds
}

func makeCIDRs(data []byte) []map[string]int64 {
	cs := make([]map[string]int64, 0)
	data = bytes.TrimSpace(data)
	data = bytes.Replace(data, []byte{0x20}, []byte{}, -1)
	data = bytes.Replace(data, []byte{0x0d, 0x0a}, []byte{0x0a}, -1)
	ss := strings.Split(string(data), "\n")
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
		cs = append(cs, m)
	}
	return cs
}

const js = `
//
// https://github.com/txthinking/pac
//

var proxy="{{.proxy}}";

var mode = "{{.mode}}";

{{if .domains}}
var domains = {
	{{range .domains}}
	"{{.}}": 1,
	{{end}}
};
{{end}}

{{if .cidrs}}
var cidrs = [
    {{range .cidrs}}
    [{{.first}},{{.last}}],
	{{end}}
];
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
        {{if .cidrs}}
        var d = ip2decimal(host);
        var l = cidrs.length;
        var min = 0;
        var max = l;
        for(;;){
            if (min+1 > max) {
                break;
            }
            var mid = Math.floor(min+(max-min)/2);
            if(d >= cidrs[mid][0] && d <= cidrs[mid][1]){
				if(mode == "white"){
					return "DIRECT";
				}
				if(mode == "black"){
					return proxy;
				}
            }else if(d < cidrs[mid][0]){
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

    {{if .domains}}
    var a = host.split(".");
    for(var i=a.length-1; i>=0; i--){
        if (domains.hasOwnProperty(a.slice(i).join("."))){
			if(mode == "white"){
				return "DIRECT";
			}
			if(mode == "black"){
				return proxy;
			}
        }
    }
	if(mode == "white"){
		return proxy;
	}
	if(mode == "black"){
		return "DIRECT";
	}
	{{end}}

	if(mode == "global"){
		return proxy;
	}
}
`
