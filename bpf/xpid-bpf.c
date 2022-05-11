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

#include <stdio.h>
#include <stdlib.h>
#include "xpid-bpf.h"

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
  return 0;
}