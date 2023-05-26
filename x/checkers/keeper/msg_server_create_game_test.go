package keeper_test

import (
	"github.com/alice/checkers/x/checkers/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGameCreate(t *testing.T) {
	msgServer, ctx := setupMsgServer(t)
	createResponse, err := msgServer.CreateGame(ctx, &types.MsgCreateGame{
		Creator: alice,
		Black:   bob,
		Red:     carol,
	})
	require.NoError(t, err)
	require.EqualValues(t, &types.MsgCreateGameResponse{
		GameIndex: "1",
	}, createResponse)
}
