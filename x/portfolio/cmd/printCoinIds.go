/*
Copyright © 2023 Bradley Lostak <lostak@engineer.com>

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
	"encoding/json"
	"fmt"

	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"
)

// printCoinIdsCmd represents the printCoinIds command
var printCoinIdsCmd = &cobra.Command{
	Use:   "printCoinIds",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("printCoinIds called")

		const url = "https://api.coingecko.com/api/v3/coins/list"

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Print(err.Error())
			return
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Print(err.Error())
			return
		}

		defer res.Body.Close()
		body, readErr := ioutil.ReadAll(res.Body)
		if readErr != nil {
			fmt.Print(err.Error())
			return
		}

		var f interface{}
		err = json.Unmarshal(body, &f)
		if err != nil {
			fmt.Print(err.Error())
			return
		}

		response, _ := json.MarshalIndent(f, "", "\t")
		fmt.Println(string(response))
	},
}

func init() {
	rootCmd.AddCommand(printCoinIdsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// printCoinIdsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// printCoinIdsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
