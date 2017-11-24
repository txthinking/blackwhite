package blackwhite

import "net"

func init() {
	whiteNet = make([]*net.IPNet, 0)
	for _, v := range white_cidr {
		_, in, err := net.ParseCIDR(v)
		if err != nil {
			continue
		}
		whiteNet = append(whiteNet, in)
	}
}
