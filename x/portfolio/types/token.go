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

func createToken(name string, amount uint32) *Token {
	return &Token{
		Name:    name,
		Amounts: createAmountHistory(amount),
	}
}

func createTokens(names []string, amounts []uint32) []*Token {
	var tokens []*Token

	for i := range names {
		tokens[i] = createToken(names[i], amounts[i])
	}

	return tokens
}

func createTokenWithHistory(name string, amountHistory *AmountHistory) *Token {
	return &Token{
		Name:    name,
		Amounts: amountHistory,
	}
}
