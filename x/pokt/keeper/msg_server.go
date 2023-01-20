package keeper

import (
	"context"
	"flag"
	"fmt"

	"github.com/lostak/pokt/store"
	"github.com/lostak/pokt/types"
)

var (
	port = flag.Int("port", 8080, "Tx server port")
)

type Keeper struct {
	UnimplementedMsgServer
}

func (k *Keeper) CreatePortfolio(ctx context.Context, msg *MsgCreatePortfolio) (*MsgCreatePortfolioResponse, error) {
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
func (k *Keeper) CreateToken(ctx context.Context, msg *MsgCreateToken) (*MsgCreateTokenResponse, error) {
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
func (k *Keeper) CreateAmount(ctx context.Context, msg *MsgCreateAmount) (*MsgCreateAmountResponse, error) {
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
func (k *Keeper) UpdateCoinGeckoId(ctx context.Context, msg *MsgUpdateCoinGeckoId) (*MsgUpdateCoinGeckoIdResponse, error) {
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
func (k *Keeper) ClearPortfolio(ctx context.Context, msg *MsgClearPortfolio) (*MsgClearPortfolioResponse, error) {
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
func (k *Keeper) ClearToken(ctx context.Context, msg *MsgClearToken) (*MsgClearTokenResponse, error) {
	fmt.Printf("Received: %v\n", msg.String())
	portfolio, err := store.GetPortfolio()
	if err != nil {
		fmt.Println(err.Error())
		return &MsgClearTokenResponse{}, err
	}

	err = portfolio.ClearTokenHistory(args[0], args[1], args[2])
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
func (k *Keeper) DeletePortfolio(ctx context.Context, msg *MsgDeletePortfolio) (*MsgDeletePortfolioResponse, error) {
	fmt.Printf("Received: %v\n", msg.String())

	return &MsgDeletePortfolioResponse{}, nil
}
func (k *Keeper) DeleteAccount(ctx context.Context, msg *MsgDeleteAccount) (*MsgDeleteAccountResponse, error) {
	fmt.Printf("Received: %v\n", msg.String())

	return &MsgDeleteAccountResponse{}, nil
}
func (k *Keeper) DeleteChain(ctx context.Context, msg *MsgDeleteChain) (*MsgDeleteChainResponse, error) {
	fmt.Printf("Received: %v\n", msg.String())

	return &MsgDeleteChainResponse{}, nil
}
func (k *Keeper) DeleteToken(ctx context.Context, msg *MsgDeleteToken) (*MsgDeleteTokenResponse, error) {
	fmt.Printf("Received: %v\n", msg.String())

	return &MsgDeleteTokenResponse{}, nil
}
