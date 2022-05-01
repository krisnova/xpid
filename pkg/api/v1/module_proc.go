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
	"os/user"
	"strconv"
	"strings"

	"github.com/kris-nova/xpid/pkg/procfs"

	"github.com/kris-nova/xpid/pkg/libxpid"
)

var _ ProcessExplorerModule = &ProcModule{}

type ProcModule struct {
	Opendir int
	Chdir   int
	Dent    int
	Comm    string
	Cmdline string
	Status  string
}

func NewProcModule() *ProcModule {
	return &ProcModule{}
}

func (m *ProcModule) Meta() *Meta {
	return &Meta{
		Name:        "Proc module",
		Description: "Search proc(5) filesystems for pid information. Will do an in depth scan and search for obfuscated directories.",
		Authors: []string{
			"Kris Nóva <kris@nivenly.com>",
		},
	}
}

func (m *ProcModule) Execute(p *Process) error {
	// Module specific (correlated)

	// Process visibility
	p.ProcessVisible.Opendir = libxpid.ProcDirOpendir(p.PID)
	p.ProcessVisible.Chdir = libxpid.ProcDirChdir(p.PID)
	p.ProcessVisible.Dent = libxpid.ProcDirDent(p.PID)

	// Meta (comm. cmdline, status)
	procfshandle := procfs.NewProcFileSystem(procfs.Proc())
	comm, _ := procfshandle.ContentsPID(p.PID, "comm")
	p.ProcModule.Comm = strings.TrimSpace(comm)
	p.Name = p.ProcModule.Comm
	cmdline, _ := procfshandle.ContentsPID(p.PID, "cmdline")
	p.ProcModule.Cmdline = strings.TrimSpace(cmdline)
	p.CommandLine = p.ProcModule.Cmdline
	status, _ := procfshandle.ContentsPID(p.PID, "status")
	p.Status = strings.TrimSpace(status)

	// User
	uidMap, _ := procfshandle.ContentsPID(p.PID, "uid_map")
	p.User.ID = IDFromMap(uidMap)
	u, _ := user.LookupId(IDFromMapString(uidMap))
	if u != nil {
		p.User.User = *u
		p.User.Name = u.Username
		p.User.Name = u.Username
	}

	// Group
	gidMap, _ := procfshandle.ContentsPID(p.PID, "gid_map")
	p.User.Group.ID = IDFromMap(gidMap)
	g, _ := user.LookupGroupId(IDFromMapString(gidMap))
	if g != nil {
		p.User.Group.Name = g.Name
	}

	// Higher level process (blind)
	// Map to higher abstraction here
	p.Thread = StatusFileIsThread(p.Status)

	return nil
}

// IDFromMap returns the first value in uid_map and gid_map in /proc
func IDFromMap(mp string) int {
	retstr := ""
	mp = strings.TrimSpace(mp)
	spl := strings.Split(mp, " ")
	retstr = spl[0] // Should always be the first value
	reti, err := strconv.Atoi(retstr)
	if err != nil {
		return -1
	}
	return reti
}

func IDFromMapString(mp string) string {
	mp = strings.TrimSpace(mp)
	spl := strings.Split(mp, " ")
	return spl[0]
}

func StatusFileIsThread(status string) bool {
	tgid := procfs.FileKeyValue(status, "Tgid")
	pid := procfs.FileKeyValue(status, "Pid")
	if tgid != "" && pid != "" {
		if tgid != pid {
			return true
		}
	}
	return false
}
