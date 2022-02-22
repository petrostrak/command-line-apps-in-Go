package main

type exceptionStep struct {
	step
}

func newExceptionStep(name, exe, msg, proj string, args []string) exceptionStep {
	s := exceptionStep{}

	s.step = newStep(name, exe, msg, proj, args)

	return s
}
