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
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

// findCmd represents the find command
var findCmd = &cobra.Command{
	Use:   "find",
	Short: "Find data objects in AWS primary data store",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		res := apiGwFindQuery(fmt.Sprintf("%s&rowsPerPage=5000", args[0]))
		parseFindQueryResults(res)
	},
}

func apiGwFindQuery(query string) string {
	// sess := session.Must(session.NewSession(aws.NewConfig().WithMaxRetries(3)))
	// svc := apigateway.New(sess)
	// req, resp := svc.Client.NewRequest()

	// err := req.Send()
	// if err == nil { // resp is now filled
	// 	fmt.Println(resp)
	// }

	// resp, err := http.Get(fmt.Sprintf("%s/dev/files?query=%s", viper.Get("aws_ummcr_api_endpoint"), query))
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// return string(body)

	// res := request.New(aws.Config(credentials.))
	//return res
	return "WIP"
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
		fmt.Printf("%s\n", results.Rows.DataRows[i][2])
	}
}

func init() {
	rootCmd.AddCommand(findCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// findCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// findCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
