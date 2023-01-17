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
	"fmt"

	"github.com/lostak/pokt/types"
	"github.com/spf13/cobra"
)

// deleteChainHistoryCmd represents the deleteChainHistory command
var deleteChainHistoryCmd = &cobra.Command{
	Use:   "deleteChainHistory",
	Short: "Clear the history of every token on the chain",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("deleteChainHistory called")

		portfolio, err := types.GetPortfolio()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		err = portfolio.ClearChainHistory(args[0], args[1])
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		if err := types.SetPortfolio(portfolio); err != nil {
			fmt.Println(err.Error())
			return
		}

		portfolio.Println()
	},
}

func init() {
	rootCmd.AddCommand(deleteChainHistoryCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteChainHistoryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteChainHistoryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
