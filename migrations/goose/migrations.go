package migrations

import (
	"context"
	"fmt"
	"time"

	"github.com/pressly/goose/v3"
)

func New(provider *goose.Provider) *Migrator {

	return &Migrator{provider: provider}
}

type Migrator struct {
	provider *goose.Provider
}

func (mt *Migrator) Stairway(stepTimeout time.Duration) error {
	for range mt.provider.ListSources() {
		ctx, done := context.WithTimeout(context.Background(), stepTimeout)
		defer done()

		if err := mt.Step(ctx); err != nil {
			return err
		}
	}

	return nil
}

func (mt *Migrator) Step(ctx context.Context) error {
	if _, err := mt.provider.UpByOne(ctx); err != nil {
		return fmt.Errorf("cannot make first up step : %w", err)

	}

	if _, err := mt.provider.Down(ctx); err != nil {
		return fmt.Errorf("cannot make down step : %w", err)
	}

	if _, err := mt.provider.UpByOne(ctx); err != nil {
		return fmt.Errorf("cannot make second up step : %w", err)
	}

	return nil
}
