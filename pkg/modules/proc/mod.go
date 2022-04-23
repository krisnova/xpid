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

// #cgo LDFLAGS: -lxpid
//
// #include "xpid.h"
import "C"

import (
	api "github.com/kris-nova/xpid/pkg/api/v1"
	module "github.com/kris-nova/xpid/pkg/modules"
	"github.com/kris-nova/xpid/pkg/procx"
)

var _ procx.ProcessExplorerModule = &ProcModule{}

type ProcModule struct {
}

var _ procx.ProcessExplorerResult = &ProcModuleResult{}

type ProcModuleResult struct {
	pid *api.Process
	ls  bool
	nav bool
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

func (m *ProcModule) Execute(process *api.Process) (procx.ProcessExplorerResult, error) {

	result := &ProcModuleResult{}

	// See if ls
	result.ls = proc_dir_ls(process.PID)

	// See if nav
	result.nav = proc_dir_nav(process.PID)

	return nil, nil
}
