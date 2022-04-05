package parse

import (
	"net"
	"strings"
)

func ConvertToIPSlice(sl []string) (ips []net.IP) {
	for i := range sl {
		ips = append(ips, net.ParseIP(sl[i]))
	}
	return
}

func ConvertToStringSlice(s string) (sl []string) {
	for _, ss := range strings.Split(s, ",") {
		sl = append(sl, strings.TrimSpace(ss))
	}
	return
}
