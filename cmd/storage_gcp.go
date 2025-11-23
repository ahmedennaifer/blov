/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var storageGcpCmd = &cobra.Command{
	Use:   "gcs",
	Short: "Handles operations for google cloud storage",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("storageGcp called")
	},
}

func init() {
	storageCmd.AddCommand(storageGcpCmd)
}
