package types

func (t *Token) deleteHistory() {
	for entry := range t.GetAmounts() {
		delete(t.Amounts, entry)
	}
}

func (e *TokenEntry) deleteHistory() {
	e.GetValue().deleteHistory()
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

func (a *Account) clearTokenHistory(chainName, tokenName string) error {
	chain, err := a.getChain(chainName)
	if err != nil {
		return err
	}

	return chain.clearTokenHistory(tokenName)
}

func (a *Account) clearChainHistory(chainName string) error {
	chain, err := a.getChain(chainName)
	if err != nil {
		return err
	}

	chain.deleteHistory()
	return nil
}

func (a *Account) deleteHistory() {
	for _, entry := range a.GetChains() {
		chain := entry.GetValue()
		if chain == nil {
			continue
		}

		chain.deleteHistory()
	}
}

func (p *Portfolio) ClearTokenHistory(accountName, chainName, tokenName string) error {
	account, err := p.GetAccount(accountName)
	if err != nil {
		return err
	}

	return account.clearTokenHistory(chainName, tokenName)
}

func (p *Portfolio) ClearChainHistory(accountName, chainName string) error {
	account, err := p.GetAccount(accountName)
	if err != nil {
		return err
	}

	return account.clearChainHistory(chainName)
}

func (p *Portfolio) ClearAccountHistory(accountName string) error {
	account, err := p.GetAccount(accountName)
	if err != nil {
		return err
	}

	account.deleteHistory()
	return nil
}

func (p *Portfolio) deleteHistory() {
	for _, entry := range p.GetAccounts() {
		account := entry.GetValue()
		if account == nil {
			continue
		}

		account.deleteHistory()
	}
}

func (p *Portfolio) ClearHistory() {
	p.deleteHistory()
}
