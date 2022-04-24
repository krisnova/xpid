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

package json

import (
	"encoding/json"
	"fmt"

	filter "github.com/kris-nova/xpid/pkg/filters"

	api "github.com/kris-nova/xpid/pkg/api/v1"
	"github.com/kris-nova/xpid/pkg/procx"
)

var _ procx.ProcessExplorerEncoder = &JSONEncoder{}

type JSONEncoder struct {
	filters []filter.ProcessFilter
}

func (j *JSONEncoder) Encode(p *api.Process) ([]byte, error) {
	for _, f := range j.filters {
		if !f(p) {
			return []byte(""), fmt.Errorf("filtered")
		}
	}
	return json.Marshal(p)
}

func (j *JSONEncoder) AddFilter(f filter.ProcessFilter) {
	j.filters = append(j.filters, f)
}

func NewJSONEncoder() *JSONEncoder {
	return &JSONEncoder{}
}
