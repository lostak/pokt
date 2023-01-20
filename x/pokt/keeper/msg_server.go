package keeper

import (
	"context"
	"flag"
	"fmt"

	"github.com/lostak/pokt/types"
)

var (
	port = flag.Int("port", 8080, "Tx server port")
)

type Keeper struct {
	UnimplementedMsgServer
}

func (k *Keeper) CreatePortfolio(ctx context.Context, msg *MsgCreatePortfolio) (*MsgCreatePortfolioResponse, error) {
	fmt.Printf("Received: %v", msg.GetName())

	portfolio := types.CreateBlankPortfolio(msg.GetName())

	err := SetPortfolio(portfolio)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("Portfolio created with name: %s\n", portfolio.Name)

	portfolio.Println()

	return &MsgCreatePortfolioResponse{}, nil
}

func (k *Keeper) CreateAccount(ctx context.Context, msg *MsgCreateAccount) (*MsgCreateAccountResponse, error) {
	return &MsgCreateAccountResponse{}, nil
}
func (k *Keeper) CreateChain(ctx context.Context, msg *MsgCreateChain) (*MsgCreateChainResponse, error) {
	return &MsgCreateChainResponse{}, nil
}
func (k *Keeper) CreateToken(ctx context.Context, msg *MsgCreateToken) (*MsgCreateTokenResponse, error) {
	return &MsgCreateTokenResponse{}, nil
}
func (k *Keeper) CreateAmount(ctx context.Context, msg *MsgCreateAmount) (*MsgCreateAmountResponse, error) {
	return &MsgCreateAmountResponse{}, nil
}
func (k *Keeper) UpdateCoinGeckoId(ctx context.Context, msg *MsgUpdateCoinGeckoId) (*MsgUpdateCoinGeckoIdResponse, error) {
	return &MsgUpdateCoinGeckoIdResponse{}, nil
}
func (k *Keeper) ClearPortfolio(ctx context.Context, msg *MsgClearPortfolio) (*MsgClearPortfolioResponse, error) {
	return &MsgClearPortfolioResponse{}, nil
}
func (k *Keeper) ClearAccount(ctx context.Context, msg *MsgClearAccount) (*MsgClearAccountResponse, error) {
	return &MsgClearAccountResponse{}, nil
}
func (k *Keeper) ClearChain(ctx context.Context, msg *MsgClearChain) (*MsgClearChainResponse, error) {
	return &MsgClearChainResponse{}, nil
}
func (k *Keeper) ClearToken(ctx context.Context, msg *MsgClearToken) (*MsgClearTokenResponse, error) {
	return &MsgClearTokenResponse{}, nil
}
func (k *Keeper) DeletePortfolio(ctx context.Context, msg *MsgDeletePortfolio) (*MsgDeletePortfolioResponse, error) {
	return &MsgDeletePortfolioResponse{}, nil
}
func (k *Keeper) DeleteAccount(ctx context.Context, msg *MsgDeleteAccount) (*MsgDeleteAccountResponse, error) {
	return &MsgDeleteAccountResponse{}, nil
}
func (k *Keeper) DeleteChain(ctx context.Context, msg *MsgDeleteChain) (*MsgDeleteChainResponse, error) {
	return &MsgDeleteChainResponse{}, nil
}
func (k *Keeper) DeleteToken(ctx context.Context, msg *MsgDeleteToken) (*MsgDeleteTokenResponse, error) {
	return &MsgDeleteTokenResponse{}, nil
}
