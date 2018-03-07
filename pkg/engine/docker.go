package engine

import (
	"fmt"
)

// Think about injecting the `os` module to these engines.
type DockerEngine struct {
}

func (d DockerEngine) Start() {
	fmt.Printf("start")
}

func (d DockerEngine) Ssh() {
	fmt.Printf("ssh")
}

func (d DockerEngine) Stop() {
	fmt.Printf("stop")
}

func (d DockerEngine) Rm() {
	fmt.Printf("rm")
}

func (d DockerEngine) Validate() {
	fmt.Printf("validate")
}
