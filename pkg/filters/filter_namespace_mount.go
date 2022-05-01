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

package filter

import (
	api "github.com/kris-nova/xpid/pkg/api/v1"
)

// 	NamespaceMount  string = "mnt"
//	NamespaceIPC    string = "ipc"
//	NamespaceCgroup string = "cgroup"
//	NamespacePid    string = "pid"
//	NamespaceNet    string = "net"
//	NamespaceUTS    string = "uts"
//	NamespaceTime   string = "time"

var NamespaceFilterSet_Mount = ""

func RetainNamespaceIn_Mount(p *api.Process) bool {
	if p.NamespaceModule.Mount == NamespaceFilterSet_Mount {
		return true
	}
	return false
}

func RetainNamespaceOut_Mount(p *api.Process) bool {
	if p.NamespaceModule.Mount == "" {
		return false
	}
	if p.NamespaceModule.Mount != NamespaceFilterSet_Mount {
		return true
	}
	return false
}
