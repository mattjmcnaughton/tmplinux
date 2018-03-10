package engine

import (
	"github.com/mattjmcnaughton/tmplinux/pkg/executor"
	"github.com/mattjmcnaughton/tmplinux/pkg/reporter"
)

const (
	ContainerImage = "ubuntu:16.04"
	CmdName        = "docker"
)

type DockerEngine struct {
	exec executor.Executor
	reporter.Reporter
}

func NewDockerEngine() DockerEngine {
	return DockerEngine{
		executor.ShellExecutor{},
		reporter.FmtReporter{},
	}
}

func (d DockerEngine) Start() {
	cmdArgs := []string{"run", "--name", d.getContainerName(), "-d", ContainerImage, "tail", "-f", "/dev/null"}
	err := d.exec.RunWithBoundOutput(CmdName, cmdArgs...)

	d.ReportIfError(err, "Failed to start the docker container")
}

func (d DockerEngine) Ssh() {
	cmdArgs := []string{"exec", "-it", d.getContainerName(), "/bin/bash"}
	err := d.exec.RunWithBoundInputOutput(CmdName, cmdArgs...)

	d.ReportIfError(err, "Failed to ssh into the docker container")
}

func (d DockerEngine) Stop() {
	cmdArgs := []string{"kill", d.getContainerName()}
	err := d.exec.RunWithBoundOutput(CmdName, cmdArgs...)

	d.ReportIfError(err, "Failed to stop the docker container")
}

func (d DockerEngine) Rm() {
	cmdArgs := []string{"rm", d.getContainerName()}
	err := d.exec.RunWithBoundOutput(CmdName, cmdArgs...)

	d.ReportIfError(err, "Failed to remove the docker container")
}

func (d DockerEngine) Validate() {
	cmdArgs := []string{"ps"}
	err := d.exec.Run(CmdName, cmdArgs...)

	d.ReportIfError(err, "You do not have the necessary deps to run docker")
}

func (d DockerEngine) getContainerName() string {
	return "tmplinuxcontainer"
}
