// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

// shareCmd represents the share command
var shareCmd = &cobra.Command{
	Use:   "share",
	Short: "Data sharing helper",
	Long:  `Assists in data sharing process for UMCCR employees`,
	Run: func(cmd *cobra.Command, args []string) {
		collaborator, _ := cmd.Flags().GetString("collab")
		if collaborator == "" {
			log.Fatal("Please specify a recipient for your data sharing, it MUST be a keybase username, i.e: ohofmann (as found in https://keybase.io/ohofmann)")
		}

		inputFile, _ := cmd.Flags().GetString("input")
		if inputFile == "" {
			log.Fatal("Please specify an input, cleartext file")
		}

		outputFile, _ := cmd.Flags().GetString("output")
		if collaborator == "" {
			log.Fatal("Please specify an output, encrypted file")
		}

		//XXX: Presign all urls
		//util.aws.PresignURL(args[0])

		// Use PGP instead of saltpack since it is more user friendly
		shCmd := exec.Command("keybase", "pgp", "encrypt", collaborator, "-i", inputFile, "-o", outputFile)

		if err := shCmd.Run(); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(shareCmd)

	shareCmd.Flags().StringP("collab", "c", "", "Specify the recipient of the datasharing")
	shareCmd.Flags().StringP("input", "i", "", "Specify the input file for datasharing")
	shareCmd.Flags().StringP("output", "o", "", "Specify the output file for datasharing")
}
