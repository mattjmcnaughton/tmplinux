package virtualhost

import (
	"github.com/mattjmcnaughton/tmplinux/pkg/engine"
)

// VirtualHost is a wrapper around a specific engine we use to create a tmp
// linux environment.
type VirtualHost struct {
	engine.Engine
}

// NewContainerVirtualHost returns a new VirtualHost with the default engine for
// creating a tmp linux environment using a container.
func NewContainerVirtualHost() *VirtualHost {
	return &VirtualHost{
		engine.NewDockerEngine(),
	}
}

// NewVMVirtualHost returns a new VirtualHost with the default engine for
// creating a tmp linux environment using a vm.
func NewVMVirtualHost() *VirtualHost {
	return &VirtualHost{
		engine.NewVagrantEngine(),
	}
}
