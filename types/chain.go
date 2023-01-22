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

func (c *Chain) getToken(tokenName string) (*Token, error) {
	for _, token := range c.GetTokens() {
		if token.GetName() == tokenName {
			return token, nil
		}
	}

	return nil, fmt.Errorf("Token name: %s not found on account: %s", tokenName, c.GetName())
}

func (c *Chain) updateName(name string) {
	c.Name = name
}

func (c *Chain) updateAddress(address string) {
	c.Addr = address
}

func (c *Chain) updateTokenName(tokenName, newName string) error {
	token, err := c.getToken(tokenName)
	if err != nil {
		return err
	}

	token.updateName(newName)
	return nil
}

func (c *Chain) addToken(tokenName string) error {
	// Check for existence
	_, err := c.getToken(tokenName)
	if err == nil {
		return fmt.Errorf("Token: %s already exists on account: %s", tokenName, c.GetName())
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
	token, err := c.getToken(tokenName)
	if err != nil {
		return err
	}

	token.GeckoId = geckoId
	return nil
}

func (c *Chain) addTokenAmount(tokenName string, amount float64) error {
	token, err := c.getToken(tokenName)
	if err != nil {
		return err
	}

	/*
		TODO: add support for any base denom
	*/

	amounts := token.GetAmounts()
	amounts.addAmount(amount, token.GetGeckoId(), "usd")
	return nil
}

func (c *Chain) clearTokenHistory(tokenName string) error {
	token, err := c.getToken(tokenName)
	if err != nil {
		return err
	}

	token.GetAmounts().deleteHistory()
	return nil
}

func (c *Chain) deleteHistory() {
	for _, tokens := range c.GetTokens() {
		tokens.GetAmounts().deleteHistory()
	}
}

func (c *Chain) nestedPrint(indent, incr string) {
	fmt.Printf("%sChain: %s\n", indent, c.GetName())
	indent += incr

	for _, token := range c.GetTokens() {
		token.nestedPrint(indent, incr)
	}
}
