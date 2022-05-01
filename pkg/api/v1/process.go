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

import "os/user"

// arch_status      cpu_resctrl_groups  latency     net/           root@         statm
// attr/            cpuset              limits      ns/            sched         status
// autogroup        cwd@                loginuid    numa_maps      schedstat     syscall
// auxv             environ             map_files/  oom_adj        sessionid     task/
// cgroup           exe@                maps        oom_score      setgroups     timens_offsets
// clear_refs       fd/                 mem         oom_score_adj  smaps         timers
// cmdline          fdinfo/             mountinfo   pagemap        smaps_rollup  timerslack_ns
// comm             gid_map             mounts      personality    stack         uid_map
// coredump_filter  io                  mountstats  projid_map     stat          wchan

// Process is a Linux process abstraction.
// By design this will have fields that are available in /proc and much more.
//
// Some of these fields will be calculated at runtime based on certain situations
// in the system.
//
// This data structure and the logic that populates it will be a substantial part
// of the xpid library, and the xpid API.
type Process struct {

	// The process unique ID.
	PID int64

	// ProcessVisible is a combination of the values
	// we get from libxpid that will determine if a process
	// running in Linux is visible or not
	ProcessVisible

	// User is the user associated with the process
	//
	// Group is a subset of User
	User

	// Name (proc/[pid]/comm)
	// This file exposes the process's comm value—that is, the
	// command name associated with the process.
	Name string

	// CommandLine (/proc/[pid]/cmdline)
	// This read-only file holds the complete command line for
	// the process, unless the process is a zombie.  In the
	// latter case, there is nothing in this file: that is, a
	// read on this file will return 0 characters
	CommandLine string

	// EBPF is an xpid specific field that attempts to detect
	// if a specific PID is running any eBPF related probes or maps.
	//
	// Will be set to true is eBPF is detected.
	EBPF bool

	// Container is an xpid specific field that attempts to detect
	// if a specific PID is running as a container
	//
	// Will be set to true if we suspect this is a container PID
	Container bool

	// Thread is a bool that will be set if the process is part of
	// a thread pool, or has a TGID that is unique from PID
	//
	// Reading from /proc/[pid]/status
	Thread bool

	// Modules are embedded directly here

	EBPFModule
	ProcModule
	ContainerModule
	NamespaceModule
}

// User is a user from the filesystem
// This can be populated from process details, or from
// the current user context.
type User struct {
	user.User
	Group
	ID   int
	Name string
}

type Group struct {
	ID   int
	Name string
}

type ProcessVisible struct {

	// Opendir is if the /proc/[pid] directory can be "opened" or "listed".
	// Failing Opendir is a sign that the process may be attempted to being
	// obfuscated to the user at runtime.
	Opendir int

	// Chdir is if the /proc/[pid] directory can be "navigated" or "changed to".
	// Failing chdir is a sign that the current user has invalid permission,
	// or that something in the kernel is preventing the user from open the directory.
	Chdir int

	// After opendir we see if we can "list" files inside of the directory.
	// This call happens at a higher level and will see if a directory
	// within proc can be found by opening it's parent directory for listing.
	//
	// In this case /proc is typically opened, and then the pid directories are
	// matched against!
	Dent int
}

func ProcessPID(pid int64) *Process {
	return &Process{
		ProcessVisible:  ProcessVisible{},
		User:            User{},
		PID:             pid,
		EBPFModule:      EBPFModule{},
		ProcModule:      ProcModule{},
		ContainerModule: ContainerModule{},
		NamespaceModule: NamespaceModule{},
	}
}
