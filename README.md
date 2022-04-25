# xpid - It's `nmap` but for pids. 🤓

`xpid` gives a user the ability to "probe" for process details on a Linux system.

For example a sleeping thread will have a directory `/proc/[pid]` directoy that can be navigated to, but not listed.

`xpid` will check many different places in the kernel for details about a pid. 
By searching subsets of possible pids `xpid` will be able to check for pid details in many places in the kernel.

```bash
xpid [flags] -o [output]

Investigate pid 123 and write the report to out.txt
   xpid 123 > out.txt

Find all possible pids, and investigate each one (slow). The --all flag is default.
   xpid > out.txt 
   xpid --all > out.txt

Investigate all pids from 0 to 1000 and write the report to out.json
   xpid 0-1000 -o json > out.json

Find all eBPF pids at runtime (fast).
   xpid --ebpf

Find all proc pids at runtime (fast).
   xpid --proc

Investigate pid 123 using the "--proc" module only.
   xpid --proc 123 > out.txt
```

## Obfuscated PIDs

Because of the flexibility with kernel modules, and eBPF in the kernel it can be possible to prevent the `proc(5)` filesystem from listing pid details in traditional ways.

`xpid` uses other tactics to search for pids in the same way `nmap` will use multiple tactics to probe a target.

## Go runtime

`xpid` is a Go runtime utility that depends on `libxpid`.
Install `libxpid` and then compile the Go runtime.

```bash
git clone https://github.com/kris-nova/xpid.git
cd xpid
make
sudo make install
```

## Xpid C library (libxpid)

```bash 
git clone https://github.com/kris-nova/xpid.git
cd xpid/libxpid
./configure
cd build
make
sudo make install
```

