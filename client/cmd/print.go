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

	"github.com/lostak/pokt/store"
	"github.com/spf13/cobra"
)

// printCmd represents the print command
var printCmd = &cobra.Command{
	Use:   "print",
	Short: "Print the saved portfolio",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("print called")

		portfolio, err := store.GetPortfolio()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		portfolio.Println()

		/* TODO: Add query & print GetPortfolio response
		conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		defer conn.Close()
		c := server.NewMsgClient(conn)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err := c.GetPortfolio(ctx, &server.MsgPrintPortfolio{})
		if err != nil {
			fmt.Printf("Could not update portfolio: %v\n", err)
			return
		}

		fmt.Println("Updated Portfolio:")
		r.GetPortfolio().Println()
		*/
	},
}

func init() {
	rootCmd.AddCommand(printCmd)
}
