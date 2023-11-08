package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/Scr3amz/URLshorter/config"
	desc "github.com/Scr3amz/URLshorter/internal/api/proto"
	"github.com/Scr3amz/URLshorter/internal/database/urls"
	postgres "github.com/Scr3amz/URLshorter/internal/database/urls/postgres"

	ser "github.com/Scr3amz/URLshorter/internal/server"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var Repository urls.Repository



type server struct {
	desc.UnimplementedURLshorterServer
}

func main() {
	config, err := config.LoadConfig("./", "data", "env")
	if err != nil {
		log.Fatalf("Error occured while reading the config file\nError: %v\n", err)
	}

	dsn := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v",
		config.Host,
		config.DbPort,
		config.DbUser,
		config.DbPassword,
		config.DbName,
	)

	db, err := sqlx.Connect("pgx", dsn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	Repository = postgres.NewRepository(db)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.GrpcPort))
	if err != nil {
		log.Fatalf("Failed to listen\nError: %v\n", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterURLshorterServer(s, &server{})

	log.Printf("Server listening at localhost:%d", config.GrpcPort)

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve\nError: %v\n", err)
	}
}

func (s *server) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	log.Println("Я вернул оригинальную ссылку")
	return &desc.GetResponse{
		LongURL: "https://www.google.com",
	}, nil
}

func (s *server) Post(ctx context.Context, req *desc.PostRequest) (*desc.PostResponse, error) {
	longURL := req.GetLongURL()
	shortURL := ser.ShortURL(longURL)
	url:= urls.URLs {
		LongURL: longURL,
		ShortURL: shortURL,
	}
	
	if err := Repository.Create(ctx, &url ); err != nil {
		return nil, err
	}

	log.Println("Я положил ссылку в БД и вернул сокращённую")
	return &desc.PostResponse{
		ShortURL: shortURL,
	}, nil
}
