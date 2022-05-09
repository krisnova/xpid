/*===========================================================================*\
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

package libxpid

// #cgo LDFLAGS: -lxpid
//
// #include "xpid.h"
// #include "stdlib.h"
import "C"
import (
	"sync"
	"unsafe"
)

const (
	Empty8   string = "        "
	Empty16  string = Empty8 + Empty8
	Empty32  string = Empty16 + Empty16
	Empty64  string = Empty32 + Empty32
	Empty128 string = Empty64 + Empty64
	Empty256 string = Empty128 + Empty128
)

// TODO We need to see what pid details we can get out of the kernel

var libxpidbpfenummtx sync.Mutex

func BPFMapType(mapType int) string {
	libxpidbpfenummtx.Lock()
	defer libxpidbpfenummtx.Unlock()
	name := Empty256
	cname := C.CString(name)
	C.bpf_map_type_enum(C.int(mapType), cname)
	defer C.free(unsafe.Pointer(cname))
	return C.GoString(cname)
}

//
//func ProcPidComm(pid int64) string {

//	defer C.free(unsafe.Pointer(cdata))
//	xint := int(x)
//	if xint == 1 {
//		retstr := strings.ReplaceAll(C.GoString(cdata), "\n", "")
//		return retstr
//	}
//	return ""
//}
//
//func ProcPidCmdline(pid int64) string {
//	var data string
//	cdata := C.CString(data)
//	defer C.free(unsafe.Pointer(cdata))
//	x := C.proc_pid_cmdline(C.int(int(pid)), cdata)
//	xint := int(x)
//	if xint == 1 {
//		retstr := strings.ReplaceAll(C.GoString(cdata), "\n", "")
//		return retstr
//	}
//	return ""
//}
