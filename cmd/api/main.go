package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/Vadim-12/tszh-backend/pkg/app"
	config "github.com/Vadim-12/tszh-backend/pkg/config/utils"
	"github.com/Vadim-12/tszh-backend/pkg/handler"
	"github.com/Vadim-12/tszh-backend/pkg/repository"
	"github.com/Vadim-12/tszh-backend/pkg/service"
	"github.com/Vadim-12/tszh-backend/pkg/utils"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := godotenv.Load(); err != nil {
		logrus.Info("no .env file found, using environment variables")
	}

	ctx := context.Background()
	gdb, sqlDB, err := repository.NewPostgresDB(ctx, os.Getenv("DB_DSN"))
	if err != nil {
		logrus.Fatalf("failed to initialize DB: %s", err.Error())
	}
	defer sqlDB.Close()

	accessTTL, err := time.ParseDuration(os.Getenv("JWT_ACCESS_TTL"))
	if err != nil {
		logrus.Fatalf("failed to parse JWT_ACCESS_TTL from env: %s", err)
	}
	refreshTTL, err := time.ParseDuration(os.Getenv("JWT_REFRESH_TTL"))
	if err != nil {
		logrus.Fatalf("failed to parse JWT_REFRESH_TTL from env: %s", err)
	}

	utils := utils.NewUtils(
		config.BcryptHasher{
			Cost: 0,
		},
		config.JWTSigner{
			AccessSecret:  []byte(os.Getenv("JWT_ACCESS_SECRET")),
			AccessTTL:     accessTTL,
			RefreshSecret: []byte(os.Getenv("JWT_REFRESH_SECRET")),
			RefreshTTL:    refreshTTL,
		},
	)
	log.Println("utils", utils)
	repos := repository.NewRepository(gdb)
	log.Println("repos", repos)
	services := service.NewService(repos, utils)
	log.Println("services", services)
	handlers := handler.NewHandler(services)
	log.Println("handlers", handlers)
	router := handlers.InitRoutes()

	logrus.Infof("server is starting on port %s", os.Getenv("PORT"))
	srv := new(app.Server)
	if err := srv.Run(os.Getenv("PORT"), router); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}
}
