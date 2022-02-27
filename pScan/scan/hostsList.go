// Package scan provides types and funcitons to perform TCP port
// scans on a list of hosts
package scan

import (
	"errors"
	"sort"
)

var (
	ErrExists    = errors.New("host already in the list")
	ErrNotExists = errors.New("host not in the list")
)

// HostsList represents a list of hosts to run port scan
type HostsList struct {
	Hosts []string
}

// search rearches for hosts in the list
func (hl *HostsList) search(host string) (bool, int) {

	// sorts list alphabetically
	sort.Strings(hl.Hosts)

	// search for the host in the list
	i := sort.SearchStrings(hl.Hosts, host)
	if i < len(hl.Hosts) && hl.Hosts[i] == host {
		return true, i
	}

	return false, -1
}
