package virtualhost

import (
	"fmt"

	"github.com/mattjmcnaughton/tmplinux/pkg/engine"
)

type VirtualHost struct {
	engine.Engine
}

func NewContainerVirtualHost() *VirtualHost {
	return &VirtualHost{
		engine.DockerEngine{},
	}
}

// func NewVMVirtualHost() *VirtualHost {
// 	return &VirtualHost{
// 		engine.VagrantEngine{},
// 	}
// }

func (v *VirtualHost) UniqueVirtualHostMethod() {
	fmt.Printf("hi")
}
