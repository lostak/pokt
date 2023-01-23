package types

func (a *Account) updateAddress(chainName, address string) error {
	chain, err := a.getChain(chainName)
	if err != nil {
		return err
	}

	chain.updateAddress(address)
	return nil
}

func (c *Chain) updateAddress(address string) {
	c.Addr = address
}
