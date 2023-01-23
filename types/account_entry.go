package types

import "fmt"

func createAccountEntry(key string) *AccountEntry {
	return &AccountEntry{
		Key:   key,
		Value: createAccount(),
	}
}

func (e *AccountEntry) getChainEntry(name string) (*ChainEntry, error) {
	account := e.GetValue()
	if account == nil {
		return nil, fmt.Errorf("Account entry w/ key: %s has a null value", e.GetKey())
	}

	chains := account.GetChains()
	if chains == nil {
		return nil, fmt.Errorf("Account entry w/ key: %s's chain map has not been allocated", e.GetKey())
	}

	return chains[name], nil
}

func (e *AccountEntry) addChain(name, address string) error {
	c := e.GetValue()
	if c == nil {
		e.updateValue(createAccount())
	}

	return e.GetValue().addChain(name, address)
}

func (e *AccountEntry) updateKey(name string) {
	e.Key = name
}

func (e *AccountEntry) updateValue(value *Account) {
	e.Value = value
}
