package keeper

import (
	"context"
	"fmt"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/zale144/whichnumber/x/whichnumber/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetGame(goCtx context.Context, request *types.QueryGetGameRequest) (*types.QueryGetGameResponse, error) {
	if request == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	game, found := k.GetStoredGame(ctx, request.Id)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetGameResponse{
		Game: types.StoredGameToQueryGameResponse(game),
	}, nil
}

func (k Keeper) GetGames(goCtx context.Context, request *types.QueryGetGamesRequest) (*types.QueryGetGamesResponse, error) {
	if request == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	store := ctx.KVStore(k.storeKey)
	storedGameStore := prefix.NewStore(store, types.KeyPrefix(types.StoredGameKeyPrefix))

	var games []*types.QueryGameResponse
	pageRes, err := query.Paginate(storedGameStore, request.Pagination, func(key []byte, value []byte) error {
		var storedGame types.Game
		if err := k.cdc.Unmarshal(value, &storedGame); err != nil {
			return fmt.Errorf("failed to unmarshal game: %w", err)
		}

		gameResponse := types.StoredGameToQueryGameResponse(storedGame)
		games = append(games, &gameResponse)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryGetGamesResponse{
		Games:      games,
		Pagination: pageRes,
	}, nil
}
