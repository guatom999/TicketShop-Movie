package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/guatom999/TicketShop-Movie/config"
	"github.com/guatom999/TicketShop-Movie/modules/middlewares/middlewareHandlers"
	"github.com/guatom999/TicketShop-Movie/modules/middlewares/middlewareRepositories"
	"github.com/guatom999/TicketShop-Movie/modules/middlewares/middlewareUseCases"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/omise/omise-go"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/gomail.v2"
)

type (
	Server interface {
		Start(pctx context.Context)
	}

	server struct {
		app    *echo.Echo
		db     *mongo.Client
		cfg    *config.Config
		omise  *omise.Client
		redis  *redis.Client
		mailer *gomail.Dialer
		// middleware middlewareHandlers.MiddlewareHandlerInterface
	}
)

func NewMiddleware(cfg *config.Config) middlewareHandlers.MiddlewareHandlerInterface {
	middlwareRepository := middlewareRepositories.NewMiddlewareRepository()
	middlewareUseCase := middlewareUseCases.NewMiddlwareUseCase(middlwareRepository, cfg)
	middlewareHandlers := middlewareHandlers.NewMiddlewareHandler(middlewareUseCase)

	return middlewareHandlers

}

func NewEchoServer(
	db *mongo.Client,
	cfg *config.Config,
	omise *omise.Client,
	redis *redis.Client,
	mailer *gomail.Dialer,
) Server {
	return &server{
		app:    echo.New(),
		db:     db,
		cfg:    cfg,
		omise:  omise,
		redis:  redis,
		mailer: mailer,
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
	authMiddleware := NewMiddleware(s.cfg)

	// s = &server{
	// 	middleware: authMiddleware,
	// }

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
	case "inventory":
		s.InventoryModule()
	case "payment":
		s.PaymentModule()
	case "customer":
		s.CustomerModules(authMiddleware)
	}

	if err := s.app.Start(s.cfg.App.Port); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Failed to shutdown:%v", err)

	}

}
