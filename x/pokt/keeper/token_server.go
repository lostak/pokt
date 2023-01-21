package keeper

import (
	"context"
	"fmt"

	"github.com/lostak/pokt/store"
)

func (p *PoktServer) CreateToken(ctx context.Context, msg *MsgCreateToken) (*MsgCreateTokenResponse, error) {
	fmt.Printf("Received: %v\n", msg.String())

	portfolio, err := store.GetPortfolio()
	if err != nil {
		fmt.Println(err.Error())
		return &MsgCreateTokenResponse{}, err
	}

	err = portfolio.AddToken(msg.GetAccount(), msg.GetChain(), msg.GetToken())
	if err != nil {
		fmt.Println(err.Error())
		return &MsgCreateTokenResponse{}, err
	}

	if err := store.SetPortfolio(portfolio); err != nil {
		fmt.Println(err.Error())
		return &MsgCreateTokenResponse{}, err
	}

	portfolio.Println()
	return &MsgCreateTokenResponse{Portfolio: portfolio}, nil
}

func (p *PoktServer) UpdateCoinGeckoId(ctx context.Context, msg *MsgUpdateCoinGeckoId) (*MsgUpdateCoinGeckoIdResponse, error) {
	fmt.Printf("Received: %v\n", msg.String())

	portfolio, err := store.GetPortfolio()
	if err != nil {
		fmt.Println(err.Error())
		return &MsgUpdateCoinGeckoIdResponse{}, err

	}

	err = portfolio.UpdateTokenGeckoId(msg.GetAccount(), msg.GetChain(), msg.GetToken(), msg.GetCoinGeckoId())
	if err != nil {
		fmt.Println(err.Error())
		return &MsgUpdateCoinGeckoIdResponse{Portfolio: portfolio}, err

	}

	if err := store.SetPortfolio(portfolio); err != nil {
		fmt.Println(err.Error())
		return &MsgUpdateCoinGeckoIdResponse{Portfolio: portfolio}, err
	}

	portfolio.Println()
	return &MsgUpdateCoinGeckoIdResponse{Portfolio: portfolio}, nil
}

func (p *PoktServer) ClearToken(ctx context.Context, msg *MsgClearToken) (*MsgClearTokenResponse, error) {
	fmt.Printf("Received: %v\n", msg.String())
	portfolio, err := store.GetPortfolio()
	if err != nil {
		fmt.Println(err.Error())
		return &MsgClearTokenResponse{}, err
	}

	err = portfolio.ClearTokenHistory(msg.GetAccount(), msg.GetChain(), msg.GetToken())
	if err != nil {
		fmt.Println(err.Error())
		return &MsgClearTokenResponse{}, err
	}

	if err := store.SetPortfolio(portfolio); err != nil {
		fmt.Println(err.Error())
		return &MsgClearTokenResponse{}, err
	}

	portfolio.Println()
	return &MsgClearTokenResponse{Portfolio: portfolio}, nil
}

func (p *PoktServer) DeleteToken(ctx context.Context, msg *MsgDeleteToken) (*MsgDeleteTokenResponse, error) {
	fmt.Printf("Received: %v\n", msg.String())

	portfolio, err := store.GetPortfolio()
	if err != nil {
		fmt.Println(err.Error())
		return &MsgDeleteTokenResponse{}, err
	}

	err = portfolio.RemoveToken(msg.GetAccount(), msg.GetChain(), msg.GetToken())
	if err != nil {
		fmt.Println(err.Error())
		return &MsgDeleteTokenResponse{}, err
	}

	if err := store.SetPortfolio(portfolio); err != nil {
		fmt.Println(err.Error())
		return &MsgDeleteTokenResponse{}, err
	}

	portfolio.Println()
	return &MsgDeleteTokenResponse{Portfolio: portfolio}, nil
}
