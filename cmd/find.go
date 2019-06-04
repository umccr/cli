// Copyright © 2019 Roman Valls Guimera <brainstorm at nopcode org>
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
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/mitchellh/go-homedir"

	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// findCmd represents the find command
var findCmd = &cobra.Command{
	Use:   "find",
	Short: "Find objects in AWS data store",
	Long:  `Uses the indexed S3 bucket listings database to lookup for filenames and metadata.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// XXX: Handle pagination gracefully, perhaps tailored to usecases?
		res := apiGwFindQuery(fmt.Sprintf("/dev/files?query=%s&rowsPerPage=5000", args[0]))
		parseFindQueryResults(res)
	},
}

func apiGwFindQuery(query string) string {

	cfg, err := external.LoadDefaultAWSConfig(
		external.WithSharedConfigProfile("default"),
	)

	if err != nil {
		fmt.Println("unable to create an AWS session for the provided profile")
	}

	req, _ := http.NewRequest(http.MethodGet, viper.GetString("aws_apigw_endpoint")+query, nil)
	signer := v4.NewSigner(cfg.Credentials)
	_, err = signer.Sign(req, nil, "execute-api", cfg.Region, time.Now())
	if err != nil {
		fmt.Printf("failed to sign request: (%v)\n", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("failed to call remote service: (%v)\n", err)
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		fmt.Printf("service returned a status not 200: (%d)\n", res.StatusCode)
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(res.Body)
	newStr := buf.String()

	// XXX
	fmt.Println("WARNING: (Auto)-Pagination has not been implemented yet, the results above are a subset")

	return newStr
}

func parseFindQueryResults(jsonTxt string) {

	// http://json2struct.mervine.net/
	type findResults struct {
		Meta struct {
			Page       int `json:"page"`
			Size       int `json:"size"`
			Start      int `json:"start"`
			TotalPages int `json:"totalPages"`
			TotalRows  int `json:"totalRows"`
		} `json:"meta"`
		Rows struct {
			DataRows  [][]string `json:"dataRows"`
			HeaderRow []struct {
				Key      string `json:"key"`
				Sortable bool   `json:"sortable"`
			} `json:"headerRow"`
		} `json:"rows"`
	}

	var results findResults

	// XXX:
	// * implement different find flags such as filesize and timestamp
	// * {"message":"Missing Authentication Token"}... substitute for "please run umccr login" message
	json.Unmarshal([]byte(jsonTxt), &results)
	for i := range results.Rows.DataRows {
		fmt.Printf("%s\n", results.Rows.DataRows[i][2]) // keynames ("paths", filenames, extensions... object names really)
	}
}

func init() {
	rootCmd.AddCommand(findCmd)

	// XXX: Centralise config reading business
	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	viper.AddConfigPath(home)
	viper.SetConfigName(".umccr")
	viper.ReadInConfig()

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// findCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// findCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
