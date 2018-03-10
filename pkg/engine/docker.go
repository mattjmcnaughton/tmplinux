package engine

import (
	"github.com/mattjmcnaughton/tmplinux/pkg/executor"
	"github.com/mattjmcnaughton/tmplinux/pkg/reporter"
)

const (
	containerImage = "ubuntu:16.04"
	cmdName        = "docker"
)

// DockerEngine conforms to the `Engine` interface, and provides all the methods
// for creating a tmplinux environment using docker.
type DockerEngine struct {
	exec executor.Executor
	reporter.Reporter
}

// NewDockerEngine returns a DockerEngine, to which we've injected production
// (i.e. non-testing) executors and reporters.
func NewDockerEngine() DockerEngine {
	return DockerEngine{
		executor.ShellExecutor{},
		reporter.FmtReporter{},
	}
}

// Start starts a tmp linux environment using Docker.
func (d DockerEngine) Start() {
	cmdArgs := []string{"run", "--name", d.getContainerName(), "-d", containerImage, "tail", "-f", "/dev/null"}
	err := d.exec.RunWithBoundOutput(cmdName, cmdArgs...)

	d.ReportIfError(err, "Failed to start the docker container")
}

// SSH gives the user a terminal in the tmp linux environment.
func (d DockerEngine) SSH() {
	cmdArgs := []string{"exec", "-it", d.getContainerName(), "/bin/bash"}
	err := d.exec.RunWithBoundInputOutput(cmdName, cmdArgs...)

	d.ReportIfError(err, "Failed to ssh into the docker container")
}

// Stop stops the tmp linux environment. It can be restarted.
// @TODO(mattjmcnaughton) Either get rid of `Stop` or write a command for
// restarting the tmp linux environment.
func (d DockerEngine) Stop() {
	cmdArgs := []string{"kill", d.getContainerName()}
	err := d.exec.RunWithBoundOutput(cmdName, cmdArgs...)

	d.ReportIfError(err, "Failed to stop the docker container")
}

// Rm deletes the tmp linux environment.
func (d DockerEngine) Rm() {
	cmdArgs := []string{"rm", d.getContainerName()}
	err := d.exec.RunWithBoundOutput(cmdName, cmdArgs...)

	d.ReportIfError(err, "Failed to remove the docker container")
}

// Validate ensures the docker engine has access to the deps it needs (i.e.
// docker is installed.)
func (d DockerEngine) Validate() {
	cmdArgs := []string{"ps"}
	err := d.exec.Run(cmdName, cmdArgs...)

	d.ReportIfError(err, "You do not have the necessary deps to run docker")
}

func (d DockerEngine) getContainerName() string {
	return "tmplinuxcontainer"
}
