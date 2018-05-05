package engine

import (
	"github.com/mattjmcnaughton/wutils/pkg/executor"
	"github.com/mattjmcnaughton/wutils/pkg/reporter"
)

const (
	vagrantBox     = "ubuntu/bionic64"
	vagrantCmdName = "vagrant"
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

// NewCustomVagrantEngine creates a new VagrantEngine, where the caller can
// specify the injected executor and reporter.
func NewCustomVagrantEngine(e executor.Executor, r reporter.Reporter) VagrantEngine {
	return VagrantEngine{e, r}
}

// Start starts a tmp linux environment using Vagrant.
func (v VagrantEngine) Start() {
	initCmdArgs := []string{"init", vagrantBox}
	err := v.exec.RunInDirWithBoundOutput(v.getTmpDirPath(), vagrantCmdName, initCmdArgs...)
	v.reporter.ReportIfError(err, "Failed to init vagrant vm")

	upCmdArgs := []string{"up"}
	err = v.exec.RunInDirWithBoundOutput(v.getTmpDirPath(), vagrantCmdName, upCmdArgs...)
	v.reporter.ReportIfError(err, "Failed to start vagrant vm")
}

// SSH ssh's the user into the tmp linux environment.
func (v VagrantEngine) SSH() {
	sshCmdArgs := []string{"ssh"}
	err := v.exec.RunInDirWithBoundInputOutput(v.getTmpDirPath(), vagrantCmdName, sshCmdArgs...)
	v.reporter.ReportIfError(err, "Failed to ssh into vagrant vm")
}

// Stop stops the users tmp linux environment.
func (v VagrantEngine) Stop() {
	stopCmdArgs := []string{"suspend"}
	err := v.exec.RunInDirWithBoundOutput(v.getTmpDirPath(), vagrantCmdName, stopCmdArgs...)
	v.reporter.ReportIfError(err, "Failed to stop vagrant vm")
}

// Rm stops the tmp linux environment.
func (v VagrantEngine) Rm() {
	rmCmdArgs := []string{"destroy", "-f"}
	err := v.exec.RunInDirWithBoundOutput(v.getTmpDirPath(), vagrantCmdName, rmCmdArgs...)
	v.reporter.ReportIfError(err, "Failed to remove vagrant vm")

	err = v.exec.RunInDir(v.getTmpDirPath(), "rm", "Vagrantfile")
	v.reporter.ReportIfError(err, "Failed to clean up Vagrantfile")
}

// Validate ensures all the necessary dependencies are in place.
func (v VagrantEngine) Validate() {
	validateCmdArgs := []string{"global-status"}
	err := v.exec.RunInDir(v.getTmpDirPath(), vagrantCmdName, validateCmdArgs...)
	v.reporter.ReportIfError(err, "Vagrant environment not properly setup")
}

func (v VagrantEngine) getTmpDirPath() string {
	return "/tmp"
}
