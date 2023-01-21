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

var updateChainName = &cobra.Command{
	Use:   "updateChainName",
	Short: "Update the chain name of an existing chain on an account",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("updateAccountName called")
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

		r, err := c.UpdateChainName(ctx, &server.MsgUpdateChainName{Account: args[0], Chain: args[1], NewName: args[2]})
		if err != nil {
			fmt.Printf("Could not update portfolio: %v\n", err)
			return
		}
		fmt.Println("Updated Portfolio:")
		r.GetPortfolio().Println()
	},
}
