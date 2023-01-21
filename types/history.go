package types

import (
	"fmt"

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

func (h *AmountHistory) nestedPrint(indent, incr, symbol string) {
	fmt.Printf("%sHistory:\n", indent)
	indent += incr

	for i := range h.GetAmount() {
		fmt.Printf("%sEntry #%d @ %s:  %d %s\n", indent, i, h.GetUpdateTimes()[i], h.GetAmount()[i], symbol)
	}
}

func (h *AmountHistory) deleteHistory() {
	amounts := make([]uint32, 0)
	times := make([]*timestamppb.Timestamp, 0)

	h.Amount = append(amounts, 0)
	h.UpdateTimes = append(times, timestamppb.Now())
}
