package main

import (
	"flag"
	"fmt"
	"net"

	"github.com/lostak/pokt/server"
	"google.golang.org/grpc"
)

const port = ":8080"

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", port)
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
}
