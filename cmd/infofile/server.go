package main

import (
	"fmt"
	"github.com/mblancoa/go-rpc/conf"
	"github.com/mblancoa/go-rpc/internal/adapters/mongodb"
	"github.com/mblancoa/go-rpc/internal/core"
	"github.com/mblancoa/go-rpc/internal/errors"
	"github.com/mblancoa/go-rpc/rpc"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type rpcConfiguration struct {
	Rpc struct {
		Server struct {
			Host string `yaml:"host"`
			Port int    `yaml:"port"`
		} `yaml:"server"`
	} `yaml:"rpc"`
}

func main() {
	mongodb.SetupMongodbConfiguration()
	core.SetupCoreConfiguration()
	SetupRpcServer()
}

func SetupRpcServer() {
	log.Info().Msg("Initializing rpc server configuration")
	var c rpcConfiguration
	conf.LoadYamlConfiguration(conf.GetConfigFile(), &c)

	address := fmt.Sprintf("%s:%d", c.Rpc.Server.Host, c.Rpc.Server.Port)
	log.Info().Msgf("running server on %s", address)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal().Msgf("Failed to listen %v", err)
	}

	server := rpc.NewInfoFileServiceServer(core.Context.InfoFileService)
	s := grpc.NewServer()
	rpc.RegisterInfoFileServiceServer(s, server)

	reflection.Register(s)
	err = s.Serve(lis)
	errors.ManageErrorPanic(err)
}
