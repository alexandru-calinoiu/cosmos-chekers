package types

import (
	"testing"

	"github.com/alice/checkers/testutil/sample"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgPlayMove_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgPlayMove
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgPlayMove{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "invalid game index",
			msg: MsgPlayMove{
				Creator:   sample.AccAddress(),
				GameIndex: "-1",
			},
			err: ErrInvalidGameIndex,
		},
		{
			name: "invalid from",
			msg: MsgPlayMove{
				Creator:   sample.AccAddress(),
				GameIndex: "1",
				FromX:     uint64(42),
				FromY:     uint64(42),
			},
			err: ErrInvalidPositionIndex,
		}, {
			name: "no move",
			msg: MsgPlayMove{
				Creator:   sample.AccAddress(),
				GameIndex: "1",
				FromX:     1,
				FromY:     1,
				ToX:       1,
				ToY:       1,
			},
			err: ErrMoveAbsent,
		}, {
			name: "valid address",
			msg: MsgPlayMove{
				Creator:   sample.AccAddress(),
				GameIndex: "1",
				FromX:     1,
				FromY:     1,
				ToX:       2,
				ToY:       2,
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
