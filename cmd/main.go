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

	// Concurrent will toggle the concurrent pool buffer for the process explorer.
	// If enabled will run concurrently, or "fast".
	Concurrent bool

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

	// Available namespaces: mnt, pid, net, ipc, cgroup
	NamespaceInPid1  cli.StringSlice
	NamespaceOutPid1 cli.StringSlice
	NamespaceInUser  cli.StringSlice
	NamespaceOutUser cli.StringSlice

	// TableFormatters
	ShowTableNamespaces bool

	ProcListing bool
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

	// EXAMPLE: Override a template
	cli.AppHelpTemplate = fmt.Sprintf(`%s
{{.Usage}}

{{.UsageText}}
Options
   {{range .VisibleFlags}}{{.}}
   {{end}}
`, xpid.Banner())

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

Investigate all pids
	xpid

Investigate pid 1
	xpid 1

Find all container processes on a system
	xpid -c

Find all processes in the same namespace(s) as pid 1
	xpid --ns-in [mnt, net, pid, ipc, cgroup]

Find all processes not in the same namespace(s) as current user
	xpid --ns-out-user [mnt, net, pid, ipc, cgroup]

Find all processes running with eBPF programs as JSON
	xpid --ebpf -o json <pid-query>

Find all processes running with eBPF programs, in a container, in /proc
	xpid -b -c -p

Find all processes between specific values (Query syntax)
	xpid <flags> +100      # Search pids up to 100
	xpid <flags> 100-2000  # Search pids between 100-2000 
	xpid <flags> 65000+    # Search pids 65000 or above

Find all hidden processes on a system (slow)
	xpid -x <pid-query>
`,
		Commands: []*cli.Command{},
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "verbose",
				Aliases:     []string{"v"},
				Destination: &cfg.Verbose,
				Usage:       "Toggle verbose mode",
				DefaultText: "",
			},
			&cli.StringFlag{
				Name:        "out",
				Aliases:     []string{"o"},
				Destination: &cfg.Output,
				Usage:       "Set output encoder (json, table, raw)",
			},
			&cli.BoolFlag{
				Name:        "proc",
				Aliases:     []string{"p", "pl"},
				Destination: &cfg.ProcListing,
				Usage:       "List only pids in /proc (Fast).",
				DefaultText: "",
			},
			&cli.StringSliceFlag{
				Name:        "ns-in",
				Aliases:     []string{"N"},
				Destination: &cfg.NamespaceInPid1,
				Usage:       "Show pids in namespace(s) as pid 1.",
				DefaultText: "",
			},
			&cli.StringSliceFlag{
				Name:        "ns-out",
				Aliases:     []string{"O"},
				Destination: &cfg.NamespaceOutPid1,
				Usage:       "Reject pids in namespace(s) as pid 1.",
				DefaultText: "",
			},
			&cli.StringSliceFlag{
				Name:        "ns-in-user",
				Destination: &cfg.NamespaceInPid1,
				Usage:       "Show pids in namespace(s) as user.",
				DefaultText: "",
			},
			&cli.StringSliceFlag{
				Name:        "ns-out-user",
				Destination: &cfg.NamespaceOutPid1,
				Usage:       "Reject pids in namespace(s) as user.",
				DefaultText: "",
			},
			&cli.BoolFlag{
				Name:        "concurrent",
				Aliases:     []string{"C", "fast"},
				Destination: &cfg.Concurrent,
				Value:       true,
				Usage:       "Run concurrently (heavy CPU).",
				// Leave default text alone
			},
			&cli.BoolFlag{
				Name:        "ebpf",
				Aliases:     []string{"bpf", "b"}, // The "B" stands for Berkeley, Bitches.
				Destination: &cfg.Ebpf,
				Value:       false,
				Usage:       "Show pids with BPF programs attached.",
				DefaultText: "",
			},
			&cli.BoolFlag{
				Name:        "hidden",
				Aliases:     []string{"x"},
				Destination: &cfg.Hidden,
				Value:       false,
				Usage:       "Useful for finding rootkits.",
				DefaultText: "",
			},
			&cli.BoolFlag{
				Name:        "threads",
				Aliases:     []string{"t", "thread"},
				Destination: &cfg.Threads,
				Value:       false,
				Usage:       "Show threads and parent pids",
				DefaultText: "",
			},
			&cli.BoolFlag{
				Name:        "container",
				Aliases:     []string{"c"},
				Destination: &cfg.Container,
				Value:       false,
				Usage:       "Show pids with unique cgroup namespace.",
				DefaultText: "",
			},
			&cli.BoolFlag{
				Name:        "n",
				Aliases:     []string{"namespaces"},
				Destination: &cfg.ShowTableNamespaces,
				Value:       false,
				Usage:       "Display system namespaces.",
				DefaultText: "",
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

			// Validation
			if cfg.Hidden && cfg.ProcListing {
				return fmt.Errorf("unable to find hidden pids by only looking in procfs")
			}
			if cfg.Ebpf {
				table.TableFmtBPF = true
			}

			// Initialize the explorer based on flags
			if cfg.ProcListing {
				pids = procx.ProcListingQuery(query)
			} else {
				pids = procx.PIDQuery(query)
			}

			if pids == nil {
				return fmt.Errorf("invalid pid query: %s", query)
			}
			logrus.Infof("Query : %s\n", query)
			x := procx.NewProcessExplorer(pids)

			// Fast
			x.SetFast(cfg.Concurrent)

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
				rawcolor := raw.NewRawEncoder()
				rawcolor.SetFormat(raw.ColorFormatter)
				encoder = rawcolor
			default:
				table.TableFmtNS = cfg.ShowTableNamespaces
				encoder = table.NewTableEncoder()
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
				return fmt.Errorf("unable to encode current upid: %v", err)
			}
			fmt.Print(string(bytes))

			// Finally our pid
			upid := v1.ProcessPID(int64(os.Getpid()))
			v1.NewProcModule().Execute(upid)
			v1.NewNamespaceModule().Execute(upid)

			upid.DrawLineAfter = true
			bytes, err = encoder.Encode(upid)
			if err != nil {
				return fmt.Errorf("unable to encode current upid: %v", err)
			}
			fmt.Print(string(bytes))

			// Filters
			filter.PidOne = pid1
			filter.CurrentUser = currentUser()

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

			// Namespace in Pid 1
			nsIns := cfg.NamespaceInPid1.Value()
			if len(nsIns) > 0 {
				for _, nsIn := range nsIns {
					switch nsIn {
					case v1.NamespaceMount:
						filter.NamespaceFilterSet_Mount = pid1.NamespaceModule.Mount
						encoder.AddFilter(filter.RetainNamespaceIn_Mount)
						break
					case v1.NamespaceIPC:
						filter.NamespaceFilterSet_IPC = pid1.NamespaceModule.IPC
						encoder.AddFilter(filter.RetainNamespaceIn_IPC)
						break
					case v1.NamespaceNet:
						filter.NamespaceFilterSet_Net = pid1.NamespaceModule.Net
						encoder.AddFilter(filter.RetainNamespaceIn_Net)
						break
					case v1.NamespaceCgroup:
						filter.NamespaceFilterSet_Cgroup = pid1.NamespaceModule.Cgroup
						encoder.AddFilter(filter.RetainNamespaceIn_Cgroup)
						break
					case v1.NamespacePid:
						filter.NamespaceFilterSet_PID = pid1.NamespaceModule.PID
						encoder.AddFilter(filter.RetainNamespaceIn_PID)
						break
					default:
						logrus.Errorf("invalid namespace-in: %s", nsIn)
						os.Exit(ExitCode_InvalidNamespace)
					}
				}
			}

			// Namespace in User
			nsInsU := cfg.NamespaceInUser.Value()
			if len(nsInsU) > 0 {
				for _, nsIn := range nsInsU {
					switch nsIn {
					case v1.NamespaceMount:
						filter.NamespaceFilterSet_Mount = upid.NamespaceModule.Mount
						encoder.AddFilter(filter.RetainNamespaceIn_Mount)
						break
					case v1.NamespaceIPC:
						filter.NamespaceFilterSet_IPC = upid.NamespaceModule.IPC
						encoder.AddFilter(filter.RetainNamespaceIn_IPC)
						break
					case v1.NamespaceNet:
						filter.NamespaceFilterSet_Net = upid.NamespaceModule.Net
						encoder.AddFilter(filter.RetainNamespaceIn_Net)
						break
					case v1.NamespaceCgroup:
						filter.NamespaceFilterSet_Cgroup = upid.NamespaceModule.Cgroup
						encoder.AddFilter(filter.RetainNamespaceIn_Cgroup)
						break
					case v1.NamespacePid:
						filter.NamespaceFilterSet_PID = upid.NamespaceModule.PID
						encoder.AddFilter(filter.RetainNamespaceIn_PID)
						break
					default:
						logrus.Errorf("invalid namespace-in-user: %s", nsIn)
						os.Exit(ExitCode_InvalidNamespace)
					}
				}
			}

			// Namespace out Pid 1
			nsOuts := cfg.NamespaceOutPid1.Value()
			if len(nsOuts) > 0 {
				for _, nsOut := range nsOuts {
					switch nsOut {

					case v1.NamespaceMount:
						filter.NamespaceFilterSet_Mount = pid1.NamespaceModule.Mount
						encoder.AddFilter(filter.RetainNamespaceOut_Mount)
						break
					case v1.NamespaceIPC:
						filter.NamespaceFilterSet_IPC = pid1.NamespaceModule.IPC
						encoder.AddFilter(filter.RetainNamespaceOut_IPC)
						break
					case v1.NamespaceNet:
						filter.NamespaceFilterSet_Net = pid1.NamespaceModule.Net
						encoder.AddFilter(filter.RetainNamespaceOut_Net)
						break
					case v1.NamespaceCgroup:
						filter.NamespaceFilterSet_Cgroup = pid1.NamespaceModule.Cgroup
						encoder.AddFilter(filter.RetainNamespaceOut_Cgroup)
						break
					case v1.NamespacePid:
						filter.NamespaceFilterSet_PID = pid1.NamespaceModule.PID
						encoder.AddFilter(filter.RetainNamespaceOut_PID)
						break
					default:
						logrus.Errorf("invalid namespace-out: %s", nsOut)
						os.Exit(ExitCode_InvalidNamespace)
					}
				}
			}

			// Namespace out User
			nsOutsU := cfg.NamespaceOutUser.Value()
			if len(nsOutsU) > 0 {
				for _, nsOut := range nsOutsU {
					switch nsOut {

					case v1.NamespaceMount:
						filter.NamespaceFilterSet_Mount = upid.NamespaceModule.Mount
						encoder.AddFilter(filter.RetainNamespaceOut_Mount)
						break
					case v1.NamespaceIPC:
						filter.NamespaceFilterSet_IPC = upid.NamespaceModule.IPC
						encoder.AddFilter(filter.RetainNamespaceOut_IPC)
						break
					case v1.NamespaceNet:
						filter.NamespaceFilterSet_Net = upid.NamespaceModule.Net
						encoder.AddFilter(filter.RetainNamespaceOut_Net)
						break
					case v1.NamespaceCgroup:
						filter.NamespaceFilterSet_Cgroup = upid.NamespaceModule.Cgroup
						encoder.AddFilter(filter.RetainNamespaceOut_Cgroup)
						break
					case v1.NamespacePid:
						filter.NamespaceFilterSet_PID = upid.NamespaceModule.PID
						encoder.AddFilter(filter.RetainNamespaceOut_PID)
						break
					default:
						logrus.Errorf("invalid namespace-out-user: %s", nsOut)
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
