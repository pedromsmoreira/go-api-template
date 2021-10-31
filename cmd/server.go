package cmd

import (
	"fmt"
	"go-api-template/internal/logs"
	"go-api-template/internal/templates"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

// Server struct, should have the needed dependencies
type Server struct {
	app *fiber.App
	wg  sync.WaitGroup
}

// NewServer method to init a new server instance
func NewServer() *Server {
	app := fiber.New(fiber.Config{
		ReadTimeout: 5 * time.Second, // to make sure keep alive connections are closed on Shutdown()
	})

	app.Use(cors.New())
	app.Use(logger.New(logger.Config{
		Format: "${time} [${status}] ${method} ${path} (${latency})",
		Output: &logs.FiberLogWriter{},
	}))
	app.Get("/dashboard", monitor.New())
	app.Get("/ping", func(ctx *fiber.Ctx) error {
		return ctx.Send([]byte("ping"))
	})

	return &Server{
		app: app,
		wg:  sync.WaitGroup{},
	}
}

// Start method to start the server configurations
func (s *Server) Start(address string, port string) error {
	var err error
	tpltSvc := templates.NewService()
	api := s.app.Group("/api")
	templates.AddRoutesTo(api, tpltSvc)

	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		zap.S().Infof("API Server listening on %v:%v", address, port)
		err := s.app.Listen(fmt.Sprintf("%v:%v", address, port))
		if err != nil {
			zap.S().Errorf("API Server stopped listening due to %v")
		} else {
			zap.S().Info("API Server stopped listening")
		}
	}()

	return err
}

func (s *Server) Shutdown() error {
	err := s.app.Shutdown()

	if err == nil {
		s.wg.Wait()
	}

	return err
}
