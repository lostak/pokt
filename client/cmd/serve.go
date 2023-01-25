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

	"github.com/lostak/pokt/server"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

const portGRPC = ":8080"
const portHTTP = ":3333"

// serveCmd represents the serve command
var ServeGrpcCmd = &cobra.Command{
	Use:   "serveGrpc",
	Short: "start portfolio msg gRPC server - WILL OVERWRITE EXISTING PORTFOLIO",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("serveMsgGrpcS called")

		flag.Parse()
		lis, err := net.Listen("tcp", portGRPC)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		s := grpc.NewServer()
		server.RegisterMsgServiceServer(s, &server.PoktServer{})
		server.RegisterQueryServiceServer(s, &server.PoktServer{})
		fmt.Printf("server listening at: %v\n", lis.Addr())
		if err := s.Serve(lis); err != nil {
			fmt.Printf("Failed to serve: %v\n", err)
		}

	},
}

var ServeHTTPCmd = &cobra.Command{
	Use:   "serveHTTP",
	Short: "start HTTP server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("serveHTTP w/ port :3333")
		err := server.ServeHTTP(portHTTP)
		if err != nil {
			fmt.Println("Failed to start server", err)
			return
		}
	},
}
