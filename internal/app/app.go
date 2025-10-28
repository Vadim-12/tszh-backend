package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	httpx "github.com/Vadim-12/tszh-backend/internal/adapters/http"
	pgx "github.com/Vadim-12/tszh-backend/internal/adapters/postgres"
	"github.com/Vadim-12/tszh-backend/internal/domains/buildings"
	"github.com/Vadim-12/tszh-backend/internal/domains/users"
	"github.com/Vadim-12/tszh-backend/internal/platform/config"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Run() error {
	config := loadConfig()

	ctx := context.Background()
	gdb, sqlDB, err := pgx.Open(ctx, config.DSN)
	if err != nil {
		log.Fatal("cannot connect DB:", err)
	}
	defer sqlDB.Close()

	// Авто-создание таблицы для модели
	migrateEntitesDB(gdb)
	router := setupDomainsWithRouter(gdb)
	server := &http.Server{Addr: ":" + config.Port, Handler: router}

	// graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Println("server started on port", config.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("server error:", err)
		}
	}()

	<-stop
	log.Println("shutting down...")

	ctxTimeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctxTimeout); err != nil {
		log.Println("shutdown error:", err)
	}
	log.Println("server stopped")
	return nil
}

func loadConfig() *config.Config {
	config := config.Load()
	if config.DSN == "" {
		log.Fatal("DB_DSN is empty")
	}
	return &config
}

func migrateEntitesDB(gdb *gorm.DB) {
	if err := gdb.AutoMigrate(&buildings.Building{}); err != nil {
		log.Fatal("auto-migrate:", err)
	}
}

func setupDomainsWithRouter(gdb *gorm.DB) *gin.Engine {
	buildingService := &buildings.Service{DB: gdb}
	buildingHandlers := &httpx.BuildingHandlers{Service: buildingService}

	userService := &users.Service{DB: gdb}
	userHandlers := &httpx.UserHandlers{Service: userService}

	r := httpx.NewRouter(httpx.Deps{Buildings: buildingHandlers, Users: userHandlers})
	return r
}
