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

#include "xpid-bpf.h"

#include <bpf/libbpf.h>
#include <stdio.h>
#include <stdlib.h>

void usage() {
  printf("xpid-bpf\n");
  printf("\n");
  printf("eBPF programs used for exercising xpid.\n");
  printf("\n");
  exit(0);
}

struct config {
  char probe_tracepoints[PATH_MAX];
} cfg;

void cmdsetup(int argc, char **argv) {
  // Set defaults here
  sprintf(cfg.probe_tracepoints, PROBE_DEFAULT_FMT, PROBE_TRACEPOINT);
  for (int i = 0; i < argc; i++) {
    if (argv[i][0] == '-') {
      switch (argv[i][1]) {
        case 'h':
          usage();
          break;
      }
    }
  }
}

int main(int argc, char **argv) {
  cmdsetup(argc, argv);

  // Probe Tracepoints
  struct bpf_object   *probetracepoints_obj;
  struct bpf_program  *probetracepoints_prog = NULL;

  int loaded;
  printf(" ->   eBPF Open %s\n", cfg.probe_tracepoints);
  probetracepoints_obj = bpf_object__open(cfg.probe_tracepoints);
  if (!probetracepoints_obj) {
    printf(" XX   Unable to open %s\n", cfg.probe_tracepoints);
    return 1;
  }
  bpf_object__load(probetracepoints_obj);
  bpf_object__next_map(probetracepoints_obj, NULL);
  bpf_object__for_each_program(probetracepoints_prog, probetracepoints_obj) {
    const char *progname = bpf_program__name(probetracepoints_prog);
    const char *progsecname = bpf_program__section_name(probetracepoints_prog);
    struct bpf_link *link = bpf_program__attach(probetracepoints_prog);
    if (!link) {
      return 1;
    }
    printf(" ->   eBPF Program Attached (Name) : %s\n", progname);
    printf(" ->   eBPF Program Attached (Sec)  : %s\n", progsecname);
  }
  printf(" ->   xpid map  : BPF_MAP_TYPE_HASH\n");
  printf(" ->   xpid prog : perf\n");


  printf("Hanging...\n");
  while(1){}

  return 0;
}