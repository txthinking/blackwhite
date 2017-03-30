package blackwhite

import (
	"net"
	"strings"
)

// IsWhite determine whether host is in white list, host like www.google.com or ip
func IsWhite(host string) bool {
	ip := net.ParseIP(host)
	if ip != nil {
		if _, ok := white_list[host]; ok {
			return true
		}
		return false
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
