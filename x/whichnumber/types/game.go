package types

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func GetPlayerAddress(address string) (sdk.AccAddress, error) {
	player, errPl := sdk.AccAddressFromBech32(address)
	return player, sdkerrors.Wrapf(errPl, sdkerrors.ErrInvalidAddress.Error(), address)
}

func CalculateNumberCommit(number int64, salt string) string {
	h := sha256.Sum256([]byte(fmt.Sprintf("%d:%s", number, salt)))
	return hex.EncodeToString(h[:])
}

func StoredGameToQueryGameResponse(storedGame Game) QueryGameResponse {
	var revealTimeout *time.Time
	if !storedGame.RevealTimeout.IsZero() {
		revealTimeout = &storedGame.RevealTimeout
	}

	playerReveals := make([]*QueryPlayerReveal, len(storedGame.PlayerReveals))
	for i, player := range storedGame.PlayerReveals {
		playerReveals[i] = &QueryPlayerReveal{
			PlayerAddress: player.PlayerAddress,
			CreatedAt:     player.CreatedAt,
			Proximity:     player.Proximity,
		}
	}

	return QueryGameResponse{
		Id:            storedGame.Id,
		Creator:       storedGame.Creator,
		PlayerCommits: storedGame.PlayerCommits,
		PlayerReveals: playerReveals,
		Reward:        storedGame.Reward.String(),
		EntryFee:      storedGame.EntryFee.String(),
		CommitTimeout: storedGame.CommitTimeout,
		RevealTimeout: revealTimeout,
		Status:        storedGame.Status,
		Winners:       storedGame.Winners,
	}
}

func (g *Game) Validate() error {
	if g.EntryFee.IsZero() {
		return fmt.Errorf("entry fee cannot be zero")
	}

	if g.Reward.IsZero() {
		return fmt.Errorf("reward cannot be zero")
	}

	if g.CommitTimeout.Second() <= 0 {
		return fmt.Errorf("commit timeout must be positive")
	}

	return nil
}
