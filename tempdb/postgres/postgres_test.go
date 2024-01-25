package postgres

import (
	"context"
	"fmt"
	"net/url"
	"testing"
	"time"

	"github.com/Alviner/drillerfy/utils_test"
	"github.com/Alviner/drillerfy/utils_test/postgres"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPostgres(t *testing.T) {
	t.Parallel()

	require := require.New(t)
	assert := assert.New(t)

	dbUrl, err := url.Parse(utils_test.PostgresDNS(t))
	require.NoError(err)

	t.Run("create", func(t *testing.T) {
		t.Parallel()
		// arrange
		dbName := t.Name()
		ctx, done := context.WithTimeout(context.Background(), 2*time.Second)
		defer done()

		pg, err := New(dbUrl)
		require.NoError(err)
		//act
		require.NoError(pg.CreateDatabase(ctx, dbName, ""))
		defer func() { require.NoError(pg.DeleteDatabase(ctx, dbName)) }()
		// assert
		names, err := postgres.DBNames(t, dbUrl)
		require.NoError(err)
		assert.Contains(names, dbName)
	})
	t.Run("create from template", func(t *testing.T) {
		t.Parallel()
		// arrange
		dbName := t.Name()
		templateName := fmt.Sprintf("%s-template", t.Name())
		ctx, done := context.WithTimeout(context.Background(), 5*time.Second)
		defer done()

		pg, err := New(dbUrl)
		require.NoError(err)

		require.NoError(pg.CreateDatabase(ctx, templateName, ""))
		defer func() { require.NoError(pg.DeleteDatabase(ctx, templateName)) }()
		//act
		require.NoError(pg.CreateDatabase(ctx, dbName, templateName))
		defer func() { require.NoError(pg.DeleteDatabase(ctx, dbName)) }()

		// assert
		names, err := postgres.DBNames(t, dbUrl)
		require.NoError(err)
		assert.Contains(names, dbName)
	})

	t.Run("clear", func(t *testing.T) {
		t.Parallel()
		// arrange
		dbName := t.Name()
		ctx, done := context.WithTimeout(context.Background(), 5*time.Second)
		defer done()

		pg, err := New(dbUrl)
		require.NoError(err)

		require.NoError(pg.CreateDatabase(ctx, dbName, ""))
		//act
		require.NoError(pg.DeleteDatabase(ctx, dbName))

		// assert
		names, err := postgres.DBNames(t, dbUrl)
		require.NoError(err)
		assert.NotContains(names, dbName)
	})
}
