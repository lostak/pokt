package types

import (
	"fmt"
)

/*
	TODO:
		- Add Chain CRUD
*/

func createBlankChain(address string) *Chain {
	return &Chain{
		Addr:   address,
		Tokens: make(map[string]*TokenEntry),
	}
}

func (c *Chain) getToken(tokenName string) (*TokenEntry, error) {
	entries := c.GetTokens()
	if entries == nil {
		c.Tokens = make(map[string]*TokenEntry)
		return nil, fmt.Errorf("Token entry map was not allocated - now is")
	}

	entry := entries[tokenName]
	if entry == nil {
		return nil, fmt.Errorf("Token entry w/ key: %s is nil", tokenName)
	}

	return entry, nil
}

func (c *Chain) updateAddress(address string) {
	c.Addr = address
}

func (c *Chain) addToken(tokenName string) error {
	if c.GetTokens() == nil {
		c.Tokens = make(map[string]*TokenEntry)
	}

	// Create and add new Token
	c.Tokens[tokenName] = createTokenEntry(tokenName)

	fmt.Println("Token added")
	return nil
}

func (c *Chain) removeToken(tokenName string) error {
	_, err := c.getToken(tokenName)
	if err != nil {
		return err
	}

	delete(c.Tokens, tokenName)

	return nil
}

func (c *Chain) updateTokenName(tokenName, newName string) error {
	entries := c.GetTokens()
	if entries == nil {
		return fmt.Errorf("Chain's token map has not been allocated")
	}

	entry := entries[tokenName]
	if entry == nil {
		return fmt.Errorf("Token w/ key: %s is nil", tokenName)
	}

	entry.updateKey(newName)
	return nil
}

func (c *Chain) addTokenAmount(tokenName string, amount float64) error {
	t, err := c.getToken(tokenName)
	if err != nil {
		return err
	}

	/*
		TODO: add support for any base denom
	*/

	return t.addTokenAmount(amount)
}

func (c *Chain) clearTokenHistory(tokenName string) error {
	token, err := c.getToken(tokenName)
	if err != nil {
		return err
	}

	token.deleteHistory()
	return nil
}

func (c *Chain) deleteHistory() {
	for _, entry := range c.GetTokens() {
		token := entry.GetValue()
		if token == nil {
			continue
		}

		token.deleteHistory()
	}
}
