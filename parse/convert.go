package parse

import (
	"net"
	"strings"
)

// ConvertToIPSlice iterates over a slice of strings and returns a slice of type net.IP
func ConvertToIPSlice(sl []string) (ips []net.IP) {
	for i := range sl {
		ips = append(ips, net.ParseIP(sl[i]))
	}
	return
}

// ConvertToStringSlice receives a comma-delimited string and returns a slice of strings containing each delimited value
func ConvertToStringSlice(s string) (sl []string) {
	for _, ss := range strings.Split(s, ",") {
		sl = append(sl, strings.TrimSpace(ss))
	}
	return
}
