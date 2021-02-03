package tui

import (
	"fmt"
	"os"
	"strconv"

	"github.com/absam-io/absam-cli/models"
	"github.com/absam-io/absam-cli/utils"
	"github.com/olekukonko/tablewriter"
)

func PrintFirewallMessage(message string) {
	fmt.Println(message)
	os.Exit(utils.SUCCESS)
}

func PrintFirewallStatus(status models.FirewallStatus) {
	fmt.Printf("Status: ")
	if status.Enabled {
		fmt.Println("On")
	} else {
		fmt.Println("Off")
	}
	os.Exit(utils.SUCCESS)
}

func PrintRules(rules []models.Rule) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{
		"Position", "Protocol", "Port", "IP or Range", "Comment", "Expires (hour)",
	})

	for _, rule := range rules {
		position := strconv.Itoa(rule.Pos)
		data := []string{position, rule.Proto, rule.Port, rule.IP, rule.Comment, rule.Expires}
		table.Append(data)
	}

	table.Render()
}

func PrintAllRules(rules models.FirewallRules) {
	fmt.Println("Inbound Rules")
	PrintRules(rules.Rules.In)
	fmt.Println()

	fmt.Println("Outbound Rules")
	PrintRules(rules.Rules.Out)

	os.Exit(utils.SUCCESS)
}
