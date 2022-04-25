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

	filter "github.com/kris-nova/xpid/pkg/filters"

	api "github.com/kris-nova/xpid/pkg/api/v1"
	"github.com/kris-nova/xpid/pkg/procx"
)

var _ procx.ProcessExplorerEncoder = &RawEncoder{}

type RawEncoder struct {
	filters []filter.ProcessFilter
	format  Formatter
}

type Formatter func(p *api.Process) string

var _ Formatter = DefaultFormatter

func DefaultFormatter(p *api.Process) string {
	return fmt.Sprintf("[%d] %s (%s)\n", p.PID, p.Name, p.CommandLine)
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
