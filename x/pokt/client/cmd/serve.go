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
	"flag"
	"fmt"
	"net"

	"github.com/lostak/pokt/keeper"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var ()

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Initial setup command for a portfolio - WILL OVERWRITE EXISTING PORTFOLIO",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("serve called")

		flag.Parse()
		lis, err := net.Listen("tcp", ":8080")
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		server := grpc.NewServer()
		keeper.RegisterMsgServer(server, &keeper.Keeper{})
		fmt.Printf("server listening at: %v\n", lis.Addr())
		if err := server.Serve(lis); err != nil {
			fmt.Printf("Failed to serve: %v\n", err)
		}

	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
