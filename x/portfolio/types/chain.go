package types

import "fmt"

/*
	TODO:
		- Add Chain CRUD
*/

func createBlankChain(chainName, address string) *Chain {
	return &Chain{
		Name: chainName,
		Addr: address,
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

func (c *Chain) updateTokenGeckoId(tokenName, geckoId string) error {
	for _, token := range c.GetTokens() {
		if token.GetName() == tokenName {
			token.GeckoId = geckoId
			return nil
		}
	}
	return fmt.Errorf("Token: %s not found in chain: %s", tokenName, c.GetName())
}

func (c *Chain) addTokenAmount(tokenName string, amount uint32) error {
	for _, token := range c.GetTokens() {
		if token.GetName() == tokenName {
			amounts := token.GetAmounts()
			amounts.addAmount(amount)
			return nil
		}
	}

	return fmt.Errorf("Token: %s not found in chain: %s", tokenName, c.GetName())
}

func (c *Chain) nestedPrint() {
	fmt.Printf("\n\t\tChain:\n\t\t\t%s\n", c.GetName())
	for _, token := range c.GetTokens() {
		token.nestedPrint()
	}
}
