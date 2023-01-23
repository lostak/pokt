package types

/*
	TODO:
		- Add Token CRUD
*/

func createBlankToken(name string) *Token {
	/*
		TODO: add support for any base denom
	*/
	amount := make(map[int64]*AmountEntry)

	return &Token{
		GeckoId: name,
		Amounts: amount,
	}
}

func (t *Token) setAmount(key int64, value *AmountEntry) {
	t.Amounts[key] = value
}

func (t *Token) deleteHistory() {
	for entry, _ := range t.GetAmounts() {
		delete(t.Amounts, entry)
	}
}
