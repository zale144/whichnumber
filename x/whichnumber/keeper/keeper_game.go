package keeper

import (
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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

func (k Keeper) IterateStoredGames(ctx sdk.Context, headId int64, cb func(game types.Game) (nextId int64)) {
	gameId := headId

	for gameId != types.NoFifoId {
		game, found := k.GetStoredGame(ctx, gameId)
		if !found {
			panic("Game not found")
		}

		gameId = cb(game)
		// TODO: is there a case where the iteration should stop prematurely?
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

func (k Keeper) removeFromFifo(ctx sdk.Context, game *types.Game, info *types.SystemInfo) {
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

func (k Keeper) sendToFifoTail(ctx sdk.Context, game *types.Game, info *types.SystemInfo) {
	if info.FifoHeadId == types.NoFifoId && info.FifoTailId == types.NoFifoId {
		game.BeforeId = types.NoFifoId
		game.AfterId = types.NoFifoId
		info.FifoHeadId = game.Id
		info.FifoTailId = game.Id
	} else if info.FifoHeadId == types.NoFifoId || info.FifoTailId == types.NoFifoId {
		panic("Fifo should have both head and tail or none")
	} else if info.FifoTailId != game.Id {
		// Snip game out
		k.removeFromFifo(ctx, game, info)

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

func (k Keeper) sendCoinsToPlayer(ctx sdk.Context, player string, coins sdk.Coin) error {
	if !coins.IsPositive() {
		return nil
	}
	creatorAddress, err := types.GetPlayerAddress(player)
	if err != nil {
		return fmt.Errorf("invalid creator address: %w", err)
	}

	if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, creatorAddress, sdk.NewCoins(coins)); err != nil {
		return fmt.Errorf("failed to send coins to player %s: %w", player, err)
	}
	return nil
}

func (k Keeper) sendCoinsToModule(ctx sdk.Context, player string, coins sdk.Coin) error {
	if !coins.IsPositive() {
		return nil
	}
	playerAddress, err := types.GetPlayerAddress(player)
	if err != nil {
		return sdkerrors.Wrapf(types.ErrFailedToSendCoinsToModule, "error: %s", err)
	}

	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, playerAddress, types.ModuleName, sdk.NewCoins(coins)); err != nil {
		return sdkerrors.Wrapf(types.ErrFailedToSendCoinsToModule, "error: %s", err)
	}
	return nil
}
