package engine

import (
	"fmt"

	"github.com/mattjmcnaughton/tmplinux/pkg/executor"
	"github.com/mattjmcnaughton/tmplinux/pkg/reporter"
)

type VagrantEngine struct {
	exec executor.Executor
	reporter.Reporter
}

func NewVagrantEngine() VagrantEngine {
	return VagrantEngine{
		executor.ShellExecutor{},
		reporter.FmtReporter{},
	}
}

func (v VagrantEngine) Start() {
	fmt.Printf("start")
}

func (v VagrantEngine) Ssh() {
	fmt.Printf("ssh")
}

func (v VagrantEngine) Stop() {
	fmt.Printf("stop")
}

func (v VagrantEngine) Rm() {
	fmt.Printf("rm")
}

func (v VagrantEngine) Validate() {
	fmt.Printf("validate")
}
