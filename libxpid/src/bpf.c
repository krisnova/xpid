/* ==========================================================================*\
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

#include <string.h>

void bpf_map_type_enum_linux_5_17(int i, char *name) {
  switch (i) {
    case 0 : strncpy(name, "BPF_MAP_TYPE_UNSPEC", 1024);
    case 1 : strncpy(name, "BPF_MAP_TYPE_HASH", 1024);
    case 2 : strncpy(name, "BPF_MAP_TYPE_ARRAY", 1024);
    case 3 : strncpy(name, "BPF_MAP_TYPE_PROG_ARRAY", 1024);
    case 4 : strncpy(name, "BPF_MAP_TYPE_PERF_EVENT_ARRAY", 1024);
    case 5 : strncpy(name, "BPF_MAP_TYPE_PERCPU_HASH", 1024);
    case 6 : strncpy(name, "BPF_MAP_TYPE_PERCPU_ARRAY", 1024);
    case 7 : strncpy(name, "BPF_MAP_TYPE_STACK_TRACE", 1024);
    case 8 : strncpy(name, "BPF_MAP_TYPE_CGROUP_ARRAY", 1024);
    case 9 : strncpy(name, "BPF_MAP_TYPE_LRU_HASH", 1024);
    case 10 : strncpy(name, "BPF_MAP_TYPE_LRU_PERCPU_HASH", 1024);
    case 11 : strncpy(name, "BPF_MAP_TYPE_LPM_TRIE", 1024);
    case 12 : strncpy(name, "BPF_MAP_TYPE_ARRAY_OF_MAPS", 1024);
    case 13 : strncpy(name, "BPF_MAP_TYPE_HASH_OF_MAPS", 1024);
    case 14 : strncpy(name, "BPF_MAP_TYPE_DEVMAP", 1024);
    case 15 : strncpy(name, "BPF_MAP_TYPE_SOCKMAP", 1024);
    case 16 : strncpy(name, "BPF_MAP_TYPE_CPUMAP", 1024);
    case 17 : strncpy(name, "BPF_MAP_TYPE_XSKMAP", 1024);
    case 18 : strncpy(name, "BPF_MAP_TYPE_SOCKHASH", 1024);
    case 19 : strncpy(name, "BPF_MAP_TYPE_CGROUP_STORAGE", 1024);
    case 20 : strncpy(name, "BPF_MAP_TYPE_REUSEPORT_SOCKARRAY", 1024);
    case 21 : strncpy(name, "BPF_MAP_TYPE_PERCPU_CGROUP_STORAGE", 1024);
    case 22 : strncpy(name, "BPF_MAP_TYPE_QUEUE", 1024);
    case 23 : strncpy(name, "BPF_MAP_TYPE_STACK", 1024);
    case 24 : strncpy(name, "BPF_MAP_TYPE_SK_STORAGE", 1024);
    case 25 : strncpy(name, "BPF_MAP_TYPE_DEVMAP_HASH", 1024);
    case 26 : strncpy(name, "BPF_MAP_TYPE_STRUCT_OPS", 1024);
    case 27 : strncpy(name, "BPF_MAP_TYPE_RINGBUF", 1024);
    case 28 : strncpy(name, "BPF_MAP_TYPE_INODE_STORAGE", 1024);
    case 29 : strncpy(name, "BPF_MAP_TYPE_TASK_STORAGE", 1024);
    case 30 : strncpy(name, "BPF_MAP_TYPE_BLOOM_FILTER", 1024);
    case 31 : strncpy(name, "BPF_MAP_TYPE_STRUCT_OPS", 1024);
    default:
      strncpy(name, "UNKNOWN", 1024);
  }
}

void bpf_map_type_enum(int i, char *name) {
  return bpf_map_type_enum_linux_5_17(i, name);
}