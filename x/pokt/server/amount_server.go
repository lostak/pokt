package server

import (
	"context"
	"fmt"

	"github.com/lostak/pokt/store"
)

func (p *PoktServer) CreateAmount(ctx context.Context, msg *MsgCreateAmount) (*MsgCreateAmountResponse, error) {
	fmt.Printf("Received: %v\n", msg.String())

	portfolio, err := store.GetPortfolio()
	if err != nil {
		fmt.Println(err.Error())
		return &MsgCreateAmountResponse{}, err
	}

	err = portfolio.AddTokenAmount(msg.GetAccount(), msg.GetChain(), msg.GetToken(), msg.GetAmount())
	if err != nil {
		fmt.Println(err.Error())
		return &MsgCreateAmountResponse{}, err
	}

	if err := store.SetPortfolio(portfolio); err != nil {
		fmt.Println(err.Error())
		return &MsgCreateAmountResponse{}, err
	}

	portfolio.Println()
	return &MsgCreateAmountResponse{Portfolio: portfolio}, nil
}
