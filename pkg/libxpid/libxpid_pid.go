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
	"strings"
	"unsafe"
)

func ProcPidComm(pid int64) string {
	var data string
	cdata := C.CString(data)
	x := C.proc_pid_comm(C.int(int(pid)), cdata)
	xint := int(x)
	if xint == 1 {
		retstr := strings.ReplaceAll(C.GoString(cdata), "\n", "")
		C.free(unsafe.Pointer(cdata))
		return retstr
	}
	return ""
}

func ProcPidCmdline(pid int64) string {
	var data string
	cdata := (*C.char)(C.CString(data))
	defer C.free(unsafe.Pointer(cdata))
	x := C.proc_pid_cmdline(C.int(int(pid)), cdata)
	xint := int(x)
	if xint == 1 {
		retstr := strings.ReplaceAll(C.GoString(cdata), "\n", "")
		C.free(unsafe.Pointer(cdata))
		return retstr
	}
	return ""
}
