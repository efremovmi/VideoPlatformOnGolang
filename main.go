package main

import (
	"context"
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/genridarkbkru/registration/internal/app/apiserver"
	pb "github.com/genridarkbkru/registration/internal/app/grpc_server/github.com/genridarkbkru/registration"
	"github.com/genridarkbkru/registration/internal/app/store/sqlstore"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"os"
)

var (
	configPath string
)

const (
	grpcPort           = ":5000"
	grpcServerEndpoint = "localhost:5000"
)

// Взять значения из файла apiserver.toml для конфигурации сервиса.
func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

// Запуск сервиса.
func run() error {
	flag.Parse()
	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)

	if err != nil {
		return err
	}

	lis, err := net.Listen("tcp", grpcPort)
	if err != nil {
		return err
	}
	s := grpc.NewServer()

	databaseURL := os.Getenv("DATABASE_URL") + "?sslmode=" + os.Getenv("DATABASE_SSL_MODE")
	db, err := apiserver.NewDB(databaseURL)
	if err != nil {
		return err
	}
	defer db.Close()

	store := sqlstore.New(db)

	srv := apiserver.NewServer(store)

	pb.RegisterServiceRAServer(s, srv)
	log.Printf("Starting gRPC listener on port " + grpcPort)
	if err := s.Serve(lis); err != nil {
		return err
	}
	return nil
}

func runJSON() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := pb.RegisterServiceRAHandlerFromEndpoint(ctx, mux, grpcServerEndpoint, opts)
	if err != nil {
		panic(err)
	}

	err = http.ListenAndServe(":5001", mux)
	if err != nil {
		panic(err)
	}
}

func main() {
	go runJSON()
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
