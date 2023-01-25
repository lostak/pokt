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

	"github.com/lostak/pokt/client/grpc"
	"github.com/spf13/cobra"
)

// getPortfolioCmd represents the getPortfolio command
var getPortfolioCmd = &cobra.Command{
	Use:   "getPortfolio",
	Short: "Gets the saved portfolio as a JSON string",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("getPortfolio called")

		_, err := grpc.GetPortfolioJSON()
		if err != nil {
			fmt.Printf("Could not get portfolio JSON: %v\n", err.Error())
			return
		}
	},
}

// getAccountCmd represents the getAccount command
var getAccountCmd = &cobra.Command{
	Use:   "getAccount",
	Short: "Gets the saved account as a JSON string",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("getAccount called")

		_, err := grpc.GetAccountJSON(args[0])
		if err != nil {
			fmt.Printf("Could not get account JSON: %v\n", err.Error())
			return
		}
	},
}

// getChainCmd represents the getChain command
var getChainCmd = &cobra.Command{
	Use:   "getChain",
	Short: "Gets the saved account as a JSON string",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("getChain called")

		_, err := grpc.GetChainJSON(args[0], args[1])
		if err != nil {
			fmt.Printf("Could not get account JSON: %v\n", err.Error())
			return
		}
	},
}

// getTokenCmd represents the getToken command
var getTokenCmd = &cobra.Command{
	Use:   "getToken",
	Short: "Gets the saved account as a JSON string",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("getToken called")

		_, err := grpc.GetTokenJSON(args[0], args[1], args[2])
		if err != nil {
			fmt.Printf("Could not get account JSON: %v\n", err.Error())
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(getPortfolioCmd)
	rootCmd.AddCommand(getAccountCmd)
	rootCmd.AddCommand(getChainCmd)
	rootCmd.AddCommand(getTokenCmd)
}
