package virtualhost

import (
	"github.com/mattjmcnaughton/tmplinux/pkg/engine"
)

type VirtualHost struct {
	engine.Engine
}

func NewContainerVirtualHost() *VirtualHost {
	return &VirtualHost{
		engine.NewDockerEngine(),
	}
}

func NewVMVirtualHost() *VirtualHost {
	return &VirtualHost{
		engine.NewVagrantEngine(),
	}
}
