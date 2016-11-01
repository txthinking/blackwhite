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

var wl = [{{range .wl}}
"{{.}}",{{end}}];
var bl = [{{range .bl}}
"{{.}}",{{end}}];
var cm = { {{range $k, $v := .cm}}
"{{$k}}":"{{$v}}",{{end}} };

function FindProxyForURL(url, host){
    if(cm.hasOwnProperty(host)){
        return cm[host];
    }

    if(/\d+\.\d+\.\d+\.\d+/.test(host)){
        if (isInNet(dnsResolve(host), "10.0.0.0", "255.0.0.0") ||
                isInNet(dnsResolve(host), "172.16.0.0",  "255.240.0.0") ||
                isInNet(dnsResolve(host), "192.168.0.0", "255.255.0.0") ||
                isInNet(dnsResolve(host), "127.0.0.0", "255.255.255.0")){
            return "DIRECT";
        }
    }
    if (isPlainHostName(host)){
        return "DIRECT";
    }


    if(wl.length !== 0){
        var l = wl.length;
        for(var i=0;i<l;i++){
            if(dnsDomainIs(host, wl[i])){
                return 'DIRECT';
            }
        }
        return proxy;
    }

    if(bl.length !== 0){
        var l = wl.length;
        for(var i=0;i<l;i++){
            if(dnsDomainIs(host, bl[i])){
                return proxy;
            }
        }
        return 'DIRECT';
    }

    return proxy;
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
