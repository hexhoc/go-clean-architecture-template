package app

import (
	"fmt"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/hexhoc/go-mall-api/config"
	handler "github.com/hexhoc/go-mall-api/internal/controller/http/v1"
	"github.com/hexhoc/go-mall-api/internal/controller/middleware"
	"github.com/hexhoc/go-mall-api/internal/repository"
	"github.com/hexhoc/go-mall-api/internal/usecase/book"
	"github.com/hexhoc/go-mall-api/internal/usecase/loan"
	"github.com/hexhoc/go-mall-api/internal/usecase/user"
	"github.com/hexhoc/go-mall-api/pkg/datasource"
	"github.com/hexhoc/go-mall-api/pkg/httpserver"
	"github.com/hexhoc/go-mall-api/pkg/logger"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func Run(cfg *config.Config) {
	log := logger.New(cfg.Log.Level)
	var err error

	// Create connection to db
	db, err := datasource.NewPostgresConnection(&cfg.Datasource)
	if err != nil {
		log.Fatal(fmt.Errorf("app - Run - datasource.New: %w", err))
	}
	defer db.Close()

	// repository
	bookRepo := repository.NewBookPostgres(db)
	bookService := book.NewService(bookRepo)

	userRepo := repository.NewUserPostgres(db)
	userService := user.NewService(userRepo)

	loanUseCase := loan.NewService(userService, bookService)

	//metricService, err := metric.NewPrometheusService()
	//if err != nil {
	//	log.Fatal(err.Error())
	//}

	r := mux.NewRouter()
	//handlers
	n := negroni.New(
		negroni.HandlerFunc(middleware.Cors),
		//negroni.HandlerFunc(middleware.Metrics(metricService)),
		negroni.NewLogger(),
	)
	// HTTP Server
	//book
	handler.MakeBookHandlers(r, *n, bookService)

	//user
	handler.MakeUserHandlers(r, *n, userService)

	//loan
	handler.MakeLoanHandlers(r, *n, bookService, userService, loanUseCase)

	http.Handle("/", r)
	http.Handle("/metrics", promhttp.Handler())
	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	httpServer := httpserver.New(context.ClearHandler(http.DefaultServeMux), httpserver.Port(cfg.HTTP.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		log.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		log.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}

}
