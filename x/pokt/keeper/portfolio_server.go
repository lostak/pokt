package keeper

import (
	"github.com/lostak/pokt/types"
)

func NewPortfolio(name string) types.Portfolio {
	portfolio := &types.Portfolio{
		Name: name,
	}

	return *portfolio
}
