package keeper

import (
	"fmt"

	"cosmossdk.io/collections"
	"cosmossdk.io/core/address"
	storetypes "cosmossdk.io/core/store"
	"github.com/alice/checkers"
	"github.com/cosmos/cosmos-sdk/codec"
)

type Keeper struct {
    cdc          codec.BinaryCodec
    addressCodec address.Codec

    // authority is the address capable of executing a MsgUpdateParams and other authority-gated message.
    // typically, this should be the x/gov module account.
    authority string

    // state management
    Schema collections.Schema
    Params collections.Item[checkers.Params]
}

// NewKeeper creates a new Keeper instance
func NewKeeper(cdc codec.BinaryCodec, addressCodec address.Codec, storeService storetypes.KVStoreService, authority string) Keeper {
    if _, err := addressCodec.StringToBytes(authority); err != nil {
        panic(fmt.Errorf("invalid authority address: %w", err))
    }

    sb := collections.NewSchemaBuilder(storeService)
    k := Keeper{
        cdc:          cdc,
        addressCodec: addressCodec,
        authority:    authority,
        Params:       collections.NewItem(sb, checkers.ParamsKey, "params", codec.CollValue[checkers.Params](cdc)),
    }

    schema, err := sb.Build()
    if err != nil {
        panic(err)
    }

    k.Schema = schema

    return k
}
