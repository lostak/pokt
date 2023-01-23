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

func (t *Token) setAmount(key int64, value *AmountEntry) {
	if t.GetAmounts() == nil {
		t.Amounts = make(map[int64]*AmountEntry)
	}

	t.Amounts[key] = value
}

func (t *Token) deleteHistory() {
	for entry, _ := range t.GetAmounts() {
		delete(t.Amounts, entry)
	}
}
