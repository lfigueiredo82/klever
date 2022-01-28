package main

import (
	"context"
	"flag"
	"log"
	"time"

	assets "github.com/lfigueiredo82/klever/cmd/vote_client/add_crypto/dependencies"
	pb "github.com/lfigueiredo82/klever/internal/pkg/core/model"
	"github.com/lfigueiredo82/klever/pkg/vote_server/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", config.DefaultHost()+":50051", "the address to connect to")
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
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*5)
	defer cancel()

	AddBitcoinCurrency(ctx, c)
	AddEthereumCurrency(ctx, c)
	AddLitecoinCurrency(ctx, c)

}

func AddBitcoinCurrency(ctx context.Context, client pb.VoteServiceClient) {

	bitcoinCrypto := pb.CryptoCurrency{}

	bitcoinCrypto.Code = "BTC"
	bitcoinCrypto.Name = "Bitcoin"
	bitcoinCrypto.Decimals = 8
	btcImage, _ := assets.AssetsBitcoinPngBytes()
	bitcoinCrypto.ImageSymbol = btcImage

	r, err := client.AddCryptoCurrency(ctx, &bitcoinCrypto)
	if err != nil {
		log.Fatalf("could not add: %v", err)
	}
	log.Printf("Currency Added: %s", r.Name)
}

func AddEthereumCurrency(ctx context.Context, client pb.VoteServiceClient) {

	ethCrypto := pb.CryptoCurrency{}

	ethCrypto.Code = "ETH"
	ethCrypto.Name = "Ethereum"
	ethCrypto.Decimals = 8
	ethImage, _ := assets.AssetsEthereumPngBytes()
	ethCrypto.ImageSymbol = ethImage

	r, err := client.AddCryptoCurrency(ctx, &ethCrypto)
	if err != nil {
		log.Fatalf("could not add: %v", err)
	}
	log.Printf("Currency Added: %s", r.Name)
}

func AddLitecoinCurrency(ctx context.Context, client pb.VoteServiceClient) {

	litecoinCrypto := pb.CryptoCurrency{}

	litecoinCrypto.Code = "LTC"
	litecoinCrypto.Name = "Litecoin"
	litecoinCrypto.Decimals = 8
	ltcImage, _ := assets.AssetsLitecoinPngBytes()
	litecoinCrypto.ImageSymbol = ltcImage

	r, err := client.AddCryptoCurrency(ctx, &litecoinCrypto)
	if err != nil {
		log.Fatalf("could not add: %v", err)
	}
	log.Printf("Currency Added: %s", r.Name)
}
