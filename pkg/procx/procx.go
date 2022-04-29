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

const (
	// PidPoolLimit is the limit of concurrent pids to investigate
	// concurrently. Default to 2^19 = 524288
	PidPoolLimit int = 524288
)

type ProcessExplorer struct {
	processes []*api.Process
	modules   []ProcessExplorerModule
	encoder   ProcessExplorerEncoder
	writer    io.Writer
	fast      bool
	ftx       PidPool
}

func NewProcessExplorer(processes []*api.Process) *ProcessExplorer {
	return &ProcessExplorer{
		processes: processes,
		ftx:       NewPidPool(PidPoolLimit),
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

	// Safe to do pid 1 again
	for _, process := range x.processes {
		for _, module := range x.modules {
			x.ftx.Add()
			if x.fast {
				go x.walk(process, module)
			} else {
				x.walk(process, module)
			}
		}
	}
	for x.ftx.Cur() != 0 {
	}
	return nil
}

// Walk ignores errors and will walk a process and a module
//
// Walk may be ran concurrently if needed
func (x *ProcessExplorer) walk(p *api.Process, module ProcessExplorerModule) {
	module.Execute(p)
	r, _ := x.encoder.Encode(p)
	x.writer.Write(r)
	x.ftx.Sub()
}

type PidPool struct {
	sync.Mutex
	cur   int
	limit int
}

func NewPidPool(limit int) PidPool {
	return PidPool{
		limit: limit,
		cur:   0,
		Mutex: sync.Mutex{},
	}
}

func (x *PidPool) Add() {
	for x.Cur() >= x.limit {
	}
	x.Lock()
	defer x.Unlock()
	x.cur++
}

func (x *PidPool) Sub() {
	x.Lock()
	defer x.Unlock()
	x.cur--
}

func (x *PidPool) Cur() int {
	x.Lock()
	defer x.Unlock()
	return x.cur
}
