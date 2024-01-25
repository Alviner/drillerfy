package migoose

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pressly/goose/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func newTxFn(query string) func(context.Context, *sql.Tx) error {
	return func(ctx context.Context, tx *sql.Tx) error {
		_, err := tx.ExecContext(ctx, query)

		return err
	}
}

func tempDB(t *testing.T) (*sql.DB, func()) {
	t.Helper()

	f, err := os.CreateTemp("", "test-*.db")
	if err != nil {
		t.Fatal(err)
	}

	db, err := sql.Open("sqlite3", fmt.Sprintf("file:%s", f.Name()))
	if err != nil {
		t.Fatal(err)
	}

	return db, func() {
		if err := db.Close(); err != nil {
			t.Fatal(err)
		}

		if err := os.Remove(f.Name()); err != nil {
			t.Fatal(err)
		}
	}
}

func tempProvider(t *testing.T, migrations ...*goose.Migration) (*goose.Provider, func(), error) {
	t.Helper()
	db, closer := tempDB(t)

	provider, err := goose.NewProvider(goose.DialectSQLite3, db, nil,
		goose.WithGoMigrations(migrations...),
	)
	if err != nil {
		return nil, nil, err
	}

	return provider, closer, nil
}

func TestMigrations(t *testing.T) {
	t.Parallel()
	require := require.New(t)
	assert := assert.New(t)

	register := []*goose.Migration{
		goose.NewGoMigration(
			1,
			&goose.GoFunc{RunTx: newTxFn("CREATE TABLE users (id INTEGER)")},
			&goose.GoFunc{RunTx: newTxFn("DROP TABLE users")},
		),
		goose.NewGoMigration(
			2,
			&goose.GoFunc{RunTx: newTxFn("ALTER TABLE users ADD COLUMN name text")},
			&goose.GoFunc{RunTx: newTxFn("ALTER TABLE users DROP COLUMN name")},
		),
	}

	t.Run("step", func(t *testing.T) {
		t.Parallel()
		// arrange
		provider, closer, err := tempProvider(t, register...)
		require.NoError(err)
		defer closer()

		migrations := New(provider)
		ctx, done := context.WithTimeout(context.Background(), 2*time.Second)
		defer done()

		// act
		require.NoError(migrations.Step(ctx))
		version, err := provider.GetDBVersion(ctx)
		require.NoError(err)
		// assert
		assert.Equal(register[0].Version, version)
	})
	t.Run("stairway", func(t *testing.T) {
		t.Parallel()
		// arrange
		provider, closer, err := tempProvider(t, register...)
		require.NoError(err)
		defer closer()

		ctx, done := context.WithTimeout(context.Background(), 2*time.Second)
		defer done()

		migrations := New(provider)

		// act
		require.NoError(migrations.Stairway(2 * time.Second))

		// assert
		version, err := provider.GetDBVersion(ctx)
		require.NoError(err)
		assert.Equal(register[len(register)-1].Version, version)
	})
}
