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

var _ ProcessExplorerModule = &NamespaceModule{}

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

func (m *NamespaceModule) Meta() *Meta {
	return &Meta{
		Name:        "Namespace module",
		Description: "Search proc(5) filesystems for namespace meta.",
		Authors: []string{
			"Kris Nóva <kris@nivenly.com>",
		},
	}
}

func (m *NamespaceModule) Execute(p *Process) error {
	// Module specific (correlated)
	return nil
}
