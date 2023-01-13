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

	"github.com/spf13/cobra"
)

// addTokenStateCmd represents the addTokenState command
var addTokenStateCmd = &cobra.Command{
	Use:   "addTokenState",
	Args:  cobra.ExactArgs(5), // [account] [chain] [token] [status] [amount]
	Short: "Add a new status for an existing token on the designated chain for the designated account",
	Long:  "Add a new status for an existing token on the designated chain for the designated account",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("addTokenState called")
	},
}

func init() {
	rootCmd.AddCommand(addTokenStateCmd)
}
