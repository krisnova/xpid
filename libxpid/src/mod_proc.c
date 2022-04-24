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
#include <string.h>
#include <stdio.h>

#include "proc.h"
#include "xpid.h"

// proc_opendir will attempt to opendir(/proc/[pid])
// and report.
int proc_opendir(int pid){
  //printf("pid: %d\n", pid);
  struct dirent *dent;
  char *p = malloc(PROCFS_PATH_MAX);
  procfs_pid(p, pid);
  DIR *dir = opendir(p);
  free(p);
  if (dir == NULL) {
    closedir(dir);
    //printf("opendir: 0\n");
    return 0;
  }
  dent = readdir(dir);
  if (dent) {
    closedir(dir);
    //printf("opendir: 1\n");
    return 1;
  }
  closedir(dir);
  //printf("opendir: 2\n");
  return 2;
}

// proc_chdir will attempt to chdir(/proc/[pid])
// and report.
int proc_chdir(int pid) {
  int ret;
  char *p = malloc(PROCFS_PATH_MAX);
  procfs_pid(p, pid);
  ret = chdir(p);
  free(p);
  if (ret == -1) {
    //printf("chdir: 0\n");
    return 0;
  }
  if (ret == 0) {
    //printf("chdir: 1\n");
    return 1;
  }
  //printf("chdir: 2\n");
  return 2;
}


// proc_dent will attempt to opendir(/proc/[pid])
// and list (dent) files and report.
int proc_dent(int pid){
  char *pproc = malloc(PROCFS_PATH_MAX);
  char *pprocpid = malloc(PROCFS_PATH_MAX);
  procfs(pproc);
  procfs_pid(pprocpid, pid);
  struct dirent *dent;
  DIR *dir = opendir(pproc);

  if (dir == NULL) {
    closedir(dir);
    //printf("opendir: 0
    free(pproc);
    free(pprocpid);
    return 0;
  }
  while ((dent = readdir(dir)) != NULL) {
    char pidstr[PROCFS_PATH_MAX];
    sprintf(pidstr, "%d", pid);
    if (strncmp(dent->d_name, pidstr, sizeof dent->d_name) == 0) {
      closedir(dir);
      free(pproc);
      free(pprocpid);
      return 1;
    }
  }

  closedir(dir);
  //printf("opendir: 2\n");
  free(pproc);
  free(pprocpid);
  return 0;
}

