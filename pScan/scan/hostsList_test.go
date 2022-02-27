package scan

import (
	"errors"
	"testing"
)

func TestAdd(t *testing.T) {
	testCases := []struct {
		name      string
		host      string
		expectLen int
		expectErr error
	}{
		{"AddNew", "host2", 2, nil},
		{"AddExisting", "host1", 1, ErrExists},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			hl := &HostsList{}

			// Initialize list
			if err := hl.Add("host1"); err != nil {
				t.Fatal(err)
			}

			err := hl.Add(tc.host)

			if tc.expectErr != nil {
				if err == nil {
					t.Fatalf("expected error, got nil instead\n")
				}

				if !errors.Is(err, tc.expectErr) {
					t.Errorf("expected error %q, got %q instead\n", tc.expectErr, err)
				}

				return
			}

			if err != nil {
				t.Fatalf("expected no error, got %q instead\n", err)
			}

			if len(hl.Hosts) != tc.expectLen {
				t.Errorf("expected list length %d, got %d instead\n", tc.expectLen, len(hl.Hosts))
			}

			if hl.Hosts[1] != tc.host {
				t.Errorf("expected host name %q as index 1, got %q instead\n", tc.host, hl.Hosts[1])
			}
		})
	}
}
