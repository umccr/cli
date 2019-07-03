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
	"time"

	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/umccr/cli/util"

	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Athena/APIGW data structure in JSON
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

// findCmd represents the find command
var findCmd = &cobra.Command{
	Use:   "find",
	Short: "Find objects in AWS data store",
	Long:  `Uses the indexed S3 bucket listings database to lookup for filenames and metadata.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var res findResults

		count, _ := cmd.Flags().GetBool("count")
		if count {
			// Only interested in metadata for the total count, so only one page to retrieve
			req := apiGwFindQuery(fmt.Sprintf("/dev/files?query=%s&rowsPerPage=1", args[0]))
			json.Unmarshal([]byte(req), &res)
			fmt.Printf("%d\n", res.Meta.TotalRows)
		} else {
			// Regular request without parameters
			// XXX: implement limit for amount of rows returned
			req := apiGwFindQuery(fmt.Sprintf("/dev/files?query=%s&rowsPerPage=1", args[0]))
			json.Unmarshal([]byte(req), &res)

			totalPages := res.Meta.TotalPages
			for page := 0; page < totalPages; page++ {
				req := apiGwFindQuery(fmt.Sprintf("/dev/files?query=%s&page=%d&rowsPerPage=1000", args[0], page))
				json.Unmarshal([]byte(req), &res)
				for row := range res.Rows.DataRows {
					fmt.Printf("%s\n", res.Rows.DataRows[row][2])
				}
			}
		}
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
		fmt.Printf("ERROR: (%d): (%s)\n", res.StatusCode, res.Status)
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(res.Body)
	newStr := buf.String()

	return newStr
}

func init() {
	rootCmd.AddCommand(findCmd)

	viper.AddConfigPath(util.FindHome())
	viper.SetConfigName(".umccr")
	viper.ReadInConfig()

	findCmd.Flags().BoolP("count", "c", false, "Count number of search results only")
	//	findCmd.Flags().BoolP("sort", "s", true, "Sort the results")
}
