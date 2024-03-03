package keeper

import (
	"add-evm/x/addevm/types"
)

var _ types.QueryServer = Keeper{}
