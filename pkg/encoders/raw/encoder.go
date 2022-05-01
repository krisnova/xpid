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

package raw

import (
	"fmt"

	encoder "github.com/kris-nova/xpid/pkg/encoders"

	filter "github.com/kris-nova/xpid/pkg/filters"

	api "github.com/kris-nova/xpid/pkg/api/v1"
)

var _ encoder.ProcessExplorerEncoder = &RawEncoder{}

type RawEncoder struct {
	filters []filter.ProcessFilter
	format  Formatter
}

func (r *RawEncoder) EncodeAll(p *api.Process) ([]byte, error) {
	return r.Encode(p)
}

func (r *RawEncoder) EncodeUser(u *api.User) ([]byte, error) {
	str := fmt.Sprintf("\n[%s] %d [%s] %d\n\n",
		u.Name,
		u.ID,
		u.Group.Name,
		u.Group.ID)
	return []byte(str), nil
}

type Formatter func(p *api.Process) string

var _ Formatter = DefaultFormatter

func DefaultFormatter(p *api.Process) string {
	return fmt.Sprintf("[%d] %s(%d):%s(%d) %s (%s)\n",
		p.PID,
		p.User.Name,
		p.User.ID,
		p.User.Group.Name,
		p.User.Group.ID,
		p.Name,
		p.CommandLine)
}

func (r *RawEncoder) SetFormat(f Formatter) {
	r.format = f
}

func (r *RawEncoder) Encode(p *api.Process) ([]byte, error) {
	for _, f := range r.filters {
		x := f(p)
		if !x {
			return []byte(""), fmt.Errorf("filtered")
		}
	}
	return []byte(r.format(p)), nil
}

func (r *RawEncoder) AddFilter(f filter.ProcessFilter) {
	r.filters = append(r.filters, f)
}

func NewRawEncoder() *RawEncoder {
	return &RawEncoder{
		format: DefaultFormatter,
	}
}
