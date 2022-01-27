package test

import (
	"context"
	"testing"

	pb "github.com/lfigueiredo82/klever/internal/pkg/core/model"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

func TestReadCryptoCurrencies(t *testing.T) {
	const address = "localhost:50051"
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewVoteServiceClient(conn)

	t.Run("ReadCrypto", func(t *testing.T) {
		r, err := c.ReadCryptoCurrencies(context.Background(), &emptypb.Empty{})
		if err != nil {
			t.Fatalf("could not read: %v", err)
		}
		t.Logf("Cryptos: %s", r.Criptos)

	})
}

func TestUpdateCryptoCurrency(t *testing.T) {
	const address = "localhost:50051"
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewVoteServiceClient(conn)

	t.Run("UpdateCrypto", func(t *testing.T) {
		_, err := c.UpdateCryptoCurrency(context.Background(), &pb.CryptoCurrency{
			Code:     "ETH",
			Name:     "ETHEREUM",
			Decimals: 2,
		})
		if err != nil {
			t.Fatalf("could not read: %v", err)
		}

	})
}

func TestAddCryptoCurrency(t *testing.T) {
	const address = "localhost:50051"
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewVoteServiceClient(conn)

	t.Run("AddCrypto", func(t *testing.T) {
		_, err := c.AddCryptoCurrency(context.Background(), &pb.CryptoCurrency{
			Code:     "BTC",
			Name:     "Bitcoin",
			Decimals: 8,
		})
		if err != nil {
			t.Fatalf("could not read: %v", err)
		}

	})
}

func TestRemoveCryptoCurrency(t *testing.T) {
	const address = "localhost:50051"
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewVoteServiceClient(conn)

	t.Run("RemoveCrypto", func(t *testing.T) {
		_, err := c.RemoveCryptoCurrency(context.Background(), &pb.CryptoSymbol{Symbol: "BTC"})
		if err != nil {
			t.Fatalf("could not read: %v", err)
		}

	})
}

func TestUpvoteCryptoCurrency(t *testing.T) {
	const address = "localhost:50051"
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewVoteServiceClient(conn)

	t.Run("UpvoteCrypto", func(t *testing.T) {
		_, err := c.Upvote(context.Background(), &pb.Vote{
			CryptoCurrency: &pb.CryptoCurrency{Code: "BTC"},
		})
		if err != nil {
			t.Fatalf("could not read: %v", err)
		}

	})
}

func TestDownvoteCryptoCurrency(t *testing.T) {
	const address = "localhost:50051"
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewVoteServiceClient(conn)

	t.Run("DownvoteCrypto", func(t *testing.T) {
		_, err := c.Downvote(context.Background(), &pb.Vote{
			CryptoCurrency: &pb.CryptoCurrency{Code: "BTC"},
		})
		if err != nil {
			t.Fatalf("could not read: %v", err)
		}

	})
}

func TestSumVotesFromCryptoCurrency(t *testing.T) {
	const address = "localhost:50051"
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewVoteServiceClient(conn)

	t.Run("SumVotes", func(t *testing.T) {
		r, err := c.SumVotesFromCryptoCurrency(context.Background(), &pb.CryptoSymbol{
			Symbol: "BTC",
		})
		if err != nil {
			t.Fatalf("could not read: %v", err)
		}
		if r.Votes < 0 {
			t.Fail()
		}

	})
}
