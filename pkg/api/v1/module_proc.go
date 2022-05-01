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

	procfs "github.com/kris-nova/xpid/pkg/procfs"

	"github.com/kris-nova/xpid/pkg/libxpid"
)

var _ ProcessExplorerModule = &ProcModule{}

type ProcModule struct {
	Opendir   int
	Chdir     int
	Dent      int
	Comm      string
	Cmdline   string
	Status    string
	UID       int
	GID       int
	Username  string
	Groupname string
	User      *user.User
	Group     *user.Group
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
	result := &ProcModule{}

	// Process visibility
	result.Opendir = libxpid.ProcDirOpendir(p.PID)
	result.Chdir = libxpid.ProcDirChdir(p.PID)
	result.Dent = libxpid.ProcDirDent(p.PID)

	// Meta (comm. cmdline, status)
	procfshandle := procfs.NewProcFileSystem(procfs.Proc())
	comm, _ := procfshandle.ContentsPID(p.PID, "comm")
	result.Comm = strings.TrimSpace(comm)
	cmdline, _ := procfshandle.ContentsPID(p.PID, "cmdline")
	result.Cmdline = strings.TrimSpace(cmdline)
	status, _ := procfshandle.ContentsPID(p.PID, "status")
	result.Status = strings.TrimSpace(status)

	// User
	uidMap, _ := procfshandle.ContentsPID(p.PID, "uid_map")
	result.UID = IDFromMap(uidMap)
	u, _ := user.LookupId(IDFromMapString(uidMap))
	if u != nil {
		result.User = u
		result.Username = u.Username
	}

	// Group
	gidMap, _ := procfshandle.ContentsPID(p.PID, "gid_map")
	result.GID = IDFromMap(gidMap)
	g, _ := user.LookupGroupId(IDFromMapString(gidMap))
	if g != nil {
		result.Group = g
		result.Groupname = g.Name
	}

	// Higher level process (blind)
	// Map to higher abstraction here
	p.ProcessVisible.Chdir = result.Chdir
	p.ProcessVisible.Dent = result.Dent
	p.ProcessVisible.Opendir = result.Opendir
	p.CommandLine = result.Cmdline
	p.Name = result.Comm
	p.Thread = StatusFileIsThread(result.Status)

	p.User.ID = result.UID
	p.User.Name = result.Username
	p.User.Group.ID = result.GID
	p.User.Group.Name = result.Groupname

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
