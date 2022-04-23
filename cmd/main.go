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

package main

import (
	"fmt"
	"os"
	"time"

	v1 "github.com/kris-nova/xpid/pkg/api/v1"

	"github.com/kris-nova/xpid/pkg/encoders/raw"

	modebpf "github.com/kris-nova/xpid/pkg/modules/ebpf"
	modproc "github.com/kris-nova/xpid/pkg/modules/proc"

	"github.com/kris-nova/xpid/pkg/procx"

	"github.com/kris-nova/xpid"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

var cfg = &AppOptions{}

type AppOptions struct {
	Verbose  bool
	PIDQuery string // 0, 1, 5-10

	// Encoders
	Output string

	// Modules
	All  bool
	EBPF bool
	Proc bool
}

func main() {
	/* Change version to -V */
	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"V"},
		Usage:   "The version of the program.",
	}
	app := &cli.App{
		Name:     xpid.Name,
		Version:  xpid.Version,
		Compiled: time.Now(),
		Authors: []*cli.Author{
			&cli.Author{
				Name:  xpid.AuthorName,
				Email: xpid.AuthorEmail,
			},
		},
		Copyright: xpid.Copyright,
		HelpName:  xpid.Copyright,
		Usage:     "Linux Process Discovery.",
		UsageText: `xpid [flags] -o [output]

Investigate pid 123 and write the report to out.txt
	xpid -p 123 > out.txt

Find all possible pids, and investigate each one (slow). The --all flag is default.
	xpid > out.txt 
	xpid --all > out.txt

Investigate all pids from 0 to 1000 and write the report to out.json
	xpid -p 0-1000 -o json > out.json

Find all eBPF pids at runtime (fast).
	xpid --ebpf

Find all proc pids at runtime (fast).
	xpid --proc

Investigate pid 123 using the "--proc" module only.
	xpid --proc -p 123 > out.txt

`,
		Commands: []*cli.Command{
			&cli.Command{},
		},
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "verbose",
				Aliases:     []string{"v"},
				Destination: &cfg.Verbose,
			},
			&cli.StringFlag{
				Name:        "pid",
				Aliases:     []string{"pids", "p"},
				Destination: &cfg.PIDQuery,
			},
			&cli.StringFlag{
				Name:        "output",
				Aliases:     []string{"o", "out"},
				Destination: &cfg.Output,
			},
			// Modules should have capital single letter flags!
			&cli.BoolFlag{
				Name:        "all",
				Aliases:     []string{"A"},
				Destination: &cfg.All,
				Value:       true,
			},
			&cli.BoolFlag{
				Name:        "ebpf",
				Aliases:     []string{"E"},
				Destination: &cfg.EBPF,
			},
			&cli.BoolFlag{
				Name:        "proc",
				Aliases:     []string{"P"},
				Destination: &cfg.Proc,
			},
		},
		EnableBashCompletion: false,
		HideHelp:             false,
		HideVersion:          true,
		Before: func(c *cli.Context) error {
			Preloader()
			fmt.Fprintf(c.App.Writer, xpid.Banner())
			return nil
		},
		After: func(c *cli.Context) error {
			// Destruct
			return nil
		},
		Action: func(c *cli.Context) error {
			var pids []*v1.Process
			if cfg.PIDQuery != "" {
				pids = procx.PIDQuery(cfg.PIDQuery)
				if pids == nil {
					return fmt.Errorf("invalid pid query: %s", cfg.PIDQuery)
				}
			} else {
				max := procx.MaxPid()
				if max == -1 {
					return fmt.Errorf("unable to read from /proc")
				}
				query := fmt.Sprintf("1-%d", max)
				fmt.Println(query)
				pids = procx.PIDQuery(query)
				if pids == nil {
					return fmt.Errorf("invalid pid query: %s", cfg.PIDQuery)
				}
			}

			// Initialize the explorer based on flags
			x := procx.NewProcessExplorer(pids)

			// Encoder
			encoder := raw.NewRawEncoder()
			x.SetEncoder(encoder) // Default raw encoder for now

			// Modules
			if cfg.All {
				cfg.EBPF = true
				cfg.Proc = true
			}
			if cfg.Proc {
				pmod := modproc.NewProcModule()
				x.AddModule(pmod)
			}
			if cfg.EBPF {
				emod := modebpf.NewEBPFModule()
				x.AddModule(emod)
			}
			x.SetWriter(os.Stdout)

			// Execute
			return x.Execute()
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		logrus.Errorf("execution error: %v", err)
		os.Exit(1)
	}
	os.Exit(0)
}

// Preloader will run for ALL commands, and is used
// to initalize the runtime environments of the program.
func Preloader() {
	/* Flag parsing */
	if cfg.Verbose {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}
}
