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

/**
 * /proc/[pid]/comm (since Linux 2.6.33)
 *
 *  This file exposes the process's comm value—that is, the
 * command name associated with the process.  Different
 * threads in the same process may have different comm
 * values, accessible via /proc/[pid]/task/[tid]/comm.  A
 * thread may modify its comm value, or that of any of other
 * thread in the same thread group (see the discussion of
 * CLONE_THREAD in clone(2)), by writing to the file
 * /proc/self/task/[tid]/comm.  Strings longer than
 * TASK_COMM_LEN (16) characters (including the terminating
 * null byte) are silently truncated.
 *
 * This file provides a superset of the prctl(2) PR_SET_NAME
 * and PR_GET_NAME operations, and is employed by
 * pthread_setname_np(3) when used to rename threads other
 * than the caller.  The value in this file is used for the
 * %e specifier in /proc/sys/kernel/core_pattern; see
 * core(5).
 *
 * @param pid
 * @param data
 * @return
 */
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

/**
 * /proc/[pid]/cmdline
 * This read-only file holds the complete command line for
 * the process, unless the process is a zombie.  In the
 * latter case, there is nothing in this file: that is, a
 * read on this file will return 0 characters.  The command-
 * line arguments appear in this file as a set of strings
 * separated by null bytes ('\0'), with a further null byte
 * after the last string.
 *
 * If, after an execve(2), the process modifies its argv
 * strings, those changes will show up here.  This is not the
 * same thing as modifying the argv array.
 *
 * Furthermore, a process may change the memory location that
 * this file refers via prctl(2) operations such as
 * PR_SET_MM_ARG_START.
 *
 * Think of this file as the command line that the process
 * wants you to see.
 * 
 * @param pid
 * @param data
 * @return
 */
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