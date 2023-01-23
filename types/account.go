package types

import "fmt"

/*
	TODO:
		- Add Account CRUD
*/

func createAccount() *Account {
	chains := make(map[string]*ChainEntry)

	return &Account{
		Chains: chains,
	}
}

func (a *Account) getChain(chainName string) (*Chain, error) {
	entry := a.GetChains()[chainName]
	if entry == nil {
		return nil, fmt.Errorf("Chain name: %s not found on account", chainName)
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

func (a *Account) addChainEntry(chain *ChainEntry) error {
	entries := a.GetChains()
	if entries == nil {
		return fmt.Errorf("Chain map not allocated")
	}

	entry := entries[chain.GetKey()]
	if entry != nil {
		return fmt.Errorf("Chain w/ key %s already exists", chain.GetKey())
	}

	entries[chain.GetKey()] = chain
	return nil
}

func (a *Account) addChain(chainName, address string) error {
	// Check for existence
	if _, err := a.getChain(chainName); err == nil {
		return fmt.Errorf("Chain name: %s already exists on account", chainName)
	}

	// Create and add new Chain
	a.Chains[chainName] = createChainEntry(chainName)

	fmt.Println("Chain added")
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

func (a *Account) updateTokenGeckoId(chainName, tokenName, geckoId string) error {
	chain, err := a.getChain(chainName)
	if err != nil {
		return err
	}

	return chain.updateTokenGeckoId(tokenName, geckoId)
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
