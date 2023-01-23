package types

import "fmt"

func createTokenEntry(name string) *TokenEntry {
	return &TokenEntry{
		Key:   name,
		Value: createBlankToken(name),
	}
}

func (e *TokenEntry) getAmountData(time int64) (*AmountData, error) {
	token := e.GetValue()
	if token == nil {
		return nil, fmt.Errorf("Token entry w/ key %s has a null value", e.GetKey())
	}

	amounts := token.GetAmounts()
	if amounts == nil {
		return nil, fmt.Errorf("Token entry w/ key %s's amount map has not been allocated")
	}

	return amounts[time].GetValue(), nil
}

func (e *TokenEntry) updateKey(name string) {
	e.Key = name
}

func (e *TokenEntry) updateValue(value *Token) {
	e.Value = value
}
