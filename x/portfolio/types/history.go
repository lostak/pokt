package types

/*
	TODO:
		- Add PriceHistory CRUD
		- Add StateHistory CRUD
*/

func createAmountHistory(amount uint32) []*AmountHistory {
	var history []*AmountHistory
	var amounts []uint32
	amounts[0] = amount

	// TODO: Add timestamp

	history[0] = &AmountHistory{
		Amount: amounts,
	}

	return history
}

func createPriceHistory(name, geckoId, baseDenomId string, firstPrice float64) *PriceHistory {
	var prices []*SpotPrice
	prices[0] = createSpotPrice(firstPrice)

	return &PriceHistory{
		TokenName:   name,
		GeckoId:     geckoId,
		BaseDenomId: baseDenomId,
		Prices:      prices,
	}
}
