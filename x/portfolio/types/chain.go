package types

import "fmt"

/*
	TODO:
		- Add Chain CRUD
*/

func CreateBlankChain(chainName, address string) *Chain {
	return &Chain{
		Name: chainName,
		Addr: address,
	}
}

func createChainWithToken(chainName string, address string, tokenName string, state States, amount uint32) *Chain {
	var tokens []*Token
	tokens[0] = createToken(tokenName, state, amount)

	// TODO: Add timestamp
	return &Chain{
		Name:   chainName,
		Addr:   address,
		Tokens: tokens,
	}
}

func (c *Chain) addBlankToken(tokenName string) error {
	// Check for existence
	for _, token := range c.GetTokens() {
		if token.GetName() == tokenName {
			return fmt.Errorf("Token: %s already exists in state: %s on chain: %s", tokenName, token.GetStates()[0], c.GetName())
		}
	}

	// Create and add new Token
	token := createBlankToken(tokenName)
	c.Tokens = append(c.Tokens, token)

	fmt.Println("Token added")
	return nil
}
