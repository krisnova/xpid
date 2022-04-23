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

package xpid

import "fmt"

var (
	Name        string
	Version     string
	Copyright   string
	License     string
	AuthorName  string
	AuthorEmail string
)

func Banner() string {
	var banner string
	banner += fmt.Sprintf("\n")
	banner += fmt.Sprintf(" ┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓\n")
	banner += fmt.Sprintf(" ┃  ██╗  ██╗██████╗ ██╗██████╗  ┃\n")
	banner += fmt.Sprintf(" ┃  ╚██╗██╔╝██╔══██╗██║██╔══██╗ ┃\n")
	banner += fmt.Sprintf(" ┃   ╚███╔╝ ██████╔╝██║██║  ██║ ┃\n")
	banner += fmt.Sprintf(" ┃   ██╔██╗ ██╔═══╝ ██║██║  ██║ ┃\n")
	banner += fmt.Sprintf(" ┃  ██╔╝ ██╗██║     ██║██████╔╝ ┃\n")
	banner += fmt.Sprintf(" ┃  ╚═╝  ╚═╝╚═╝     ╚═╝╚═════╝  ┃\n")
	banner += fmt.Sprintf(" ┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛\n")
	banner += fmt.Sprintf("   %s <%s>\n", AuthorName, AuthorEmail)
	return banner
}
