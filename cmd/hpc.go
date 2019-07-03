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
	"log"
	"os"
	"path/filepath"
	"regexp"

	"github.com/eugenmayer/go-sshclient/sshwrapper"
	"github.com/kevinburke/ssh_config"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

// hpcCmd represents the hpc command
var hpcCmd = &cobra.Command{
	Use:   "hpc",
	Short: "Convenient hpc-related operations",
	Long:  `Assists in day-to-day HPC tasks, with emphasis on migrating away to Cloud environments.`,
	Run: func(cmd *cobra.Command, args []string) {
		//sshHpc("spartan.hpc.unimelb.edu.au", "brainstorm", "")
		readSSHConfig()
	},
}

func syncLocalAWSConfig(sshAPI *sshwrapper.SshApi, localHomedir string) {
	// XXX: Refactor to only copy the assumed role.
	log.Println("Syncronizing STS AWS creds to HPC cluster")
	sshAPI.CopyToRemote(localHomedir+"/.aws/credentials", ".aws/credentials")
	sshAPI.CopyToRemote(localHomedir+"/.aws/config", "aws/credentials")

	// expose on the CLI as:
	// umccr hpc sync --aws
}

func syncLocalIlluminaConfig(sshAPI *sshwrapper.SshApi, localHomedir string) {
	log.Println("Syncronizing Illumina creds to HPC cluster")
	sshAPI.CopyToRemote(localHomedir+"/.igp/.session.yaml", ".igp/.session.yaml")
}

func syncHpcToAws() {
	// expose on the CLI as:
	// umccr hpc sync <hpc_filesystem_path> [--env [prod | dev]]. Sensible default: dev?
}

func setupRRSync() {
	//XXX: Facilitate rrsync (restricted rsync) as shown here: https://opus.nci.org.au/display/Help/Using+SSH+keys

	// expose on the CLI as:
	// umccr hpc share <hpc_filesystem_path>
}

func sshHpc(hpcHost string, hpcUser string, cmd string) {
	// XXX: Refactor out of here to util lib
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// XXX: Scan ~/.ssh/config and discover priv/pubkey settings
	sshAPI, err := sshwrapper.DefaultSshApiSetup(hpcHost, 22, hpcUser, home+"/.ssh/id_ed25519")
	if err != nil {
		fmt.Print(err)
	}

	if cmd != "" {
		stdout, stderr, err := sshAPI.Run(cmd)
		if err != nil {
			log.Print(stdout)
			log.Print(stderr)
			log.Fatal(err)
		}
	}

	syncLocalAWSConfig(sshAPI, home)
	syncLocalIlluminaConfig(sshAPI, home)

	if err != nil {
		log.Fatal(err)
	}

	//log.Printf("your ssh host '%s' returned:\n %s", sshAPI.Host, stdout)
}

func readSSHConfig() {
	f, _ := os.Open(filepath.Join(os.Getenv("HOME"), ".ssh", "config"))
	cfg, _ := ssh_config.Decode(f)
	for _, host := range cfg.Hosts {
		fmt.Println("patterns:", host.Patterns)
		for _, node := range host.Nodes {
			// Manipulate the nodes as you see fit, or use a type switch to
			// distinguish between Empty, KV, and Include nodes.
			//fmt.Println(node.String())
			if node.String() == "spartan" {
				fmt.Println("FOOO")
			}
		}
	}

	// Print the config to stdout:
	findHpcSSHClusters(cfg.String())
}

func findHpcSSHClusters(sshConfig string) {
	var validID = regexp.MustCompile(`^[a-z]+\[[0-9]+\]$`)
	if validID.MatchString("spartan") {
		fmt.Println(validID.FindStringSubmatch("spartan"))
	}
}

func init() {
	rootCmd.AddCommand(hpcCmd)
}
