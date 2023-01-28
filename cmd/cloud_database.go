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
	cloudDatabaseCmd = &cobra.Command{
		Use:   "cloud-database [all] [id [status|start|shutdown|restart|stop]]",
		Short: "Manage servers",
		Long:  "Manage servers",

		Run: func(cmd *cobra.Command, args []string) {
			err := manageCloudDatabase(cmd, args)
			if err != nil {
				utils.Die(err)
			}
		},
	}
)

func validateCloudDatabaseArgs(cmd *cobra.Command, args []string) {
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

func parseCloudDatabaseArgs(cmd *cobra.Command, args []string) {
	validateArgs(cmd, args)

	switch args[SERVICE_ID] {
	case "all":
		servers := api.GetCloudDatabase()
		tui.PrintAllServers(servers)
	default:
		if len(args) > 1 {
			if args[SERVER_ACTION] == "status" {
				status := api.GetCloudDatabaseStatus(args[0])
				tui.PrintServerStatus(status)
			} else if args[SERVER_ACTION] == "start" {
				result := api.ManageCloudDatabase("start", args[SERVICE_ID])
				tui.PrintResult(result)
			} else if args[SERVER_ACTION] == "shutdown" {
				result := api.ManageCloudDatabase("shutdown", args[SERVICE_ID])
				tui.PrintResult(result)
			} else if args[SERVER_ACTION] == "restart" {
				result := api.ManageCloudDatabase("reset", args[SERVICE_ID])
				tui.PrintResult(result)
			} else if args[SERVER_ACTION] == "stop" {
				result := api.ManageCloudDatabase("stop", args[SERVICE_ID])
				tui.PrintResult(result)
			}
		}
		info := api.GetCloudDatabaseInfo(args[SERVICE_ID])
		tui.PrintSingleServer(info)
	}
}

func manageCloudDatabase(cmd *cobra.Command, args []string) error {
	parseCloudDatabaseArgs(cmd, args)
	return nil
}
