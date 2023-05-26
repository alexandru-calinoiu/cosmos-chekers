package keeper_test

import (
	"context"
	keepertest "github.com/alice/checkers/testutil/keeper"
	"github.com/alice/checkers/testutil/sample"
	"github.com/alice/checkers/x/checkers"
	"github.com/alice/checkers/x/checkers/keeper"
	"github.com/alice/checkers/x/checkers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGameCreate(t *testing.T) {
	msgServer, _, ctx := setupMessageServerCreateGame(t)
	createResponse, err := msgServer.CreateGame(ctx, &types.MsgCreateGame{
		Creator: sample.AccAddress(),
		Black:   sample.AccAddress(),
		Red:     sample.AccAddress(),
	})
	require.NoError(t, err)
	require.EqualValues(t, &types.MsgCreateGameResponse{
		GameIndex: "1",
	}, createResponse)
}

func TestGameCreateGameWasSaved(t *testing.T) {
	black := sample.AccAddress()
	red := sample.AccAddress()
	msgServer, k, ctx := setupMessageServerCreateGame(t)
	_, err := msgServer.CreateGame(ctx, &types.MsgCreateGame{
		Creator: sample.AccAddress(),
		Black:   black,
		Red:     red,
	})
	require.NoError(t, err)

	systemInfo, found := k.GetSystemInfo(sdk.UnwrapSDKContext(ctx))
	require.True(t, found)
	require.Equal(t, uint64(2), systemInfo.GetNextId())

	storedGame, found := k.GetStoredGame(sdk.UnwrapSDKContext(ctx), "1")
	require.True(t, found)
	require.EqualValues(t, types.StoredGame{
		Index: "1",
		Board: "*b*b*b*b|b*b*b*b*|*b*b*b*b|********|********|r*r*r*r*|*r*r*r*r|r*r*r*r*",
		Turn:  "b",
		Black: black,
		Red:   red,
	}, storedGame)
}

func setupMessageServerCreateGame(t *testing.T) (types.MsgServer, keeper.Keeper, context.Context) {
	k, ctx := keepertest.CheckersKeeper(t)
	checkers.InitGenesis(ctx, *k, *types.DefaultGenesis())
	return keeper.NewMsgServerImpl(*k), *k, sdk.WrapSDKContext(ctx)
}
