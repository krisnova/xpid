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

package modproc

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const (
	DefaultProcRoot = "/proc"
)

type ProcFileSystem struct {
	rootPath string
}

var p *ProcFileSystem

func NewProcFileSystem(root string) *ProcFileSystem {
	if p == nil {
		p = &ProcFileSystem{}
	}
	return p
}

func (p *ProcFileSystem) Contents(file string) (string, error) {
	bytes, err := ioutil.ReadFile(filepath.Join(p.rootPath, file))
	return string(bytes), err
}

func (p *ProcFileSystem) ContentsPID(pid int64, file string) (string, error) {
	file = fmt.Sprintf("%d/%s", pid, file)
	return p.Contents(file)
}

func (p *ProcFileSystem) Dir(dir string) ([]fs.FileInfo, error) {
	return ioutil.ReadDir(filepath.Join(p.rootPath, dir))
}

func (p *ProcFileSystem) DirPID(pid int64, dir string) ([]fs.FileInfo, error) {
	dir = fmt.Sprintf("%d/%s", pid, dir)
	return p.Dir(dir)
}

func (p *ProcFileSystem) ReadlinkPID(pid int64, file string) (string, error) {
	file = fmt.Sprintf("%d/%s", pid, file)
	return p.Readlink(file)
}

func (p *ProcFileSystem) Readlink(file string) (string, error) {
	return os.Readlink(filepath.Join(p.rootPath, file))
}

func Proc() string {
	// TODO Look up procfs!
	return DefaultProcRoot
}

func FileKeyValue(content, key string) string {
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		spl := strings.Split(line, ":")
		if len(spl) != 2 {
			continue
		}
		k := strings.TrimSpace(spl[0])
		v := strings.TrimSpace(spl[1])
		if key == k {
			return v
		}
	}
	return ""
}
