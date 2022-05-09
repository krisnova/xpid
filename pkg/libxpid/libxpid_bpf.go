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
	"sync"
	"unsafe"
)

const (

	// Memory in C is hard

	Empty1    string = " "
	Empty2    string = Empty1 + Empty1
	Empty4    string = Empty2 + Empty2
	Empty8    string = Empty4 + Empty4
	Empty16   string = Empty8 + Empty8
	Empty32   string = Empty16 + Empty16
	Empty64   string = Empty32 + Empty32
	Empty128  string = Empty64 + Empty64
	Empty256  string = Empty128 + Empty128
	Empty512  string = Empty256 + Empty256
	Empty1024 string = Empty512 + Empty512
)

var libxpidbpfenummtx sync.Mutex

func BPFProgramSections(progId int) []string {
	// void bpf_program_details(__u32 id, char *sec);
	secstr := Empty1024
	csecstr := C.CString(secstr)
	defer C.free(unsafe.Pointer(csecstr))
	C.bpf_program_details(C.__u32(progId), csecstr)
	secstr = C.GoString(csecstr)
	return strings.Split(secstr, ":")
}

func BPFMapType(mapType int) string {
	libxpidbpfenummtx.Lock()
	defer libxpidbpfenummtx.Unlock()
	name := Empty256
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	C.bpf_map_type_enum(C.int(mapType), cname)
	return C.GoString(cname)
}
