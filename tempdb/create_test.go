package tempdb

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/Alviner/drillerfy/tempdb/postgres"
)

func TestCreate(t *testing.T) {
	t.Parallel()

	require := require.New(t)
	assert := assert.New(t)

	t.Run("postgresql", func(t *testing.T) {
		t.Parallel()
		// act
		db, err := New("postgresql://pguser:pgpass@localhost:5432/pgdb")
		require.NoError(err)
		// assert
		assert.IsType(new(postgres.Postgres), db)
	})
	t.Run("postgres", func(t *testing.T) {
		t.Parallel()
		// act
		db, err := New("postgres://pguser:pgpass@localhost:5432/pgdb")
		require.NoError(err)
		// assert
		assert.IsType(new(postgres.Postgres), db)
	})
	t.Run("unknown", func(t *testing.T) {
		t.Parallel()
		// act
		_, err := New("unknown://database.url")
		// assert
		require.EqualErrorf(
			err,
			"cannot create querier: unknown dialect: \"unknown\"",
			"cannot create querier: unknown dialect: %q",
			"unknown",
		)
	})
}
