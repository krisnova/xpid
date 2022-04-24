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
	module "github.com/kris-nova/xpid/pkg/modules"
	"github.com/kris-nova/xpid/pkg/procx"
	"github.com/sirupsen/logrus"
)

var _ procx.ProcessExplorerModule = &ProcModule{}

type ProcModule struct {
}

var _ procx.ProcessExplorerResult = &ProcModuleResult{}

type ProcModuleResult struct {
	pid     *api.Process
	opendir int
	chdir   int
	dent    int
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
	result := &ProcModuleResult{}
	result.opendir = proc_opendir(p.PID)
	result.chdir = proc_chdir(p.PID)
	result.dent = proc_dent(p.PID)
	p.ProcessVisible.Chdir = result.chdir
	p.ProcessVisible.Dent = result.dent
	p.ProcessVisible.Opendir = result.opendir
	if p.ProcessVisible.Opendir != p.ProcessVisible.Dent {
		logrus.Infof("Hidden PID : %d\n", p.PID)
		//logrus.Infof("Chdir      : %d\n", p.ProcessVisible.Chdir)
		//logrus.Infof("Opendir    : %d\n", p.ProcessVisible.Opendir)
		//logrus.Infof("Dent       : %d\n", p.ProcessVisible.Dent)
	}
	return result, nil
}
