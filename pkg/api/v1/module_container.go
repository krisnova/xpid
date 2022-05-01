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

package v1

import (
	"fmt"

	"github.com/kris-nova/xpid/pkg/procfs"
)

var _ ProcessExplorerModule = &ContainerModule{}

type ContainerModule struct {
	// NamespaceCgroupLink /proc/[pid]/ns/@cgroup
	NamespaceCgroupLink string
}

func NewContainerModule() *ContainerModule {
	m := &ContainerModule{}
	// We always need a pid 1
	m.Execute(ProcessPID(1))
	return m
}

func (m *ContainerModule) Meta() *Meta {
	return &Meta{
		Name:        "Container module",
		Description: "Find container meta information at runtime.",
		Authors: []string{
			"Kris Nóva <kris@nivenly.com>",
		},
	}
}

var pidone *ContainerModule

func (m *ContainerModule) Execute(p *Process) error {
	// Module specific (correlated)
	result := &ContainerModule{}

	procfshandle := procfs.NewProcFileSystem(procfs.Proc())
	nscgroup, _ := procfshandle.ReadlinkPID(p.PID, "ns/cgroup")
	result.NamespaceCgroupLink = nscgroup

	// If it's pid 1  we can just move on, there is nothing to compare
	if p.PID == 1 {
		p.Container = false
		pidone = result
		return nil
	}
	if pidone == nil {
		return fmt.Errorf("pid one not initialized")
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

	return nil
}
