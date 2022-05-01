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

package modnamespace

import (
	api "github.com/kris-nova/xpid/pkg/api/v1"
	module "github.com/kris-nova/xpid/pkg/modules"
	"github.com/kris-nova/xpid/pkg/procx"
)

var _ procx.ProcessExplorerModule = &NamespaceModule{}

const (
	NamespaceMount  string = "mnt"
	NamespaceIPC    string = "ipc"
	NamespaceCgroup string = "cgroup"
	NamespacePid    string = "pid"
	NamespaceNet    string = "net"
)

// see lsns for more

type NamespaceModule struct {
}

func NewNamespaceModule() *NamespaceModule {
	return &NamespaceModule{}
}

type NamespaceModuleResult struct {
	pid    *api.Process
	Mounts string
}

func (m *NamespaceModule) Meta() *module.Meta {
	return &module.Meta{
		Name:        "Namespace module",
		Description: "Search proc(5) filesystems for namespace meta.",
		Authors: []string{
			"Kris Nóva <kris@nivenly.com>",
		},
	}
}

func (m *NamespaceModule) Execute(p *api.Process) (procx.ProcessExplorerResult, error) {
	// Module specific (correlated)
	result := &NamespaceModuleResult{}
	return result, nil
}
