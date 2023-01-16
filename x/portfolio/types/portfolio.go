package types

import (
	"fmt"
)

/*
	TODO: Add Portfolio CRUD
*/

func CreateBlankPortfolio(name string) *Portfolio {
	return &Portfolio{
		Name: name,
	}
}

func createPortfolioWithAccount(portfolioName, accountName, chainName, tokenName, address string, state States, amount uint32) *Portfolio {
	var accounts []*Account
	accounts[0] = createAccountWithChain(accountName, address, chainName, tokenName, state, amount)

	// TODO: add price histroy
	return &Portfolio{
		Name:     portfolioName,
		Accounts: accounts,
	}
}

func (p *Portfolio) AddAccount(accountName string) error {
	for _, account := range p.GetAccounts() {
		if account.GetName() == accountName {
			return fmt.Errorf("Account name: %s already exists", accountName)
		}
	}

	account := createBlankAccount(accountName)
	p.Accounts = append(p.Accounts, account)

	return nil
}

func (p *Portfolio) RemoveAccount(accountName string) bool {
	var accounts []*Account
	found := false

	for _, account := range p.GetAccounts() {
		if account.GetName() == accountName {
			found = true
			fmt.Println("Account removed")
		} else {
			accounts = append(accounts, account)
		}
	}

	p.Accounts = accounts
	return found
}

func (p *Portfolio) GetAccount(accountName string) (bool, *Account) {
	for _, account := range p.GetAccounts() {
		if account.GetName() == accountName {
			return true, account
		}
	}

	return false, nil
}

func (p *Portfolio) AddChain(accountName, chainName, address string) error {
	found, account := p.GetAccount(accountName)
	if !found {
		return fmt.Errorf("Account: %s not found\n", accountName)
	}

	err := account.AddChain(chainName, address)
	if err != nil {
		return err
	}

	return nil
}

func (p *Portfolio) PrintAccounts() {
	fmt.Printf("Portfolio:\n\t%s\n", p.GetName())
	for _, account := range p.GetAccounts() {
		fmt.Printf("\tAccount:\n\t\t%s\n", account.GetName())
	}
}

func (p *Portfolio) PrintChains() {
	fmt.Printf("Portfolio:\n\t%s\n", p.GetName())
	for _, account := range p.GetAccounts() {
		fmt.Printf("\tAccount:\n\t\t%s\n", account.GetName())
		for _, chain := range account.GetChains() {
			fmt.Printf("\t\tChain:\n\t\t\t%s\n", chain.GetName())
		}
	}
}
