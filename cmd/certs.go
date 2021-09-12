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
	"fmt"
	"log"
	"os"

	"github.com/dylanrhysscott/terrarium/internal/ssl"
	"github.com/spf13/cobra"
)

var dir string
var commonName string
var name string
var dns []string
var bits int
var force bool

// certsCmd represents the certs command
var certsCmd = &cobra.Command{
	Use:   "certs",
	Short: "Generates self signed certs for the Terraform login protocol",
	Long:  `Generates self signed certs for the Terraform login protocol`,
	Run: func(cmd *cobra.Command, args []string) {
		if force {
			err := os.RemoveAll(dir)
			if err != nil {
				log.Fatal(err)
			}
		}
		keyName := fmt.Sprintf("%s.key", name)
		certName := fmt.Sprintf("%s.pem", name)
		_, err := os.Stat(dir)
		if err == nil {
			log.Println(".certs directory exists. Use --force to regenerate certs")
			return
		}
		ca := ssl.NewCA(4096)
		fmt.Println("Creating Root CA...")
		err = ca.GenerateRootCA(commonName, dir)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Creating Client Certs...")
		key, cert, err := ca.CreateClientCertificate(commonName, dns, bits)
		if err != nil {
			log.Fatal(err)
		}
		err = ssl.Write(dir, keyName, ssl.GetPrivateKey(key))
		if err != nil {
			log.Fatal(err)
		}
		err = ssl.Write(dir, certName, ssl.GetCertificate(cert.Raw))
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(certsCmd)
	certsCmd.Flags().StringVarP(&name, "name", "n", "terrarium", "Name of the generated certs file")
	certsCmd.Flags().StringVarP(&dir, "directory", "d", ".certs", "Directory for the generated certs")
	certsCmd.Flags().IntVarP(&bits, "bits", "b", 4096, "Key size")
	certsCmd.Flags().StringVarP(&commonName, "common-name", "x", "localhost", "CA and Cert CN")
	certsCmd.Flags().StringArrayVarP(&dns, "dns", "u", []string{"localhost"}, "Cert SANS")
	certsCmd.Flags().BoolVarP(&force, "force", "f", false, "Forces recreation of CA and certs")
}
