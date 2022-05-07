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
	"io/ioutil"
	"strconv"
	"strings"

	api "github.com/kris-nova/xpid/pkg/api/v1"
)

func ProcListingQuery(raw string) []*api.Process {
	var processes []*api.Process
	raw = strings.TrimSpace(raw)
	left, right := queryLeftRight(raw)
	if left == -1 || right == -1 {
		return nil
	}
	files, err := ioutil.ReadDir(ProcPath())
	if err != nil {
		return nil
	}
	for _, file := range files {
		if file.IsDir() {
			i, e := strconv.Atoi(file.Name())
			if e == nil {
				// We have a dir and its name is a number :)
				if i >= left && i <= right {
					processes = append(processes, &api.Process{
						PID: int64(i),
					})
				}
			}
		}
	}
	return processes
}

// PIDQuery will take an nmap like query for linux pids
//
// PIDQuery will return ALL POSSIBLE PIDS!
// PIDQuery will not pull a list from /proc but rather build a list of potential pids.
func PIDQuery(raw string) []*api.Process {
	var processes []*api.Process
	raw = strings.TrimSpace(raw)
	left, right := queryLeftRight(raw)
	if left == -1 || right == -1 {
		return nil
	}
	for i := left; i <= right; i++ {
		p := api.ProcessPID(int64(i))
		processes = append(processes, p)
	}
	return processes
}

func queryLeftRight(raw string) (left, right int) {
	if strings.HasPrefix(raw, "+") {
		raw = strings.TrimPrefix(raw, "+")
		pid, err := strconv.Atoi(raw)
		if err != nil {
			return -1, -1
		}
		left := 0
		right := pid
		return left, right
	} else if strings.Contains(raw, "-") {
		spl := strings.Split(raw, "-")
		if len(spl) != 2 {
			return -1, -1
		}
		left, err := strconv.Atoi(spl[0])
		if err != nil {
			return -1, -1
		}
		right, err := strconv.Atoi(spl[1])
		if err != nil {
			return -1, -1
		}
		if left > right {
			sw := left
			left = right
			right = sw
		}
		return left, right
		//for i := left; i <= right; i++ {
		//	p := api.ProcessPID(int64(i))
		//	processes = append(processes, p)
		//}
	} else if strings.HasSuffix(raw, "+") {
		raw = strings.TrimSuffix(raw, "+")
		pid, err := strconv.Atoi(raw)
		if err != nil {
			return -1, -1
		}
		left := pid
		right := int(MaxPid())
		return left, right
		//for i := left; i <= right; i++ {
		//	p := api.ProcessPID(int64(i))
		//	processes = append(processes, p)
		//}
	} else {
		pid, err := strconv.Atoi(raw)
		if err != nil {
			return -1, -1
		}
		return pid, pid
	}
}
