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

	"github.com/kris-nova/xpid/pkg/encoders/table"

	"github.com/kris-nova/xpid/pkg/encoders/raw"

	encoder "github.com/kris-nova/xpid/pkg/encoders"

	filter "github.com/kris-nova/xpid/pkg/filters"

	"github.com/kris-nova/xpid/pkg/encoders/json"

	v1 "github.com/kris-nova/xpid/pkg/api/v1"

	"github.com/kris-nova/xpid/pkg/procx"

	"github.com/kris-nova/xpid"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

var cfg = &AppOptions{}

type AppOptions struct {

	// Verbose will toggle verbosity (breaks encoding)
	Verbose bool

	// Fast will toggle the concurrent pool buffer for the process explorer.
	// If enabled will run concurrently, or "fast".
	Fast bool

	// Encoders
	//
	// Support: raw, json, color
	Output string

	// Hidden will only show "hidden" pids
	Hidden bool

	// Threads will toggle showing threads in a result
	Threads bool

	// Ebpf works like containers. Will attempt to show Ebpf programs.
	Ebpf bool

	// Containers
	Container bool

	// Namespace flags can be used to filter on pids that are in/out of the
	// current namespace context.
	//
	// For example if the user was running in a given mount namespace and
	// wanted to list all pids in the same mount namespace they would use
	// 	xpid --ns-in mount
	//
	// Available namespaces: mnt, pid, net, ipc, cgroup
	NamespaceIn  cli.StringSlice
	NamespaceOut cli.StringSlice
}

const (
	ExitCode_PermissionDenied int = 99
	ExitCode_InvalidNamespace int = 80
)

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
		Usage:     "Linux Process Discovery. Like nmap, but for pids.",
		UsageText: `xpid [flags] -o [output] <query>

Investigate pid 123 and write the report to out.txt
	xpid 123 > out.txt

Find all container processes on a system 
	# Looks for /proc/[pid]/ns/cgroup != /proc/1/ns/cgroup 
	xpid -c <query>

Find all processes running with eBPF programs at runtime.
	# Looks for /proc/[pid]/fdinfo and correlates to /sys/fs/bpf
	xpid --ebpf <query>

Find all processes between specific values
	xpid <flags> +100      # Search pids up to 100
	xpid <flags> 100-2000  # Search pids between 100-2000 
	xpid <flags> 65000+    # Search pids 65000 or above

Find all "hidden" processes on a system
	# Looks for chdir, opendir, and dent in /proc
	xpid -x <query>

Find all possible pids on a system, and investigate each one (slow).
	xpid > out.txt 

Investigate all pids from 0 to 1000 and write the report to out.json
	xpid -o json 0-1000 > out.json

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
				Name:        "output",
				Aliases:     []string{"o", "out"},
				Destination: &cfg.Output,
			},
			&cli.StringSliceFlag{
				Name:        "ns-in",
				Aliases:     []string{"namespace-in", "N"},
				Destination: &cfg.NamespaceIn,
				Usage:       "Return only pids in [mnt, net, pid, ipc, cgroup] namespace that matches the user.",
			},
			&cli.StringSliceFlag{
				Name:        "ns-out",
				Aliases:     []string{"namespace-out", "O"},
				Destination: &cfg.NamespaceOut,
				Usage:       "Reject pids from [mnt, net, pid, ipc, cgroup] namespace that matches the user.",
			},
			&cli.BoolFlag{
				Name:        "fast",
				Aliases:     []string{"f"},
				Destination: &cfg.Fast,
				Value:       true,
			},
			&cli.BoolFlag{
				Name:        "ebpf",
				Aliases:     []string{"bpf", "b"},
				Destination: &cfg.Ebpf,
				Value:       false,
			},
			&cli.BoolFlag{
				Name:        "hidden",
				Aliases:     []string{"x"},
				Destination: &cfg.Hidden,
				Value:       false,
			},
			&cli.BoolFlag{
				Name:        "threads",
				Aliases:     []string{"t", "thread"},
				Destination: &cfg.Threads,
				Value:       false,
			},
			&cli.BoolFlag{
				Name:        "container",
				Aliases:     []string{"c", "containers"},
				Destination: &cfg.Container,
				Value:       false,
			},
		},
		EnableBashCompletion: false,
		HideHelp:             false,
		HideVersion:          true,
		Before: func(c *cli.Context) error {
			Preloader()
			return nil
		},
		After: func(c *cli.Context) error {
			// Destruct
			return nil
		},
		Action: func(c *cli.Context) error {
			var pids []*v1.Process

			query := c.Args().Get(0)
			if query == "" {
				max := procx.MaxPid()
				if max == -1 {
					return fmt.Errorf("unable to read from /proc")
				}
				query = fmt.Sprintf("1-%d", max)
			}

			// Initialize the explorer based on flags
			pids = procx.PIDQuery(query)
			if pids == nil {
				return fmt.Errorf("invalid pid query: %s", query)
			}
			logrus.Infof("Query : %s\n", query)
			x := procx.NewProcessExplorer(pids)

			// Fast
			x.SetFast(cfg.Fast)

			// Encoder
			var encoder encoder.ProcessExplorerEncoder
			switch cfg.Output {
			case "json":
				encoder = json.NewJSONEncoder()
				break
			case "table":
				encoder = table.NewTableEncoder()
				break
			case "raw":
				encoder = raw.NewRawEncoder()
				break
			case "color":
			default:
				rawcolor := raw.NewRawEncoder()
				rawcolor.SetFormat(raw.ColorFormatter)
				encoder = rawcolor
			}

			// First the current user
			bytes, err := encoder.EncodeUser(currentUser())
			if err != nil {
				return fmt.Errorf("unable to find current user: %v", err)
			}
			fmt.Print(string(bytes))

			// Next pid one
			pid1 := v1.ProcessPID(1)
			pid1.ShowHeader = true
			v1.NewProcModule().Execute(pid1)
			v1.NewNamespaceModule().Execute(pid1)

			bytes, err = encoder.Encode(pid1)
			if err != nil {
				return fmt.Errorf("unable to find current pid 1: %v", err)
			}
			fmt.Print(string(bytes))

			// Filters
			filter.PidOne = pid1
			filter.CurrentUser = currentUser()

			// Todo namespace type switch

			encoder.AddFilter(filter.RetainOnlyNamed)
			if cfg.Hidden {
				encoder.AddFilter(filter.RetainOnlyHidden)
			}
			if !cfg.Threads {
				encoder.AddFilter(filter.RejectThreads)
			}

			// Always load "proc" module
			x.AddModule(v1.NewProcModule())

			// Always load "namespace" module
			x.AddModule(v1.NewNamespaceModule())

			// Check for EBPF
			if cfg.Ebpf {
				x.AddModule(v1.NewEBPFModule())
				encoder.AddFilter(filter.RetainOnlyEBPF)
			}

			// Check for container
			if cfg.Container {
				x.AddModule(v1.NewContainerModule())
				encoder.AddFilter(filter.RetainOnlyContainers)
			}

			// Namespace [Out]
			nsOuts := cfg.NamespaceOut.Value()
			if len(nsOuts) > 0 {
				for _, nsIn := range nsOuts {
					switch nsIn {
					case v1.NamespaceMount:
						break
					case v1.NamespaceIPC:
						break
					case v1.NamespaceNet:
						break
					case v1.NamespaceCgroup:
						break
					case v1.NamespacePid:
						break
					default:
						logrus.Errorf("invalid namespace-in: %s", nsIn)
						os.Exit(ExitCode_InvalidNamespace)
					}
				}
			}

			// Namespace [In]
			nsIns := cfg.NamespaceIn.Value()
			if len(nsIns) > 0 {
				for _, nsIn := range nsIns {
					switch nsIn {
					case v1.NamespaceMount:
						break
					case v1.NamespaceIPC:
						break
					case v1.NamespaceNet:
						break
					case v1.NamespaceCgroup:
						break
					case v1.NamespacePid:
						break
					default:
						logrus.Errorf("invalid namespace-in: %s", nsIn)
						os.Exit(ExitCode_InvalidNamespace)
					}
				}
			}

			// Execute
			x.SetEncoder(encoder)
			x.SetWriter(os.Stdout)
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
		logrus.SetLevel(logrus.WarnLevel)
	}

	if cfg.Container {
		if !isuid(0) {
			logrus.Errorf("Permission denied. UID=0 Required for [containers].")
			os.Exit(ExitCode_PermissionDenied)
		}
	}

	if cfg.Ebpf {
		if !isuid(0) {
			logrus.Errorf("Permission denied. UID=0 Required for [ebpf].")
			os.Exit(ExitCode_PermissionDenied)
		}
	}
}
