package keeper_test

import (
	"context"
	"fmt"
	"github.com/alice/checkers/testutil/sample"
	"github.com/alice/checkers/x/checkers/keeper"
	"github.com/alice/checkers/x/checkers/rules"
	"github.com/alice/checkers/x/checkers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGameNotFound(t *testing.T) {
	msgServer, _, ctx := SetupMessageServerCreateGame(t)
	_, err := msgServer.PlayMove(ctx, &types.MsgPlayMove{
		GameIndex: "1",
	})
	require.Equal(t, "1: game by id not found", err.Error())
}

func TestCreatorNotPlayer(t *testing.T) {
	msgServer, k, ctx := SetupMessageServerCreateGame(t)
	createStoredGame(k, ctx, defaultStoredGame())
	creatorAddress := sample.AccAddress()
	_, err := msgServer.PlayMove(ctx, &types.MsgPlayMove{
		GameIndex: "1",
		Creator:   creatorAddress,
	})
	require.Equal(t, fmt.Sprintf("%s: message creator is not a player", creatorAddress), err.Error())
}

func TestNotPlayerTurn(t *testing.T) {
	msgServer, k, ctx := SetupMessageServerCreateGame(t)
	creatorAddress := sample.AccAddress()
	cases := []struct {
		gameIndex string
		turn      string
		black     string
		red       string
	}{
		{
			gameIndex: "1",
			turn:      "r",
			red:       "",
			black:     creatorAddress,
		},
		{
			gameIndex: "2",
			turn:      "b",
			red:       creatorAddress,
			black:     "",
		},
	}

	for _, c := range cases {
		createStoredGame(k, ctx, types.StoredGame{
			Index: c.gameIndex,
			Turn:  c.turn,
			Black: c.black,
			Red:   c.red,
		})
		_, err := msgServer.PlayMove(ctx, &types.MsgPlayMove{
			GameIndex: c.gameIndex,
			Creator:   creatorAddress,
		})
		require.Equal(t, fmt.Sprintf("%s: player tried to play out of turn", creatorAddress), err.Error())
	}
}

func TestWrongMove(t *testing.T) {
	msgServer, k, ctx := SetupMessageServerCreateGame(t)
	creatorAddress := sample.AccAddress()
	createStoredGame(k, ctx, types.StoredGame{
		Index: "1",
		Turn:  "r",
		Red:   creatorAddress,
	})
	_, err := msgServer.PlayMove(ctx, &types.MsgPlayMove{
		GameIndex: "1",
		Creator:   creatorAddress,
		FromX:     1,
		FromY:     1,
		ToX:       1,
		ToY:       1,
	})
	require.Equal(t, "No piece at source position: {1 1}: wrong move", err.Error())
}

func TestWillStoreTheNewGame(t *testing.T) {
	msgServer, k, ctx := SetupMessageServerCreateGame(t)
	creatorAddress := sample.AccAddress()
	createStoredGame(k, ctx, types.StoredGame{
		Index: "1",
		Turn:  "b",
		Black: creatorAddress,
	})
	response, err := msgServer.PlayMove(ctx, &types.MsgPlayMove{
		GameIndex: "1",
		Creator:   creatorAddress,
		FromX:     1,
		FromY:     2,
		ToX:       2,
		ToY:       3,
	})
	require.NoError(t, err)
	storedGame, _ := k.GetStoredGame(sdk.UnwrapSDKContext(ctx), "1")
	require.Equal(t, "r", storedGame.Turn)
	require.Equal(t, "*b*b*b*b|b*b*b*b*|***b*b*b|**b*****|********|r*r*r*r*|*r*r*r*r|r*r*r*r*", storedGame.Board)
	require.EqualValues(t, types.MsgPlayMoveResponse{
		CapturedX: -1,
		CapturedY: -1,
		Winner:    "*",
	}, *response)
}

func createStoredGame(k keeper.Keeper, ctx context.Context, storedGame types.StoredGame) {
	storedGame.Board = rules.New().String()
	k.SetStoredGame(sdk.UnwrapSDKContext(ctx), storedGame)
}

func defaultStoredGame() types.StoredGame {
	return types.StoredGame{
		Index: "1",
		Board: "",
		Turn:  "",
		Black: sample.AccAddress(),
		Red:   sample.AccAddress(),
	}
}
