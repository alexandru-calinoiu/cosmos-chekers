package types

import (
	"testing"

	"github.com/alice/checkers/testutil/sample"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateGame_ValidateBasic(t *testing.T) {
	alice := sample.AccAddress()
	bob := sample.AccAddress()
	carol := sample.AccAddress()

	tests := []struct {
		name string
		msg  MsgCreateGame
		err  error
	}{
		{
			name: "invalid creator address",
			msg: MsgCreateGame{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "invalid black address",
			msg: MsgCreateGame{
				Creator: alice,
				Black:   "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "invalid red address",
			msg: MsgCreateGame{
				Creator: alice,
				Black:   bob,
				Red:     "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateGame{
				Creator: alice,
				Black:   bob,
				Red:     carol,
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
