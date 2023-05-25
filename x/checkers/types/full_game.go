package types

import (
	"fmt"
	"github.com/alice/checkers/x/checkers/rules"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (storedGame *StoredGame) GetBlackAddress() (sdk.AccAddress, error) {
	black, err := sdk.AccAddressFromBech32(storedGame.Black)
	return black, sdkerrors.Wrapf(err, ErrInvalidBlack.Error(), storedGame.Black)
}

func (storedGame *StoredGame) GetRedAddress() (sdk.AccAddress, error) {
	red, err := sdk.AccAddressFromBech32(storedGame.Red)
	return red, sdkerrors.Wrapf(err, ErrInvalidRed.Error(), storedGame.Red)
}

func (storedGame *StoredGame) ParseGame() (*rules.Game, error) {
	board, err := rules.Parse(storedGame.Board)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, ErrGameNotParsable.Error())
	}

	board.Turn = rules.StringPieces[storedGame.Turn].Player
	if board.Turn.Color == "" {
		return nil, sdkerrors.Wrapf(fmt.Errorf("turn: %s", storedGame.Turn), ErrGameNotParsable.Error())
	}

	return board, nil
}

func (storedGame *StoredGame) Validate() error {
	_, err := storedGame.GetBlackAddress()
	if err != nil {
		return err
	}

	_, err = storedGame.GetRedAddress()
	if err != nil {
		return err
	}

	_, err = storedGame.ParseGame()
	if err != nil {
		return err
	}

	return nil
}
