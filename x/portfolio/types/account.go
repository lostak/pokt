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

func createAccountWithChain(accountName, address, chainName, tokenName string, state States, amount uint32) *Account {
	var chains []*Chain
	chains[0] = createChainWithToken(chainName, address, tokenName, state, amount)

	return &Account{
		Name:   accountName,
		Chains: chains,
	}
}

func (a *Account) AddChain(chainName, address string) error {
	for _, chain := range a.GetChains() {
		if chain.Name == chainName {
			return fmt.Errorf("Chain name: %s already exists on account: %s", chainName, a.GetName())
		}
	}

	chain := CreateBlankChain(chainName, address)
	a.Chains = append(a.GetChains(), chain)

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
