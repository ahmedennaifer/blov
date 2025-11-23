/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"time"

	config "github.com/ahmedennaifer/blov/internal/config/gcp"
	"github.com/ahmedennaifer/blov/internal/storage/gcp"
	"github.com/spf13/cobra"
)

// storageGcpListCmd represents the storageGcpList command
var storageGcpListCmd = &cobra.Command{
	Use:   "list-all",
	Short: "Lists all available buckets in the config's region",
	Run: func(cmd *cobra.Command, args []string) {
		// turn into struct, or embed into gcpStorage struct
		ctx, cancel := context.WithTimeout(context.Background(), 40*time.Second)
		defer cancel()
		gcpConfig := config.NewGCPConfig()
		gcpConfig.Read()
		gcpStorage, err := gcp.NewGCPStorageFromConfig(ctx, *gcpConfig)
		if err != nil {
			fmt.Printf("\033[31m%v\033[0m", err)
			return
		}
		buckets, err := gcpStorage.ListAll(ctx, args)
		if err != nil {
			fmt.Printf("error when llisting blobs: %v\n", err)
		}
		for _, bucket := range buckets {
			fmt.Println(bucket)
		}
	},
}

func init() {
	storageGcpCmd.AddCommand(storageGcpListCmd)
}
