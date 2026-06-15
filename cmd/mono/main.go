package main

import (
	"context"
	"fmt"
	"log"
	"mono/internal/config"
	"mono/internal/runner"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	app := &cli.Command{
		Name:  "mono",
		Usage: "Language-agnostic dev runner for monorepos",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Usage:   "Path to the configuration file",
				Value:   "mono.yaml",
			},
			&cli.StringSliceFlag{
				Name:    "only",
				Usage:   "Run only the specified tasks (space-separated)",
				Aliases: []string{"o"},
			},
		},
		Action: run,
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

}

func run(ctx context.Context, c *cli.Command) error {
	cfgPath := c.String("config")
	onlyTasks := c.StringSlice("only")

	fmt.Printf("Loading configuration from %s", cfgPath)
	cfg, err := config.Load(cfgPath)
	if err != nil {
		return fmt.Errorf("failed to load configuration: %v", err)
	}

	if len(onlyTasks) > 0 {
		fmt.Printf("Running only tasks: %v", onlyTasks)
		// Filter cfg.Services based on onlyTasks
		cfg = config.FilterServices(cfg, onlyTasks)
	}
	r := runner.NewRunner(cfg)
	return r.Start(ctx)
}
