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

// removeTokenCmd represents the removeToken command
var removeTokenCmd = &cobra.Command{
	Use:   "removeToken",
	Short: "Remove token from a chain on account",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("removeToken called")

		conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		defer conn.Close()
		c := keeper.NewMsgClient(conn)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err := c.DeleteToken(ctx, &keeper.MsgDeleteToken{Account: args[0], Chain: args[1], Token: args[2]})
		if err != nil {
			fmt.Printf("Could not update portfolio: %v\n", err)
			return
		}
		fmt.Println("Updated Portfolio:")
		r.GetPortfolio().Println()
	},
}

func init() {
	rootCmd.AddCommand(removeTokenCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// removeTokenCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// removeTokenCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
