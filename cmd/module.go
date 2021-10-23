/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"log"

	"github.com/dylanrhysscott/terrarium/api"
	"github.com/dylanrhysscott/terrarium/internal/terrariummongo"
	"github.com/dylanrhysscott/terrarium/pkg/types"
	"github.com/spf13/cobra"
)

var storageBackend string
var databaseHost string
var databaseUser string
var databasePassword string
var databaseName string
var databaseSSLMode string

// moduleCmd represents the module command
var moduleCmd = &cobra.Command{
	Use:   "module",
	Short: "Starts the Terrarium Module API",
	Long:  `The Terrarium Module API allows users to manage Terraform modules in a private registry using Terrarium`,
	Run: func(cmd *cobra.Command, args []string) {
		var driver types.TerrariumDriver
		var err error
		if storageBackend == "mongo" {
			driver, err = terrariummongo.New(databaseHost, databaseUser, databasePassword, databaseName)
		}
		if err != nil {
			log.Fatalf("Error initialising database - %s", err.Error())
		}
		if driver == nil {
			log.Fatalf("Unsupported database driver: %s", storageBackend)
		}
		terrarium := api.NewTerrarium(3000, driver, &types.TerrariumAPIResponseWriter{}, &types.TerrariumAPIErrorHandler{})
		err = terrarium.Serve()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	serveCmd.AddCommand(moduleCmd)
	moduleCmd.Flags().StringVarP(&storageBackend, "storage-backend", "s", "mongo", "Controls the database storage backend. Available backends: 'mongo'")
	moduleCmd.Flags().StringVarP(&databaseHost, "database-host", "", "", "Database Host")
	moduleCmd.Flags().StringVarP(&databaseName, "database", "", "terrarium", "Database Name")
	moduleCmd.Flags().StringVarP(&databaseUser, "database-user", "", "terrarium", "Database User")
	moduleCmd.Flags().StringVarP(&databasePassword, "database-password", "", "", "Database Password")
	moduleCmd.Flags().StringVarP(&databaseSSLMode, "database-sslmode", "", "require", "Database SSL Mode")
}
