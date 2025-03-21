package keeper

import (
	"context"

	"github.com/alice/checkers"
)

// InitGenesis initializes the module state from a genesis state.
func (k *Keeper) InitGenesis(ctx context.Context, data *checkers.GenesisState) error {
    if err := k.Params.Set(ctx, data.Params); err != nil {
        return err
    }

    for _, indexedStoredGame := range data.IndexedStoredGameList {
        if err := k.StoredGames.Set(ctx, indexedStoredGame.Index, indexedStoredGame.StoredGame); err != nil {
            return err
        }
    }

    return nil
}

// ExportGenesis exports the module state to a genesis state.
func (k *Keeper) ExportGenesis(ctx context.Context) (*checkers.GenesisState, error) {
    params, err := k.Params.Get(ctx)
    if err != nil {
        return nil, err
    }

    var indexedStoredGameList []checkers.IndexedStoredGame
    if err := k.StoredGames.Walk(ctx, nil, func(index string, storedGame checkers.StoredGame) (bool, error) {
        indexedStoredGameList = append(indexedStoredGameList, checkers.IndexedStoredGame{
            Index: index,
            StoredGame: storedGame,
        })
        return true, nil
    }); err != nil {
        return nil, err
    }

    return &checkers.GenesisState{
        Params: params,
        IndexedStoredGameList: indexedStoredGameList,
    }, nil
}