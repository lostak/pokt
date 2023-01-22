package types

import (
	"fmt"

	"github.com/lostak/pokt/client/gecko"
	"google.golang.org/protobuf/types/known/timestamppb"
) /*
	TODO:
		- Add PriceHistory CRUD
		- Add StateHistory CRUD
*/

func createAmountHistory(amount float64, coinId, baseDenom string) (h *AmountHistory) {
	amounts := make([]float64, 0)
	times := make([]*timestamppb.Timestamp, 0)
	prices := make([]float64, 0)

	price, err := gecko.GetTokenPrice(coinId, baseDenom)
	if err != nil {
		price = 0
	}

	amounts = append(amounts, amount)
	times = append(times, timestamppb.Now())
	prices = append(prices, price)

	return &AmountHistory{
		Amount:      amounts,
		UpdateTimes: times,
		Price:       prices,
	}
}

func (h *AmountHistory) addAmount(amount float64, coinId, baseDenom string) error {
	amounts := h.GetAmount()
	price, err := gecko.GetTokenPrice(coinId, baseDenom)
	if err != nil {
		return err
	}

	h.Amount = append(amounts, amount)
	h.UpdateTimes = append(h.UpdateTimes, timestamppb.Now())
	h.Price = append(h.GetPrice(), price)
	return nil
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
