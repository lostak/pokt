package types

import "fmt"

/*
	TODO:
		- Add Token CRUD
*/

func createBlankToken(name string) *Token {
	/*
		TODO: add support for any base denom
	*/
	amount := createAmountHistory(0, name, "usd")

	return &Token{
		Name:    name,
		GeckoId: name,
		Amounts: amount,
	}
}

func (t *Token) updateName(name string) {
	t.Name = name
}

func (t *Token) nestedPrint(indent, incr string) {
	nextIndent := indent + incr
	fmt.Printf("%sToken: %s\n%sCurrent Amount: %d %s\n%sCoinGecko Id: %s\n", indent, t.GetName(), nextIndent, t.GetAmounts().Amount[len(t.GetAmounts().Amount)-1], t.GetName(), nextIndent, t.GetGeckoId())
	t.GetAmounts().nestedPrint(nextIndent, " - ", t.GetName())
}
