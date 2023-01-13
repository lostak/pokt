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

// tokenPriceCmd represents the tokenPrice command
var tokenPriceCmd = &cobra.Command{
	Use:   "tokenPrice",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("\ntokenPrice called")

		const baseURL = "https://api.coingecko.com/api/v3/simple/price?ids=%s&vs_currencies=%s"
		url := fmt.Sprintf(baseURL, args[0], args[1])

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

		// Unmarshal using a generic interface
		var f interface{}
		err = json.Unmarshal(body, &f)
		if err != nil {
			fmt.Println("Error parsing JSON: ", err)
		}

		// JSON object parses into a map with string keys
		chainIdMap := f.(map[string]interface{})
		var price float64

		for _, v := range chainIdMap {
			// type assert value is JSON
			switch jsonObj := v.(type) {
			case interface{}:
				for baseDenom, amount := range jsonObj.(map[string]interface{}) {
					switch baseDenom {
					case args[1]:
						// make sure value is a number
						switch amount := amount.(type) {
						case float64:
							price = float64(amount)
							break
						default:
							fmt.Println("Incorrect type for ", baseDenom)
							break
						}
						break
					default:
						fmt.Println("Unkown key")
					}
				}
				break
			default:
				fmt.Println("Expected JSON object, got something else")
			}
		}

		prettyJSON, _ := json.MarshalIndent(f, "", "  ")

		fmt.Printf("\nResponse:\n\n%s\n\nPrice: %f\n\n", prettyJSON, price)

	},
}

func init() {
	rootCmd.AddCommand(tokenPriceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tokenPriceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// tokenPriceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
