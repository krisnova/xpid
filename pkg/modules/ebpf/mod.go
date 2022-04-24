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

package modebpf

import (
	"strings"

	"github.com/kris-nova/xpid/pkg/libxpid"

	api "github.com/kris-nova/xpid/pkg/api/v1"
	module "github.com/kris-nova/xpid/pkg/modules"
	"github.com/kris-nova/xpid/pkg/procx"
)

var _ procx.ProcessExplorerModule = &EBPFModule{}

const (
	EBPF_FullMount   string = "bpf /sys/fs/bpf bpf"
	EBPF_SYS_FSMount string = "/sys/fs/bpf"
)

type EBPFModule struct {
}

func NewEBPFModule() *EBPFModule {
	return &EBPFModule{}
}

type EBPFModuleResult struct {
	pid    *api.Process
	Mounts string
}

func (m *EBPFModule) Meta() *module.Meta {
	return &module.Meta{
		Name:        "eBPF module",
		Description: "Search proc(5) filesystems for eBPF programs. Will do an in depth scan and search for obfuscated directories.",
		Authors: []string{
			"Kris Nóva <kris@nivenly.com>",
		},
	}
}

func (m *EBPFModule) Execute(p *api.Process) (procx.ProcessExplorerResult, error) {
	// Module specific (correlated)
	result := &EBPFModuleResult{}
	result.Mounts = libxpid.ProcPidMounts(p.PID)

	// Higher level process (blind)
	if strings.Contains(result.Mounts, EBPF_FullMount) {
		p.EBPF = true
	}
	if strings.Contains(result.Mounts, EBPF_SYS_FSMount) {
		p.EBPF = true
	}

	return result, nil
}
