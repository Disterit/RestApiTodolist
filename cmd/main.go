package main

import (
	RestApiTodolist "ResTApiTodolist"
	"ResTApiTodolist/pkg/handler"
	"ResTApiTodolist/pkg/repository"
	"ResTApiTodolist/pkg/service"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"

	_ "github.com/lib/pq"
)

func main() {

	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initial config: %v", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading .env file: %v", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetInt("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: viper.GetString("db.database"),
		SSLMode:  viper.GetString("db.SSLMode"),
	})

	if err != nil {
		logrus.Fatalf("error initial db: %v", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(RestApiTodolist.Server)

	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while runing http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
