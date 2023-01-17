package types

import (
	"google.golang.org/protobuf/types/known/timestamppb"
) /*
	TODO:
		- Add PriceHistory CRUD
		- Add StateHistory CRUD
*/

func createAmountHistory(amount uint32) (h *AmountHistory) {
	amounts := make([]uint32, 0)
	times := make([]*timestamppb.Timestamp, 0)

	amounts = append(amounts, amount)
	times = append(times, timestamppb.Now())

	return &AmountHistory{
		Amount:      amounts,
		UpdateTimes: times,
	}
}

func (h *AmountHistory) addAmount(amount uint32) {
	amounts := h.GetAmount()
	h.Amount = append(amounts, amount)
	h.UpdateTimes = append(h.UpdateTimes, timestamppb.Now())
}
