package server

import (
	"context"
	"fmt"

	"github.com/lostak/pokt/store"
)

func (k *Keeper) CreateChain(ctx context.Context, msg *MsgCreateChain) (*MsgCreateChainResponse, error) {
	fmt.Printf("Received: %v\n", msg.String())

	portfolio, err := store.GetPortfolio()
	if err != nil {
		fmt.Println(err.Error())
		return &MsgCreateChainResponse{}, err
	}

	err = portfolio.AddChain(msg.Account, msg.Chain, msg.Address)
	if err != nil {
		fmt.Println(err.Error())
		return &MsgCreateChainResponse{}, err
	}

	err = store.SetPortfolio(portfolio)
	if err != nil {
		fmt.Println(err.Error())
		return &MsgCreateChainResponse{}, err
	}
	portfolio.Println()

	return &MsgCreateChainResponse{Portfolio: portfolio}, nil
}

func (k *Keeper) ClearChain(ctx context.Context, msg *MsgClearChain) (*MsgClearChainResponse, error) {
	fmt.Printf("Received: %v\n", msg.String())

	portfolio, err := store.GetPortfolio()
	if err != nil {
		fmt.Println(err.Error())
		return &MsgClearChainResponse{}, err
	}

	err = portfolio.ClearChainHistory(msg.GetAccount(), msg.GetChain())
	if err != nil {
		fmt.Println(err.Error())
		return &MsgClearChainResponse{}, err
	}

	if err := store.SetPortfolio(portfolio); err != nil {
		fmt.Println(err.Error())
		return &MsgClearChainResponse{}, err
	}

	portfolio.Println()
	return &MsgClearChainResponse{Portfolio: portfolio}, nil
}

func (k *Keeper) DeleteChain(ctx context.Context, msg *MsgDeleteChain) (*MsgDeleteChainResponse, error) {
	fmt.Printf("Received: %v\n", msg.String())

	portfolio, err := store.GetPortfolio()
	if err != nil {
		fmt.Println(err.Error())
		return &MsgDeleteChainResponse{}, err
	}

	err = portfolio.RemoveChain(msg.GetAccount(), msg.GetChain())
	if err != nil {
		fmt.Println(err.Error())
		return &MsgDeleteChainResponse{}, err
	}

	if err := store.SetPortfolio(portfolio); err != nil {
		fmt.Println(err.Error())
		return &MsgDeleteChainResponse{}, err
	}

	portfolio.Println()
	return &MsgDeleteChainResponse{Portfolio: portfolio}, nil
}
