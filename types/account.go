package types

import "fmt"

/*
	TODO:
		- Add Account CRUD
*/

func createAccount() *Account {
	return &Account{
		Chains: make(map[string]*ChainEntry),
	}
}

func (a *Account) getChain(chainName string) (*Chain, error) {
	entries := a.GetChains()
	if entries == nil {
		a.Chains = make(map[string]*ChainEntry)
		return nil, fmt.Errorf("Chain map was not allocated in account")
	}

	entry := entries[chainName]
	if entry == nil {
		return nil, fmt.Errorf("Chain name: %s not found", chainName)
	}

	chain := entry.GetValue()
	if chain == nil {
		return nil, fmt.Errorf("Chain entry w/ key: %s's value is nil", chainName)
	}

	return chain, nil
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

func (a *Account) updateAddress(chainName, address string) error {
	chain, err := a.getChain(chainName)
	if err != nil {
		return err
	}

	chain.updateAddress(address)
	return nil
}

func (a *Account) addChainEntry(chain *ChainEntry) {
	if x := a.GetChains(); x == nil {
		a.Chains = make(map[string]*ChainEntry)
	}

	a.Chains[chain.GetKey()] = chain
}

func (a *Account) addChain(chainName, address string) error {
	if a.GetChains() == nil {
		a.Chains = make(map[string]*ChainEntry)
	}

	c, err := a.getChain(chainName)
	if err == nil && c != nil {
		return fmt.Errorf("Chain already exists on account")
	}

	a.Chains[chainName] = createChainEntry(chainName)
	return nil
}

func (a *Account) removeChain(chainName string) error {
	_, err := a.getChain(chainName)
	if err != nil {
		return err
	}

	delete(a.Chains, chainName)
	return nil
}

func (a *Account) addToken(chainName, tokenName string) error {
	chain, err := a.getChain(chainName)
	if err != nil {
		return err
	}

	return chain.addToken(tokenName)
}

func (a *Account) removeToken(chainName, tokenName string) error {
	chain, err := a.getChain(chainName)
	if err != nil {
		return err
	}

	return chain.removeToken(tokenName)
}

func (a *Account) addTokenAmount(chainName, tokenName string, amount float64) error {
	chain, err := a.getChain(chainName)
	if err != nil {
		return err
	}

	return chain.addTokenAmount(tokenName, amount)
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
