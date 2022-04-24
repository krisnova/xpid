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
#include <stdio.h>
#include "proc.h"

int proc_pid_comm(pid_t pid, char *data){
  char *p = malloc(PROCFS_PATH_MAX);
  procfs_pid_file(p, pid, "comm");
  FILE *f;
  char ch;
  f = fopen(p, "r");
  if (f == NULL){
    free(p);
    return -1;
  }
  while (( ch = fgetc(f)) != EOF){
    sprintf(data,"%s%c", data, ch);
  }
  free(p);
  return 1;
}

int proc_pid_cmdline(pid_t pid, char *data){
  char *p = malloc(PROCFS_PATH_MAX);
  procfs_pid_file(p, pid, "cmdline");
  FILE *f;
  char ch;
  f = fopen(p, "r");
  if (f == NULL){
    free(p);
    return -1;
  }
  while (( ch = fgetc(f)) != EOF){
    sprintf(data,"%s%c", data, ch);
  }
  free(p);
  return 1;
}