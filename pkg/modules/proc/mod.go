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
/*
 * The proc module is responsible for the vast majority of the PID meta detail.
 * In most situations, this module will need to be executed in order to retrieve
 * the meta information about a process.
 */

package modproc

import (
	"os/user"
	"strconv"
	"strings"

	api "github.com/kris-nova/xpid/pkg/api/v1"
	"github.com/kris-nova/xpid/pkg/libxpid"
	module "github.com/kris-nova/xpid/pkg/modules"
	"github.com/kris-nova/xpid/pkg/procx"
)

var _ procx.ProcessExplorerModule = &ProcModule{}

type ProcModule struct {
}

var _ procx.ProcessExplorerResult = &ProcModuleResult{}

type ProcModuleResult struct {
	pid       *api.Process
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

func (m *ProcModule) Meta() *module.Meta {
	return &module.Meta{
		Name:        "Proc module",
		Description: "Search proc(5) filesystems for pid information. Will do an in depth scan and search for obfuscated directories.",
		Authors: []string{
			"Kris Nóva <kris@nivenly.com>",
		},
	}
}

func (m *ProcModule) Execute(p *api.Process) (procx.ProcessExplorerResult, error) {
	// Module specific (correlated)
	result := &ProcModuleResult{}

	// Process visibility
	result.Opendir = libxpid.ProcDirOpendir(p.PID)
	result.Chdir = libxpid.ProcDirChdir(p.PID)
	result.Dent = libxpid.ProcDirDent(p.PID)

	// Meta (comm. cmdline, status)
	procfs := NewProcFileSystem(Proc())
	comm, _ := procfs.ContentsPID(p.PID, "comm")
	result.Comm = strings.TrimSpace(comm)
	cmdline, _ := procfs.ContentsPID(p.PID, "cmdline")
	result.Cmdline = strings.TrimSpace(cmdline)
	status, _ := procfs.ContentsPID(p.PID, "status")
	result.Status = strings.TrimSpace(status)

	// User
	uidMap, _ := procfs.ContentsPID(p.PID, "uid_map")
	result.UID = IDFromMap(uidMap)
	u, _ := user.LookupId(IDFromMapString(uidMap))
	if u != nil {
		result.User = u
		result.Username = u.Username
	}

	// Group
	gidMap, _ := procfs.ContentsPID(p.PID, "gid_map")
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
	p.Group.ID = result.GID
	p.Group.Name = result.Groupname

	return result, nil
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
	tgid := FileKeyValue(status, "Tgid")
	pid := FileKeyValue(status, "Pid")
	if tgid != "" && pid != "" {
		if tgid != pid {
			return true
		}
	}
	return false
}
