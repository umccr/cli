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

	"github.com/spf13/cobra"
)

// shareCmd represents the share command
var shareCmd = &cobra.Command{
	Use:   "share",
	Short: "Data sharing helper",
	Long:  `Assists in data sharing process for UMCCR employees`,
	Run: func(cmd *cobra.Command, args []string) {
		receiver, _ := cmd.Flags().GetString("receiver")
		if receiver == "" {
			log.Fatal("Please specify a receiver for your data sharing, it MUST be a keybase username, i.e: ohofmann (as found in https://keybase.io/ohofmann)")
		}

		inputFile, _ := cmd.Flags().GetString("input")
		if inputFile == "" {
			log.Fatal("Please specify an input, cleartext stream")
		}

		outputFile, _ := cmd.Flags().GetString("output")
		if outputFile == "" {
			outputFile = "share_encrypted.gpg"
		}

		// Presign all urls
		// file, err := os.Open(inputFile)
		// if err != nil {
		// 	fmt.Println(err)
		// 	return
		// }
		// defer file.Close()

		// scanner := bufio.NewScanner(file)
		// scanner.Split(bufio.ScanLines)

		// var presigned []string

		// for scanner.Scan() {
		// 	presigned = append(presigned, util.PresignURL(scanner.Text()))
		// }

		// // Use PGP instead of saltpack since it is more user friendly
		// // XXX: Use keybase Go client instead: keybase1.PGPClient()... public keybase api is not very clean :/
		// shCmd := exec.Command("keybase", "pgp", "encrypt", receiver, "-i", inputFile, "-o", outputFile)

		// if err := shCmd.Run(); err != nil {
		// 	log.Fatal(err)
		// }

		// pr, pw := io.Pipe()
		// defer pw.Close()

		// // tell the command to write to our pipe
		// cmd := exec.Command("keybase", "pgp", "encrypt", receiver, "-i", inputFile, "-o", outputFile)
		// cmd.Stdout = pw

		// go func() {
		// 	defer pr.Close()
		// 	// copy the data written to the PipeReader via the cmd to stdout
		// 	if _, err := io.Copy(os.Stdout, pr); err != nil {
		// 		log.Fatal(err)
		// 	}
		// }()

		// // run the command, which writes all output to the PipeWriter
		// // which then ends up in the PipeReader
		// if err := cmd.Run(); err != nil {
		// 	log.Fatal(err)
		// }
	},
}

func init() {
	rootCmd.AddCommand(shareCmd)

	shareCmd.Flags().StringP("receiver", "r", "", "Specify the receiver of the datasharing")
	shareCmd.Flags().StringP("input", "i", "", "Specify the input file for datasharing")
	shareCmd.Flags().StringP("output", "o", "", "Specify the output file for datasharing")
}
