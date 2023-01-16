package types

import (
	"google.golang.org/protobuf/types/known/timestamppb"
) /*
	TODO:
		- Add PriceHistory CRUD
		- Add StateHistory CRUD
*/

func createAmountHistory(amount uint32) (h *AmountHistory) {
	return &AmountHistory{
		Amount:      append(h.Amount, amount),
		UpdateTimes: append(h.UpdateTimes, timestamppb.Now()),
	}
}

func (h *AmountHistory) addAmount(amount uint32) {
	h.Amount = append(h.Amount, amount)
	h.UpdateTimes = append(h.UpdateTimes, timestamppb.Now())
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
