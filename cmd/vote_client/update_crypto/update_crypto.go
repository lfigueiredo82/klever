package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/joho/godotenv"
	pb "github.com/lfigueiredo82/klever/internal/pkg/core/model"
	"github.com/lfigueiredo82/klever/pkg/vote_server/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	godotenv.Load()
	addr := flag.String("addr", config.DefaultHost()+":50051", "the address to connect to")
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewVoteServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	bitcoinCrypto := pb.CryptoCurrency{}

	bitcoinCrypto.Code = "BTC"
	bitcoinCrypto.Name = "Bitcoin_XXX"
	bitcoinCrypto.Decimals = 8

	r, err := c.UpdateCryptoCurrency(ctx, &bitcoinCrypto)
	if err != nil {
		log.Fatalf("could not update: %v", err)
	}
	log.Printf("Currency updated: %s", r.Name)

}
