//go:build linux
// +build linux

package gateway

import (
	"net"
)

func discoverGatewayOSSpecificIPv4() (ip net.IP, err error) {
	ipstr := execCmd("sh", "-c", "route -n | grep 'UG[ \t]' | awk 'NR==1{print $2}'")
	ipv4 := net.ParseIP(ipstr)
	if ipv4 == nil {
		// For debian if called by a no-root user
		ipstr = execCmd("sh", "-c", "/usr/sbin/route -n | grep 'UG[ \t]' | awk 'NR==1{print $2}'")
		ipv4 = net.ParseIP(ipstr)
		if ipv4 == nil {
			return nil, errCantParse
		}
		//return nil, errCantParse
	}
	return ipv4, nil
}

func discoverGatewayOSSpecificIPv6() (ip net.IP, err error) {
	ipstr := execCmd("sh", "-c", "route -6 -n | grep 'UG[ \t]' | awk 'NR==1{print $2}'")
	ipv6 := net.ParseIP(ipstr)
	if ipv6 == nil {
		return nil, errCantParse
	}
	return ipv6, nil
}
