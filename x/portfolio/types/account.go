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

func createAccountWithChain(accountName, address, chainName, tokenName string, amount uint32) *Account {
	var chains []*Chain
	chains[0] = createChainWithToken(chainName, address, tokenName, amount)

	return &Account{
		Name:   accountName,
		Chains: chains,
	}
}

func (a *Account) AddChain(chainName, address string) error {
	// Check for existence
	for _, chain := range a.GetChains() {
		if chain.GetName() == chainName {
			return fmt.Errorf("Chain name: %s already exists on account: %s", chainName, a.GetName())
		}
	}

	// Create and add new Chain
	chain := CreateBlankChain(chainName, address)
	a.Chains = append(a.Chains, chain)

	fmt.Println("Chain added")
	return nil
}

func (a *Account) RemoveChain(chainName string) error {
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

func (a *Account) AddToken(chainName, tokenName string) error {
	for _, chain := range a.GetChains() {
		if chain.GetName() == chainName {
			return chain.addBlankToken(tokenName)
		}
	}

	return fmt.Errorf("Chain: %s not found in account: %s", chainName, a.GetName())
}

func (a *Account) RemoveToken(chainName, tokenName string) error {
	for _, chain := range a.GetChains() {
		if chain.GetName() == chainName {
			return chain.removeToken(tokenName)
		}
	}

	return fmt.Errorf("Chain: %s not found in account: %s", chainName, a.GetName())
}
