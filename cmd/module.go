package cmd

import (
	"log"

	"github.com/terrariumcloud/terrarium/api"

	"github.com/spf13/cobra"
	fs_db "github.com/terrariumcloud/terrarium/internal/database/filesystem"
	"github.com/terrariumcloud/terrarium/internal/responder"
	fs_storage "github.com/terrariumcloud/terrarium/internal/storage/filesystem"
	"github.com/terrariumcloud/terrarium/pkg/registry/drivers"
)

var storageFilesystemRootPath string

// moduleCmd represents the module command
var moduleCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts the Terrarium Module API",
	Long:  `The Terrarium Module API allows users to manage Terraform modules in a private registry using Terrarium`,
	Run: func(cmd *cobra.Command, args []string) {
		var driver drivers.TerrariumDatabaseDriver
		var storage drivers.TerrariumStorageDriver
		var err error

		if storageFilesystemRootPath == "" {
			log.Fatal("Error: No root path specified")
		}

		driver, err = fs_db.New(storageFilesystemRootPath)
		if err != nil {
			log.Fatalf("Error initializing the filesystem database driver - %s", err.Error())
		}

		storage, err = fs_storage.New(storageFilesystemRootPath)
		if err != nil {
			log.Fatalf("Error initialising filesystem storage backend - %s", err.Error())
		}

		terrarium := api.NewTerrarium(443, driver, storage, &responder.TerrariumAPIResponseWriter{}, &responder.TerrariumAPIErrorHandler{})
		err = terrarium.Serve()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(moduleCmd)
	moduleCmd.Flags().StringVarP(&storageFilesystemRootPath, "filesystem-storage-root", "", "/terrarium/store", "Path to the storage for the filesystem storage")
}
