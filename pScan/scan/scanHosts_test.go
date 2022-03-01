package scan_test

import (
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
