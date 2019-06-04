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

	"github.com/eugenmayer/go-sshclient/sshwrapper"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

// hpcCmd represents the hpc command
var hpcCmd = &cobra.Command{
	Use:   "hpc",
	Short: "Convenient hpc-related operations",
	Long:  `Assists in day-to-day HPC tasks, with emphasis on migrating away to Cloud environments.`,
	Run: func(cmd *cobra.Command, args []string) {
		sshHpc("ls")
	},
}

func syncLocalAWSConfig() {
	// XXX: Copies recently acquired ~/.aws credentials to the correspoding HPC user for seamless AWS operation(s).
	// It can also sync some other temporary tokens such as Illumina's `~/.igp/.session.yaml`, among others.
}

// func sshCmd(cmd string) {
// 	key, err := ioutil.ReadFile("~/.ssh/id_rsa")
// 	if err != nil {
// 		log.Fatalf("unable to read private key: %v", err)
// 	}

// 	signer, err := ssh.ParsePrivateKey(key)
// 	if err != nil {
// 		log.Fatalf("unable to parse private key: %v", err)
// 	}

// 	config := &ssh.ClientConfig{
// 		User: "brainstorm",
// 		Auth: []ssh.AuthMethod{
// 			ssh.PublicKeys(signer),
// 		},
// 		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
// 	}

// 	conn, err := ssh.Dial("tcp", "spartan.hpc.unimelb.edu.au", config)
// 	defer conn.Close()

// 	sess, err := conn.NewSession()
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer sess.Close()
// 	sessStdOut, err := sess.StdoutPipe()
// 	if err != nil {
// 		panic(err)
// 	}
// 	go io.Copy(os.Stdout, sessStdOut)
// 	sessStderr, err := sess.StderrPipe()
// 	if err != nil {
// 		panic(err)
// 	}
// 	go io.Copy(os.Stderr, sessStderr)
// 	err = sess.Run(cmd) // eg., /usr/bin/whoami
// 	if err != nil {
// 		panic(err)
// 	}
// }

func sshHpc(cmd string) {
	// XXX: Refactor out
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	sshAPI, err := sshwrapper.DefaultSshApiSetup("spartan.hpc.unimelb.edu.au", 22, "brainstorm", home+"/.ssh/id_ed25519")
	if err != nil {
		fmt.Print(err)
	}

	stdout, stderr, err := sshAPI.Run(cmd)
	if err != nil {
		log.Print(stdout)
		log.Print(stderr)
		log.Fatal(err)
	}

	log.Print(fmt.Sprintf("your ssh host '%s' returned:\n %s", sshAPI.Host, stdout))
}

func init() {
	rootCmd.AddCommand(hpcCmd)
}
