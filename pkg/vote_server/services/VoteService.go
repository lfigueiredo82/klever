package services

import (
	"context"
	"encoding/json"
	"log"
	"strconv"
	"time"

	"github.com/go-redis/redis"
	pb "github.com/lfigueiredo82/klever/internal/pkg/core/model"
	r "github.com/lfigueiredo82/klever/internal/pkg/core/redis"
	"github.com/lfigueiredo82/klever/pkg/vote_server/config"
	"google.golang.org/protobuf/types/known/emptypb"
)

type VoteService struct {
	pb.UnimplementedVoteServiceServer
}

func (s *VoteService) ReadCryptoCurrencies(ctx context.Context, in *emptypb.Empty) (*pb.CryptoReply, error) {

	client := connectOnRedis()

	result := r.ReadAllCryptoCurrency(client)
	cryptos := make([]*pb.CryptoCurrency, 0)
	for _, key := range result {
		crypto := pb.CryptoCurrency{}
		js := r.GetCryptoCurrency(client, key)
		json.Unmarshal([]byte(js), &crypto)
		cryptos = append(cryptos, &crypto)
	}
	return &pb.CryptoReply{Criptos: cryptos}, nil

}

func (s *VoteService) AddCryptoCurrency(ctx context.Context, cripto *pb.CryptoCurrency) (*pb.CryptoCurrency, error) {

	client := connectOnRedis()

	js, _ := json.Marshal(cripto)

	r.CreateCryptoCurrency(client, cripto.Code, js)

	return cripto, nil
}

func (s *VoteService) UpdateCryptoCurrency(ctx context.Context, cripto *pb.CryptoCurrency) (*pb.CryptoCurrency, error) {

	client := connectOnRedis()
	js, _ := json.Marshal(cripto)
	r.UpdateCryptoCurrency(client, cripto.Code, cripto.Code, js)
	return cripto, nil
}

func (s *VoteService) RemoveCryptoCurrency(ctx context.Context, cripto *pb.CryptoSymbol) (*emptypb.Empty, error) {

	client := connectOnRedis()

	r.DeleteCryptoCurrency(client, cripto.Symbol)

	return nil, nil
}

func (s *VoteService) Upvote(ctx context.Context, vote *pb.Vote) (*emptypb.Empty, error) {
	client := connectOnRedis()
	r.UpvoteCrytoCurrency(client, vote.CryptoCurrency.Code)

	return &emptypb.Empty{}, nil
}

func (s *VoteService) Downvote(ctx context.Context, vote *pb.Vote) (*emptypb.Empty, error) {
	client := connectOnRedis()
	r.DownvoteCrytoCurrency(client, vote.CryptoCurrency.Code)

	return &emptypb.Empty{}, nil
}

func (s *VoteService) SumVotesFromCryptoCurrency(ctx context.Context, cryptoSymbol *pb.CryptoSymbol) (*pb.TotalVotes, error) {
	client := connectOnRedis()

	value := r.GetCryptoCurrency(client, cryptoSymbol.Symbol)
	if value != "" {
		total, _ := strconv.ParseFloat(value, 64)
		return &pb.TotalVotes{Votes: total}, nil
	}
	return &pb.TotalVotes{Votes: 0}, nil

}

func connectOnRedis() *redis.Client {
	log.Println("Connecting on Redis...")
	client := redis.NewClient(&redis.Options{
		Addr:         config.RedisAddress(),
		Password:     config.RedisPassword(),
		DB:           0,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	})

	_, err := client.Ping().Result()
	if err != nil {
		log.Panic(err)
	}
	return client
}
