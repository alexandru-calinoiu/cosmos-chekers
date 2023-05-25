package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "github.com/alice/checkers/testutil/keeper"
	"github.com/alice/checkers/testutil/nullify"
	"github.com/alice/checkers/x/checkers/keeper"
	"github.com/alice/checkers/x/checkers/types"
)

func createTestSystemInfo(keeper *keeper.Keeper, ctx sdk.Context) types.SystemInfo {
	item := types.SystemInfo{NextId: 42}
	keeper.SetSystemInfo(ctx, item)
	return item
}

func TestSystemInfoGet(t *testing.T) {
	kp, ctx := keepertest.CheckersKeeper(t)
	item := createTestSystemInfo(kp, ctx)
	rst, found := kp.GetSystemInfo(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
	require.Equal(t, uint64(42), rst.GetNextId())
}

func TestSystemInfoRemove(t *testing.T) {
	kp, ctx := keepertest.CheckersKeeper(t)
	createTestSystemInfo(kp, ctx)
	kp.RemoveSystemInfo(ctx)
	_, found := kp.GetSystemInfo(ctx)
	require.False(t, found)
}
