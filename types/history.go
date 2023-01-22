package types

import (
	"fmt"

	"google.golang.org/protobuf/types/known/timestamppb"
) /*
	TODO:
		- Add PriceHistory CRUD
		- Add StateHistory CRUD
*/

func createAmountHistory(amount float64) (h *AmountHistory) {
	amounts := make([]float64, 0)
	times := make([]*timestamppb.Timestamp, 0)

	amounts = append(amounts, amount)
	times = append(times, timestamppb.Now())

	return &AmountHistory{
		Amount:      amounts,
		UpdateTimes: times,
	}
}

func (h *AmountHistory) addAmount(amount float64) {
	amounts := h.GetAmount()
	h.Amount = append(amounts, amount)
	h.UpdateTimes = append(h.UpdateTimes, timestamppb.Now())
}

func (h *AmountHistory) nestedPrint(indent, incr, symbol string) {
	fmt.Printf("%sHistory:\n", indent)
	indent += incr

	for i := range h.GetAmount() {
		fmt.Printf("%sEntry #%d @ %s:  %f %s\n", indent, i, h.GetUpdateTimes()[i], h.GetAmount()[i], symbol)
	}
}

func (h *AmountHistory) deleteHistory() {
	amounts := make([]float64, 0)
	times := make([]*timestamppb.Timestamp, 0)

	h.Amount = append(amounts, 0)
	h.UpdateTimes = append(times, timestamppb.Now())
}
