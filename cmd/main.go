package main

import (
	"context"
	"os"
	"time"

	"github.com/urfave/cli"

	"github.com/yashwantsinghcode/go_backend/api"
	"github.com/yashwantsinghcode/go_backend/constants"
	"github.com/yashwantsinghcode/go_backend/models"
)

func main() {
	app := &cli.App{
		Name:   constants.APP_NAME,
		Usage:  constants.APP_USAGE,
		Action: RunServer,
	}

	err := app.Run(os.Args)
	if err != nil {
		os.Exit(1)
	}
}

func RunServer(c *cli.Context) error {
	// Request Setup
	ctx := context.Background()
	_, cancel := context.WithTimeout(ctx, time.Second*90)
	defer cancel()

	// Setup configs
	// To Do: get the address from the flag
	cfg := models.NewConfig("0.0.0.0:8080")

	// Setup Database
	// To Do

	// Initialize Service
	serviceV1 := api.NewService(cfg)
	// Start Http server
	serviceV1.RunHttpServer()

	return nil
}
