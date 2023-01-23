package types

/*
	TODO:
		- Add set base denom
*/

func createBlankToken(name string) *Token {
	return &Token{
		GeckoId:   name,
		BaseDenom: "usd",
		Amounts:   make(map[int64]*AmountEntry),
	}
}

func (t *Token) setAmount(amount float64) {
	if t.GetAmounts() == nil {
		t.Amounts = make(map[int64]*AmountEntry)
	}

	a := createAmountEntry(amount, t.GetGeckoId(), t.GetGeckoId())

	t.Amounts[a.GetKey()] = a
}

func (t *Token) deleteHistory() {
	for entry, _ := range t.GetAmounts() {
		delete(t.Amounts, entry)
	}
}
