package keeper

import (
	"github.com/govm-net/phoenix/x/phoenix/types"
)

var _ types.QueryServer = Keeper{}
