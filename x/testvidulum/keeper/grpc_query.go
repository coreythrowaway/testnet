package keeper

import (
	"github.com/foo/testvidulum/x/testvidulum/types"
)

var _ types.QueryServer = Keeper{}
