package types

import (
	"fmt"
)

/*
	TODO: Add Portfolio CRUD
*/

func CreateBlankPortfolio(name string) *Portfolio {
	accounts := make(map[string]*AccountEntry)

	return &Portfolio{
		Name:     name,
		Accounts: accounts,
	}
}

func (p *Portfolio) GetAccountEntry(accountName string) (*AccountEntry, error) {
	account := p.GetAccounts()[accountName]
	if account == nil {
		return nil, fmt.Errorf("Account %s not found\n", accountName)
	}

	return account, nil
}

func (p *Portfolio) GetAccount(name string) (*Account, error) {
	accounts := p.GetAccounts()
	if accounts == nil {
		return nil, fmt.Errorf("Portfolio w/ name: %s's account is not allocated", p.GetName())
	}

	account := accounts[name]
	if account == nil {
		return nil, fmt.Errorf("Account w/ key: %s's value is null", name)
	}

	return account.GetValue(), nil
}

func (p *Portfolio) setAccount(name string, account *AccountEntry) {
	p.Accounts[name] = account
}

func (p *Portfolio) GetChainEntry(accountName, chainName string) (*ChainEntry, error) {
	entry, err := p.GetAccountEntry(accountName)
	if err != nil {
		return nil, err
	}

	chain, err := entry.getChainEntry(chainName)
	if err != nil {
		return nil, err
	}

	return chain, nil
}

func (p *Portfolio) GetTokenEntry(accountName, chainName, tokenName string) (*TokenEntry, error) {
	entry, err := p.GetChainEntry(accountName, chainName)
	if err != nil {
		return nil, err
	}

	token, err := entry.getTokenEntry(tokenName)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (p *Portfolio) GetAmountData(accountName, chainName, tokenName string, time int64) (*AmountData, error) {
	token, err := p.GetTokenEntry(accountName, chainName, tokenName)
	if err != nil {
		return nil, err
	}

	entry, err := token.getAmountData(time)
	if err != nil {
		return nil, err
	}

	return entry, nil
}

func (p *Portfolio) AddAccount(accountName string) error {
	if _, err := p.GetAccount(accountName); err == nil {
		return fmt.Errorf("Account already exits with name: %s on portfolio: %s", accountName, p.GetName())
	}

	account := createAccountEntry(accountName)
	p.setAccount(accountName, account)

	fmt.Println("Account added")
	return nil
}

func (p *Portfolio) RemoveAccount(accountName string) error {
	_, err := p.GetAccount(accountName)
	if err != nil {
		return err
	}

	delete(p.Accounts, accountName)

	return nil
}

func (p *Portfolio) UpdateAccountName(accountName, newName string) error {
	account, err := p.GetAccountEntry(accountName)
	if err != nil {
		return err
	}

	oldKey := account.GetKey()
	account.updateKey(newName)
	p.setAccount(newName, account)
	delete(p.Accounts, oldKey)

	return nil
}

func (p *Portfolio) AddChain(accountName, chainName, address string) error {
	account, err := p.GetAccount(accountName)
	if err != nil {
		return err
	}

	return account.addChain(chainName, address)
}

func (p *Portfolio) RemoveChain(accountName, chainName string) error {
	account, err := p.GetAccount(accountName)
	if err != nil {
		return err
	}

	return account.removeChain(chainName)
}

func (p *Portfolio) AddToken(accountName, chainName, tokenName string) error {
	account, err := p.GetAccount(accountName)
	if err != nil {
		return err
	}

	return account.addToken(chainName, tokenName)
}

func (p *Portfolio) RemoveToken(accountName, chainName, tokenName string) error {
	account, err := p.GetAccount(accountName)
	if err != nil {
		return err
	}

	return account.removeToken(chainName, tokenName)
}

func (p *Portfolio) UpdateTokenGeckoId(accountName, chainName, tokenName, geckoId string) error {
	account, err := p.GetAccount(accountName)
	if err != nil {
		return err
	}

	return account.updateTokenGeckoId(chainName, tokenName, geckoId)
}

func (p *Portfolio) AddTokenAmount(accountName, chainName, tokenName string, amount float64) error {
	account, err := p.GetAccount(accountName)
	if err != nil {
		return err
	}

	return account.addTokenAmount(chainName, tokenName, amount)
}

func (p *Portfolio) ClearTokenHistory(accountName, chainName, tokenName string) error {
	account, err := p.GetAccount(accountName)
	if err != nil {
		return err
	}

	return account.clearTokenHistory(chainName, tokenName)
}

func (p *Portfolio) ClearChainHistory(accountName, chainName string) error {
	account, err := p.GetAccount(accountName)
	if err != nil {
		return err
	}

	return account.clearChainHistory(chainName)
}

func (p *Portfolio) ClearAccountHistory(accountName string) error {
	account, err := p.GetAccount(accountName)
	if err != nil {
		return err
	}

	account.deleteHistory()
	return nil
}

func (p *Portfolio) deleteHistory() {
	for _, entry := range p.GetAccounts() {
		account := entry.GetValue()
		if account == nil {
			continue
		}

		account.deleteHistory()
	}
}

func (p *Portfolio) ClearHistory() {
	p.deleteHistory()
}

func (p *Portfolio) nestedPrint(indent, incr string) {
	fmt.Printf("Portfolio: %s\n", p.GetName())

	for _, entry := range p.GetAccounts() {
		account := entry.GetValue()
		if account == nil {
			continue
		}

		account.nestedPrint(indent, incr)
	}
}

func (p *Portfolio) Println() {
	p.nestedPrint(" | ", " | ")
}
