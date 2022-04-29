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
	"sync"

	api "github.com/kris-nova/xpid/pkg/api/v1"
)

type ProcessExplorer struct {
	processes []*api.Process
	modules   []ProcessExplorerModule
	encoder   ProcessExplorerEncoder
	writer    io.Writer
	fast      bool
	waitGroup *sync.WaitGroup
}

func NewProcessExplorer(processes []*api.Process) *ProcessExplorer {
	return &ProcessExplorer{
		processes: processes,
		waitGroup: &sync.WaitGroup{},
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

func (x *ProcessExplorer) SetFast(f bool) {
	x.fast = f
}

// Execute will run the process explorer.
//
// The function is O(m*p) where the runtime complexity grows withs the amount
// of modules and pids to execute.
func (x *ProcessExplorer) Execute() error {

	// Validation
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

	// Main execution loops
	for _, process := range x.processes {
		for _, module := range x.modules {
			x.waitGroup.Add(1)
			if x.fast {
				go x.walk(process, module)
			} else {
				x.walk(process, module)
			}
		}
	}
	x.waitGroup.Wait()
	return nil
}

// Walk ignores errors and will walk a process and a module
//
// Walk may be ran concurrently if needed
func (x *ProcessExplorer) walk(p *api.Process, module ProcessExplorerModule) {
	module.Execute(p)
	r, _ := x.encoder.Encode(p)
	x.writer.Write(r)
	x.waitGroup.Done()
}
