package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hexhoc/go-mall-api/config"
	"github.com/hexhoc/go-mall-api/pkg/httpserver"
	"github.com/hexhoc/go-mall-api/pkg/logger"
	"os"
	"os/signal"
	"syscall"
)

func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	// HTTP Server
	gin.SetMode(gin.ReleaseMode)
	handler := gin.New()
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	var err error

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}

}
