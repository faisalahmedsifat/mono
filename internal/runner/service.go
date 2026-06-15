package runner

import (
	"context"
	"mono/internal/config"
)

type Runner struct {
	config *config.Config
}

func NewRunner(cfg *config.Config) *Runner {
	return &Runner{
		config: cfg,
	}
}

func (r *Runner) Start(ctx context.Context) error {
	//TODO
	return nil
}
