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

	"github.com/lostak/pokt/keeper"
	"github.com/spf13/cobra"
)

// removeChainCmd represents the removeChain command
var removeChainCmd = &cobra.Command{
	Use:   "removeChain",
	Short: "From move a chain from an account",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("removeChain called")

		portfolio, err := keeper.GetPortfolio()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		err = portfolio.RemoveChain(args[0], args[1])
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		if err := keeper.SetPortfolio(portfolio); err != nil {
			fmt.Println(err.Error())
			return
		}

		portfolio.Println()
	},
}

func init() {
	rootCmd.AddCommand(removeChainCmd)
}