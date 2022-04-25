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
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

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
