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
	"flag"
	"fmt"
	"strconv"
	"time"

	"github.com/lostak/pokt/keeper"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// addAmountCmd represents the addAmount command
var addAmountCmd = &cobra.Command{
	Use:   "addAmount",
	Short: "Adds a new amount to a token's amount history",
	Args:  cobra.ExactArgs(4),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("addAmount called")
		flag.Parse()

		conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		defer conn.Close()
		c := keeper.NewMsgClient(conn)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		amount, err := strconv.ParseUint(args[3], 10, 32)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		r, err := c.CreateAmount(ctx, &keeper.MsgCreateAmount{Account: args[0], Chain: args[1], Token: args[2], Amount: uint32(amount)})
		if err != nil {
			fmt.Printf("Could not update portfolio: %v\n", err)
			return
		}
		fmt.Println("Updated Portfolio:")
		r.GetPortfolio().Println()

	},
}

func init() {
	rootCmd.AddCommand(addAmountCmd)
}
