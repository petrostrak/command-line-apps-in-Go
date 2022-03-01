// Package scan provides types and functions to perform TCP port
// scans on a list of hosts.
package scan

import (
	"fmt"
	"net"
	"time"
)

// PortState represents the state of a single TCP port.
type PortState struct {
	Port int
	Open state
}

// state is a wrapper on the bool type. It indicates whether a
// port is open or closed.
type state bool

// Results represents the scan results for a single host
type Results struct {
	Host       string
	NotFound   bool
	PortStates []PortState
}

// String converts the boolean value of state to a human readable string
func (s state) String() string {
	if s {
		return "open"
	}

	return "closed"
}

// scanPort performs a port scan on a single TCP port
func scanPort(host string, port int) PortState {
	p := PortState{
		Port: port,
	}

	// net.JoinHostPort function is recommended over concatenating
	// the values directly at it takes care of corner cases, such as
	// the IPv6 value.
	address := net.JoinHostPort(host, fmt.Sprintf("%d", port))

	scanConn, err := net.DialTimeout("tcp", address, 1*time.Second)
	if err != nil {
		return p
	}

	// when the connection succeeds, close the connection
	scanConn.Close()
	p.Open = true
	return p
}

// Run performs a port scan on the hosts list
func Run(hl *HostsList, ports []int) []Results {
	res := make([]Results, 0, len(hl.Hosts))

	for _, h := range hl.Hosts {
		r := Results{
			Host: h,
		}

		if _, err := net.LookupHost(h); err != nil {
			r.NotFound = true
			res = append(res, r)
			continue
		}

		for _, p := range ports {
			r.PortStates = append(r.PortStates, scanPort(h, p))
		}

		res = append(res, r)
	}

	return res
}
