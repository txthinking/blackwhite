//
// https://github.com/txthinking/pac
//
package main

import (
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

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

var cm = { {{range $k, $v := .cm}}
"{{$k}}":"{{$v}}",{{end}} };

function FindProxyForURL(url, host){
	// customized
    if(cm.hasOwnProperty(host)){
        return cm[host];
    }

	// internal
    if(/\d+\.\d+\.\d+\.\d+/.test(host)){
        if (isInNet(dnsResolve(host), "10.0.0.0", "255.0.0.0") ||
                isInNet(dnsResolve(host), "172.16.0.0",  "255.240.0.0") ||
                isInNet(dnsResolve(host), "192.168.0.0", "255.255.0.0") ||
                isInNet(dnsResolve(host), "127.0.0.0", "255.255.255.0")){
            return "DIRECT";
        }
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

func getData(mode, proxy string) map[string]interface{} {
	if mode == "white" {
		return map[string]interface{}{
			"proxy": proxy,
			"wl":    wl,
			"cm":    cm,
		}
	}
	if mode == "black" {
		return map[string]interface{}{
			"proxy": proxy,
			"bl":    bl,
			"cm":    cm,
		}
	}
	return map[string]interface{}{
		"proxy": proxy,
		"cm":    cm,
	}
}

func pac(w http.ResponseWriter, r *http.Request) {
	t := template.New("pac")
	t, err := t.Parse(js)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	vars := mux.Vars(r)
	if vars["mode"] != "white" && vars["mode"] != "black" && vars["mode"] != "all" {
		http.Error(w, "Unsupport mode", 400)
		return
	}
	w.Header().Set("Content-Type", "application/x-ns-proxy-autoconfig")
	err = t.Execute(w, getData(vars["mode"], vars["proxy"]))
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}
