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
	"time"

	"github.com/lostak/pokt/server"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// updateGeckoIdCmd represents the updateGeckoId command
var updateGeckoIdCmd = &cobra.Command{
	Use:   "updateGeckoId",
	Short: "Add/Update the GeckoId for a named token on a chain for an account",
	Args:  cobra.ExactArgs(4),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("updateGeckoId called")
		flag.Parse()

		conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		defer conn.Close()
		c := server.NewMsgClient(conn)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		r, err := c.UpdateCoinGeckoId(ctx, &server.MsgUpdateCoinGeckoId{Account: args[0], Chain: args[1], Token: args[2], CoinGeckoId: args[3]})
		if err != nil {
			fmt.Printf("Could not update portfolio: %v\n", err)
			return
		}
		fmt.Println("Updated Portfolio:")
		r.GetPortfolio().Println()
	},
}

func init() {
	rootCmd.AddCommand(updateGeckoIdCmd)
}
