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

package main

import (
	"os/user"
	"strconv"

	"github.com/fatih/color"

	"golang.org/x/term"

	api "github.com/kris-nova/xpid/pkg/api/v1"
)

func isuid(check int) bool {
	u, _ := user.Current()
	if u == nil {
		return false
	}
	i, _ := strconv.Atoi(u.Uid)
	return check == i
}

func currentUser() *api.User {
	gouser, _ := user.Current()
	gogroup, _ := user.LookupGroupId(gouser.Gid)
	uid, _ := strconv.Atoi(gouser.Uid)
	gid, _ := strconv.Atoi(gouser.Gid)
	u := &api.User{
		Group: api.Group{
			ID:   gid,
			Name: gogroup.Name,
		},
		User: *gouser,
		ID:   uid,
		Name: gouser.Username,
	}
	return u
}

func drawLine() string {
	if cfg.Output == "json" {
		return ""
	}
	y, _, _ := term.GetSize(0)
	if y == 0 {
		return ""
	}
	var str string
	for i := 0; i < y; i++ {
		str += color.GreenString("=")
	}
	str += "\n"
	return str
}
