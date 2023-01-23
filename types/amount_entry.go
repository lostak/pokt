package types

import (
	"fmt"

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

func (e *AmountEntry) nestedPrint(indent, incr, symbol string, num uint32) {

	var nextIndent string

	if len(indent) > 2 {
		nextIndent = fmt.Sprintf("%s %s ", indent, indent[1])
	} else {
		nextIndent = indent + incr
	}

	data := e.GetValue()
	if data == nil {
		return
	}

	nextIndent = indent + incr
	fmt.Printf("%sEntry #%d:\n%sTime: %s \n%sAmount:%f %s\n%sPrice:%f\n", indent, num, nextIndent, e.GetKey(), nextIndent, data.GetAmount(), symbol, data.GetPrice())
}
