/******************************************************************************
* MIT License
* Copyright (c) 2022 Kris Nóva <kris@nivenly.com>
*
* ┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓
* ┃   ███╗   ██╗ ██████╗ ██╗   ██╗ █████╗  ┃
* ┃   ████╗  ██║██╔═████╗██║   ██║██╔══██╗ ┃
* ┃   ██╔██╗ ██║██║██╔██║██║   ██║███████║ ┃
* ┃   ██║╚██╗██║████╔╝██║╚██╗ ██╔╝██╔══██║ ┃
* ┃   ██║ ╚████║╚██████╔╝ ╚████╔╝ ██║  ██║ ┃
* ┃   ╚═╝  ╚═══╝ ╚═════╝   ╚═══╝  ╚═╝  ╚═╝ ┃
* ┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛
*
*****************************************************************************/

package procx

import (
	api "github.com/kris-nova/xpid/pkg/api/v1"
	"github.com/sirupsen/logrus"
)

type ProcessExplorer struct {
	processes []*api.Process
	modules   []ProcessExplorerModule
	encoder   ProcessExplorerEncoder
	writer    ProcessExplorerWriter
}

func NewProcessExplorer(processes []*api.Process) *ProcessExplorer {
	return &ProcessExplorer{
		processes: processes,
	}
}

func (x *ProcessExplorer) AddModule(m ProcessExplorerModule) {
	x.modules = append(x.modules, m)
}

func (x *ProcessExplorer) SetEncoder(e ProcessExplorerEncoder) {
	x.encoder = e
}

func (x *ProcessExplorer) Execute() error {
	for _, process := range x.processes {
		logrus.Infof("Process: %d\n", process.PID)
	}
	return nil
}
