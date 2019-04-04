// Connect to storage serve to get latest 4 scores
// If the trend is worse bestscore then return larger size of shape to make the game less difficult
// If the trend is improved bestscore then return smaller size of shape to make the game more difficult
// Start with just assuming you have the scores from storage by randomly generating 4 numbers with certain order
// TEST if the function returns values as expected
// THEN add logic to connect to storage service once storage service is up and running

// ms-game-engine opens a server so that when it gets a GetSize request, it returns size of shape: height, width
// Only border-radius, height and width matter
// Border radius (50% - circle, 0% - rectangle) may or may not be decided on the client side
package main

import (
	"flag"
	grpcSetup "github.com/emailtovamos/ms-game-engine/internal/server/grpc"
	"github.com/rs/zerolog/log"
)

func main() {

	var addressPtr = flag.String("address", ":60051", "address where you can connect with ms-game-engine service")
	// Following is a reason why :60051 is used as address instead of localhost:60051
	// https://stackoverflow.com/questions/43911793/cannot-connect-to-go-grpc-server-running-in-local-docker-container

	flag.Parse()

	s := grpcSetup.NewServer(*addressPtr)

	// start gRPC server
	if err := s.ListenAndServe(); err != nil {
		log.Fatal().Err(err).Msg("failed to start gRPC server")
	}

}
