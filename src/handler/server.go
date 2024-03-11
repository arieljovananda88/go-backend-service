package handler

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"go-backend-service/graduit-be/src/middleware"

	"github.com/labstack/echo"
)

func (s *Service) InitializeRoutes() *echo.Echo {
	e := echo.New()
	internalGroup := e.Group("api/internal")
	internalGroup.Use(middleware.BasicAuth())
	s.InternalSystemLogHandler.MountInternal(internalGroup)

	//belum ada middleware
	adminGroup := e.Group("api/admin")
	adminGroup.Use(middleware.BearerAuth("TIMTA"))
	s.AdminSytemLogHandler.MountAdmin(adminGroup)
	return e
}

func (s *Service) StartServer() {
	server := s.InitializeRoutes()
	listenerPort := fmt.Sprintf(":%v", 8080)
	if err := server.StartServer(&http.Server{
		Addr:         listenerPort,
		ReadTimeout:  120 * time.Second,
		WriteTimeout: 120 * time.Second,
	}); err != nil {
		server.Logger.Fatal(err.Error())
	}
}

func (s *Service) ShutdownServer() {
	server := s.InitializeRoutes()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		server.Logger.Fatal(err.Error())
	}
}
