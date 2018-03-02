package main

import (
	"fmt"
	"os"
	"os/exec"
)

// Start tmplinux
// Attach tmplinux
// Remove tmplinux

const (
	ContainerName = "tmplinuxcontainer"
	CmdName       = "docker"
)

func main() {
	action := os.Args[1]

	switch action {
	case "start":
		start()
	case "attach":
		attach()
	case "remove":
		remove()
	default:
		fmt.Fprintf(os.Stderr, "Not an option: %s", action)
	}
}

func start() {
	cmdArgs := []string{"run", "--name", ContainerName, "-d", "ubuntu:16.04", "tail", "-f", "/dev/null"}

	cmd := exec.Command(CmdName, cmdArgs...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	cmd.Run()
}

func attach() {
	cmdArgs := []string{"exec", "-it", ContainerName, "/bin/bash"}

	cmd := exec.Command(CmdName, cmdArgs...)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	cmd.Run()
}

func remove() {
	// What is the difference between `kill` and `stop`?
	killCmdArgs := []string{"kill", ContainerName}
	killCmd := exec.Command(CmdName, killCmdArgs...)

	rmCmdArgs := []string{"rm", ContainerName}
	rmCmd := exec.Command(CmdName, rmCmdArgs...)

	killCmd.Run()
	rmCmd.Run()
}
