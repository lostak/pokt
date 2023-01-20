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
	"context"
	"fmt"
	"time"

	"github.com/lostak/pokt/keeper"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// removeAccountCmd represents the removeAccount command
var removeAccountCmd = &cobra.Command{
	Use:   "removeAccount",
	Short: "Removes all accounts with the given account from the portfolio",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("removeAccount called")

		conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		defer conn.Close()
		c := keeper.NewMsgClient(conn)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err := c.DeleteAccount(ctx, &keeper.MsgDeleteAccount{Account: args[0]})
		if err != nil {
			fmt.Printf("Could not update portfolio: %v\n", err)
			return
		}
		fmt.Println("Updated Portfolio:")
		r.GetPortfolio().Println()
	},
}

func init() {
	rootCmd.AddCommand(removeAccountCmd)
}
