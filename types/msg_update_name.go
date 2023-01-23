package types

import "fmt"

func (p *Portfolio) UpdateAccountName(accountName, newName string) error {
	entry, err := p.GetAccountEntry(accountName)
	if err != nil {
		return err
	}

	oldKey := entry.GetKey()
	entry.updateKey(newName)
	p.setAccount(newName, entry)
	delete(p.Accounts, oldKey)

	return nil
}

func (a *Account) updateChainName(chainName, newName string) error {
	entries := a.GetChains()
	if entries == nil {
		return fmt.Errorf("Account's chain map has not been allocated")
	}

	entry := entries[chainName]
	if entry == nil {
		return fmt.Errorf("Chain w/ key: %s is nil", chainName)
	}

	entry.updateKey(newName)
	return nil
}

func (a *Account) updateTokenName(chainName, token, newName string) error {
	entries := a.GetChains()
	if entries == nil {
		return fmt.Errorf("Account's chain map has not been allocated")
	}

	entry := entries[chainName]
	if entry == nil {
		return fmt.Errorf("Chain w/ key: %s is nil", chainName)
	}

	chain := entry.GetValue()
	if chain == nil {
		return fmt.Errorf("Chain w/ key: %s's value is nil", chainName)
	}

	err := chain.updateTokenName(token, newName)
	if err != nil {
		return err
	}

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
