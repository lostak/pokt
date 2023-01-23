package types

import "fmt"

func (d *AmountData) nestedPrint(indent, incr, symbol string, time int64) {
	fmt.Printf("%sTime: %d \n%sAmount:%f %s\n%sPrice:%f\n", indent, time, indent, d.GetAmount(), symbol, indent, d.GetPrice())
}

func (e *AmountEntry) nestedPrint(indent, incr, symbol string, num uint32) {

	var nextIndent string

	if len(indent) > 2 {
		nextIndent = fmt.Sprintf("%s %c ", indent, indent[1])
	} else {
		nextIndent = indent + incr
	}

	data := e.GetValue()
	if data == nil {
		return
	}

	nextIndent = indent + incr
	fmt.Printf("%sEntry #%d:\n", indent, num)

	data.nestedPrint(nextIndent, incr, symbol, e.GetKey())
}

func (t *Token) nestedPrint(indent, incr, id, symbol string) {
	nextIndent := indent + incr
	fmt.Printf("%sCoinGecko Id: %s\n", indent, id)
	var i uint32
	i = 0

	fmt.Printf("%sHistory:\n", nextIndent)
	for _, amount := range t.GetAmounts() {
		amount.nestedPrint(nextIndent, " - ", symbol, i)
		i++
	}
}

func (e *TokenEntry) nestedPrint(indent, incr string) {
	nextIndent := indent + incr
	fmt.Printf("%sToken: %s\n", indent, e.GetKey())

	token := e.GetValue()
	if token == nil {
		return
	}

	token.nestedPrint(nextIndent, incr, e.GetKey(), e.GetKey())
}

func (c *Chain) nestedPrint(indent, incr string) {
	indent += incr

	for _, token := range c.GetTokens() {
		token.nestedPrint(indent, incr)
	}
}

func (e *ChainEntry) nestedPrint(indent, incr string) {
	fmt.Printf("%sChain: %s\n", indent, e.GetKey())

	chain := e.GetValue()
	if chain == nil {
		return
	}

	chain.nestedPrint(indent, incr)
}

func (a *Account) nestedPrint(indent, incr string) {
	indent += incr

	for _, entry := range a.GetChains() {
		entry.nestedPrint(indent, incr)
	}
}

func (e *AccountEntry) nestedPrint(indent, incr string) {
	fmt.Printf("%sAccount: %s\n", indent, e.GetKey())
	account := e.GetValue()
	if account == nil {
		return
	}

	account.nestedPrint(indent, incr)
}

func (p *Portfolio) nestedPrint(indent, incr string) {
	fmt.Printf("Portfolio: %s\n", p.GetName())

	for _, entry := range p.GetAccounts() {
		entry.nestedPrint(indent, incr)
	}
}

func (p *Portfolio) Println() {
	p.nestedPrint(" | ", " | ")
}
