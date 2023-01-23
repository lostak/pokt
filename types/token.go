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

func (t *Token) nestedPrint(indent, incr, symbol string) {
	nextIndent := indent + incr
	fmt.Printf("%sCurrent Amount: %d %s\n%sCoinGecko Id: %s\n", indent)
	var i uint32
	i = 0

	fmt.Printf("%sHistory:\n", nextIndent)
	for _, amount := range t.GetAmounts() {
		amount.nestedPrint(nextIndent, " - ", symbol, i)
		i++
	}
}
