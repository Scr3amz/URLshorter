package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/Scr3amz/URLshorter/config"
	desc "github.com/Scr3amz/URLshorter/internal/api/proto"
	ser "github.com/Scr3amz/URLshorter/internal/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const grpcPort = 50051

type server struct {
	desc.UnimplementedURLshorterServer
}

func (s *server) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	log.Println("Я вернул оригинальную ссылку")
	return &desc.GetResponse{
		LongURL: "https://www.google.com",
	}, nil
} 

func (s *server) Post(ctx context.Context, req *desc.PostRequest) (*desc.PostResponse, error) {
	log.Println("Я положил ссылку в БД и вернул сокращённую")
	return &desc.PostResponse{
		ShortURL: ser.ShortURL(req.GetLongURL()),
	}, nil
}

func main() {
	config, err := config.LoadConfig("./", "data", "env")
	if err != nil {
		log.Fatalf("Error occured while reading the config file\nError: %v\n", err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.GrpcPort))
	if err != nil {
		log.Fatalf("Failed to listen\nError: %v\n", err)
	}

	s:= grpc.NewServer()
	reflection.Register(s)
	desc.RegisterURLshorterServer(s, &server{})

	log.Printf("Server listening at localhost:%d", config.GrpcPort)

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve\nError: %v\n", err)
	}
}