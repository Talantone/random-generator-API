package main

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"random-generator-API"
	"random-generator-API/pkg/generator"
	"random-generator-API/pkg/handler"
	"random-generator-API/pkg/usecase"
	"syscall"
)

// @title GeneratorApp API
// @version 1.0
// @description API Server for Generator Application

// @host localhost:8000
// @BasePath /generator

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing config: %s", err.Error())
	}
	gen := generator.NewRandomGenerator()
	usecases := usecase.NewGeneratorUseCase(gen)
	handlers := handler.NewHandler(usecases)
	server := new(random_generator_API.Server)
	go func() {
		if err := server.Run(viper.GetString("port"), handlers.RegisterHTTPEndpoints()); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()
	logrus.Print("App Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("App Shutting Down")

	if err := server.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
