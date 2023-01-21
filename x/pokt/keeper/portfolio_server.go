package keeper

import (
	"context"
	"fmt"

	"github.com/lostak/pokt/store"
	"github.com/lostak/pokt/types"
)

func (p *PoktServer) CreatePortfolio(ctx context.Context, msg *MsgCreatePortfolio) (*MsgCreatePortfolioResponse, error) {
	fmt.Printf("Received: %v\n", msg.String())

	portfolio := types.CreateBlankPortfolio(msg.GetName())

	err := store.SetPortfolio(portfolio)
	if err != nil {
		fmt.Println(err.Error())
		return &MsgCreatePortfolioResponse{}, err
	}

	fmt.Printf("Portfolio created with name: %s\n", portfolio.Name)

	portfolio.Println()

	return &MsgCreatePortfolioResponse{Portfolio: portfolio}, nil
}

func (p *PoktServer) ClearPortfolio(ctx context.Context, msg *MsgClearPortfolio) (*MsgClearPortfolioResponse, error) {
	fmt.Printf("Received: %v\n", msg.String())
	portfolio, err := store.GetPortfolio()
	if err != nil {
		fmt.Println(err.Error())
		return &MsgClearPortfolioResponse{}, err
	}

	portfolio.ClearHistory()

	if err := store.SetPortfolio(portfolio); err != nil {
		fmt.Println(err.Error())
		return &MsgClearPortfolioResponse{}, err
	}

	portfolio.Println()
	return &MsgClearPortfolioResponse{Portfolio: portfolio}, nil
}

func (p *PoktServer) DeletePortfolio(ctx context.Context, msg *MsgDeletePortfolio) (*MsgDeletePortfolioResponse, error) {
	fmt.Printf("Received: %v\n", msg.String())

	reRoute := &MsgClearPortfolio{Portfolio: msg.GetPortfolio()}
	resp, err := p.ClearPortfolio(ctx, reRoute)
	if err != nil {
		return &MsgDeletePortfolioResponse{}, err
	}

	return &MsgDeletePortfolioResponse{Portfolio: resp.GetPortfolio()}, nil
}
