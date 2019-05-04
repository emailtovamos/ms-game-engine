package grpc

import (
	"context"
	pbgameengine "github.com/emailtovamos/ms-apis/ms-game-engine/v1"
	"github.com/emailtovamos/ms-game-engine/internal/server/logic"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"net"
)

type Grpc struct {
	address string
	srv     *grpc.Server
}

var HighScore = 4445555.0
var Size = 300.0

// NewServer creates a new instance of a gRPC server
func NewServer(address string) *Grpc {
	return &Grpc{
		address: address,
	}
}

// GetHighScore returns the highscore from the HighScore variable
func (g *Grpc) GetSize(ctx context.Context, input *pbgameengine.GetSizeRequest) (*pbgameengine.GetSizeResponse, error) {
	log.Info().Msg("GetSize in ms-game-engine called")
	sizeLogic := logic.GetSize()
	return &pbgameengine.GetSizeResponse{
		Size: sizeLogic, // For now this is a test size to see if connection happens correctly when the frontend calls it
	}, nil

}

// SetScore saves a score in ms-game-engine
func (g *Grpc) SetScore(ctx context.Context, input *pbgameengine.SetScoreRequest) (*pbgameengine.SetScoreResponse, error) {
	log.Info().Msg("GetSize in ms-game-engine called")
	set := logic.SetScore(input.Score)
	return &pbgameengine.SetScoreResponse{
		Set: set,
	}, nil

}

// ListenAndServe starts the gRPC server on the given address
func (g *Grpc) ListenAndServe() error {
	// open tcp port to listen for incoming connections on
	lis, err := net.Listen("tcp", g.address)
	if err != nil {
		return errors.Wrap(err, "failed to open socket")
	}

	serverOpts := []grpc.ServerOption{}

	// create the server with the specified options
	g.srv = grpc.NewServer(serverOpts...)

	pbgameengine.RegisterGameEngineServer(g.srv, g)

	log.Info().Str("addr", g.address).Msg("starting gRPC server")

	// start listening on the given address
	if err := g.srv.Serve(lis); err != nil {
		return errors.Wrap(err, "failed to start gRPC server")
	}

	return nil
}

// Create a package which will do the work of getting recent 4 scores and recommending size based on that.
