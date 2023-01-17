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

func (h *AmountHistory) nestedPrint(indent, symbol string) {
	fmt.Printf("%sHistory:\n", indent)
	indent += "  "

	for i := range h.GetAmount() {
		fmt.Printf("%s%s:\t%d %s\n", indent, h.GetUpdateTimes()[i], h.GetAmount()[i], symbol)
	}
}
