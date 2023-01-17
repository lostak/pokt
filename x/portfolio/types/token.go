package types

/*
	TODO:
		- Add Token CRUD
*/

func createBlankToken(name string) *Token {
	amount := createAmountHistory(0)

	return &Token{
		Name:    name,
		GeckoId: name,
		Amounts: amount,
	}
}
