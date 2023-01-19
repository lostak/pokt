package keeper

import (
	"context"
	"flag"
)

var (
	port = flag.Int("port", 8080, "Tx server port")
)

type server struct {
	UnimplementedMsgServer
}

func (s *server) CreatePortfolio(ctx context.Context, msg *MsgCreatePortfolio) (*MsgCreatePortfolioResponse, error) {

	return &MsgCreatePortfolioResponse{}, nil
}

func (s *server) CreateAccount(ctx context.Context, msg *MsgCreateAccount) (*MsgCreateAccountResponse, error) {
	return &MsgCreateAccountResponse{}, nil
}
func (s *server) CreateChain(ctx context.Context, msg *MsgCreateChain) (*MsgCreateChainResponse, error) {
	return &MsgCreateChainResponse{}, nil
}
func (s *server) CreateToken(ctx context.Context, msg *MsgCreateToken) (*MsgCreateTokenResponse, error) {
	return &MsgCreateTokenResponse{}, nil
}
func (s *server) CreateAmount(ctx context.Context, msg *MsgCreateAmount) (*MsgCreateAmountResponse, error) {
	return &MsgCreateAmountResponse{}, nil
}
func (s *server) UpdateCoinGeckoId(ctx context.Context, msg *MsgUpdateCoinGeckoId) (*MsgUpdateCoinGeckoIdResponse, error) {
	return &MsgUpdateCoinGeckoIdResponse{}, nil
}
func (s *server) ClearPortfolio(ctx context.Context, msg *MsgClearPortfolio) (*MsgClearPortfolioResponse, error) {
	return &MsgClearPortfolioResponse{}, nil
}
func (s *server) ClearAccount(ctx context.Context, msg *MsgClearAccount) (*MsgClearAccountResponse, error) {
	return &MsgClearAccountResponse{}, nil
}
func (s *server) ClearChain(ctx context.Context, msg *MsgClearChain) (*MsgClearChainResponse, error) {
	return &MsgClearChainResponse{}, nil
}
func (s *server) ClearToken(ctx context.Context, msg *MsgClearToken) (*MsgClearTokenResponse, error) {
	return &MsgClearTokenResponse{}, nil
}
func (s *server) DeletePortfolio(ctx context.Context, msg *MsgDeletePortfolio) (*MsgDeletePortfolioResponse, error) {
	return &MsgDeletePortfolioResponse{}, nil
}
func (s *server) DeleteAccount(ctx context.Context, msg *MsgDeleteAccount) (*MsgDeleteAccountResponse, error) {
	return &MsgDeleteAccountResponse{}, nil
}
func (s *server) DeleteChain(ctx context.Context, msg *MsgDeleteChain) (*MsgDeleteChainResponse, error) {
	return &MsgDeleteChainResponse{}, nil
}
func (s *server) DeleteToken(ctx context.Context, msg *MsgDeleteToken) (*MsgDeleteTokenResponse, error) {
	return &MsgDeleteTokenResponse{}, nil
}
