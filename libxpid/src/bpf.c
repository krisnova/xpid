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
      strncpy(name, "BPF_MAP_TYPE_UNSPEC", strlen(name));
      return;
    case 1:
      strncpy(name, "BPF_MAP_TYPE_HASH", strlen(name));
      return;
    case 2:
      strncpy(name, "BPF_MAP_TYPE_ARRAY", strlen(name));
      return;
    case 3:
      strncpy(name, "BPF_MAP_TYPE_PROG_ARRAY", strlen(name));
      return;
    case 4:
      strncpy(name, "BPF_MAP_TYPE_PERF_EVENT_ARRAY", strlen(name));
      return;
    case 5:
      strncpy(name, "BPF_MAP_TYPE_PERCPU_HASH", strlen(name));
      return;
    case 6:
      strncpy(name, "BPF_MAP_TYPE_PERCPU_ARRAY", strlen(name));
      return;
    case 7:
      strncpy(name, "BPF_MAP_TYPE_STACK_TRACE", strlen(name));
      return;
    case 8:
      strncpy(name, "BPF_MAP_TYPE_CGROUP_ARRAY", strlen(name));
      return;
    case 9:
      strncpy(name, "BPF_MAP_TYPE_LRU_HASH", strlen(name));
      return;
    case 10:
      strncpy(name, "BPF_MAP_TYPE_LRU_PERCPU_HASH", strlen(name));
      return;
    case 11:
      strncpy(name, "BPF_MAP_TYPE_LPM_TRIE", strlen(name));
      return;
    case 12:
      strncpy(name, "BPF_MAP_TYPE_ARRAY_OF_MAPS", strlen(name));
      return;
    case 13:
      strncpy(name, "BPF_MAP_TYPE_HASH_OF_MAPS", strlen(name));
      return;
    case 14:
      strncpy(name, "BPF_MAP_TYPE_DEVMAP", strlen(name));
      return;
    case 15:
      strncpy(name, "BPF_MAP_TYPE_SOCKMAP", strlen(name));
      return;
    case 16:
      strncpy(name, "BPF_MAP_TYPE_CPUMAP", strlen(name));
      return;
    case 17:
      strncpy(name, "BPF_MAP_TYPE_XSKMAP", strlen(name));
      return;
    case 18:
      strncpy(name, "BPF_MAP_TYPE_SOCKHASH", strlen(name));
      return;
    case 19:
      strncpy(name, "BPF_MAP_TYPE_CGROUP_STORAGE", strlen(name));
      return;
    case 20:
      strncpy(name, "BPF_MAP_TYPE_REUSEPORT_SOCKARRAY", strlen(name));
      return;
    case 21:
      strncpy(name, "BPF_MAP_TYPE_PERCPU_CGROUP_STORAGE", strlen(name));
      return;
    case 22:
      strncpy(name, "BPF_MAP_TYPE_QUEUE", strlen(name));
      return;
    case 23:
      strncpy(name, "BPF_MAP_TYPE_STACK", strlen(name));
      return;
    case 24:
      strncpy(name, "BPF_MAP_TYPE_SK_STORAGE", strlen(name));
      return;
    case 25:
      strncpy(name, "BPF_MAP_TYPE_DEVMAP_HASH", strlen(name));
      return;
    case 26:
      strncpy(name, "BPF_MAP_TYPE_STRUCT_OPS", strlen(name));
      return;
    case 27:
      strncpy(name, "BPF_MAP_TYPE_RINGBUF", strlen(name));
      return;
    case 28:
      strncpy(name, "BPF_MAP_TYPE_INODE_STORAGE", strlen(name));
      return;
    case 29:
      strncpy(name, "BPF_MAP_TYPE_TASK_STORAGE", strlen(name));
      return;
    case 30:
      strncpy(name, "BPF_MAP_TYPE_BLOOM_FILTER", strlen(name));
      return;
    case 31:
      strncpy(name, "BPF_MAP_TYPE_STRUCT_OPS", strlen(name));
      return;
    default:
      strncpy(name, "UNKNOWN", strlen(name));
      return;
  }
}

void bpf_map_type_enum(int i, char *name) {
  bpf_map_type_enum_linux_5_17(i, name);
}