package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/jomluz/jomtx/x/jomtx/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) TxnAll(goCtx context.Context, req *types.QueryAllTxnRequest) (*types.QueryAllTxnResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var txns []types.Txn
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	txnStore := prefix.NewStore(store, types.KeyPrefix(types.TxnKey))

	pageRes, err := query.Paginate(txnStore, req.Pagination, func(key []byte, value []byte) error {
		var txn types.Txn
		if err := k.cdc.Unmarshal(value, &txn); err != nil {
			return err
		}

		txns = append(txns, txn)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllTxnResponse{Txn: txns, Pagination: pageRes}, nil
}

func (k Keeper) Txn(goCtx context.Context, req *types.QueryGetTxnRequest) (*types.QueryGetTxnResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	txn, found := k.GetTxn(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetTxnResponse{Txn: txn}, nil
}
