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

	return account.addChain(chainName, address)
}

func (p *Portfolio) RemoveChain(accountName, chainName string) error {
	err, account := p.GetAccount(accountName)
	if err != nil {
		return err
	}

	return account.removeChain(chainName)
}

func (p *Portfolio) AddToken(accountName, chainName, tokenName string) error {
	err, account := p.GetAccount(accountName)
	if err != nil {
		return err
	}

	return account.addToken(chainName, tokenName)
}

func (p *Portfolio) RemoveToken(accountName, chainName, tokenName string) error {
	err, account := p.GetAccount(accountName)
	if err != nil {
		return err
	}

	return account.removeToken(chainName, tokenName)
}

func (p *Portfolio) UpdateTokenGeckoId(accountName, chainName, tokenName, geckoId string) error {
	err, account := p.GetAccount(accountName)
	if err != nil {
		return err
	}

	return account.updateTokenGeckoId(chainName, tokenName, geckoId)
}

func (p *Portfolio) AddTokenAmount(accountName, chainName, tokenName string, amount uint32) error {
	err, account := p.GetAccount(accountName)
	if err != nil {
		return err
	}

	return account.addTokenAmount(chainName, tokenName, amount)
}

func (p *Portfolio) ClearTokenHistory(accountName, chainName, tokenName string) error {
	err, account := p.GetAccount(accountName)
	if err != nil {
		return err
	}

	return account.clearTokenHistory(chainName, tokenName)
}

func (p *Portfolio) ClearChainHistory(accountName, chainName string) error {
	err, account := p.GetAccount(accountName)
	if err != nil {
		return err
	}

	return account.clearChainHistory(chainName)
}

func (p *Portfolio) nestedPrint(indent string) {
	fmt.Printf("Portfolio:\n%s%s\n", indent, p.GetName())
	for _, account := range p.GetAccounts() {
		account.nestedPrint(indent + "  ")
	}
}

func (p *Portfolio) Println() {
	p.nestedPrint("  ")
	fmt.Print("\n")
}
