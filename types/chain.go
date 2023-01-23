package types

import (
	"fmt"

	"google.golang.org/protobuf/types/known/timestamppb"
)

/*
	TODO:
		- Add Chain CRUD
*/

func createBlankChain(address string) *Chain {
	tokens := make(map[string]*TokenEntry)

	return &Chain{
		Addr:   address,
		Tokens: tokens,
	}
}

func (c *Chain) getToken(tokenName string) (*Token, error) {
	entries := c.GetTokens()
	if entries == nil {
		return nil, fmt.Errorf("Token entry map is not allocated")
	}

	entry := entries[tokenName]
	if entry == nil {
		return nil, fmt.Errorf("Token entry w/ key: %s is nil", tokenName)
	}

	token := entry.GetValue()
	if token == nil {
		return nil, fmt.Errorf("Token entry w/ key: %s's value is nil", tokenName)
	}

	return token, nil
}

func (c *Chain) updateAddress(address string) {
	c.Addr = address
}

func (c *Chain) addToken(tokenName string) error {
	// Check for existence
	_, err := c.getToken(tokenName)
	if err == nil {
		return fmt.Errorf("Token: %s already exists on account", tokenName)
	}

	// Create and add new Token
	token := createTokenEntry(tokenName)
	c.Tokens[tokenName] = token

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

func (c *Chain) updateTokenGeckoId(tokenName, geckoId string) error {
	token, err := c.getToken(tokenName)
	if err != nil {
		return err
	}

	token.GeckoId = geckoId
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
	token, err := c.getToken(tokenName)
	if err != nil {
		return err
	}

	/*
		TODO: add support for any base denom
	*/

	entry := &AmountEntry{
		Key:   timestamppb.Now().Seconds,
		Value: createAmountData(amount, tokenName, "usd"),
	}

	token.setAmount(entry.GetKey(), entry)

	return nil
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

func (c *Chain) nestedPrint(indent, incr, name string) {
	fmt.Printf("%sChain: \n", indent, name)
	indent += incr

	for _, token := range c.GetTokens() {
		token.nestedPrint(indent, incr)
	}
}
