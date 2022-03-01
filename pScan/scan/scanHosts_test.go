package scan_test

import (
	"net"
	"strconv"
	"testing"

	"github.com/petrostrak/command-line-apps-in-Go/pScan/scan"
)

func TestStateString(t *testing.T) {
	ps := scan.PortState{}

	if ps.Open.String() != "closed" {
		t.Errorf("expected %q, got %q instead\n", "closed", ps.Open.String())
	}

	ps.Open = true

	if ps.Open.String() != "open" {
		t.Errorf("expected %q, got %q instead\n", "open", ps.Open.String())
	}
}

func TestRunHostFound(t *testing.T) {
	testCases := []struct {
		name          string
		expectedState string
	}{
		{"OpenPort", "open"},
		{"ClosedPort", "closed"},
	}

	host := "localhost"
	hl := &scan.HostsList{}

	hl.Add(host)

	ports := []int{}

	// init ports, 1 open, 1 closed
	for _, tc := range testCases {
		ln, err := net.Listen("tcp", net.JoinHostPort(host, "0"))
		if err != nil {
			t.Fatal(err)
		}
		defer ln.Close()

		_, portStr, err := net.SplitHostPort(ln.Addr().String())
		if err != nil {
			t.Fatal(err)
		}

		port, err := strconv.Atoi(portStr)
		if err != nil {
			t.Fatal(err)
		}

		ports = append(ports, port)

		if tc.name == "ClosedPort" {
			ln.Close()
		}
	}

	res := scan.Run(hl, ports)

	// verify results for HostFound test
	if len(res) != 1 {
		t.Fatalf("expected 1 results, got %d instead\n", len(res))
	}

	if res[0].Host != host {
		t.Errorf("expected host %q, got %q instead\n", host, res[0].Host)
	}

	if res[0].NotFound {
		t.Errorf("expected host %q to be found\n", host)
	}

	if len(res[0].PortStates) != 2 {
		t.Fatalf("expected 2 port states, got %d instead\n", len(res[0].PortStates))
	}

	for i, tc := range testCases {
		if res[0].PortStates[i].Port != ports[i] {
			t.Errorf("expected port %d, got %d instead\n", ports[0], res[0].PortStates[i].Port)
		}

		if res[0].PortStates[i].Open.String() != tc.expectedState {
			t.Errorf("Expected port %d to be %s\n", ports[i], tc.expectedState)
		}
	}
}
