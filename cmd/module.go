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

	"github.com/dylanrhysscott/terrarium/internal/database/dynamo"
	"github.com/dylanrhysscott/terrarium/internal/database/mongodb"
	"github.com/dylanrhysscott/terrarium/internal/responder"
	"github.com/dylanrhysscott/terrarium/internal/sourcecontrol"
	"github.com/dylanrhysscott/terrarium/internal/storage/s3objects"
	"github.com/dylanrhysscott/terrarium/pkg/registry/drivers"
	"github.com/spf13/cobra"
)

var awsRegion string
var storageBackend string
var storageBackendName string
var databaseBackend string
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
		var driver drivers.TerrariumDatabaseDriver
		var storage drivers.TerrariumStorageDriver
		var err error
		if databaseBackend == "mongo" {
			driver, err = mongodb.New(databaseHost, databaseUser, databasePassword, databaseName)
			if err != nil {
				log.Fatalf("Error initialising database - %s", err.Error())
			}
		}
		if databaseBackend == "dynamodb" {
			if awsRegion == "" {
				log.Fatal("Error: No AWS Region Set")
			}
			driver, err = dynamo.New(awsRegion)
			if err != nil {
				log.Fatalf("Error initialising DynamoDB - %s", err.Error())
			}
		}
		if storageBackend == "s3" {
			if awsRegion == "" {
				log.Fatal("Error: No AWS Region Set")
			}
			storage, err = s3objects.New(awsRegion, storageBackendName)
			if err != nil {
				log.Fatalf("Error initialising S3 storage backend - %s", err.Error())
			}
		}
		if driver == nil {
			log.Fatalf("Unsupported database driver: %s", databaseBackend)
		}
		if storage == nil {
			log.Fatalf("Unsupported storage driver: %s", storageBackend)
		}
		source := &sourcecontrol.TerrariumSourceControl{}
		terrarium := api.NewTerrarium(3000, driver, storage, source, &responder.TerrariumAPIResponseWriter{}, &responder.TerrariumAPIErrorHandler{})
		err = terrarium.Serve()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	serveCmd.AddCommand(moduleCmd)
	moduleCmd.Flags().StringVarP(&databaseBackend, "database-backend", "d", "mongo", "Controls the database storage backend. Available backends: 'mongo', 'dynamodb'")
	moduleCmd.Flags().StringVarP(&storageBackend, "storage-backend", "s", "s3", "Controls the file storage backend. Available backends: 's3'")
	moduleCmd.Flags().StringVarP(&storageBackendName, "storage-backend-name", "", "terrarium-dev", "Controls the name of the storage backend. For example in the case of s3 it will be a bucket name")
	moduleCmd.Flags().StringVarP(&awsRegion, "aws-region", "", "eu-west-2", "AWS Region (required if S3 backend is used")
	moduleCmd.Flags().StringVarP(&databaseHost, "database-host", "", "", "Database Host")
	moduleCmd.Flags().StringVarP(&databaseName, "database", "", "terrarium", "Database Name")
	moduleCmd.Flags().StringVarP(&databaseUser, "database-user", "", "terrarium", "Database User")
	moduleCmd.Flags().StringVarP(&databasePassword, "database-password", "", "", "Database Password")
	moduleCmd.Flags().StringVarP(&databaseSSLMode, "database-sslmode", "", "require", "Database SSL Mode")
}
