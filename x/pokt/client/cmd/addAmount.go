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
	"strconv"

	"github.com/lostak/pokt/store"
	"github.com/spf13/cobra"
)

// addAmountCmd represents the addAmount command
var addAmountCmd = &cobra.Command{
	Use:   "addAmount",
	Short: "Adds a new amount to a token's amount history",
	Args:  cobra.ExactArgs(4),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("addAmount called")

		amount, err := strconv.ParseUint(args[3], 10, 32)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		portfolio, err := store.GetPortfolio()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		err = portfolio.AddTokenAmount(args[0], args[1], args[2], uint32(amount))
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		if err := store.SetPortfolio(portfolio); err != nil {
			fmt.Println(err.Error())
			return
		}

		portfolio.Println()
	},
}

func init() {
	rootCmd.AddCommand(addAmountCmd)
}
