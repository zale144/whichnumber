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
		Game: game,
	}, nil
}

func (k Keeper) GetGames(goCtx context.Context, request *types.QueryGetGamesRequest) (*types.QueryGetGamesResponse, error) {
	if request == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	store := ctx.KVStore(k.storeKey)
	storedGameStore := prefix.NewStore(store, types.KeyPrefix(types.StoredGameKeyPrefix))

	var games []types.Game
	pageRes, err := query.Paginate(storedGameStore, request.Pagination, func(key []byte, value []byte) error {
		var storedGame types.Game
		if err := k.cdc.Unmarshal(value, &storedGame); err != nil {
			return fmt.Errorf("failed to unmarshal game: %w", err)
		}

		games = append(games, storedGame)
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

func (k Keeper) GetSystemInfo(goCtx context.Context, _ *types.QueryGetSystemInfoRequest) (*types.QueryGetSystemInfoResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	systemInfo, found := k.GetStoredSystemInfo(ctx)
	if !found {
		return nil, status.Error(codes.NotFound, "system info not found")
	}

	return &types.QueryGetSystemInfoResponse{
		SystemInfo: systemInfo,
	}, nil
}

func (k Keeper) GetParams(c context.Context, _ *types.QueryGetParamsRequest) (*types.QueryGetParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	params := k.GetStoredParams(ctx)

	return &types.QueryGetParamsResponse{
		Params: params,
	}, nil
}

var _ types.QueryServer = Keeper{}
