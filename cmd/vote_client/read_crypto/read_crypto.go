package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/joho/godotenv"
	pb "github.com/lfigueiredo82/klever/internal/pkg/core/model"
	"github.com/lfigueiredo82/klever/pkg/vote_server/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
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
	r, err := c.ReadCryptoCurrencies(ctx, &emptypb.Empty{})
	if err != nil {
		log.Fatalf("could not get Cryptos: %v", err)
	}
	for _, c := range r.GetCriptos() {
		fmt.Println(c.Name)
	}
}
