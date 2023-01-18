package keeper

import (
	"fmt"

	"github.com/lostak/pokt/types"
)

func MsgCreatePortfolio(name string) {
	portfolio := types.CreateBlankPortfolio(name)

	err := SetPortfolio(portfolio)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("Portfolio created with name: %s\n", portfolio.Name)

	portfolio.Println()
}
