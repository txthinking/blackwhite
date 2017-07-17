package blackwhite

import (
	"net"
	"strings"
)

// IsWhite determine whether host is in white list, host like www.google.com or ip
func IsWhite(host string) bool {
	ip := net.ParseIP(host)
	if ip != nil {
		return IsWhiteIP(ip)
	}
	ss := strings.Split(host, ".")
	var s string
	for i := len(ss) - 1; i >= 0; i-- {
		if s == "" {
			s = ss[i]
		} else {
			s = ss[i] + "." + s
		}
		if _, ok := white_list[s]; ok {
			return true
		}
	}
	return false
}

var chinaNet []*net.IPNet = make([]*net.IPNet, 0)

func IsWhiteIP(ip net.IP) bool {
	for _, v := range chinaNet {
		if v.Contains(ip) {
			return true
		}
	}
	return false
}

func GetWhiteAPP() (string, error) {
	return strings.TrimSpace(white_app), nil
}
