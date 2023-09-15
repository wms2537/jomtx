package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/jomtx module sentinel errors
var (
	ErrInvalidCreator = sdkerrors.Register(ModuleName, 1100, "creator address is invalid: %s")
)
