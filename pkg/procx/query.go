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
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"

	api "github.com/kris-nova/xpid/pkg/api/v1"
)

// PIDQuery will take an nmap like query for linux pids
func PIDQuery(raw string) []*api.Process {
	var processes []*api.Process
	raw = strings.TrimSpace(raw)
	if strings.Contains(raw, "-") {
		spl := strings.Split(raw, "-")
		if len(spl) != 2 {
			return nil
		}
		left, err := strconv.Atoi(spl[0])
		if err != nil {
			logrus.Warnf("invalid pid query: %v\n", err)
			return nil
		}
		right, err := strconv.Atoi(spl[1])
		if err != nil {
			logrus.Warnf("invalid pid query: %v\n", err)
			return nil
		}
		if left > right {
			sw := left
			left = right
			right = sw
		}
		for i := left; i <= right; i++ {
			p := api.ProcessPID(int64(i))
			processes = append(processes, p)
		}
	} else if strings.HasSuffix(raw, "+") {
		raw = strings.TrimSuffix(raw, "+")
		pid, err := strconv.Atoi(raw)
		if err != nil {
			logrus.Warnf("invalid pid query: %v\n", err)
			return nil
		}
		left := pid
		right := int(MaxPid())
		for i := left; i <= right; i++ {
			p := api.ProcessPID(int64(i))
			processes = append(processes, p)
		}
	} else {
		pid, err := strconv.Atoi(raw)
		if err != nil {
			logrus.Warnf("invalid pid query: %v\n", err)
			return nil
		}
		p := api.ProcessPID(int64(pid))
		processes = append(processes, p)
	}
	return processes
}
