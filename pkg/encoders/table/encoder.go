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

package table

import (
	"fmt"

	encoder "github.com/kris-nova/xpid/pkg/encoders"

	filter "github.com/kris-nova/xpid/pkg/filters"

	api "github.com/kris-nova/xpid/pkg/api/v1"
)

var _ encoder.ProcessExplorerEncoder = &TableEncoder{}

type TableEncoder struct {
	filters []filter.ProcessFilter
}

func (j *TableEncoder) EncodeAll(p *api.Process) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (j *TableEncoder) EncodeUser(u *api.User) ([]byte, error) {
	var str string

	// Header
	str += fmt.Sprintf("\n")

	str += fmt.Sprintf("%-*s", len(u.Name)+3, "USER")
	str += fmt.Sprintf("%-*s", 5, "UID")
	str += fmt.Sprintf("%-*s", len(u.Group.Name)+3, "GROUP")
	str += fmt.Sprintf("%-*s", 5, "GID")
	str += fmt.Sprintf("\n")

	// First line
	str += fmt.Sprintf("%-*s", len(u.Name)+3, u.Name)
	str += fmt.Sprintf("%-*d", 5, u.ID)
	str += fmt.Sprintf("%-*s", len(u.Group.Name)+3, u.Group.Name)
	str += fmt.Sprintf("%-*d", 5, u.Group.ID)
	str += fmt.Sprintf("\n")
	str += fmt.Sprintf("\n")

	return []byte(str), nil
}

func (j *TableEncoder) Encode(p *api.Process) ([]byte, error) {
	for _, f := range j.filters {
		if !f(p) {
			return []byte(""), fmt.Errorf(filter.Filtered)
		}
	}

	var str string

	if p.ShowHeader {
		// Header
		str += fmt.Sprintf("%-6s", "PID")
		str += fmt.Sprintf("%-9s", "USER")
		str += fmt.Sprintf("%-9s", "GROUP")
		str += fmt.Sprintf("%-9s", "CMD")
		str += fmt.Sprintf("%-12s", "NSPID")     // Compute
		str += fmt.Sprintf("%-12s", "NS-CGROUP") // Compute
		str += fmt.Sprintf("%-12s", "NS-NET")    // Network
		str += fmt.Sprintf("%-12s", "NS-MNT")    // Storage
		str += fmt.Sprintf("\n")
	}

	// First line
	str += fmt.Sprintf("%-6d", p.PID)
	str += fmt.Sprintf("%-9s", p.User.Name)
	str += fmt.Sprintf("%-9s", p.User.Group.Name)
	str += fmt.Sprintf("%-9s", p.ProcModule.Comm)
	str += fmt.Sprintf("%-12s", p.NamespaceModule.PID)
	str += fmt.Sprintf("%-12s", p.NamespaceModule.Cgroup)
	str += fmt.Sprintf("%-12s", p.NamespaceModule.Net)
	str += fmt.Sprintf("%-12s", p.NamespaceModule.Mount)
	str += fmt.Sprintf("\n")

	return []byte(str), nil

}

func (j *TableEncoder) AddFilter(f filter.ProcessFilter) {
	j.filters = append(j.filters, f)
}

func NewTableEncoder() *TableEncoder {
	return &TableEncoder{}
}