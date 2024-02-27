package keeper

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zale144/whichnumber/x/whichnumber/types"
)

func (k Keeper) GetStoredGame(ctx sdk.Context, id int64) (val types.Game, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StoredGameKeyPrefix))

	b := store.Get(types.StoredGameKey(strconv.FormatInt(id, 10)))
	if b == nil {
		return
	}

	found = true
	k.cdc.MustUnmarshal(b, &val)
	return
}

func (k Keeper) GetStoredGames(ctx sdk.Context) (list []types.Game) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StoredGameKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, nil)

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var game types.Game
		k.cdc.MustUnmarshal(iterator.Value(), &game)
		list = append(list, game)
	}

	return
}

func (k Keeper) IterateStoredGames(ctx sdk.Context, headId int64, cb func(game types.Game) (stop bool, nextId int64)) {
	gameId := headId
	stop := false

	for gameId != types.NoFifoId {
		game, found := k.GetStoredGame(ctx, gameId)
		if !found {
			panic("Game not found")
		}

		stop, gameId = cb(game)
		if stop {
			break
		}
	}
}

func (k Keeper) SetStoredGame(ctx sdk.Context, storedGame types.Game) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StoredGameKeyPrefix))
	b := k.cdc.MustMarshal(&storedGame)
	store.Set(types.StoredGameKey(strconv.FormatInt(storedGame.Id, 10)), b)
}

func (k Keeper) RemoveStoredGame(ctx sdk.Context, id int64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StoredGameKeyPrefix))
	store.Delete(types.StoredGameKey(strconv.FormatInt(id, 10)))
}

func (k Keeper) GetNumberCommit(game types.Game, player string) (*types.NumberCommit, bool) {
	for _, commit := range game.PlayerCommits {
		if commit.PlayerAddress == player {
			return commit, true
		}
	}

	return nil, false
}

func (k Keeper) GetNumberReveal(game types.Game, player string) (*types.NumberReveal, bool) {
	for _, reveal := range game.PlayerReveals {
		if reveal.PlayerAddress == player {
			return reveal, true
		}
	}

	return nil, false
}
func (k Keeper) RemoveFromFifo(ctx sdk.Context, game *types.Game, info *types.SystemInfo) {
	// Does it have a predecessor?
	if game.BeforeId != types.NoFifoId {
		beforeGame, found := k.GetStoredGame(ctx, game.BeforeId)
		if !found {
			panic("Element before in Fifo was not found")
		}

		beforeGame.AfterId = game.AfterId
		k.SetStoredGame(ctx, beforeGame)

		if game.AfterId == types.NoFifoId {
			info.FifoTailId = game.Id
		}
		// Is it at the FIFO head?
	} else if info.FifoHeadId == game.Id {
		info.FifoHeadId = game.AfterId
	}

	// Does it have a successor?
	if game.AfterId != types.NoFifoId {
		afterGame, found := k.GetStoredGame(ctx, game.AfterId)
		if !found {
			panic("Element after in Fifo was not found")
		}

		afterGame.BeforeId = game.BeforeId
		k.SetStoredGame(ctx, afterGame)

		if game.BeforeId == types.NoFifoId {
			info.FifoHeadId = afterGame.Id
		}
		// Is it at the FIFO tail?
	} else if info.FifoTailId == game.Id {
		info.FifoTailId = game.BeforeId
	}
	game.BeforeId = types.NoFifoId
	game.AfterId = types.NoFifoId
}

func (k Keeper) SendToFifoTail(ctx sdk.Context, game *types.Game, info *types.SystemInfo) {
	if info.FifoHeadId == types.NoFifoId && info.FifoTailId == types.NoFifoId {
		game.BeforeId = types.NoFifoId
		game.AfterId = types.NoFifoId
		info.FifoHeadId = game.Id
		info.FifoTailId = game.Id
	} else if info.FifoHeadId == types.NoFifoId || info.FifoTailId == types.NoFifoId {
		panic("Fifo should have both head and tail or none")
	} else if info.FifoTailId != game.Id {
		// Snip game out
		k.RemoveFromFifo(ctx, game, info)

		// Now add to tail
		currentTail, found := k.GetStoredGame(ctx, info.FifoTailId)
		if !found {
			panic("Current Fifo tail was not found")
		}
		currentTail.AfterId = game.Id
		k.SetStoredGame(ctx, currentTail)

		game.BeforeId = currentTail.Id
		info.FifoTailId = game.Id
	}
}
