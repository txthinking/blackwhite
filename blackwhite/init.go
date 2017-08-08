package blackwhite

import "net"

func init() {
	for _, v := range china_cidr {
		_, in, err := net.ParseCIDR(v)
		if err != nil {
			continue
		}
		chinaNet = append(chinaNet, in)
	}
}
