
//
// https://github.com/txthinking/pac
//

var proxy="SOCKS5 [::1]:1090; SOCKS [::1]:1090; DIRECT";

var mode = "global";





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
        
    }

    if (isPlainHostName(host)){
        return "DIRECT";
    }

    

	if(mode == "global"){
		return proxy;
	}
}
