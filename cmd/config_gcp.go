/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
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
			fmt.Printf("error: Subcommand %v is not supported for command `gcp`\nUsage: blov config gcp set region-id your-region-id\n", gcpCmd)
			return
		}
		cfg := config.NewProviderConfig("gcp")
		switch gcpCmd {
		case "list":
			if err := cfg.Config.Read(); err != nil {
				fmt.Printf("error: make sure you setup your project-id and region correctly: %v\n", err)
				return
			}
			if gcpConfig, ok := cfg.Config.(*gcp.GoogleCloudConfig); ok {
				fmt.Printf("%v\n%v\n", gcpConfig.ProjectId, gcpConfig.Region)
				return
			}

		case "set":
			attribute := args[1]
			value := args[2]

			// Read existing config first
			cfg.Config.Read() // Ignore error if file doesn't exist

			switch attribute {
			case "project-id":
				if err := cfg.Config.SetProjectOrSubscription(value); err != nil {
					fmt.Printf("%v", err)
				}
				fmt.Printf("GCP ProjectId set to %v\n", value)
				cfg.Config.Save()
				return

			case "region":
				if err := cfg.Config.SetRegionOrLocation(value); err != nil {
					fmt.Printf("error: %v", err)
				}
				fmt.Printf("GCP Region set to %v\n", value)
				cfg.Config.Save()
				return

			default:
				fmt.Printf("error: command %v does not exist. Choose one of project-id or region for GCP", attribute)
				return
			}
		default:
			fmt.Printf("error: command %v does not exists. Choose one of: list or set for GCP", gcpCmd)
		}
	},
}

func init() {
	configCmd.AddCommand(gcpCmd)
}
