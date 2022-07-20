# xpid 

It's [`nmap`](https://nmap.org/) but for pids. 🤓

---

Please help me become an independent programmer by donating directly to my `ko-fi` site below.

[![ko-fi](https://ko-fi.com/img/githubbutton_sm.svg)](https://ko-fi.com/D1D8CXLHZ) 

---

`xpid` gives a user the ability to "investigate" for process details on a Linux system.

For example a sleeping thread will have a directory `/proc/[pid]` that can be navigated to, but not listed.

`xpid` will check many different places in the kernel for details about a pid. 
By searching subsets of possible pids `xpid` will be able to check for pid details in many places in the kernel.

```
xpid [flags] -o [output] <query>

Investigate all pids
  xpid

Investigate pid 1
  xpid 1

Find all container processes on a system
  xpid -c

Find all processes in the same namespace(s) as pid 1
  xpid --ns-in [mnt, net, pid, ipc, cgroup]

Find all processes not in the same namespace(s) as current user
  xpid --ns-out-user [mnt, net, pid, ipc, cgroup]

Find all processes running with eBPF programs as JSON
  xpid --ebpf -o json <pid-query>

Find all processes running with eBPF programs, in a container, in /proc
  xpid -b -c -p

Find all processes between specific values (Query syntax)
  xpid <flags> +100      # Search pids up to 100
  xpid <flags> 100-2000  # Search pids between 100-2000 
  xpid <flags> 65000+    # Search pids 65000 or above

Find all hidden processes on a system (slow)
  xpid -x <pid-query>

```

## Container pids (xpid -c) 📦

`xpid` will lookup container processes at runtime. 🎉

This works by reading the link in [`/proc/[pid]/ns/@cgroup`](https://man7.org/linux/man-pages/man7/namespaces.7.html#:~:text=/proc/%5Bpid%5D/ns/cgroup) and correlating it back to the value in `/proc/1/[pid]/ns/@cgroup`.

Regardless of the pid namespace context, if there is a "container" that is unique from the current pid 1, `xpid` will find it.

## eBPF pids (xpid -b) 🐝

`xpid` will find pids that have eBPF programs loaded at runtime.

This works by correlating the file descriptor info from [`/proc/[pid]/fdinfo/*`](https://man7.org/linux/man-pages/man5/proc.5.html#:~:text=file%20descriptor%200.-,/proc/%5Bpid%5D/fdinfo/,-(since%20Linux%202.6.22)) back to `/sys/fs/bpf/progs.debug`. 
If a pid has an eBPF program loaded, `xpid` will find it.

## Hidden pids (xpid -x) 🙈

Because of the flexibility with kernel modules and eBPF in the kernel, it can be possible to prevent the [`proc(5)`](https://man7.org/linux/man-pages/man5/proc.5.html) filesystem from listing pid details in traditional ways.

`xpid` uses a variety of tactics to search for pids in the same way `nmap` will use different tactics to port scan a target.

## Go runtime

`xpid` is a Go runtime utility that depends on `libxpid`.
Install `libxpid` first (below), and then compile the Go runtime.

```bash
git clone https://github.com/kris-nova/xpid.git
cd xpid
make
sudo make install
```

## Xpid C library (libxpid)

`libxpid` is written in C, as it will leverage [`ptrace(2)`](https://man7.org/linux/man-pages/man2/ptrace.2.html) and eBPF code directly. 
This means that the `xpid` executable is NOT entirely statically linked. 
You must first have `libxpid` installed on your system, before the `xpid` Go program will run.

```bash 
git clone https://github.com/kris-nova/xpid.git
cd xpid/libxpid
./configure
cd build
make
sudo make install
```

