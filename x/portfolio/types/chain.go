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

func createChainWithToken(chainName string, address string, tokenName string, amount uint32) *Chain {
	var tokens []*Token
	tokens[0] = createToken(tokenName, amount)

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
			return fmt.Errorf("Token: %s already exists on chain: %s", tokenName, c.GetName())
		}
	}

	// Create and add new Token
	token := createBlankToken(tokenName)
	c.Tokens = append(c.Tokens, token)

	fmt.Println("Token added")
	return nil
}

func (c *Chain) removeToken(tokenName string) error {
	var tokens []*Token
	found := false

	for _, token := range c.GetTokens() {
		if token.GetName() == tokenName {
			found = true
			fmt.Println("Token removed")
		} else {
			tokens = append(tokens, token)
		}
	}

	if !found {
		return fmt.Errorf("Token: %s not found in chain: %s", tokenName, c.GetName())
	}

	c.Tokens = tokens

	return nil
}

func (c *Chain) UpdateTokenGeckoId(tokenName, geckoId string) error {
	for _, token := range c.GetTokens() {
		if token.GetName() == tokenName {
			token.GeckoId = geckoId
			return nil
		}
	}
	return fmt.Errorf("Token: %s not found in chain: %s", tokenName, c.GetName())
}

func (c *Chain) AddTokenAmount(tokenName string, amount uint32) error {
	for _, token := range c.GetTokens() {
		if token.GetName() == tokenName {
			amounts := token.GetAmounts()
			amounts.addAmount(amount)
			return nil
		}
	}

	return fmt.Errorf("Token: %s not found in chain: %s", tokenName, c.GetName())
}
