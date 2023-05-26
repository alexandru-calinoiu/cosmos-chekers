package types_test

import (
	"github.com/alice/checkers/testutil/sample"
	"github.com/alice/checkers/x/checkers/rules"
	types "github.com/alice/checkers/x/checkers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func getStoredGame(alice, bob string) *types.StoredGame {
	return &types.StoredGame{
		Index: "1",
		Board: rules.New().String(),
		Turn:  "b",
		Black: alice,
		Red:   bob,
	}
}

func TestCanGetBlackAddress(t *testing.T) {
	alice := sample.AccAddress()
	bob := sample.AccAddress()

	aliceAddress, err := sdk.AccAddressFromBech32(alice)
	require.NoError(t, err)
	blackAddress, err := getStoredGame(alice, bob).GetBlackAddress()
	require.NoError(t, err)
	require.Equal(t, aliceAddress, blackAddress)
}

func TestGetWrongBlackAddress(t *testing.T) {
	alice := sample.AccAddress()
	bob := sample.AccAddress()

	storedGame := getStoredGame(alice, bob)
	storedGame.Black = "42"
	blackAddress, err := storedGame.GetBlackAddress()
	require.Nil(t, blackAddress)
	require.EqualError(t,
		err,
		"black address is invalid: 42: decoding bech32 failed: invalid bech32 string length 2",
	)
	require.EqualError(t, storedGame.Validate(), err.Error())
}
