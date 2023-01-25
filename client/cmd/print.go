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
	"encoding/json"
	"fmt"
	"time"

	"github.com/lostak/pokt/server"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"
)

func GetPortfolioJSON() (string, error) {
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return "", err
	}

	defer conn.Close()
	c := server.NewQueryServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetPortfolio(ctx, &server.MsgGetPortfolio{})
	if err != nil {
		return "", err
	}

	fmt.Println("Portfolio Recieved:")
	r.GetPortfolio().Println()

	b, err := protojson.Marshal(r)
	if err != nil {
		return "", err
	}

	var j any
	err = json.Unmarshal(b, &j)
	if err != nil {

	}

	pretty, err := json.MarshalIndent(j, "", "  ")
	if err != nil {
		return "", err
	}

	return string(pretty), nil
}

// printCmd represents the print command
var printCmd = &cobra.Command{
	Use:   "print",
	Short: "Print the saved portfolio",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("print called")

		_, err := GetPortfolioJSON()
		if err != nil {
			fmt.Printf("Could not get portfolio JSON: %v\n", err.Error())
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(printCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// printCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// printCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
