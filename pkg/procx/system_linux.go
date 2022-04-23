//go:build linux

/******************************************************************************
* MIT License
* Copyright (c) 2022 Kris Nóva <kris@nivenly.com>
*
* ┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓
* ┃   ███╗   ██╗ ██████╗ ██╗   ██╗ █████╗  ┃
* ┃   ████╗  ██║██╔═████╗██║   ██║██╔══██╗ ┃
* ┃   ██╔██╗ ██║██║██╔██║██║   ██║███████║ ┃
* ┃   ██║╚██╗██║████╔╝██║╚██╗ ██╔╝██╔══██║ ┃
* ┃   ██║ ╚████║╚██████╔╝ ╚████╔╝ ██║  ██║ ┃
* ┃   ╚═╝  ╚═══╝ ╚═════╝   ╚═══╝  ╚═╝  ╚═╝ ┃
* ┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛
*
*****************************************************************************/

package procx

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

const (
	DefaultProcLocation = "/proc"
)

func ProcPath() string {
	// TODO Logic to lookup if proc is not /proc
	return DefaultProcLocation
}

// MaxPid will return the system specific maximum PID number
//
// For Linux systems this can be found in proc(5) if it is mounted!
func MaxPid() int32 {
	maxPidFile := filepath.Join(ProcPath(), "sys/kernel/pid_max")
	bytes, err := os.ReadFile(maxPidFile)
	if err != nil {
		logrus.Warnf("err reading %s: %v", maxPidFile, err)
		return -1
	}
	v := string(bytes)
	v = strings.Replace(v, "\n", "", -1)
	vi, err := strconv.Atoi(v)
	if err != nil {
		logrus.Warnf("err reading %s: %v", maxPidFile, err)
		return -1
	}
	return int32(vi)
}
