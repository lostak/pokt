package types

import (
	"github.com/lostak/pokt/client/gecko"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func createAmountData(amount float64, coinId, baseDenom string) *AmountData {
	price, err := gecko.GetTokenPrice(coinId, baseDenom)
	if err != nil {
		price = 0
	}

	return &AmountData{
		Amount: amount,
		Price:  price,
	}
}

func createAmountEntry(amount float64, coinId, baseDenom string) *AmountEntry {
	return &AmountEntry{
		Key:   timestamppb.Now().Seconds,
		Value: createAmountData(amount, coinId, baseDenom),
	}
}

func (e *AmountEntry) updateKey(seconds int64) {
	e.Key = seconds
}

func (e *AmountEntry) updateValue(value *AmountData) {
	e.Value = value
}
