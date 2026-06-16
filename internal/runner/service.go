package runner

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"sync"

	"mono/internal/config"

	"github.com/fatih/color"
)

type Runner struct {
	config *config.Config
}

func NewRunner(cfg *config.Config) *Runner {
	return &Runner{config: cfg}
}

func (r *Runner) Start(ctx context.Context) error {
	var wg sync.WaitGroup
	errCh := make(chan error, len(r.config.Services))

	fmt.Printf("Starting %d services...\n\n", len(r.config.Services))

	for name, svc := range r.config.Services {
		wg.Add(1)
		go func(name string, svc config.Service) {
			defer wg.Done()
			if err := r.runService(ctx, name, svc); err != nil {
				errCh <- fmt.Errorf("%s: %w", name, err)
			}
		}(name, svc)
	}

	// Wait for all services
	go func() {
		wg.Wait()
		close(errCh)
	}()

	// Block until shutdown or error
	select {
	case <-ctx.Done():
		fmt.Println("\nShutting down all services...")
		return nil
	case err := <-errCh:
		if err != nil {
			return err
		}
	}
	return nil
}

// runService starts and manages one service
func (r *Runner) runService(ctx context.Context, name string, svc config.Service) error {
	cmd := exec.CommandContext(ctx, "sh", "-c", svc.Command)
	if svc.Dir != "" {
		cmd.Dir = svc.Dir
	}

	// Add environment variables
	if len(svc.Env) > 0 {
		env := os.Environ()
		for k, v := range svc.Env {
			env = append(env, fmt.Sprintf("%s=%s", k, v))
		}
		cmd.Env = env
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("stdout pipe: %w", err)
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return fmt.Errorf("stderr pipe: %w", err)
	}

	prefix := color.New(color.FgCyan).Sprint("[" + name + "]")

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start: %w", err)
	}

	fmt.Printf("%s Started (PID: %d)\n", prefix, cmd.Process.Pid)

	var outputWg sync.WaitGroup
	outputWg.Add(2)

	go streamOutput(&outputWg, stdout, prefix)
	go streamOutput(&outputWg, stderr, prefix)

	outputWg.Wait()

	return cmd.Wait()
}

// streamOutput reads from io.Reader (stdout/stderr) and prints with prefix
func streamOutput(wg *sync.WaitGroup, reader io.Reader, prefix string) {
	defer wg.Done()
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Printf("%s %s\n", prefix, scanner.Text())
	}
}
