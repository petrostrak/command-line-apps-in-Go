// Package scan provides types and funcitons to perform TCP port
// scans on a list of hosts
package scan

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
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

// Add adds a host to the list
func (hl *HostsList) Add(host string) error {
	if found, _ := hl.search(host); found {
		return fmt.Errorf("%w: %s", ErrExists, host)
	}

	hl.Hosts = append(hl.Hosts, host)
	return nil
}

// Remove deletes a host from the list
func (hl *HostsList) Remove(host string) error {
	if found, i := hl.search(host); found {

		// Theory
		//
		// [i:j] ([i:] including i), ([:j] j excluding)
		// append to the list up to i (i excluding) the rest of
		// the list (i+1 including)
		hl.Hosts = append(hl.Hosts[:i], hl.Hosts[i+1:]...)
		return nil
	}

	return fmt.Errorf("%w: %s", ErrNotExists, host)
}

// Load obtains hosts from a hosts file
func (hl *HostsList) Load(hostFile string) error {
	f, err := os.Open(hostFile)
	if err != nil {
		if errors.Is(err, ErrNotExists) {
			return nil
		}

		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		hl.Hosts = append(hl.Hosts, scanner.Text())
	}

	return nil
}

// Save saves hosts to a hosts file
func (hl *HostsList) Save(hostsFile string) error {
	output := ""

	for _, h := range hl.Hosts {
		output += fmt.Sprintln(h)
	}

	return ioutil.WriteFile(hostsFile, []byte(output), 0644)
}
