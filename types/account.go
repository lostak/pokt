package types

import "fmt"

/*
	TODO:
		- Add Account CRUD
*/

func createBlankAccount(accountName string) *Account {
	return &Account{
		Name: accountName,
	}
}

func (a *Account) getChain(chainName string) (*Chain, error) {
	for _, chain := range a.GetChains() {
		if chain.GetName() == chainName {
			return chain, nil
		}
	}

	return nil, fmt.Errorf("Chain name: %s not found on account: %s", chainName, a.GetName())
}

func (a *Account) updateName(name string) {
	a.Name = name
}

func (a *Account) updateChainName(chainName, newName string) error {
	chain, err := a.getChain(chainName)
	if err != nil {
		return err
	}

	chain.updateName(newName)
	return nil
}

func (a *Account) updateTokenName(chainName, token, newName string) error {
	chain, err := a.getChain(chainName)
	if err != nil {
		return err
	}

	err = chain.updateTokenName(token, newName)
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

func (a *Account) addChain(chainName, address string) error {
	// Check for existence
	if _, err := a.getChain(chainName); err == nil {
		return fmt.Errorf("Chain name: %s already exists on account: %s", chainName, a.GetName())
	}

	// Create and add new Chain
	chain := createBlankChain(chainName, address)
	a.Chains = append(a.Chains, chain)

	fmt.Println("Chain added")
	return nil
}

func (a *Account) removeChain(chainName string) error {
	var chains []*Chain
	found := false

	for _, chain := range a.GetChains() {
		if chain.GetName() == chainName {
			found = true
			fmt.Println("Chain removed")
		} else {
			chains = append(chains, chain)
		}
	}

	if !found {
		return fmt.Errorf("Chain: %s not found in account: %s", chainName, a.GetName())
	}

	a.Chains = chains
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
	for _, chain := range a.GetChains() {
		if chain.GetName() == chainName {
			chain.deleteHistory()
			return nil
		}
	}

	return fmt.Errorf("Chain: %s not found in account: %s", chainName, a.GetName())
}

func (a *Account) deleteHistory() {
	for _, chain := range a.GetChains() {
		chain.deleteHistory()
	}
}

func (a *Account) nestedPrint(indent, incr string) {
	fmt.Printf("%sAccount: %s\n", indent, a.GetName())
	indent += incr

	for _, chain := range a.GetChains() {
		chain.nestedPrint(indent, incr)
	}
}
