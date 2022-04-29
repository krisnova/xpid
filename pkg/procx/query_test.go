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
	"testing"
)

func TestPIDQueryDash(t *testing.T) {
	test := "1-100"
	pids := PIDQuery(test)
	if len(pids) != 100 {
		t.Errorf("invalid PID query: missing 1-100")
	}
	for i := 0; i < 100; i++ {
		pid := pids[i]
		if int(pid.PID) != i+1 {
			t.Errorf("invalid PID %d:%d", int(pid.PID), int(i))
		}
	}
}

func TestPIDQuery(t *testing.T) {
	test := "100"
	pids := PIDQuery(test)
	if len(pids) != 1 {
		t.Errorf("invalid PID query: missing 100")
	}
	if pids[0].PID != int64(100) {
		t.Errorf("invalid PID: 100")
	}
}

func TestPIDQueryDashSad(t *testing.T) {
	test := "11-da-asdf---100--"
	pids := PIDQuery(test)
	if pids != nil {
		t.Errorf("error expected")
	}
}

func TestPIDQueryDashReverse(t *testing.T) {
	test := "100-1"
	pids := PIDQuery(test)
	if len(pids) != 100 {
		t.Errorf("invalid PID query: missing 1-100")
	}
	for i := 0; i < 100; i++ {
		pid := pids[i]
		if int(pid.PID) != i+1 {
			t.Errorf("invalid PID %d:%d", int(pid.PID), int(i))
		}
	}
}

func TestPIDQueryAfter(t *testing.T) {
	test := "100+"
	pids := PIDQuery(test)
	if len(pids) <= 1 {
		t.Errorf("invalid plus query")
		t.FailNow()
	}
	if pids[0].PID != 100 {
		t.Errorf("invalid plus query starting point")
	}
}

func TestPIDQueryLeading(t *testing.T) {
	test := "+100"
	pids := PIDQuery(test)
	if len(pids) <= 1 {
		t.Errorf("invalid minus query")
		t.FailNow()
	}
	if pids[0].PID != 0 {
		t.Errorf("invalid minus query starting point")
	}
	if pids[100].PID != 100 {
		t.Errorf("invalid minus query finish point")
	}
}
