package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/absam-io/absam-cli/config"
	"github.com/absam-io/absam-cli/utils"
	"github.com/spf13/cobra"
)

var (
	access_token        = ""
	secret_access_token = ""

	initCmd = &cobra.Command{
		Use:   "init",
		Short: "Setup absam-cli",
		Long: `The command (absam-cli init) will create a config file in your home directory
containing the access_token and secret_access_token generated in your Absam.`,

		Run: func(cmd *cobra.Command, args []string) {
			err := createConfig(cmd, args)
			if err != nil {
				utils.Die(err)
			}
			fmt.Println("absam-cli is ready")
		},
	}
)

func init() {
	initCmd.Flags().StringVar(&access_token, "access-token", "", "access_token")
	initCmd.Flags().StringVar(&secret_access_token, "secret-access-token", "", "secret_access_token")
	initCmd.MarkFlagRequired("access-token")
	initCmd.MarkFlagRequired("secret-access-token")
}

func getFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	filepath := fmt.Sprintf("%s/%s", home, "absamconfig.json")
	return filepath, nil
}

func writeConfig(config []byte) error {
	filepath, err := getFilePath()
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filepath, config, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func createConfig(cmd *cobra.Command, args []string) error {
	if len(args) > 0 {
		cmd.Help()
		os.Exit(utils.FAIL)
	}

	conf := &config.Config{
		Token:  access_token,
		Secret: secret_access_token,
	}

	byteConfig, err := json.Marshal(conf)
	if err != nil {
		return errors.New("Error while setting up the config file.")
	}

	err = writeConfig(byteConfig)
	if err != nil {
		return errors.New("Error while writing the config file.")
	}

	return nil
}
