/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/ahmedennaifer/blov/internal/config"
	"github.com/ahmedennaifer/blov/internal/config/gcp"
	"github.com/spf13/cobra"
)

var gcpCmd = &cobra.Command{
	Use:   "gcp",
	Short: "Manages the config of GCP: project-id and region",
	Args:  cobra.RangeArgs(1, 3),
	Run: func(cmd *cobra.Command, args []string) {
		gcpCmd := args[0]
		if gcpCmd != "set" && gcpCmd != "get" && gcpCmd != "list" {
			fmt.Printf("\033[31merror: Subcommand %v is not supported for command `gcp`\033[0m\nUsage: blov config gcp set region-id your-region-id\n", gcpCmd)
			return
		}
		cfg := config.NewProviderConfig("gcp")
		switch gcpCmd {
		case "list":
			if err := cfg.Config.Read(); err != nil {
				fmt.Printf("\033[31merror: make sure you setup your project-id and region correctly: %v\033[0m\n", err)
				return
			}
			if gcpConfig, ok := cfg.Config.(*gcp.GoogleCloudConfig); ok {

				data, _ := json.MarshalIndent(gcpConfig, "", "  ")
				fmt.Printf("%v\n", string(data))
				return
			}

		case "set":
			attribute := args[1]
			value := args[2]

			cfg.Config.Read()

			switch attribute {
			case "project-id":
				if err := cfg.Config.SetProjectOrSubscription(value); err != nil {
					fmt.Printf("\033[31m%v\033[0m", err)
				}
				fmt.Printf("\033[32mGCP ProjectId set to %v\033[0m\n", value)
				cfg.Config.Save()
				return

			case "region":
				if err := cfg.Config.SetRegionOrLocation(value); err != nil {
					fmt.Printf("\033[31merror: %v\033[0m", err)
				}
				fmt.Printf("\033[32mGCP Region set to %v\033[0m\n", value)
				cfg.Config.Save()
				return

			default:
				fmt.Printf("\033[31merror: command %v does not exist. Choose one of project-id or region for GCP\033[0m", attribute)
				return
			}
		default:
			fmt.Printf("\033[31merror: command %v does not exists. Choose one of: list or set for GCP\033[0m", gcpCmd)
		}
	},
}

func init() {
	configCmd.AddCommand(gcpCmd)
}
