package main

import (
	"fmt"
	"go-api-template/cmd"
	"go-api-template/internal/logs"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
)

func main() {
	cfg := zap.NewDevelopmentConfig()
	cfg.DisableStacktrace = true
	cfg.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	logger, err := cfg.Build()

	if err != nil {
		fmt.Printf("failed to build logger: %v\n", err)
		os.Exit(1)
	}

	zap.ReplaceGlobals(logger)

	code := mainWithReturnCode()
	logger.Sync()
	if code != 0 {
		os.Exit(code)
	}
}

func mainWithReturnCode() int {
	port, exists := os.LookupEnv("PORT")

	if !exists {
		port = "3000"
	}

	address, exists := os.LookupEnv("ADDRESS")

	if !exists {
		address = "localhost"
	}

	sc := make(chan os.Signal, 1)

	server := cmd.NewServer()

	err := server.Start(address, port)

	if err != nil {
		zap.S().Warnf("Problems occurred when start server. Error: %v", err)
		return 1
	}

	zap.S().Info("Server is running. Ctrl-C to exit")
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
	zap.S().Info("Server is terminating!")

	err = server.Shutdown()

	if err == nil {
		zap.S().Info("Server terminated!")
		return 0
	}

	logs.LogErrorIfNotNil(err, "Failed to shutdown API server: %v")
	return 1
}
