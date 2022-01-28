package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/joho/godotenv"
	pb "github.com/lfigueiredo82/klever/internal/pkg/core/model"
	"github.com/lfigueiredo82/klever/pkg/vote_server/config"
	"github.com/lfigueiredo82/klever/pkg/vote_server/services"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	godotenv.Load()
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", config.DefaultHost(), *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterVoteServiceServer(s, &services.VoteService{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
