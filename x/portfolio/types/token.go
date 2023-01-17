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

func (t *Token) nestedPrint() {
	fmt.Printf("\n\t\t\tToken:\n\t\t\t\tAmount: %d %s\n\t\t\t\tCoinGecko Id: %s\n", t.GetAmounts().Amount[len(t.GetAmounts().Amount)-1], t.GetName(), t.GetGeckoId())
}
