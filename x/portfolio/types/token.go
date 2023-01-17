package types

import "fmt"

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

func (t *Token) nestedPrint(indent string) {
	nextIndent := indent + "  "
	fmt.Printf("%sToken: %s\n%sCurrent Amount: %d %s\n%sCoinGecko Id: %s\n", indent, t.GetName(), nextIndent, t.GetAmounts().Amount[len(t.GetAmounts().Amount)-1], t.GetName(), nextIndent, t.GetGeckoId())
	t.GetAmounts().nestedPrint(nextIndent, t.GetName())
}
