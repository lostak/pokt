package types

import "fmt"

func createChainEntry(name string, value *Chain) *ChainEntry {
	return &ChainEntry{
		Key:   name,
		Value: createBlankChain("temp"),
	}
}

func (e *ChainEntry) getTokenEntry(name string) (*TokenEntry, error) {
	chain := e.GetValue()
	if chain == nil {
		return nil, fmt.Errorf("Chain Entry w/ key %s has a null value", e.GetKey())
	}

	tokens := chain.GetTokens()
	if tokens == nil {
		return nil, fmt.Errorf("Chain entry w/ key %s's token map has not been allocated")
	}

	return tokens[name], nil
}

func (e *ChainEntry) updateKey(name string) {
	e.Key = name
}

func (e *ChainEntry) updateValue(value *Chain) {
	e.Value = value
}
