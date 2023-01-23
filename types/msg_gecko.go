package types

import "fmt"

func (t *Token) updateGeckoId(geckoId string) {
	t.GeckoId = geckoId
}

func (e *TokenEntry) updateGeckoId(geckoId string) error {
	t := e.GetValue()
	if t == nil {
		return fmt.Errorf("Token w/ key: %s's value is nil", e.GetKey())
	}

	t.updateGeckoId(geckoId)
	return nil
}

func (c *Chain) updateTokenGeckoId(tokenName, geckoId string) error {
	e, err := c.getToken(tokenName)
	if err != nil {
		return err
	}

	return e.updateGeckoId(geckoId)
}

func (e *ChainEntry) updateTokenGeckoId(tokenName, geckoId string) error {
	c := e.GetValue()
	if c == nil {
		return fmt.Errorf("Token w/ key: %s has nil val on chain w/ key: %s", tokenName, e.GetKey())
	}

	return c.updateTokenGeckoId(tokenName, geckoId)
}

func (a *Account) updateTokenGeckoId(chainName, tokenName, geckoId string) error {
	e, err := a.getChain(chainName)
	if err != nil {
		return err
	}

	return e.updateTokenGeckoId(tokenName, geckoId)
}

func (e *AccountEntry) updateTokenGeckoId(chainName, tokenName, geckoId string) error {
	a := e.GetValue()
	if a == nil {
		return fmt.Errorf("Account w/ key: %s has nil val", e.GetKey())
	}

	return a.updateTokenGeckoId(chainName, tokenName, geckoId)
}

func (p *Portfolio) UpdateTokenGeckoId(accountName, chainName, tokenName, geckoId string) error {
	account, err := p.GetAccount(accountName)
	if err != nil {
		return err
	}

	return account.updateTokenGeckoId(chainName, tokenName, geckoId)
}
