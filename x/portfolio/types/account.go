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

func (a *Account) addChain(chainName, address string) error {
	// Check for existence
	for _, chain := range a.GetChains() {
		if chain.GetName() == chainName {
			return fmt.Errorf("Chain name: %s already exists on account: %s", chainName, a.GetName())
		}
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
	for _, chain := range a.GetChains() {
		if chain.GetName() == chainName {
			return chain.addBlankToken(tokenName)
		}
	}

	return fmt.Errorf("Chain: %s not found in account: %s", chainName, a.GetName())
}

func (a *Account) removeToken(chainName, tokenName string) error {
	for _, chain := range a.GetChains() {
		if chain.GetName() == chainName {
			return chain.removeToken(tokenName)
		}
	}

	return fmt.Errorf("Chain: %s not found in account: %s", chainName, a.GetName())
}

func (a *Account) updateTokenGeckoId(chainName, tokenName, geckoId string) error {
	for _, chain := range a.GetChains() {
		if chain.GetName() == chainName {
			return chain.updateTokenGeckoId(tokenName, geckoId)
		}
	}

	return fmt.Errorf("Chain: %s not found in account: %s", chainName, a.GetName())
}

func (a *Account) addTokenAmount(chainName, tokenName string, amount uint32) error {
	for _, chain := range a.GetChains() {
		if chain.GetName() == chainName {
			return chain.addTokenAmount(tokenName, amount)
		}
	}

	return fmt.Errorf("Chain: %s not found in account: %s", chainName, a.GetName())
}

func (a *Account) nestedPrint() {
	fmt.Printf("\tAccount:\n\t\t%s\n", a.GetName())
	for _, chain := range a.GetChains() {
		chain.nestedPrint()
	}
}
