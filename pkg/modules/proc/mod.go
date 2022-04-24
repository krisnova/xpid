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

package modproc

import (
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
	pid     *api.Process
	Opendir int
	Chdir   int
	Dent    int
	Comm    string
	Cmdline string
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
	result.Opendir = libxpid.ProcDirOpendir(p.PID)
	result.Chdir = libxpid.ProcDirChdir(p.PID)
	result.Dent = libxpid.ProcDirDent(p.PID)
	result.Comm = libxpid.ProcPidComm(p.PID)
	result.Cmdline = libxpid.ProcPidCmdline(p.PID)

	// Higher level process (blind)
	p.ProcessVisible.Chdir = result.Chdir
	p.ProcessVisible.Dent = result.Dent
	p.ProcessVisible.Opendir = result.Opendir
	p.CommandLine = result.Cmdline
	p.Name = result.Comm
	return result, nil
}
