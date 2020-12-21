package main

import (
	"github.com/spf13/viper"
	"github.com/zayonvx/todo-app"
	"github.com/zayonvx/todo-app/pkg/handler"
	"github.com/zayonvx/todo-app/pkg/repository"
	"github.com/zayonvx/todo-app/pkg/service"
	"log"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
