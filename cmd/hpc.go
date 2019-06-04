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
	"fmt"

	"github.com/spf13/cobra"
)

// hpcCmd represents the hpc command
var hpcCmd = &cobra.Command{
	Use:   "hpc",
	Short: "Convenient hpc-related operations",
	Long:  `Assists in day-to-day HPC tasks, with emphasis on migrating away to Cloud environments.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hpc called")
	},
}

func syncLocalAWSConfig() {
	// XXX: Copies recently acquired ~/.aws credentials to the correspoding HPC user for seamless AWS operation(s).
	// It can also sync some other temporary tokens such as Illumina's `~/.igp/.session.yaml`, among others.
}

func sshhpc() {
	// XXX: Searching for higher abstractions of https://github.com/Scalingo/go-ssh-examples/blob/master/client.go
}

func init() {
	rootCmd.AddCommand(hpcCmd)
}
