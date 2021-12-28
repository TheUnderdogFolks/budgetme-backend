package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/TheWokeDeveloper/budgetme-backend/grpc/budgetme/proto"
	"github.com/TheWokeDeveloper/budgetme-backend/internal/config"
	serverHandler "github.com/TheWokeDeveloper/budgetme-backend/internal/handler/grpc/server"
	serverUseCase "github.com/TheWokeDeveloper/budgetme-backend/internal/usecase/server"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/newrelic/go-agent/v3/integrations/nrgrpc"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func startServer(cfg *config.Config, metrics *newrelic.Application, serverUseCase *serverUseCase.Server) {
	server := grpc.NewServer(
		grpc.UnaryInterceptor(nrgrpc.UnaryServerInterceptor(metrics)),
		grpc.StreamInterceptor(nrgrpc.StreamServerInterceptor(metrics)),
	)
	reflection.Register(server)

	serverHandler := serverHandler.New(serverUseCase)
	pb.RegisterServerServer(server, serverHandler)

	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", cfg.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("server listening at %v", listener.Addr())
	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
