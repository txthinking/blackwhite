package blackwhite

import (
	"strings"
)

// IsBlack determine whether host is in black list,
// host like www.google.com or ip
func IsBlack(host string) bool {
	ss := strings.Split(host, ".")
	var s string
	for i := len(ss) - 1; i >= 0; i-- {
		if s == "" {
			s = ss[i]
		} else {
			s = ss[i] + "." + s
		}
		if _, ok := black_list[s]; ok {
			return true
		}
	}
	return false
}
