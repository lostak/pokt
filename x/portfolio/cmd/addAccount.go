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

// addAccountCmd represents the addAccount command
var addAccountCmd = &cobra.Command{
	Use:   "addAccount",
	Short: "Add a new account for an existing portfolio",
	Long:  "Add a new account for an existing portfolio",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("addAccount called")

		portfolio, err := types.GetPortfolio()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		err = portfolio.AddAccount(args[0])
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
	rootCmd.AddCommand(addAccountCmd)
}
