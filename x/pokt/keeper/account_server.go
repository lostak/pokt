package keeper

import (
	"context"
	"fmt"

	"github.com/lostak/pokt/store"
)

func (k *Keeper) CreateAccount(ctx context.Context, msg *MsgCreateAccount) (*MsgCreateAccountResponse, error) {
	fmt.Printf("Received: %v\n", msg.String())

	portfolio, err := store.GetPortfolio()
	if err != nil {
		fmt.Println(err.Error())
		return &MsgCreateAccountResponse{}, err
	}

	err = portfolio.AddAccount(msg.GetAccount())
	if err != nil {
		fmt.Println(err.Error())
		return &MsgCreateAccountResponse{}, err
	}

	if err := store.SetPortfolio(portfolio); err != nil {
		fmt.Println(err.Error())
		return &MsgCreateAccountResponse{}, err
	}

	portfolio.Println()

	return &MsgCreateAccountResponse{Portfolio: portfolio}, nil
}

func (k *Keeper) ClearAccount(ctx context.Context, msg *MsgClearAccount) (*MsgClearAccountResponse, error) {
	fmt.Printf("Received: %v\n", msg.String())
	portfolio, err := store.GetPortfolio()
	if err != nil {
		fmt.Println(err.Error())
		return &MsgClearAccountResponse{}, err
	}

	err = portfolio.ClearAccountHistory(msg.GetAccount())
	if err != nil {
		fmt.Println(err.Error())
		return &MsgClearAccountResponse{}, err
	}

	if err := store.SetPortfolio(portfolio); err != nil {
		fmt.Println(err.Error())
		return &MsgClearAccountResponse{}, err
	}

	portfolio.Println()
	return &MsgClearAccountResponse{Portfolio: portfolio}, nil
}

func (k *Keeper) DeleteAccount(ctx context.Context, msg *MsgDeleteAccount) (*MsgDeleteAccountResponse, error) {
	fmt.Printf("Received: %v\n", msg.String())
	portfolio, err := store.GetPortfolio()
	if err != nil {
		fmt.Println(err.Error())
		return &MsgDeleteAccountResponse{}, err
	}

	err = portfolio.RemoveAccount(msg.GetAccount())
	if err != nil {
		fmt.Println(err.Error())
		return &MsgDeleteAccountResponse{}, err
	}

	if err := store.SetPortfolio(portfolio); err != nil {
		fmt.Println(err.Error())
		return &MsgDeleteAccountResponse{}, err
	}

	portfolio.Println()
	return &MsgDeleteAccountResponse{Portfolio: portfolio}, nil
}
