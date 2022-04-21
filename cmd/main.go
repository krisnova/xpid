/******************************************************************************
* MIT License
* Copyright (c) 2022 Kris Nóva <kris@nivenly.com>
*
* ┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓
* ┃   ███╗   ██╗ ██████╗ ██╗   ██╗ █████╗  ┃
* ┃   ████╗  ██║██╔═████╗██║   ██║██╔══██╗ ┃
* ┃   ██╔██╗ ██║██║██╔██║██║   ██║███████║ ┃
* ┃   ██║╚██╗██║████╔╝██║╚██╗ ██╔╝██╔══██║ ┃
* ┃   ██║ ╚████║╚██████╔╝ ╚████╔╝ ██║  ██║ ┃
* ┃   ╚═╝  ╚═══╝ ╚═════╝   ╚═══╝  ╚═╝  ╚═╝ ┃
* ┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛
*
*****************************************************************************/

package main

import (
	"fmt"
	"github.com/kris-nova/go-nova"
	x "github.com/kris-nova/go-nova/internal/nova"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"os"
	"time"
)

var cfg = &AppOptions{}

type AppOptions struct {
	verbose bool
}

func main() {
	/* Change version to -V */
	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"V"},
		Usage:   "The version of the program.",
	}
	app := &cli.App{
		Name:     nova.Name,
		Version:  nova.Version,
		Compiled: time.Now(),
		Authors: []*cli.Author{
			&cli.Author{
				Name:  nova.AuthorName,
				Email: nova.AuthorEmail,
			},
		},
		Copyright: nova.Copyright,
		HelpName:  nova.Copyright,
		Usage:     "A go program.",
		UsageText: `nova <options> <flags> 
A longer sentence, about how exactly to use this program`,
		Commands: []*cli.Command{
			&cli.Command{},
		},
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "verbose",
				Aliases:     []string{"v"},
				Destination: &cfg.verbose,
			},
		},
		EnableBashCompletion: true,
		HideHelp:             false,
		HideVersion:          false,
		Before: func(c *cli.Context) error {
			Preloader()
			fmt.Fprintf(c.App.Writer, nova.Banner())
			return nil
		},
		After: func(c *cli.Context) error {
			// Destruct
			return nil
		},
		Action: func(c *cli.Context) error {

			//
			novaObject := x.NewNova()
			return novaObject.Run()
			//

		},
	}
	app.Run(os.Args)
}

// Preloader will run for ALL commands, and is used
// to initalize the runtime environments of the program.
func Preloader() {
	/* Flag parsing */
	if cfg.verbose {
		logrus.SetLevel(logrus.InfoLevel)
	} else {
		logrus.SetLevel(logrus.WarnLevel)
	}
}
