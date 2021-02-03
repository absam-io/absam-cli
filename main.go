package main

import (
	"os"

	"github.com/absam-io/absam-cli/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
