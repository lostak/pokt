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

func (p *Portfolio) PrintAccounts() {
	fmt.Printf("Portfolio:\n\t%s\n", p.GetName())
	for _, account := range p.Accounts {
		fmt.Printf("\tAccount:\n\t\t%s\n", account.GetName())
	}
}
