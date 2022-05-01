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
