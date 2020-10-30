package main

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/mellaught/ethereum-blocks/src/app"
	"github.com/mellaught/ethereum-blocks/src/config"
	"github.com/sirupsen/logrus"
)

func main() {
	cfg := config.NewViperConfig()
	srvConfig, explorerConfig := cfg.ReadServiceConfig(), cfg.ReadExplorerConfig()
	// init logrus logger
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: time.RFC822,
	})

	// set logger level
	level, err := logrus.ParseLevel(cfg.GetString("logger-level"))
	if err != nil {
		panic(err)
	}
	logger.SetLevel(level)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		sign := <-c
		logger.Infoln("System signal: %v", sign)
		cancel()
	}()

	app := app.NewApp(logger, explorerConfig.URL)
	//run App
	app.Run(ctx, srvConfig.Host+":"+srvConfig.Port)
	logger.Infoln("Ethereum block scanner service has started on :%s\nPress ctrl + C to exit.", srvConfig.Port)
}
