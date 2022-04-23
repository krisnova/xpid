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
	"fmt"
	"io"

	api "github.com/kris-nova/xpid/pkg/api/v1"
	"github.com/sirupsen/logrus"
)

type ProcessExplorer struct {
	processes []*api.Process
	modules   []ProcessExplorerModule
	encoder   ProcessExplorerEncoder
	writer    io.Writer
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

func (x *ProcessExplorer) SetWriter(w io.Writer) {
	x.writer = w
}

// Execute will run the process explorer.
//
// The function is O(m*p) where the runtime complexity grows withs the amount
// of modules and pids to execute.
func (x *ProcessExplorer) Execute() error {

	if x.processes == nil {
		return fmt.Errorf("missing pids in process explorer")
	}
	if x.encoder == nil {
		return fmt.Errorf("missing encoder in process explorer")
	}
	if x.modules == nil {
		return fmt.Errorf("missing modules in process explorer")
	}
	if len(x.modules) < 1 {
		return fmt.Errorf("empty modules in process explorer")
	}

	for _, module := range x.modules {
		logrus.Infof("Module: %s\n", module.Meta().Name)
		for _, process := range x.processes {
			logrus.Debugf("PID: %d\n", process.PID)
			result, err := module.Execute(process)
			if err != nil {
				logrus.Warnf("%s(%d) error: %v\n", module.Meta().Name, process.PID, err)
			}
			rawResult, err := x.encoder.Encode(result)
			if err != nil {
				logrus.Warnf("%s.encode(%d) error: %v\n", module.Meta().Name, process.PID, err)
			}
			_, err = x.writer.Write(rawResult)
			if err != nil {
				logrus.Warnf("%s.write(%d) error: %v\n", module.Meta().Name, process.PID, err)
			}
		}
	}
	return nil
}
