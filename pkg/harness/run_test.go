package harness

import (
	"context"
	"errors"
	"testing"
)

func TestRunnerRun(t *testing.T) {
	r := &Runner{Model: func(ctx context.Context, prompt string) (string, error) {
		return "echo: " + prompt, nil
	}}
	out, err := r.Run(context.Background(), "hi")
	if err != nil || out != "echo: hi" {
		t.Fatalf("got %q %v", out, err)
	}
}

func TestRunnerCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	r := &Runner{Model: func(ctx context.Context, prompt string) (string, error) {
		return "x", nil
	}}
	_, err := r.Run(ctx, "hi")
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("expected canceled, got %v", err)
	}
}

func TestRunnerNil(t *testing.T) {
	var r *Runner
	_, err := r.Run(context.Background(), "hi")
	if err == nil {
		t.Fatal("expected error")
	}
}
