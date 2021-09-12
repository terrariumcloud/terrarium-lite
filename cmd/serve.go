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
	"net/http"

	"github.com/dylanrhysscott/terrarium/api/login"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var authClientID string
var authEndpoint string
var tokenEndpoint string
var certPath string
var certKeyPath string
var ports []int
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts the Terraform Registry",
	Long:  `Starts the Terraform Registry`,
	Run: func(cmd *cobra.Command, args []string) {
		loginAPI := login.NewLoginAPI(authClientID, authEndpoint, tokenEndpoint, ports)
		http.HandleFunc("/.well-known/terraform.json", loginAPI.DiscoveryHandler())
		port := ":8080"
		log.Printf("Listening on %s", port)
		log.Fatal(http.ListenAndServeTLS(port, certPath, certKeyPath, nil))
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().StringVarP(&certPath, "cert-path", "", ".certs/terrarium.pem", "Path to SSL certificate")
	serveCmd.Flags().StringVarP(&certKeyPath, "cert-key-path", "", ".certs/terrarium.key", "Path to SSL key")
	serveCmd.Flags().StringVarP(&authClientID, "client-id", "c", "", "OAuth Client OD")
	serveCmd.Flags().StringVarP(&authEndpoint, "auth-endpoint", "a", "", "OAuth Authorize Endpoint")
	serveCmd.Flags().StringVarP(&tokenEndpoint, "token-endpoint", "t", "", "OAuth Token Endpoint")
	serveCmd.Flags().IntSliceVarP(&ports, "ports", "p", []int{10000}, "OAuth Ports array allow for callback URI")
	serveCmd.MarkFlagRequired("client-id")
	serveCmd.MarkFlagRequired("auth-endpoint")
	serveCmd.MarkFlagRequired("token-endpoint")
}
