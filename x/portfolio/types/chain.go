package types

/*
	TODO:
		- Add Chain CRUD
*/

func CreateBlankChain(chainName, address string) *Chain {
	return &Chain{
		Name: chainName,
		Addr: address,
	}
}

func createChainWithToken(chainName string, address string, tokenName string, state States, amount uint32) *Chain {
	var tokens []*Token
	tokens[0] = createToken(tokenName, state, amount)

	// TODO: Add timestamp
	return &Chain{
		Name:   chainName,
		Addr:   address,
		Tokens: tokens,
	}
}
