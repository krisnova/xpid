# libxpid

The internal C library for `xpid`

The separation between `xpid` and `libxpid` is not as clean as it could be.

The long term goal of this library is to serve as a C implementation of all the `xpid` internal sourcing which will include:

 - ptrace support for pids at runtime
 - eBPF support for pids at runtime
 - glibc support for procfs
 - proc parsing
 - kernel module for interfacing with the pid table at runtime

For now this serves as a reliable hook for ad-hoc `xpid` functionality.
As the needs of `xpid` mature, I will slowly begin to migrate from Go directly into C. 
In other words, I am prototyping in Go until its obvious that the feature needs to be backported to C.