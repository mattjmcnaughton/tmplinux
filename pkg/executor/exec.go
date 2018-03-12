package executor

import (
	"os"
	"os/exec"
)

// Executor provides an interface for running shell commands. We do this instead
// of directly using `os/exec` so it is easier to unit test.
type Executor interface {
	Run(string, ...string) error
	RunWithBoundOutput(string, ...string) error
	RunWithBoundInputOutput(string, ...string) error
}

// ShellExecutor is a "production" executor which executes commands in the shell.
type ShellExecutor struct{}

// Run executes the given command with no input/output binding.
func (s *ShellExecutor) Run(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)

	return cmd.Run()
}

// RunWithBoundOutput executes the given command with just the output bound to
// the current shell.
func (s *ShellExecutor) RunWithBoundOutput(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

// RunWithBoundInputOutput executes the given command with the input and output
// bound to the current shell.
func (s *ShellExecutor) RunWithBoundInputOutput(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
