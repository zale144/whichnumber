package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/zale144/whichnumber/x/whichnumber/types"
)

func (k Keeper) SetStoredSystemInfo(ctx sdk.Context, systemInfo types.SystemInfo) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SystemInfoKey))
	b := k.cdc.MustMarshal(&systemInfo)
	store.Set([]byte{0}, b)
}

func (k Keeper) GetStoredSystemInfo(ctx sdk.Context) (val types.SystemInfo, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SystemInfoKey))

	b := store.Get([]byte{0})
	if b == nil {
		return
	}

	k.cdc.MustUnmarshal(b, &val)
	found = true
	return
}

func (k Keeper) RemoveStoredSystemInfo(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SystemInfoKey))
	store.Delete([]byte{0})
}
