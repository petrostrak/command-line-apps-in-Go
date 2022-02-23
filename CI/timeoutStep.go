package main

import "time"

type timeoutStep struct {
	step
	timeout time.Duration
}

func newTimeoutStep(name, exe, msg, proj string, args []string, timeout time.Duration) timeoutStep {
	s := timeoutStep{}

	s.step = newStep(name, exe, msg, proj, args)

	s.timeout = timeout
	if s.timeout == 0 {
		s.timeout = 30 * time.Second
	}

	return s
}
