package blackwhite

import "net"

var chinaNet []*net.IPNet

func IsChinaIP(ip net.IP) bool {
	for _, v := range chinaNet {
		if v.Contains(ip) {
			return true
		}
	}
	return false
}
