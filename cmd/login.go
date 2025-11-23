/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/ahmedennaifer/blov/internal/auth/gcp"
	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to the chosen provider.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		provider := args[0]
		ctx, cancel := context.WithTimeout(context.Background(), 40*time.Second)
		defer cancel()

		switch provider {
		case "gcp":

			gcpAuth := gcp.NewGCPAuthenticator()

			if err := gcpAuth.Login(ctx); err != nil {
				fmt.Printf("\033[31m%v\033[0m", err)
				return
			}

			if err := gcpAuth.Verify(ctx); err != nil {
				fmt.Printf("\033[31mCould not verify user login: %v\033[0m\n", err)
				return
			}

		case "aws":
			fmt.Println("Logging in to aws..")

		case "az":
			fmt.Println("Logging in to azure..")

		default:
			fmt.Println("Provider not recognized or not yet supported. Please specify one of : [aws, gcp, az] ")
		}
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
