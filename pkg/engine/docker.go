package engine

import (
	"github.com/mattjmcnaughton/wutils/pkg/executor"
	"github.com/mattjmcnaughton/wutils/pkg/reporter"
)

const (
	containerImage = "ubuntu:16.04"
	dockerCmdName  = "docker"
)

// DockerEngine conforms to the `Engine` interface, and provides all the methods
// for creating a tmplinux environment using docker.
type DockerEngine struct {
	exec     executor.Executor
	reporter reporter.Reporter
}

// NewDockerEngine returns a DockerEngine, to which we've injected production
// (i.e. non-testing) executors and reporters.
func NewDockerEngine() DockerEngine {
	return NewCustomDockerEngine(
		&executor.ShellExecutor{},
		&reporter.FmtReporter{},
	)
}

// NewCustomDockerEngine creates a new DockerEngine, where the caller can
// specify the injected executor and reporter.
func NewCustomDockerEngine(e executor.Executor, r reporter.Reporter) DockerEngine {
	return DockerEngine{e, r}
}

// Start starts a tmp linux environment using Docker.
func (d DockerEngine) Start() {
	cmdArgs := []string{"run", "--name", d.getContainerName(), "-d", containerImage, "tail", "-f", "/dev/null"}
	err := d.exec.RunWithBoundOutput(dockerCmdName, cmdArgs...)

	d.reporter.ReportIfError(err, "Failed to start the docker container")
}

// SSH gives the user a terminal in the tmp linux environment.
func (d DockerEngine) SSH() {
	cmdArgs := []string{"exec", "-it", d.getContainerName(), "/bin/bash"}
	err := d.exec.RunWithBoundInputOutput(dockerCmdName, cmdArgs...)

	d.reporter.ReportIfError(err, "Failed to ssh into the docker container")
}

// Stop stops the tmp linux environment. It can be restarted.
// @TODO(mattjmcnaughton) Either get rid of `Stop` or write a command for
// restarting the tmp linux environment.
func (d DockerEngine) Stop() {
	cmdArgs := []string{"kill", d.getContainerName()}
	err := d.exec.RunWithBoundOutput(dockerCmdName, cmdArgs...)

	d.reporter.ReportIfError(err, "Failed to stop the docker container")
}

// Rm deletes the tmp linux environment.
func (d DockerEngine) Rm() {
	cmdArgs := []string{"rm", d.getContainerName()}
	err := d.exec.RunWithBoundOutput(dockerCmdName, cmdArgs...)

	d.reporter.ReportIfError(err, "Failed to remove the docker container")
}

// Validate ensures the docker engine has access to the deps it needs (i.e.
// docker is installed.)
func (d DockerEngine) Validate() {
	cmdArgs := []string{"ps"}
	err := d.exec.Run(dockerCmdName, cmdArgs...)

	d.reporter.ReportIfError(err, "You do not have the necessary deps to run docker")
}

func (d DockerEngine) getContainerName() string {
	return "tmplinuxcontainer"
}
