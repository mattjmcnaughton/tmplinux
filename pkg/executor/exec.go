package executor

import (
	"os"
	"os/exec"
)

type Executor interface {
	Run(string, ...string) error
	RunWithBoundOutput(string, ...string) error
	RunWithBoundInputOutput(string, ...string) error
}

// Will define `MockShellExecutor` for use in my tests.
type ShellExecutor struct{}

func (s ShellExecutor) Run(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)

	return cmd.Run()
}

func (s ShellExecutor) RunWithBoundOutput(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func (s ShellExecutor) RunWithBoundInputOutput(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
