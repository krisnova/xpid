# Xpid

It's like `nmap` but for pids. 🤓

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

