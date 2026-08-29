package retry

import (
	"context"
	"errors"
	"time"
)

var ErrMaxAttempts = errors.New("max attempts exceeded")

type Config struct {
	MaxAttempts int
	BaseDelay   time.Duration
	MaxDelay    time.Duration
}

func DefaultConfig() Config {
	return Config{MaxAttempts: 5, BaseDelay: 100 * time.Millisecond, MaxDelay: 2 * time.Second}
}

func Do(ctx context.Context, cfg Config, fn func() error) error {
	if cfg.MaxAttempts < 1 {
		cfg.MaxAttempts = 1
	}
	delay := cfg.BaseDelay
	var err error
	for attempt := 1; attempt <= cfg.MaxAttempts; attempt++ {
		err = fn()
		if err == nil {
			return nil
		}
		if attempt == cfg.MaxAttempts {
			break
		}
		timer := time.NewTimer(delay)
		select {
		case <-ctx.Done():
			timer.Stop()
			return ctx.Err()
		case <-timer.C:
		}
		delay *= 2
		if delay > cfg.MaxDelay {
			delay = cfg.MaxDelay
		}
	}
	return errors.Join(ErrMaxAttempts, err)
}
