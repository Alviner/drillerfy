package postgres

import (
	"context"
	"database/sql"
	"net/url"
	"testing"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func DBNames(t *testing.T, url *url.URL) (map[string]bool, error) {
	t.Helper()
	names := make(map[string]bool)

	db, err := sql.Open("pgx", url.String())
	if err != nil {
		return names, err
	}
	ctx, done := context.WithTimeout(context.Background(), 2*time.Second)
	defer done()

	rows, err := db.QueryContext(ctx, "SELECT datname FROM pg_catalog.pg_database")
	if err != nil {
		return names, err
	}
	defer func() {
		if err := rows.Close(); err != nil {
			t.Errorf("failed to close rows: %v", err)
		}
	}()
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return names, err
		}
		names[name] = true
	}
	return names, nil
}
