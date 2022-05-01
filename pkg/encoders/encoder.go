package encoder

import (
	api "github.com/kris-nova/xpid/pkg/api/v1"
	filter "github.com/kris-nova/xpid/pkg/filters"
)

type ProcessExplorerEncoder interface {
	AddFilter(f filter.ProcessFilter)
	Encode(p *api.Process) ([]byte, error)
	EncodeAll(p *api.Process) ([]byte, error)
	EncodeUser(u *api.User) ([]byte, error)
}
