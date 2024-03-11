package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/guatom999/TicketShop-Movie/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	Server interface {
		Start(pctx context.Context)
	}

	server struct {
		app *echo.Echo
		db  *mongo.Client
		cfg *config.Config
	}
)

func NewEchoServer(
	db *mongo.Client,
	cfg *config.Config,
) Server {
	return &server{
		app: echo.New(),
		db:  db,
		cfg: cfg,
	}
}

func (s *server) gracefulShutdown(pctx context.Context, close <-chan os.Signal) {

	<-close

	ctx, cancel := context.WithTimeout(pctx, time.Second*20)
	defer cancel()

	if err := s.app.Shutdown(ctx); err != nil {
		log.Fatalf("Failed to shutdown server...")
		panic(err)
	}

	log.Println("Shutting Down Server......")

}

func (s *server) Start(pctx context.Context) {

	// Request Timeout
	s.app.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Skipper:      middleware.DefaultSkipper,
		ErrorMessage: "Error: Request Timeout",
		Timeout:      time.Second * 10,
	}))

	//Cors
	s.app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper:      middleware.DefaultSkipper,
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.PATCH},
	}))

	s.app.Use(middleware.Logger())

	close := make(chan os.Signal, 1)
	signal.Notify(close, syscall.SIGINT, syscall.SIGTERM)

	go s.gracefulShutdown(pctx, close)

	log.Println("Starting server...")

	switch s.cfg.App.Name {
	case "movie":
		s.MovieModule()
	case "booking":
		s.BookingModule()
	case "ticket":
		s.TicketModule()
	}

	if err := s.app.Start(fmt.Sprintf(":%d", s.cfg.App.Port)); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Failed to shutdown:%v", err)

	}

}
