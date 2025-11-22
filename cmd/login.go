/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/ahmedennaifer/blov/internal/auth"
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
			gcp := auth.NewGCPAuthenticator()
			// later call each login fn
			err := gcp.Login(ctx)
			if err != nil {
				fmt.Printf("%v", err)
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
