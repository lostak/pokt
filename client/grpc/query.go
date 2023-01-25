package grpc

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/lostak/pokt/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"
)

func bytesToPretty(b []byte) (string, error) {
	var ugly any
	err := json.Unmarshal(b, &ugly)
	if err != nil {
		return "", err
	}

	pretty, err := json.MarshalIndent(ugly, "", "  ")
	if err != nil {
		return "", err
	}

	return string(pretty), nil
}

func GetPortfolioJSON() (string, error) {
	conn, err := grpc.Dial(*adr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return "", nil
	}

	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	client := server.NewQueryServiceClient(conn)

	response, err := client.GetPortfolio(ctx, &server.MsgGetPortfolio{})
	if err != nil {
		return "", err
	}

	portfolio := response.GetPortfolio()
	if portfolio == nil {
		return "", fmt.Errorf("portfolio recieved is nil")
	}

	fmt.Println("Portfolio Recieved:")
	portfolio.Println()

	b, err := protojson.Marshal(portfolio)
	if err != nil {
		return "", err
	}

	return bytesToPretty(b)
}

func GetAccountJSON(accountName string) (string, error) {
	client, ctx, err := getQueryClient()
	if err != nil {
		return "", err
	}

	response, err := client.GetAccount(ctx, &server.MsgGetAccount{Name: accountName})
	if err != nil {
		return "", err
	}

	account := response.GetAccount()
	if account == nil {
		return "", fmt.Errorf("Account received is nil")
	}

	fmt.Println("Account received:")
	account.Println()

	b, err := protojson.Marshal(account)
	if err != nil {
		return "", err
	}

	return bytesToPretty(b)
}

func GetChainJSON(accountName, chainName string) (string, error) {
	client, ctx, err := getQueryClient()
	if err != nil {
		return "", err
	}

	response, err := client.GetChain(ctx, &server.MsgGetChain{Account: accountName, Name: chainName})
	if err != nil {
		return "", err
	}

	account := response.GetChain()
	if account == nil {
		return "", fmt.Errorf("Account received is nil")
	}

	fmt.Println("Account received:")
	account.Println()

	b, err := protojson.Marshal(account)
	if err != nil {
		return "", err
	}

	return bytesToPretty(b)
}

func GetTokenJSON(accountName, chainName, tokenName string) (string, error) {
	client, ctx, err := getQueryClient()
	if err != nil {
		return "", err
	}

	response, err := client.GetToken(ctx, &server.MsgGetToken{Account: accountName, Chain: chainName, Name: tokenName})
	if err != nil {
		return "", err
	}

	account := response.GetToken()
	if account == nil {
		return "", fmt.Errorf("Account received is nil")
	}

	fmt.Println("Account received:")
	account.Println()

	b, err := protojson.Marshal(account)
	if err != nil {
		return "", err
	}

	return bytesToPretty(b)
}
