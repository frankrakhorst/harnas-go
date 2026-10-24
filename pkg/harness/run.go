package harness

import (
	"context"
	"errors"
	"fmt"
)

// Runner executes a single harness turn.
type Runner struct {
	// Model is a stubbed model invocation; replace with real provider later.
	Model func(ctx context.Context, prompt string) (string, error)
}

// Run executes one turn with cancellation support.
func (r *Runner) Run(ctx context.Context, prompt string) (string, error) {
	if r == nil {
		return "", errors.New("harness: nil runner")
	}
	if r.Model == nil {
		return "", errors.New("harness: no model configured")
	}
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	default:
	}
	out, err := r.Model(ctx, prompt)
	if err != nil {
		return "", fmt.Errorf("harness: model: %w", err)
	}
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	default:
	}
	return out, nil
}
