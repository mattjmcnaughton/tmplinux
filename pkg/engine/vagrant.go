package engine

import (
	"fmt"

	"github.com/mattjmcnaughton/tmplinux/pkg/executor"
	"github.com/mattjmcnaughton/tmplinux/pkg/reporter"
)

// VagrantEngine conforms to the `Engine` interface, and provides all the
// methods for creating a `tmplinux` environment using vagrant.
type VagrantEngine struct {
	exec     executor.Executor
	reporter reporter.Reporter
}

// NewVagrantEngine creates a new VagrantEngine, to which we've injected
// production (i.e. non-testing) executors and reporters.
func NewVagrantEngine() VagrantEngine {
	return VagrantEngine{
		&executor.ShellExecutor{},
		&reporter.FmtReporter{},
	}
}

// Start starts a tmp linux environment using Vagrant.
func (v VagrantEngine) Start() {
	fmt.Printf("start")
}

// SSH ssh's the user into the tmp linux environment.
func (v VagrantEngine) SSH() {
	fmt.Printf("ssh")
}

// Stop stops the users tmp linux environment.
func (v VagrantEngine) Stop() {
	fmt.Printf("stop")
}

// Rm stops the tmp linux environment.
func (v VagrantEngine) Rm() {
	fmt.Printf("rm")
}

// Validate ensures all the necessary dependencies are in place.
func (v VagrantEngine) Validate() {
	fmt.Printf("validate")
}
