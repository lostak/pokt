package types

/*
	TODO:
		- Add Account CRUD
*/

func createBlankAccount(accountName string) *Account {
	return &Account{
		Name: accountName,
	}
}

func createAccountWithChain(accountName, address, chainName, tokenName string, state States, amount uint32) *Account {
	var chains []*Chain
	chains[0] = createChainWithToken(chainName, address, tokenName, state, amount)

	return &Account{
		Name:   accountName,
		Chains: chains,
	}
}
