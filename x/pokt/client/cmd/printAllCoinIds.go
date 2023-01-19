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

// printCoinIdsCmd represents the printCoinIds command
var printAllCoinIdsCmd = &cobra.Command{
	Use:   "printAllCoinIds",
	Short: "Prints the id, symbol and name of all coins supported by the CoinGecko API",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("printCoinAllIds called")

		types.PrintAllCoinIds()
	},
}

func init() {
	rootCmd.AddCommand(printAllCoinIdsCmd)
}
