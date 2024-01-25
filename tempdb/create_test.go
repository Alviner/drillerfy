package tempdb

import (
	"testing"

	"github.com/Alviner/drillerfy/tempdb/postgres"
	"github.com/Alviner/drillerfy/utils_test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	t.Parallel()

	require := require.New(t)
	assert := assert.New(t)

	t.Run("postgres", func(t *testing.T) {
		t.Parallel()
		//act
		db, err := New(utils_test.PostgresDNS(t))
		require.NoError(err)
		// assert
		assert.IsType(new(postgres.Postgres), db)
	})
	t.Run("unknown", func(t *testing.T) {
		t.Parallel()
		//act
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
