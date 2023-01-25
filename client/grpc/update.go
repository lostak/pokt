package grpc

import (
	"fmt"

	"github.com/lostak/pokt/server"
	"github.com/lostak/pokt/types"
)

func AddAccount(accountName string) (*types.Portfolio, error) {

	client, ctx, err := getMsgClientAndContext()

	response, err := client.CreateAccount(ctx, &server.MsgCreateAccount{Account: accountName})
	if err != nil {
		return &types.Portfolio{}, err
	}

	portfolio := response.GetPortfolio()
	if portfolio == nil {
		return &types.Portfolio{}, fmt.Errorf("portfolio received is nil")
	}

	fmt.Println("Portfolio updated")

	return portfolio, nil
}
