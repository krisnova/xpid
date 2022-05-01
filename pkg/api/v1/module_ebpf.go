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

package v1

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/kris-nova/xpid/pkg/procfs"
)

var _ ProcessExplorerModule = &EBPFModule{}

type EBPFModule struct {
	Mounts string
	Progs  []string
	Maps   []string
}

func NewEBPFModule() *EBPFModule {
	return &EBPFModule{}
}

const (
	EBPFFullMount  string = "bpf /sys/fs/bpf bpf"
	EBPFSYSFSMount string = "/sys/fs/bpf"
)

func (m *EBPFModule) Meta() *Meta {
	return &Meta{
		Name:        "eBPF module",
		Description: "Search proc(5) filesystems for eBPF programs. Will do an in depth scan and search for obfuscated directories.",
		Authors: []string{
			"Kris Nóva <kris@nivenly.com>",
		},
	}
}

func (m *EBPFModule) Execute(p *Process) error {
	// Module specific (correlated)

	procfshandle := procfs.NewProcFileSystem(procfs.Proc())
	mounts, _ := procfshandle.ContentsPID(p.PID, "mounts")
	p.Mounts = mounts

	e, err := NewEBPFFileSystemData()
	if err != nil {
		return fmt.Errorf("unable to read /sys/fs/bpf: %v", err)
	}

	// Compare with file descriptors in /proc
	fds, err := procfshandle.DirPID(p.PID, "fdinfo")

	if err != nil {
		return fmt.Errorf("unable to read /proc/%d/fdinfo: %v", p.PID, err)
	}
	for _, fd := range fds {
		fddata, err := procfshandle.ContentsPID(p.PID, filepath.Join("fdinfo", fd.Name()))
		if err != nil {
			continue
		}
		ebpfProgID := procfs.FileKeyValue(fddata, "prog_id")
		if ebpfProgID == "" {
			continue
		}
		for id, mp := range e.Progs {
			if id == "" {
				continue
			}
			if id == ebpfProgID {
				// We have mapped an eBPF program to a PID!
				p.EBPF = true
				p.EBPFModule.Progs = append(p.EBPFModule.Progs, mp.Name)
			}
		}
	}

	return nil
}

// EBPFFileSystemData is structured data from /sys/fs/bpf/*
type EBPFFileSystemData struct {
	Maps  map[string]*Map
	Progs map[string]*Prog
}

type Map struct {
	ID         string
	Name       string
	MaxEntries string
}
type Prog struct {
	ID       string
	Name     string
	Attached string
}

const (
	DefaultEBPFFileSystemDataDir = "/sys/fs/bpf"
)

func NewEBPFFileSystemData() (*EBPFFileSystemData, error) {
	e := &EBPFFileSystemData{
		Progs: make(map[string]*Prog),
		Maps:  make(map[string]*Map),
	}
	mapbytes, err := ioutil.ReadFile(filepath.Join(DefaultEBPFFileSystemDataDir, "maps.debug"))
	if err != nil {
		return nil, fmt.Errorf("map read: %v", err)
	}
	mapstr := string(mapbytes)
	lines := strings.Split(mapstr, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		spl := strings.Split(line, " ")
		if len(spl) < 2 {
			continue
		}
		id := strings.TrimSpace(spl[0])
		name := strings.TrimSpace(spl[1])
		if id == "id" {
			continue
		}
		mp := &Map{
			ID:   id,
			Name: name,
		}
		e.Maps[id] = mp
	}

	progbytes, err := ioutil.ReadFile(filepath.Join(DefaultEBPFFileSystemDataDir, "progs.debug"))
	if err != nil {
		return nil, fmt.Errorf("prog read: %v", err)
	}
	progstr := string(progbytes)
	lines = strings.Split(progstr, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		spl := strings.Split(line, " ")
		if len(spl) < 2 {
			continue
		}
		id := strings.TrimSpace(spl[0])
		name := strings.TrimSpace(spl[1])
		if id == "id" {
			continue
		}
		p := &Prog{
			ID:   id,
			Name: name,
		}
		e.Progs[id] = p
	}
	return e, nil
}
