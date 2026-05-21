package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"23.usersapi/internal/config"
	"23.usersapi/internal/middleware"
	"23.usersapi/internal/routes"
	"23.usersapi/internal/validation"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Module interface {
	Routes() routes.Routers
}

type Application struct {
	config *config.Config
	router *gin.Engine
}

func NewApplication(cfg *config.Config) *Application {
	go middleware.CleanUpClient()
	validation.InitialValidation()
	r := gin.Default()

	modules := []Module{
		NewUserModule(),
	}

	routes.RegisterAllRouter(r, getRoutes(modules)...)

	return &Application{
		config: cfg,
		router: r,
	}
}

func (a *Application) Run() error {
	srv := &http.Server{
		Addr:    a.config.Port,
		Handler: a.router,
	}
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	//syscall.SIGINT: ctrl+C
	//syscall.SIGTERM: kill server
	//syscall.SIGHUP: reload server
	go func() {
		log.Println("Prepare to start server.")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %s", err.Error())
		}
	}()
	<-quit
	log.Println("shutdown signal received.")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server force to shutdown: %s", err.Error())
	}
	return nil
}

func getRoutes(modules []Module) []routes.Routers {
	routes := make([]routes.Routers, len(modules))
	for i, module := range modules {
		routes[i] = module.Routes()
	}
	return routes
}

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Cảnh báo: Không tìm thấy file .env, server sẽ dùng biến môi trường mặc định")
		return
	}
}
