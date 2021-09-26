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
	"github.com/dylanrhysscott/terrarium/internal/terrariumpsql"
	"github.com/dylanrhysscott/terrarium/pkg/types"
	"github.com/spf13/cobra"
)

var storageBackend string
var postgresHost string
var postgresUser string
var postgresPassword string
var postgresDatabase string
var postgresSSLMode string

// moduleCmd represents the module command
var moduleCmd = &cobra.Command{
	Use:   "module",
	Short: "Starts the Terrarium Module API",
	Long:  `The Terrarium Module API allows users to manage Terraform modules in a private registry using Terrarium`,
	Run: func(cmd *cobra.Command, args []string) {
		var driver types.TerrariumDriver
		var err error
		if storageBackend == "postgres" {
			driver, err = terrariumpsql.New(postgresHost, postgresUser, postgresPassword, postgresDatabase, postgresSSLMode)
		}
		if err != nil {
			log.Fatalf("Error initialising database - %s", err.Error())
		}
		if driver == nil {
			log.Fatalf("Unsupported database driver: %s", storageBackend)
		}
		terrarium := api.NewTerrarium(3000, driver)
		err = terrarium.Serve()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	serveCmd.AddCommand(moduleCmd)
	moduleCmd.Flags().StringVarP(&storageBackend, "storage-backend", "s", "postgres", "Controls the database storage backend. Available backends: 'postgres'")
	moduleCmd.Flags().StringVarP(&postgresHost, "postgres-host", "", "", "Postgres Host")
	moduleCmd.Flags().StringVarP(&postgresDatabase, "postgres-database", "", "terrarium", "Postgres Database")
	moduleCmd.Flags().StringVarP(&postgresUser, "postgres-user", "", "terrarium", "Postgres User")
	moduleCmd.Flags().StringVarP(&postgresPassword, "postgres-password", "", "", "Postgres Password")
	moduleCmd.Flags().StringVarP(&postgresSSLMode, "postgres-sslmode", "", "require", "Postgres SSL Mode")
}
