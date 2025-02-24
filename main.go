package main

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"

	"github.com/itmrchow/microservice-account/delivery/grpc"
)

func main() {
	initConfig()
	initServer()
}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal().Err(err).Msg("config init error")
	}

	log.Info().Msgf("config init success")
}

func initServer() {
	log.Info().Msgf("grpc server init success , port: %s", viper.GetString("grpc.port"))
	log.Fatal().Err(grpc.RunGrpcServer(viper.GetString("grpc.port"))).Msg("grpc server init error")
}
