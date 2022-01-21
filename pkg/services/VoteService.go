package services

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis"
	pb "github.com/lfigueiredo82/klever/internal/pkg/core/model"
	r "github.com/lfigueiredo82/klever/internal/pkg/core/redis"
	"github.com/lfigueiredo82/klever/pkg/config"
	"google.golang.org/protobuf/types/known/emptypb"
)

type VoteService struct {
	pb.UnimplementedVoteServiceServer
}

func (s *VoteService) ReadCryptoCurrencies(ctx context.Context, in *emptypb.Empty) (*pb.CryptoReply, error) {

	client := connectOnRedis()

	result := r.ReadAllCryptoCurrency(client)
	cryptos := make([]*pb.CryptoCurrency, 0)
	for _, r := range result {
		fmt.Println(r)
		crypto := pb.CryptoCurrency{Name: r}
		cryptos = append(cryptos, &crypto)
	}
	return &pb.CryptoReply{Criptos: cryptos}, nil

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
