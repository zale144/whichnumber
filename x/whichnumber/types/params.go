package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var _ paramtypes.ParamSet = (*Params)(nil)

const (
	// DefaultCommitTimeoutSec is the default commit timeout in seconds
	DefaultCommitTimeoutSec uint64 = 60
	// DefaultRevealTimeoutSec is the default reveal timeout in seconds
	DefaultRevealTimeoutSec uint64 = 60
	// DefaultMaxPlayersPerGame is the default max players per game
	DefaultMaxPlayersPerGame uint64 = 10
	// DefaultMinDistanceToWin is the default min distance to win
	DefaultMinDistanceToWin uint64 = 10
	// DefaultMinReward is the default min reward
	DefaultMinReward = "1000stake"
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(commitTimeout, revealTimeout, maxPlayers, minDistanceToWin uint64, minReward sdk.Coin) Params {
	return Params{
		CommitTimeout:     commitTimeout,
		RevealTimeout:     revealTimeout,
		MaxPlayersPerGame: maxPlayers,
		MinDistanceToWin:  minDistanceToWin,
		MinReward:         minReward,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	defaultMinReward, err := sdk.ParseCoinNormalized(DefaultMinReward)
	if err != nil {
		panic(err)
	}
	return NewParams(
		DefaultCommitTimeoutSec,
		DefaultRevealTimeoutSec,
		DefaultMaxPlayersPerGame,
		DefaultMinDistanceToWin,
		defaultMinReward,
	)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair([]byte(ParamKeyCommitTimeout), &p.CommitTimeout, validateCommitTimeout),
		paramtypes.NewParamSetPair([]byte(ParamKeyRevealTimeout), &p.RevealTimeout, validateRevealTimeout),
		paramtypes.NewParamSetPair([]byte(ParamKeyMaxPlayersPerGame), &p.MaxPlayersPerGame, validateMaxPlayersPerGame),
		paramtypes.NewParamSetPair([]byte(ParamKeyMinDistanceToWin), &p.MinDistanceToWin, validateMinDistanceToWin),
		paramtypes.NewParamSetPair([]byte(ParamKeyMinReward), &p.MinReward, validateMinReward),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateCommitTimeout(p.CommitTimeout); err != nil {
		return err
	}

	if err := validateRevealTimeout(p.RevealTimeout); err != nil {
		return err
	}

	if err := validateMaxPlayersPerGame(p.MaxPlayersPerGame); err != nil {
		return err
	}

	return validateMinDistanceToWin(p.MinDistanceToWin)
}

func validateCommitTimeout(commitTimeout any) error {
	if commitTimeout.(uint64) == 0 {
		return ErrInvalidCommitTimeout
	}
	return nil
}

func validateRevealTimeout(revealTimeout any) error {
	if revealTimeout.(uint64) == 0 {
		return ErrInvalidRevealTimeout
	}
	return nil
}

func validateMaxPlayersPerGame(maxPlayers any) error {
	if maxPlayers.(uint64) == 0 {
		return ErrInvalidMaxPlayers
	}
	return nil
}

func validateMinDistanceToWin(minDistanceToWin any) error {
	if minDistanceToWin.(uint64) == 0 {
		return ErrInvalidMinDistanceToWin
	}
	return nil
}

func validateMinReward(minReward any) error {
	if minReward.(sdk.Coin).Amount.IsZero() || minReward.(sdk.Coin).Denom == "" {
		return ErrInvalidReward
	}
	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}
