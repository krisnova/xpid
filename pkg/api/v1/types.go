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

// arch_status      cpu_resctrl_groups  latency     net/           root@         statm
// attr/            cpuset              limits      ns/            sched         status
// autogroup        cwd@                loginuid    numa_maps      schedstat     syscall
// auxv             environ             map_files/  oom_adj        sessionid     task/
// cgroup           exe@                maps        oom_score      setgroups     timens_offsets
// clear_refs       fd/                 mem         oom_score_adj  smaps         timers
// cmdline          fdinfo/             mountinfo   pagemap        smaps_rollup  timerslack_ns
// comm             gid_map             mounts      personality    stack         uid_map
// coredump_filter  io                  mountstats  projid_map     stat          wchan

type Process struct {
	ProcessVisible

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

	// The process unique ID.
	PID int64
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
		ProcessVisible: ProcessVisible{},
		PID:            pid,
	}
}
