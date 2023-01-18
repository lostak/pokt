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

// addChainCmd represents the addChain command
var addChainCmd = &cobra.Command{
	Use:   "addChain",
	Short: "Add a new chain for an existing account",
	Long:  "Add a new chain for an existing account",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("addChain called")

		portfolio, err := keeper.GetPortfolio()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		err = portfolio.AddChain(args[0], args[1], args[2])
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		err = keeper.SetPortfolio(portfolio)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		portfolio.Println()

	},
}

func init() {
	rootCmd.AddCommand(addChainCmd)
}
