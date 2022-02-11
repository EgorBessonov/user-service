package main

import (
	"context"
	"github.com/EgorBessonov/user-service/internal/config"
	"github.com/EgorBessonov/user-service/internal/repository"
	"github.com/EgorBessonov/user-service/internal/service"
	"github.com/EgorBessonov/user-service/protocol"
	"github.com/caarlos0/env"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	cfg := config.Config{}
	if err := env.Parse(&cfg); err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Fatal("user service: can't parse config")
	}
	postgres, err := pgxpool.Connect(context.Background(), cfg.PostgresURL)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Fatal("user service: can't connect to the postgres database")
	}
	rps := repository.NewPostgresRepository(postgres)
	serv := service.NewService(rps)
	newgRPCServer(cfg.UserServicePort, serv)
}

func newgRPCServer(port string, s *service.Service) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Fatal("user service: can't create gRPC server")
	}
	gServer := grpc.NewServer()
	userService.RegisterUserServer(gServer, s)
	log.Printf("user service: listening at %s", lis.Addr())
	if err = gServer.Serve(lis); err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Fatal("user service: gRPC server failed")
	}
}
