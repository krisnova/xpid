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
	"encoding/json"
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

	return []byte(str), nil
}

func (j *TableEncoder) Encode(p *api.Process) ([]byte, error) {
	for _, f := range j.filters {
		if !f(p) {
			return []byte(""), fmt.Errorf(filter.Filtered)
		}
	}
	return json.Marshal(p)
}

func (j *TableEncoder) AddFilter(f filter.ProcessFilter) {
	j.filters = append(j.filters, f)
}

func NewTableEncoder() *TableEncoder {
	return &TableEncoder{}
}
