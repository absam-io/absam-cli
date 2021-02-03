package cmd

import (
	"os"
	"strconv"

	"github.com/absam-io/absam-cli/api"
	"github.com/absam-io/absam-cli/tui"
	"github.com/absam-io/absam-cli/utils"
	"github.com/spf13/cobra"
)

var (
	serverCmd = &cobra.Command{
		Use:   "server [all] [id [status|start|shutdown|restart|stop|resume]]",
		Short: "Manage servers",
		Long:  "Manage servers",

		Run: func(cmd *cobra.Command, args []string) {
			err := manageServer(cmd, args)
			if err != nil {
				utils.Die(err)
			}
		},
	}
)

const (
	EMPTY         = 0
	SERVICE_ID    = 0
	SERVER_ACTION = 1
)

func validateArgs(cmd *cobra.Command, args []string) {
	if len(args) == EMPTY {
		cmd.Help()
		os.Exit(utils.FAIL)
	}

	if args[SERVICE_ID] == "all" && len(args) > 1 {
		cmd.Help()
		os.Exit(utils.FAIL)
	} else if args[SERVICE_ID] != "all" {
		if _, err := strconv.Atoi(args[SERVICE_ID]); err != nil {
			cmd.Help()
			os.Exit(utils.FAIL)
		}
	}
}

func parseArgs(cmd *cobra.Command, args []string) {
	validateArgs(cmd, args)

	switch args[SERVICE_ID] {
	case "all":
		servers := api.GetServers()
		tui.PrintAllServers(servers)
	default:
		if len(args) > 1 {
			if args[SERVER_ACTION] == "status" {
				status := api.GetServerStatus(args[0])
				tui.PrintServerStatus(status)
			} else if args[SERVER_ACTION] == "start" {
				result := api.ManageServer("start", args[SERVICE_ID])
				tui.PrintResult(result)
			} else if args[SERVER_ACTION] == "shutdown" {
				result := api.ManageServer("shutdown", args[SERVICE_ID])
				tui.PrintResult(result)
			} else if args[SERVER_ACTION] == "restart" {
				result := api.ManageServer("reset", args[SERVICE_ID])
				tui.PrintResult(result)
			} else if args[SERVER_ACTION] == "stop" {
				result := api.ManageServer("stop", args[SERVICE_ID])
				tui.PrintResult(result)
			} else if args[SERVER_ACTION] == "resume" {
				result := api.ManageServer("resume", args[SERVICE_ID])
				tui.PrintResult(result)
			}
		}
		info := api.GetServerInfo(args[SERVICE_ID])
		tui.PrintSingleServer(info)
	}
}

func manageServer(cmd *cobra.Command, args []string) error {
	parseArgs(cmd, args)
	return nil
}
