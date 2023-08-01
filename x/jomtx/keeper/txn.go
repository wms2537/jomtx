package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jomluz/jomtx/x/jomtx/types"
)

// GetTxnCount get the total number of txn
func (k Keeper) GetTxnCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.TxnCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetTxnCount set the total number of txn
func (k Keeper) SetTxnCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.TxnCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendTxn appends a txn in the store with a new id and update the count
func (k Keeper) AppendTxn(
	ctx sdk.Context,
	txn types.Txn,
) uint64 {
	// Create the txn
	count := k.GetTxnCount(ctx)

	// Set the ID of the appended value
	txn.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TxnKey))
	appendedValue := k.cdc.MustMarshal(&txn)
	store.Set(GetTxnIDBytes(txn.Id), appendedValue)

	// Update txn count
	k.SetTxnCount(ctx, count+1)

	return count
}

// SetTxn set a specific txn in the store
func (k Keeper) SetTxn(ctx sdk.Context, txn types.Txn) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TxnKey))
	b := k.cdc.MustMarshal(&txn)
	store.Set(GetTxnIDBytes(txn.Id), b)
}

// GetTxn returns a txn from its id
func (k Keeper) GetTxn(ctx sdk.Context, id uint64) (val types.Txn, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TxnKey))
	b := store.Get(GetTxnIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveTxn removes a txn from the store
func (k Keeper) RemoveTxn(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TxnKey))
	store.Delete(GetTxnIDBytes(id))
}

// GetAllTxn returns all txn
func (k Keeper) GetAllTxn(ctx sdk.Context) (list []types.Txn) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TxnKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Txn
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetTxnIDBytes returns the byte representation of the ID
func GetTxnIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetTxnIDFromBytes returns ID in uint64 format from a byte array
func GetTxnIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
