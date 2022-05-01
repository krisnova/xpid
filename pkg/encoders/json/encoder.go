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

	encoder "github.com/kris-nova/xpid/pkg/encoders"

	filter "github.com/kris-nova/xpid/pkg/filters"

	api "github.com/kris-nova/xpid/pkg/api/v1"
)

var _ encoder.ProcessExplorerEncoder = &JSONEncoder{}

type JSONEncoder struct {
	filters []filter.ProcessFilter
}

func (j *JSONEncoder) EncodeAll(p *api.Process) ([]byte, error) {
	return j.Encode(p)
}

func (j *JSONEncoder) EncodeUser(u *api.User) ([]byte, error) {
	return json.Marshal(u)
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
