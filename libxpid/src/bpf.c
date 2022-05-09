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

#include <stdlib.h>
#include <string.h>

void bpf_map_type_enum_linux_5_17(int i, char *name) {
  switch (i) {
    case 0:
      strncpy(name, "BPF_MAP_TYPE_UNSPEC", 256);
      return;
    case 1:
      strncpy(name, "BPF_MAP_TYPE_HASH", 256);
      return;
    case 2:
      strncpy(name, "BPF_MAP_TYPE_ARRAY", 256);
      return;
    case 3:
      strncpy(name, "BPF_MAP_TYPE_PROG_ARRAY", 256);
      return;
    case 4:
      strncpy(name, "BPF_MAP_TYPE_PERF_EVENT_ARRAY", 256);
      return;
    case 5:
      strncpy(name, "BPF_MAP_TYPE_PERCPU_HASH", 256);
      return;
    case 6:
      strncpy(name, "BPF_MAP_TYPE_PERCPU_ARRAY", 256);
      return;
    case 7:
      strncpy(name, "BPF_MAP_TYPE_STACK_TRACE", 256);
      return;
    case 8:
      strncpy(name, "BPF_MAP_TYPE_CGROUP_ARRAY", 256);
      return;
    case 9:
      strncpy(name, "BPF_MAP_TYPE_LRU_HASH", 256);
      return;
    case 10:
      strncpy(name, "BPF_MAP_TYPE_LRU_PERCPU_HASH", 256);
      return;
    case 11:
      strncpy(name, "BPF_MAP_TYPE_LPM_TRIE", 256);
      return;
    case 12:
      strncpy(name, "BPF_MAP_TYPE_ARRAY_OF_MAPS", 256);
      return;
    case 13:
      strncpy(name, "BPF_MAP_TYPE_HASH_OF_MAPS", 256);
      return;
    case 14:
      strncpy(name, "BPF_MAP_TYPE_DEVMAP", 256);
      return;
    case 15:
      strncpy(name, "BPF_MAP_TYPE_SOCKMAP", 256);
      return;
    case 16:
      strncpy(name, "BPF_MAP_TYPE_CPUMAP", 256);
      return;
    case 17:
      strncpy(name, "BPF_MAP_TYPE_XSKMAP", 256);
      return;
    case 18:
      strncpy(name, "BPF_MAP_TYPE_SOCKHASH", 256);
      return;
    case 19:
      strncpy(name, "BPF_MAP_TYPE_CGROUP_STORAGE", 256);
      return;
    case 20:
      strncpy(name, "BPF_MAP_TYPE_REUSEPORT_SOCKARRAY", 256);
      return;
    case 21:
      strncpy(name, "BPF_MAP_TYPE_PERCPU_CGROUP_STORAGE", 256);
      return;
    case 22:
      strncpy(name, "BPF_MAP_TYPE_QUEUE", 256);
      return;
    case 23:
      strncpy(name, "BPF_MAP_TYPE_STACK", 256);
      return;
    case 24:
      strncpy(name, "BPF_MAP_TYPE_SK_STORAGE", 256);
      return;
    case 25:
      strncpy(name, "BPF_MAP_TYPE_DEVMAP_HASH", 256);
      return;
    case 26:
      strncpy(name, "BPF_MAP_TYPE_STRUCT_OPS", 256);
      return;
    case 27:
      strncpy(name, "BPF_MAP_TYPE_RINGBUF", 256);
      return;
    case 28:
      strncpy(name, "BPF_MAP_TYPE_INODE_STORAGE", 256);
      return;
    case 29:
      strncpy(name, "BPF_MAP_TYPE_TASK_STORAGE", 256);
      return;
    case 30:
      strncpy(name, "BPF_MAP_TYPE_BLOOM_FILTER", 256);
      return;
    case 31:
      strncpy(name, "BPF_MAP_TYPE_STRUCT_OPS", 256);
      return;
    default:
      strncpy(name, "UNKNOWN", 256);
      return;
  }
}

void bpf_map_type_enum(int i, char *name) {
  bpf_map_type_enum_linux_5_17(i, name);
}