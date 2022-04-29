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

package Raw

import (
	"fmt"

	"github.com/fatih/color"

	api "github.com/kris-nova/xpid/pkg/api/v1"
)

var _ Formatter = ColorFormatter

func ColorFormatter(p *api.Process) string {
	pidColor := color.New(color.FgBlue, color.Bold)
	cliColor := color.New(color.FgYellow, color.Italic)
	ugColor := color.New(color.FgCyan, color.Italic)

	// Default color formatter
	return fmt.Sprintf("[%s] %s(%s):%s(%s) %s (%s)\n",
		pidColor.Sprintf("%d", p.PID),
		ugColor.Sprintf("%s", p.User.Name),
		ugColor.Sprintf("%d", p.User.ID),
		ugColor.Sprintf("%s", p.Group.Name),
		ugColor.Sprintf("%d", p.Group.ID),
		color.GreenString("%s", p.Name),
		cliColor.Sprintf("%s", p.CommandLine))
}
