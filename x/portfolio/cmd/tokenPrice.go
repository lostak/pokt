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

	"github.com/pokt-labs/pokt/types"
	"github.com/spf13/cobra"
)

// tokenPriceCmd represents the tokenPrice command
var tokenPriceCmd = &cobra.Command{
	Use:   "tokenPrice",
	Short: "Prints the token price for the given coin ID pair using the CoinGecko API",
	Long:  `Prints the token price for the given coin ID pair using the CoinGecko API`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("\ntokenPrice called")

		price, err := types.GetTokenPrice(args[0], args[1])

		fmt.Printf("\nPrice: %f %s\n\n", price, args[1])

		if err != nil {
			fmt.Println(err.Error())
		}

	},
}

func init() {
	rootCmd.AddCommand(tokenPriceCmd)
	tokenPriceCmd.SetUsageTemplate("Usage:\n\tportfolio tokenPrice [coinId] [vs-currency]\n")
}
