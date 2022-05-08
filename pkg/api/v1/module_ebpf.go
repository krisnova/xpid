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

	"github.com/kris-nova/xpid/pkg/libxpid"

	"github.com/kris-nova/xpid/pkg/procfs"
)

var _ ProcessExplorerModule = &EBPFModule{}

type EBPFModule struct {
	Mounts string
	Progs  []string
	Maps   []string
}

//enum bpf_map_type {
//	BPF_MAP_TYPE_UNSPEC,			0
//	BPF_MAP_TYPE_HASH,				1
//	BPF_MAP_TYPE_ARRAY,				2
//	BPF_MAP_TYPE_PROG_ARRAY,		3
//	BPF_MAP_TYPE_PERF_EVENT_ARRAY,	4
//	BPF_MAP_TYPE_PERCPU_HASH,		5
//	BPF_MAP_TYPE_PERCPU_ARRAY,		6
//	BPF_MAP_TYPE_STACK_TRACE,		7
//	BPF_MAP_TYPE_CGROUP_ARRAY,		.. so on
//	BPF_MAP_TYPE_LRU_HASH,
//	BPF_MAP_TYPE_LRU_PERCPU_HASH,
//	BPF_MAP_TYPE_LPM_TRIE,
//	BPF_MAP_TYPE_ARRAY_OF_MAPS,
//	BPF_MAP_TYPE_HASH_OF_MAPS,
//	BPF_MAP_TYPE_DEVMAP,
//	BPF_MAP_TYPE_SOCKMAP,
//	BPF_MAP_TYPE_CPUMAP,
//	BPF_MAP_TYPE_XSKMAP,
//	BPF_MAP_TYPE_SOCKHASH,
//	BPF_MAP_TYPE_CGROUP_STORAGE,
//	BPF_MAP_TYPE_REUSEPORT_SOCKARRAY,
//	BPF_MAP_TYPE_PERCPU_CGROUP_STORAGE,
//	BPF_MAP_TYPE_QUEUE,
//	BPF_MAP_TYPE_STACK,
//	BPF_MAP_TYPE_SK_STORAGE,
//	BPF_MAP_TYPE_DEVMAP_HASH,
//	BPF_MAP_TYPE_STRUCT_OPS,
//	BPF_MAP_TYPE_RINGBUF,
//	BPF_MAP_TYPE_INODE_STORAGE,
//	BPF_MAP_TYPE_TASK_STORAGE,
//	BPF_MAP_TYPE_BLOOM_FILTER,
//};

func NewEBPFModule() *EBPFModule {
	return &EBPFModule{}
}

const (
	// Taken from <linux/bpf.h>
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/bpf.h

	FileDescriptorMapIDKey  = "map_id"
	FileDescriptorProgIDKey = "prog_id"
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

	bpfDebug, err := NewEBPFFileSystemData()
	if err != nil {
		return fmt.Errorf("unable to read /sys/fs/bpf: %v", err)
	}

	// Compare with file descriptors in fdinfo

	// [root@emily]: /proc/141735/fdinfo># cat 17
	//pos:    0
	//flags:  02000000
	//mnt_id: 15
	//ino:    10586
	//link_type:      perf
	//link_id:        19
	//prog_tag:       40bd9646d9b53ff8
	//prog_id:        106

	fds, err := procfshandle.DirPID(p.PID, "fdinfo")

	if err != nil {
		return fmt.Errorf("unable to read /proc/%d/fdinfo: %v", p.PID, err)
	}

	// File descriptor scanning
	//
	// Here we try to map the file descriptor keys (map_id, prog_id)
	// back to the established values found in the progs.debug and maps.debug
	// sys filesystem
	//
	for _, fd := range fds {
		fddata, err := procfshandle.ContentsPID(p.PID, filepath.Join("fdinfo", fd.Name()))
		if err != nil {
			continue
		}
		fdProgID := procfs.FileKeyValue(fddata, FileDescriptorProgIDKey)
		fdMapID := procfs.FileKeyValue(fddata, FileDescriptorMapIDKey)

		// Map back to /sys/fs/bpf/progs.debug
		for id, _ := range bpfDebug.Progs {
			if id == "" {
				continue
			}
			if id == fdProgID {
				// We have mapped an eBPF program to a PID!
				p.EBPF = true
				progDetails := programDetails(p, fddata)
				if !strings.Contains(strings.Join(p.EBPFModule.Progs, ""), progDetails) {
					p.EBPFModule.Progs = append(p.EBPFModule.Progs, progDetails)
				}
			}
		}

		// Map back to /sys/fs/bpf/maps.debug
		for id, _ := range bpfDebug.Maps {
			if id == "" {
				continue
			}
			if id == fdMapID {
				// We have mapped an eBPF program to a PID!
				p.EBPF = true
				mapDetails := mapDetails(p, fddata)
				if !strings.Contains(strings.Join(p.EBPFModule.Maps, ""), mapDetails) {
					p.EBPFModule.Maps = append(p.EBPFModule.Maps, mapDetails)
				}
			}
		}
	}

	// Hacking in here during the stream
	libxpid.BPFTodo()

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

// NewEBPFFileSystemData will read from /sys/fs/bpf/[maps.debug, progs.debug]
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

		// Parse the file
		spl := strings.Split(line, " ")
		var name, id string
		if len(spl) < 2 {
			name = ""
		} else {
			name = strings.TrimSpace(spl[1])
		}
		id = strings.TrimSpace(spl[0])

		// Ignore headers
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

		// Parse the file
		spl := strings.Split(line, " ")
		var name, id string
		if len(spl) < 2 {
			name = ""
		} else {
			name = strings.TrimSpace(spl[1])
		}
		id = strings.TrimSpace(spl[0])

		// Ignore headers
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

// fddata is the filedescriptor data
func programDetails(p *Process, fddata string) string {
	return "prog"
}

// fddata is the filedescriptor data
func mapDetails(p *Process, fddata string) string {
	return "map"
}
