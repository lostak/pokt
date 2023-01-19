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
	"github.com/lostak/pokt/types"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Initial setup command for a portfolio - WILL OVERWRITE EXISTING PORTFOLIO",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("serve called")
		portfolio := types.CreateBlankPortfolio(args[0])

		err := keeper.SetPortfolio(portfolio)
		if err != nil {
			fmt.Println(err.Error())
		}

		fmt.Printf("Portfolio created with name: %s\n", portfolio.Name)

		portfolio.Println()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
