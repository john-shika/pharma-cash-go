package nokocore

import "os"

type ExitCode int

const (
	ExitCodeSuccess ExitCode = iota
	ExitCodeFailure
)

type MainFunc func([]string) ExitCode

func (m MainFunc) Call(args []string) ExitCode {
	return m(args)
}

func ApplyMainFunc(mainFunc MainFunc) {
	exitCode := mainFunc.Call(os.Args)
	os.Exit(int(exitCode))
}
