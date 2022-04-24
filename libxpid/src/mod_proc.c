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

// Proc Module

#include <dirent.h>
#include <unistd.h>
#include <stdlib.h>

#include "proc.h"
#include "xpid.h"

int proc_dir_ls(int pid) {
  struct dirent *dent;
  char *p = malloc(PROCFS_PATH_MAX);
  procfs_pid(p, pid);
  DIR *dir = opendir(p);
  free(p);
  if (dir == NULL) {
    return 0;
  }
  dent = readdir(dir);
  closedir(dir);
  if (dent) {
    return 1;
  }
  return 2;
}

int proc_dir_nav(int pid) {
  int ret;
  char *p = malloc(PROCFS_PATH_MAX);
  procfs_pid(p, pid);
  ret = chdir(p);
  free(p);
  if (ret == 0) {
    return 1;
  }
  return 2;
}
