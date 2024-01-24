package database

import (
	"context"
	"database/sql"
	"os"
	"testing"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func dns(t *testing.T) string {
	t.Helper()

	value, ok := os.LookupEnv("TEST_PG_DNS")
	if !ok {
		return "postgresql://pguser:pgpass@localhost:5432/pgdb"
	}
	return value
}

func dbNames(t *testing.T, db *sql.DB) (map[string]bool, error) {
	t.Helper()

	ctx, done := context.WithTimeout(context.Background(), 2*time.Second)
	defer done()
	names := make(map[string]bool)

	rows, err := db.QueryContext(ctx, "SELECT datname FROM pg_catalog.pg_database")
	if err != nil {
		return names, err
	}
	defer rows.Close()
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return names, err
		}
		names[name] = true
	}
	return names, nil

}

func TestDatabase(t *testing.T) {
	t.Parallel()

	require := require.New(t)
	assert := assert.New(t)

	t.Run("create", func(t *testing.T) {
		t.Parallel()
		ctx, done := context.WithTimeout(context.Background(), 2*time.Second)
		defer done()

		conn, err := sql.Open("pgx", dns(t))
		require.NoError(err)
		database, err := New(WithDialect(DialectPostgres))
		require.NoError(err)
		dbName, closer, err := database.Create(ctx, "test", conn)
		require.NoError(err)
		defer closer()

		names, err := dbNames(t, conn)
		require.NoError(err)

		assert.Contains(names, dbName)

	})
	t.Run("create from template", func(t *testing.T) {
		t.Parallel()
		ctx, done := context.WithTimeout(context.Background(), 5*time.Second)
		defer done()

		conn, err := sql.Open("pgx", dns(t))
		require.NoError(err)
		template, err := New(WithDialect(DialectPostgres))
		require.NoError(err)
		templateName, templateCloser, err := template.Create(ctx, "template", conn)
		require.NoError(err)
		defer templateCloser()

		database, err := New(WithDialect(DialectPostgres), WithTemplate(templateName))
		require.NoError(err)
		dbName, closer, err := database.Create(ctx, "test", conn)
		require.NoError(err)
		defer closer()

		names, err := dbNames(t, conn)
		require.NoError(err)

		assert.Contains(names, dbName)

	})

	t.Run("clear", func(t *testing.T) {
		t.Parallel()
		ctx, done := context.WithTimeout(context.Background(), 2*time.Second)
		defer done()

		conn, err := sql.Open("pgx", dns(t))
		require.NoError(err)
		database, err := New(WithDialect(DialectPostgres))
		require.NoError(err)
		dbName, closer, err := database.Create(ctx, "test", conn)
		require.NoError(err)

		require.NoError(closer())

		names, err := dbNames(t, conn)
		require.NoError(err)
		assert.NotContains(names, dbName)
	})
}
