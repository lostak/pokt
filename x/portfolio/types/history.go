package types

/*
	TODO:
		- Add PriceHistory CRUD
		- Add StateHistory CRUD
*/

func createStateHistory(state States, amount uint32) []*StateHistory {
	var history []*StateHistory
	var amounts []uint32
	amounts[0] = amount

	// TODO: Add timestamp

	history[0] = &StateHistory{
		State:  state,
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
