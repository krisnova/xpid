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

// clang-format off
#include "vmlinux.h"
// clang-format on
#include <bpf/bpf_helpers.h>

#include "xpid-bpf.h"

struct {
  __uint(type, BPF_MAP_TYPE_HASH);
  __uint(max_entries, PROBE_MAX_ENTRIES);
  __type(key, int);
  __type(value, int);
} probe_tracepoints_map SEC(".maps");

// name: inet_sock_set_state
// ID: 1373
// format:
//        field:unsigned short common_type;       offset:0;       size:2;
//        signed:0; field:unsigned char common_flags;       offset:2; size:1;
//        signed:0; field:unsigned char common_preempt_count;       offset:3;
//        size:1; signed:0; field:int common_pid;   offset:4;       size:4;
//        signed:1;
//
//        field:const void * skaddr;      offset:8;       size:8; signed:0;
//        field:int oldstate;     offset:16;      size:4; signed:1;
//        field:int newstate;     offset:20;      size:4; signed:1;
//        field:__u16 sport;      offset:24;      size:2; signed:0;
//        field:__u16 dport;      offset:26;      size:2; signed:0;
//        field:__u16 family;     offset:28;      size:2; signed:0;
//        field:__u16 protocol;   offset:30;      size:2; signed:0;
//        field:__u8 saddr[4];    offset:32;      size:4; signed:0;
//        field:__u8 daddr[4];    offset:36;      size:4; signed:0;
//        field:__u8 saddr_v6[16];        offset:40;      size:16; signed:0;
//        field:__u8 daddr_v6[16];        offset:56;      size:16; signed:0;
//
// print fmt: "family=%s protocol=%s sport=%hu dport=%hu saddr=%pI4 daddr=%pI4
// saddrv6=%pI6c daddrv6=%pI6c oldstate=%s newstate=%s",
// __print_symbolic(REC->family, { 2, "AF_INET" }, { 10, "AF_INET6" }),
// __print_symbolic(REC->protocol, { 6, "IPPROTO_TCP" }, { 33, "IPPROTO_DCCP" },
// { 132, "IPPROTO_SCTP" }, { 262, "IPPROTO_MPTCP" }), REC->sport, REC->dport,
// REC->saddr, REC->daddr, REC->saddr_v6, REC->daddr_v6,
// __print_symbolic(REC->oldstate, { 1, "TCP_ESTABLISHED" }, { 2, "TCP_SYN_SENT"
// }, { 3, "TCP_SYN_RECV" }, { 4, "TCP_FIN_WAIT1" }, { 5, "TCP_FIN_WAIT2" }, {
// 6, "TCP_TIME_WAIT" }, { 7, "TCP_CLOSE" }, { 8, "TCP_CLOSE_WAIT" }, { 9,
// "TCP_LAST_ACK" }, { 10, "TCP_LISTEN" }, { 11, "TCP_CLOSING" }, { 12,
// "TCP_NEW_SYN_RECV" }), __print_symbolic(REC->newstate, { 1, "TCP_ESTABLISHED"
// }, { 2, "TCP_SYN_SENT" }, { 3, "TCP_SYN_RECV" }, { 4, "TCP_FIN_WAIT1" }, { 5,
// "TCP_FIN_WAIT2" }, { 6, "TCP_TIME_WAIT" }, { 7, "TCP_CLOSE" }, { 8,
// "TCP_CLOSE_WAIT" }, { 9, "TCP_LAST_ACK" }, { 10, "TCP_LISTEN" }, { 11,
// "TCP_CLOSING" }, { 12, "TCP_NEW_SYN_RECV" })
SEC("tracepoint/sock/inet_sock_set_state")
int inet_sock_set_state(void *args) { return 0; }

// name: cgroup_attach_task
// ID: 432
// format:
//        field:unsigned short common_type;       offset:0;       size:2;
//        signed:0; field:unsigned char common_flags;       offset:2; size:1;
//        signed:0; field:unsigned char common_preempt_count;       offset:3;
//        size:1; signed:0; field:int common_pid;   offset:4;       size:4;
//        signed:1;
//
//        field:int dst_root;     offset:8;       size:4; signed:1;
//        field:int dst_level;    offset:12;      size:4; signed:1;
//        field:u64 dst_id;       offset:16;      size:8; signed:0;
//        field:int pid;  offset:24;      size:4; signed:1;
//        field:__data_loc char[] dst_path;       offset:28;      size:4;
//        signed:1; field:__data_loc char[] comm;   offset:32;      size:4;
//        signed:1;
//
// print fmt: "dst_root=%d dst_id=%llu dst_level=%d dst_path=%s pid=%d comm=%s",
// REC->dst_root, REC->dst_id, REC->dst_level, __get_str(dst_path), REC->pid,
// __get_str(comm)
SEC("tracepoint/cgroup/cgroup_attach_task")
int cgroup_attach_task(void *args) { return 0; }
