package keeper

import (
	"context"
	"cosmossdk.io/collections"
	"errors"
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

func (q queryServer) GetGame(ctx context.Context, request *checkers.QueryGetGameRequest) (*checkers.QueryGetGameResponse, error) {
	game, err := q.k.StoredGames.Get(ctx, request.Index)
	if err == nil {
		return &checkers.QueryGetGameResponse{Game: &game}, nil
	}
	if errors.Is(err, collections.ErrNotFound) {
		return &checkers.QueryGetGameResponse{Game: nil}, nil
	}

	return nil, status.Error(codes.Internal, err.Error())
}
