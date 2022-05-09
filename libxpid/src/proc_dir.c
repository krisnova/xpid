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

#include <dirent.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>

#include "proc.h"
#include "xpid.h"

/**
 * proc_dir_opendir will attempt to opendir(/proc/[pid])
 * and report.
 *
 * @param pid
 * @return
 */
int proc_dir_opendir(pid_t pid) {
  struct dirent *dent;
  char *p = malloc(PROCFS_PATH_MAX);
  procfs_pid(p, pid);
  DIR *dir = opendir(p);
  free(p);
  if (dir == NULL) {
    closedir(dir);
    return 0;
  }
  dent = readdir(dir);
  if (dent) {
    closedir(dir);
    return 1;
  }
  closedir(dir);
  return 2;
}

/**
 * proc_dir_chdir will attempt to chdir(/proc/[pid])
 * and report.
 *
 * @param pid
 * @return
 */

int proc_dir_chdir(pid_t pid) {
  int ret;
  char *p = malloc(PROCFS_PATH_MAX);
  procfs_pid(p, pid);
  ret = chdir(p);
  free(p);
  if (ret == -1) {
    return 0;
  }
  if (ret == 0) {
    return 1;
  }
  return 2;
}

/**
 * proc_dir_dent will attempt to opendir(/proc/[pid])
 * and list (dent) files and report.
 *
 * @param pid
 * @return
 */
int proc_dir_dent(pid_t pid) {
  char *pproc = malloc(PROCFS_PATH_MAX);
  char *pprocpid = malloc(PROCFS_PATH_MAX);
  procfs(pproc);
  procfs_pid(pprocpid, pid);
  struct dirent *dent;
  DIR *dir = opendir(pproc);
  if (dir == NULL) {
    closedir(dir);
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
  free(pproc);
  free(pprocpid);
  return 0;
}
