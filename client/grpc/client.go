package grpc

import (
	"context"
	"flag"
	"time"

	"github.com/lostak/pokt/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:8080", "the address to connect to for grpc")
)

func getMsgClientAndContext() (server.MsgServiceClient, context.Context, error) {
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, err
	}

	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	return server.NewMsgServiceClient(conn), ctx, nil
}

func getQueryClient() (server.QueryServiceClient, context.Context, error) {
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, err
	}

	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	return server.NewQueryServiceClient(conn), ctx, nil
}
