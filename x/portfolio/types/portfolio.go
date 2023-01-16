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

	fmt.Println("Account added")
	return nil
}

func (p *Portfolio) RemoveAccount(accountName string) error {
	var accounts []*Account
	found := false

	// Iterate through accounts and omit all accounts w/ same name
	for _, account := range p.GetAccounts() {
		if account.GetName() == accountName {
			found = true
			fmt.Println("Account removed")
		} else {
			accounts = append(accounts, account)
		}
	}

	if !found {
		return fmt.Errorf("Account: %s not found", accountName)
	}

	// update portfolio
	p.Accounts = accounts
	return nil
}

func (p *Portfolio) GetAccount(accountName string) (error, *Account) {
	for _, account := range p.GetAccounts() {
		if account.GetName() == accountName {
			return nil, account
		}
	}

	return fmt.Errorf("Account: %s not found\n", accountName), nil
}

func (p *Portfolio) AddChain(accountName, chainName, address string) error {
	err, account := p.GetAccount(accountName)
	if err != nil {
		return err
	}

	return account.AddChain(chainName, address)
}

func (p *Portfolio) RemoveChain(accountName, chainName string) error {
	err, account := p.GetAccount(accountName)
	if err != nil {
		return err
	}

	err = account.RemoveChain(chainName)
	if err != nil {
		return err
	}

	return nil
}

func (p *Portfolio) AddToken(accountName, chainName, tokenName string) error {
	err, account := p.GetAccount(accountName)
	if err != nil {
		return err
	}

	return account.AddToken(chainName, tokenName)
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

func (p *Portfolio) PrintTokens() {
	fmt.Printf("Portfolio:\n\t%s\n", p.GetName())
	for _, account := range p.GetAccounts() {
		fmt.Printf("\tAccount:\n\t\t%s\n", account.GetName())
		for _, chain := range account.GetChains() {
			fmt.Printf("\t\tChain:\n\t\t\t%s\n", chain.GetName())
			for _, token := range chain.GetTokens() {
				fmt.Printf("\t\t\tToken:\n\t\t\t\t%s\n", token.GetName())
			}
		}
	}
}
