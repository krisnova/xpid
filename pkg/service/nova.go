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

package service

import (
	"github.com/kris-nova/go-nova/pkg/api"
	"github.com/sirupsen/logrus"
	"time"
)

// Compile check *Nova implements Runner interface
var _ api.Runner = &Nova{}

type Nova struct {
	// Fields
}

func NewNova() *Nova {
	return &Nova{}
}

var (
	runtimeNova bool = true
)

func (n *Nova) Run() error {
	client := api.Client{}
	server := api.Server{}
	logrus.Infof("Client: %x", client)
	logrus.Infof("Server: %x", server)
	for runtimeNova == true {
		time.Sleep(1 * time.Second)
		logrus.Infof("Sleeping...\n")
	}
	return nil
}
