package blackwhite

import (
	"net"
	"strings"
)

var whiteNet []*net.IPNet

// IsWhite determine whether host is in white list or china cidr,
// host like www.google.com or ip.
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

func GetWhiteAPP() (string, error) {
	return strings.TrimSpace(white_app), nil
}

func IsWhiteIP(ip net.IP) bool {
	for _, v := range whiteNet {
		if v.Contains(ip) {
			return true
		}
	}
	return false
}
