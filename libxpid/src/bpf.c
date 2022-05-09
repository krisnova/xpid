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
    case 0 : strncpy(name, "BPF_MAP_TYPE_UNSPEC", strlen(name));
    case 1 : strncpy(name, "BPF_MAP_TYPE_HASH", strlen(name));
    case 2 : strncpy(name, "BPF_MAP_TYPE_ARRAY", strlen(name));
    case 3 : strncpy(name, "BPF_MAP_TYPE_PROG_ARRAY", strlen(name));
    case 4 : strncpy(name, "BPF_MAP_TYPE_PERF_EVENT_ARRAY", strlen(name));
    case 5 : strncpy(name, "BPF_MAP_TYPE_PERCPU_HASH", strlen(name));
    case 6 : strncpy(name, "BPF_MAP_TYPE_PERCPU_ARRAY", strlen(name));
    case 7 : strncpy(name, "BPF_MAP_TYPE_STACK_TRACE", strlen(name));
    case 8 : strncpy(name, "BPF_MAP_TYPE_CGROUP_ARRAY", strlen(name));
    case 9 : strncpy(name, "BPF_MAP_TYPE_LRU_HASH", strlen(name));
    case 10 : strncpy(name, "BPF_MAP_TYPE_LRU_PERCPU_HASH", strlen(name));
    case 11 : strncpy(name, "BPF_MAP_TYPE_LPM_TRIE", strlen(name));
    case 12 : strncpy(name, "BPF_MAP_TYPE_ARRAY_OF_MAPS", strlen(name));
    case 13 : strncpy(name, "BPF_MAP_TYPE_HASH_OF_MAPS", strlen(name));
    case 14 : strncpy(name, "BPF_MAP_TYPE_DEVMAP", strlen(name));
    case 15 : strncpy(name, "BPF_MAP_TYPE_SOCKMAP", strlen(name));
    case 16 : strncpy(name, "BPF_MAP_TYPE_CPUMAP", strlen(name));
    case 17 : strncpy(name, "BPF_MAP_TYPE_XSKMAP", strlen(name));
    case 18 : strncpy(name, "BPF_MAP_TYPE_SOCKHASH", strlen(name));
    case 19 : strncpy(name, "BPF_MAP_TYPE_CGROUP_STORAGE", strlen(name));
    case 20 : strncpy(name, "BPF_MAP_TYPE_REUSEPORT_SOCKARRAY", strlen(name));
    case 21 : strncpy(name, "BPF_MAP_TYPE_PERCPU_CGROUP_STORAGE", strlen(name));
    case 22 : strncpy(name, "BPF_MAP_TYPE_QUEUE", strlen(name));
    case 23 : strncpy(name, "BPF_MAP_TYPE_STACK", strlen(name));
    case 24 : strncpy(name, "BPF_MAP_TYPE_SK_STORAGE", strlen(name));
    case 25 : strncpy(name, "BPF_MAP_TYPE_DEVMAP_HASH", strlen(name));
    case 26 : strncpy(name, "BPF_MAP_TYPE_STRUCT_OPS", strlen(name));
    case 27 : strncpy(name, "BPF_MAP_TYPE_RINGBUF", strlen(name));
    case 28 : strncpy(name, "BPF_MAP_TYPE_INODE_STORAGE", strlen(name));
    case 29 : strncpy(name, "BPF_MAP_TYPE_TASK_STORAGE", strlen(name));
    case 30 : strncpy(name, "BPF_MAP_TYPE_BLOOM_FILTER", strlen(name));
    case 31 : strncpy(name, "BPF_MAP_TYPE_STRUCT_OPS", strlen(name));
    default:
      strncpy(name, "UNKNOWN", strlen(name));
  }
}

void bpf_map_type_enum(int i, char *name) {
  return bpf_map_type_enum_linux_5_17(i, name);
}