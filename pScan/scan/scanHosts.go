// Package scan provides types and functions to perform TCP port
// scans on a list of hosts.
package scan

// PortState represents the state of a single TCP port.
type PortState struct {
	Port int
	Open state
}

// state is a wrapper on the bool type. It indicates whether a
// port is open or closed.
type state bool

// String converts the boolean value of state to a human readable string
func (s state) String() string {
	if s {
		return "open"
	}

	return "closed"
}
