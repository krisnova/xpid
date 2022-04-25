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

package proc

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

const (
	DefaultProcRoot = "/proc"
)

type ProcFileSystem struct {
	rootPath string
}

func NewProcFileSystem(root string) *ProcFileSystem {
	return &ProcFileSystem{
		rootPath: root,
	}
}

func (p *ProcFileSystem) Contents(file string) (string, error) {
	bytes, err := ioutil.ReadFile(filepath.Join(p.rootPath, file))
	return string(bytes), err
}

func (p *ProcFileSystem) ContentsPID(pid int64, file string) (string, error) {
	file = fmt.Sprintf("%d/%s", pid, file)
	return p.Contents(file)
}

func Proc() string {
	// TODO Look up procfs!
	return DefaultProcRoot
}
