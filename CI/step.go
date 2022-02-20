package main

import "os/exec"

type step struct {
	name    string
	exe     string
	message string
	args    []string
	proj    string
}

func newStep(name, exe, msg, proj string, args []string) step {
	return step{
		name,
		exe,
		msg,
		args,
		proj,
	}
}

func (s step) execute() (string, error) {
	cmd := exec.Command(s.exe, s.args...)
	cmd.Dir = s.proj

	if err := cmd.Run(); err != nil {
		return "", &stepErr{
			step:  s.name,
			msg:   "failed to execute",
			cause: err,
		}
	}

	return s.message, nil
}
