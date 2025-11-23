/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var storageCmd = &cobra.Command{
	Use:   "storage",
	Short: "Command for handling storage (buckets)",
}

func init() {
	rootCmd.AddCommand(storageCmd)
}
