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
	api "github.com/kris-nova/xpid/pkg/api/v1"
	module "github.com/kris-nova/xpid/pkg/modules"
	"github.com/kris-nova/xpid/pkg/procx"
)

var _ procx.ProcessExplorerModule = &ContainerModule{}

type ContainerModule struct {
}

func NewContainerModule() *ContainerModule {
	return &ContainerModule{}
}

type ContainerModuleResult struct {
	pid *api.Process
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

const (
	ContainerValueSigSuspendX86 string = "__x64_sys_rt_sigsuspend"
	ContainerValueSigSuspend    string = "sigsuspend"
)

func (m *ContainerModule) Execute(p *api.Process) (procx.ProcessExplorerResult, error) {
	// Module specific (correlated)
	result := &ContainerModuleResult{}

	//procfs := modproc.NewProcFileSystem(modproc.Proc())
	//stack, err := procfs.ContentsPID(p.PID, "stack")
	//if err != nil {
	//	return nil, fmt.Errorf("unable to read stack")
	//}

	//
	// TODO Parse proc/stack for container detail here!
	//

	return result, nil
}
