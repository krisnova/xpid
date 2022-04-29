/*===========================================================================*\
 *           MIT License Copyright (c) 2022 Kris Nóva <kris@nivenly.com>     *
 *                                                                           *
 *                ┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓                *
 *                ┃   ███╗   ██╗ ██████╗ ██╗   ██╗ █████╗   ┃                *
 *                ┃   ████╗  ██║██╔═████╗██║   ██║██╔══██╗  ┃                *
 *                ┃   ██╔██╗ ██║██║██╔██║██║   ██║███████║  ┃                *
 *                ┃   ██║╚██╗██║████╔╝██║╚██╗ ██╔╝██╔══██║  ┃                *
 *                ┃   ██║ ╚████║╚██████╔╝ ╚████╔╝ ██║  ██║  ┃                *
 *                ┃   ╚═╝  ╚═══╝ ╚═════╝   ╚═══╝  ╚═╝  ╚═╝  ┃                *
 *                ┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛                *
 *                                                                           *
 *                       This machine kills fascists.                        *
 *                                                                           *
\*===========================================================================*/

package modcontainer

import (
	"fmt"

	api "github.com/kris-nova/xpid/pkg/api/v1"
	module "github.com/kris-nova/xpid/pkg/modules"
	modproc "github.com/kris-nova/xpid/pkg/modules/proc"
	"github.com/kris-nova/xpid/pkg/procx"
)

var _ procx.ProcessExplorerModule = &ContainerModule{}

type ContainerModule struct {
}

func NewContainerModule() *ContainerModule {
	m := &ContainerModule{}
	// We always need a pid 1
	m.Execute(api.ProcessPID(1))
	return m
}

type ContainerModuleResult struct {
	pid *api.Process

	// NamespaceCgroupLink /proc/[pid]/ns/@cgroup
	NamespaceCgroupLink string
	// raw fields
}

func (m *ContainerModule) Meta() *module.Meta {
	return &module.Meta{
		Name:        "Container module",
		Description: "Find container meta information at runtime.",
		Authors: []string{
			"Kris Nóva <kris@nivenly.com>",
		},
	}
}

var pidone *ContainerModuleResult

func (m *ContainerModule) Execute(p *api.Process) (procx.ProcessExplorerResult, error) {
	// Module specific (correlated)
	result := &ContainerModuleResult{}

	procfs := modproc.NewProcFileSystem(modproc.Proc())
	nscgroup, _ := procfs.ReadlinkPID(p.PID, "ns/cgroup")
	result.NamespaceCgroupLink = nscgroup

	// Map

	// If it's pid 1  we can just move on, there is nothing to compare
	if p.PID == 1 {
		p.Container = false
		pidone = result
		return result, nil
	}
	if pidone == nil {
		return nil, fmt.Errorf("pid one not initialized")
	}

	// Research:
	//
	// As far as I can tell the majority of container environments
	// can be identified by their system.slice mounts in /sys/fs/cgroup
	// or by the ns/cgroup mapping in /proc
	//
	// For us to call something "a container" it basically needs to have
	// a unique ns/cgroup link that is different from the pid 1 in the
	// current pid namespace.
	if nscgroup != pidone.NamespaceCgroupLink {
		// We found a container
		p.Container = true
	}

	return result, nil
}
