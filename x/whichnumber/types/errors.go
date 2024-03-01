package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/whichnumber module sentinel errors

var (
	ErrPlayerNotCommitted          = sdkerrors.Register(ModuleName, 1100, "player has not committed")
	ErrGameNotFound                = sdkerrors.Register(ModuleName, 1101, "game by id not found")
	ErrCreatorNotPlayer            = sdkerrors.Register(ModuleName, 1102, "game creator is not a player")
	ErrInvalidReveal               = sdkerrors.Register(ModuleName, 1103, "number doesn't match commitment, trying to cheat?")
	ErrCannotGetSystemInfo         = sdkerrors.Register(ModuleName, 1104, "cannot get system info")
	ErrInvalidGame                 = sdkerrors.Register(ModuleName, 1105, "game is invalid")
	ErrFailedToSendCoinsToModule   = sdkerrors.Register(ModuleName, 1106, "failed to send coins to module")
	ErrInvalidReward               = sdkerrors.Register(ModuleName, 1107, "invalid reward")
	ErrGameFull                    = sdkerrors.Register(ModuleName, 1108, "game is full")
	ErrCommitTimeout               = sdkerrors.Register(ModuleName, 1109, "commit timeout has passed")
	ErrCommitNotTimeout            = sdkerrors.Register(ModuleName, 1110, "commit timeout has not passed yet")
	ErrRevealTimeout               = sdkerrors.Register(ModuleName, 1111, "reveal timeout has passed")
	ErrPlayerAlreadyCommitted      = sdkerrors.Register(ModuleName, 1112, "player has already committed")
	ErrPlayerAlreadyRevealed       = sdkerrors.Register(ModuleName, 1113, "player has already revealed")
	ErrInvalidCommitTimeout        = sdkerrors.Register(ModuleName, 1114, "invalid commit timeout")
	ErrInvalidRevealTimeout        = sdkerrors.Register(ModuleName, 1115, "invalid reveal timeout")
	ErrInvalidMaxPlayers           = sdkerrors.Register(ModuleName, 1116, "invalid max players")
	ErrInvalidMinDistanceToWin     = sdkerrors.Register(ModuleName, 1117, "invalid min distance to win")
	ErrGameNotCommitting           = sdkerrors.Register(ModuleName, 1118, "game is not committing")
	ErrGameFinished                = sdkerrors.Register(ModuleName, 1119, "game is already finished")
	ErrGameFailedToDepositReward   = sdkerrors.Register(ModuleName, 1120, "failed to deposit reward")
	ErrGameFailedToDepositEntryFee = sdkerrors.Register(ModuleName, 1121, "failed to deposit entry fee")
)
