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
	"strings"

	"github.com/kris-nova/xpid/pkg/procfs"
)

var _ ProcessExplorerModule = &NamespaceModule{}

const (

	// Only support a few namespaces for now
	// We can easily plumb more namespaces through
	// as needed

	NamespaceMount  string = "mnt"
	NamespaceIPC    string = "ipc"
	NamespaceCgroup string = "cgroup"
	NamespacePid    string = "pid"
	NamespaceNet    string = "net"
	NamespaceUTS    string = "uts"
	NamespaceTime   string = "time"
)

// [root@alice]: /proc/331236/ns># ls
// cgroup@  mnt@  pid@               time@               user@
// ipc@     net@  pid_for_children@  time_for_children@  uts@

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

type NamespaceModule struct {
	Net    string `json:"net,omitempty"`
	PID    string `json:"pid,omitempty"`
	Cgroup string `json:"cgroup,omitempty"`
	IPC    string `json:"ipc,omitempty"`
	Mount  string `json:"mnt,omitempty"`
	UTS    string `json:"uts,omitempty"`
	Time   string `json:"time,omitempty"`
}

func (m *NamespaceModule) Execute(p *Process) error {
	// Module specific (correlated)
	p.NamespaceModule.Net = nsString(p.PID, NamespaceNet)
	p.NamespaceModule.IPC = nsString(p.PID, NamespaceIPC)
	p.NamespaceModule.PID = nsString(p.PID, NamespacePid)
	p.NamespaceModule.Cgroup = nsString(p.PID, NamespaceCgroup)
	p.NamespaceModule.Mount = nsString(p.PID, NamespaceMount)
	p.NamespaceModule.UTS = nsString(p.PID, NamespaceUTS)
	p.NamespaceModule.Time = nsString(p.PID, NamespaceTime)
	return nil
}

func nsString(pid int64, ns string) string {
	procfshandle := procfs.NewProcFileSystem(procfs.Proc())
	content, _ := procfshandle.ReadlinkPID(pid, fmt.Sprintf("ns/%s", ns))
	if content == "" {
		return content
	}
	spl := strings.Split(content, ":")
	if len(spl) < 2 {
		return content
	}
	nsBracket := spl[1]
	nsBracket = strings.Replace(nsBracket, "[", "", 1)
	nsBracket = strings.Replace(nsBracket, "]", "", 1)
	return strings.TrimSpace(nsBracket)
}
