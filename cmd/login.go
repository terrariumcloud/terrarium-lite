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
	"github.com/spf13/cobra"
)

var authClientID string
var authEndpoint string
var tokenEndpoint string
var audience string
var certPath string
var certKeyPath string
var ports []int

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Starts the Terrarium Login API",
	Long:  `The Terrarium Login API allows users to use 'terraform login' with the Terrarium Registry`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	serveCmd.AddCommand(loginCmd)
	loginCmd.Flags().StringVarP(&certPath, "cert-path", "", ".certs/terrarium.pem", "Path to SSL certificate")
	loginCmd.Flags().StringVarP(&certKeyPath, "cert-key-path", "", ".certs/terrarium.key", "Path to SSL key")
	loginCmd.Flags().StringVarP(&authClientID, "client-id", "c", "", "OAuth Client OD")
	loginCmd.Flags().StringVarP(&authEndpoint, "auth-endpoint", "a", "", "OAuth Authorize Endpoint")
	loginCmd.Flags().StringVarP(&tokenEndpoint, "token-endpoint", "t", "", "OAuth Token Endpoint")
	loginCmd.Flags().StringVarP(&audience, "audience", "", "https://terrarium.dylanscott.me", "OAuth Token API Audience")
	loginCmd.Flags().IntSliceVarP(&ports, "ports", "p", []int{10000, 10001}, "OAuth Ports array allow for callback URI")
}
