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

	if t.GetBaseDenom() == "" {
		t.BaseDenom = "usd"
	}
	// TODO: add baseDenom support
	a := createAmountEntry(amount, t.GetGeckoId(), t.GetBaseDenom())

	t.Amounts[a.GetKey()] = a
}
