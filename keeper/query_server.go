package keeper

import (
	"context"
	"errors"

	"cosmossdk.io/collections"
	"github.com/alice/checkers"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ checkers.QueryServer = queryServer{}

func NewQueryServerImpl(k Keeper) checkers.QueryServer {
	return queryServer{k}
}

type queryServer struct {
	k Keeper
}

// GetGame implements checkers.QueryServer.
func (q queryServer) GetGame(ctx context.Context, req *checkers.QueryGetGameRequest) (*checkers.QueryGetGameResponse, error) {

	game, err := q.k.StoredGames.Get(ctx, req.Index);
	if errors.Is(err, collections.ErrNotFound) {
		return &checkers.QueryGetGameResponse{Game: nil}, nil
	}

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &checkers.QueryGetGameResponse{Game: &game}, nil
}
