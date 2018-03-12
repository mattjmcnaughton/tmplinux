package executor

import (
	"fmt"
)

// MockShellExecutor is a mock shell executor which records commands for
// inspection later.
type MockShellExecutor struct {
	success         bool
	executedCommand string
}

// CreateSuccessMockExecutor returns a `MockShellExecutor` where all the
// commands simulate success.
func CreateSuccessMockExecutor() MockShellExecutor {
	return MockShellExecutor{
		success: true,
	}
}

// CreateFailureMockExecutor returns a `MockShellExecutor` where all the
// commands simulate failure.
func CreateFailureMockExecutor() MockShellExecutor {
	return MockShellExecutor{
		success: false,
	}
}

// Run mocks executing the given command with no input/output binding.
func (m *MockShellExecutor) Run(name string, arg ...string) error {
	return m.mockCommand(name, arg...)
}

// RunWithBoundOutput mocks executing the given command with just the output bound to
// the current shell.
func (m *MockShellExecutor) RunWithBoundOutput(name string, arg ...string) error {
	return m.mockCommand(name, arg...)
}

// RunWithBoundInputOutput mocks executing the given command with the input and output
// bound to the current shell.
func (m *MockShellExecutor) RunWithBoundInputOutput(name string, arg ...string) error {
	return m.mockCommand(name, arg...)
}

// GetExecutedCommand is a public access on the `executedCommand`.
func (m *MockShellExecutor) GetExecutedCommand() string {
	return m.executedCommand
}

func (m *MockShellExecutor) mockCommand(name string, arg ...string) error {
	m.executedCommand = fmt.Sprintf("%s %v", name, arg)

	if m.success {
		return nil
	}

	return fmt.Errorf("Mock error")
}
