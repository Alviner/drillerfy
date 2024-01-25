package tempdb

import (
	"context"
	"fmt"
	"net/url"

	"github.com/Alviner/drillerfy/tempdb/postgres"
)

// Querier is the interface that wraps the basic methods to create a dialect
// specific queries.
type Querier interface {
	CreateDatabase(ctx context.Context, name, temlate string) error
	DeleteDatabase(ctx context.Context, name string) error
}

func New(dbUrl string) (Querier, error) {
	parsedUrl, err := url.Parse(dbUrl)
	if err != nil {
		return nil, err
	}
	querier, err := querierFromUrl(parsedUrl)
	if err != nil {
		return nil, fmt.Errorf("cannot create querier: %w", err)
	}
	return querier, nil

}

func querierFromUrl(url *url.URL) (Querier, error) {
	dialect := url.Scheme

	switch dialect {
	case "postgresql":
		q, err := postgres.New(url)
		if err != nil {
			return nil, fmt.Errorf("cannot create postgres: %w", err)
		}
		return q, nil
	default:
		return nil, fmt.Errorf("unknown dialect: %q", dialect)
	}
}
