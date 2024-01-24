package migrations

import (
	"context"
	"testing"
	"time"

	"github.com/pressly/goose/v3"
)

func New(t *testing.T, provider *goose.Provider) *MT {
	t.Helper()

	return &MT{t: t, provider: provider}
}

type MT struct {
	t        *testing.T
	provider *goose.Provider
}

func (mt *MT) Stairway(stepTimeout time.Duration) error {
	for range mt.provider.ListSources() {
		ctx, done := context.WithTimeout(context.Background(), stepTimeout)
		defer done()

		if err := mt.Step(ctx); err != nil {
			mt.t.Logf("Cannot make step: %s", err)

			return err
		}
	}

	return nil
}

func (mt *MT) Step(ctx context.Context) error {
	if res, err := mt.provider.UpByOne(ctx); err != nil {
		mt.t.Logf("Cannot make first up step: %s", res)

		return err
	}

	if res, err := mt.provider.Down(ctx); err != nil {
		mt.t.Logf("Cannot make down step: %s", res)

		return err
	}

	if res, err := mt.provider.UpByOne(ctx); err != nil {
		mt.t.Logf("Cannot make second up step: %s", res)

		return err
	}

	return nil
}
