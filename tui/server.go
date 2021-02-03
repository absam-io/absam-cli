package tui

import (
	"fmt"
	"os"
	"strconv"

	"github.com/absam-io/absam-cli/models"
	"github.com/absam-io/absam-cli/utils"
	"github.com/olekukonko/tablewriter"
)

func FormatServerInfo(id, disk int, name, status, dc, ip, price, plan, billing string) []string {
	idStr := strconv.Itoa(id)
	diskStr := strconv.Itoa(disk)

	return []string{idStr, name, status, dc, ip, plan, diskStr, price, billing}
}

func PrintSingleServer(server models.SingleServer) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{
		"Id", "Name", "Status", "DC Location", "IP", "Plan", "Disk",
		"Price", "Billing Cycle",
	})

	serverInfo := FormatServerInfo(
		server.Server.Id,
		server.Server.Disk,
		server.Server.Name,
		server.Server.ProductStatus,
		server.Server.Datacenter,
		server.Server.Ip,
		server.Server.Price,
		server.Server.Plan,
		server.Server.BillingCycle)

	table.Append(serverInfo)
	table.Render()
	os.Exit(utils.SUCCESS)
}

func PrintAllServers(servers models.Servers) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{
		"Id", "Name", "Status", "DC Location", "IP", "Plan", "Disk",
		"Price", "Billing Cycle",
	})

	for _, server := range servers.Servers {
		serverInfo := FormatServerInfo(
			server.Id,
			server.Disk,
			server.Name,
			server.ProductStatus,
			server.Datacenter,
			server.Ip,
			server.Price,
			server.Plan,
			server.BillingCycle)

		table.Append(serverInfo)
	}

	table.Render()
	os.Exit(utils.SUCCESS)
}

func PrintServerStatus(server models.Status) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Status", "CPU", "Memory", "Disk"})

	data := []string{
		server.Server.Status,
		server.Server.Consuming.CPU,
		server.Server.Consuming.Memory,
		server.Server.Consuming.Disk}

	table.Append(data)
	table.Render()
	os.Exit(utils.SUCCESS)
}

func PrintResult(result models.Result) {
	fmt.Println(result.Status)
	os.Exit(utils.SUCCESS)
}
