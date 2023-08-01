package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jomluz/jomtx/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateTxn_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateTxn
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateTxn{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateTxn{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgUpdateTxn_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateTxn
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateTxn{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateTxn{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgDeleteTxn_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteTxn
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteTxn{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteTxn{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
