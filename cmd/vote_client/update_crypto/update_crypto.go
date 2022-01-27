package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/lfigueiredo82/klever/internal/pkg/core/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func main() {

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
