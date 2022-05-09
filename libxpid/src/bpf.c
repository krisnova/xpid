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
#include <stdlib.h>

void bpf_map_type_enum_linux_5_17(int i, char *name) {
  switch (i) {
    case 0 : strncpy(name, "BPF_MAP_TYPE_UNSPEC", 128);
    case 1 : strncpy(name, "BPF_MAP_TYPE_HASH", 128);
    case 2 : strncpy(name, "BPF_MAP_TYPE_ARRAY", 128);
    case 3 : strncpy(name, "BPF_MAP_TYPE_PROG_ARRAY", 128);
    case 4 : strncpy(name, "BPF_MAP_TYPE_PERF_EVENT_ARRAY", 128);
    case 5 : strncpy(name, "BPF_MAP_TYPE_PERCPU_HASH", 128);
    case 6 : strncpy(name, "BPF_MAP_TYPE_PERCPU_ARRAY", 128);
    case 7 : strncpy(name, "BPF_MAP_TYPE_STACK_TRACE", 128);
    case 8 : strncpy(name, "BPF_MAP_TYPE_CGROUP_ARRAY", 128);
    case 9 : strncpy(name, "BPF_MAP_TYPE_LRU_HASH", 128);
    case 10 : strncpy(name, "BPF_MAP_TYPE_LRU_PERCPU_HASH", 128);
    case 11 : strncpy(name, "BPF_MAP_TYPE_LPM_TRIE", 128);
    case 12 : strncpy(name, "BPF_MAP_TYPE_ARRAY_OF_MAPS", 128);
    case 13 : strncpy(name, "BPF_MAP_TYPE_HASH_OF_MAPS", 128);
    case 14 : strncpy(name, "BPF_MAP_TYPE_DEVMAP", 128);
    case 15 : strncpy(name, "BPF_MAP_TYPE_SOCKMAP", 128);
    case 16 : strncpy(name, "BPF_MAP_TYPE_CPUMAP", 128);
    case 17 : strncpy(name, "BPF_MAP_TYPE_XSKMAP", 128);
    case 18 : strncpy(name, "BPF_MAP_TYPE_SOCKHASH", 128);
    case 19 : strncpy(name, "BPF_MAP_TYPE_CGROUP_STORAGE", 128);
    case 20 : strncpy(name, "BPF_MAP_TYPE_REUSEPORT_SOCKARRAY", 128);
    case 21 : strncpy(name, "BPF_MAP_TYPE_PERCPU_CGROUP_STORAGE", 128);
    case 22 : strncpy(name, "BPF_MAP_TYPE_QUEUE", 128);
    case 23 : strncpy(name, "BPF_MAP_TYPE_STACK", 128);
    case 24 : strncpy(name, "BPF_MAP_TYPE_SK_STORAGE", 128);
    case 25 : strncpy(name, "BPF_MAP_TYPE_DEVMAP_HASH", 128);
    case 26 : strncpy(name, "BPF_MAP_TYPE_STRUCT_OPS", 128);
    case 27 : strncpy(name, "BPF_MAP_TYPE_RINGBUF", 128);
    case 28 : strncpy(name, "BPF_MAP_TYPE_INODE_STORAGE", 128);
    case 29 : strncpy(name, "BPF_MAP_TYPE_TASK_STORAGE", 128);
    case 30 : strncpy(name, "BPF_MAP_TYPE_BLOOM_FILTER", 128);
    case 31 : strncpy(name, "BPF_MAP_TYPE_STRUCT_OPS", 128);
    default:
      strncpy(name, "UNKNOWN", 128);
  }
}

void bpf_map_type_enum(int i, char *name) {
  char namestr[128];
  bpf_map_type_enum_linux_5_17(i, namestr);
  strncpy(name, namestr, 128);
  free(namestr);
}