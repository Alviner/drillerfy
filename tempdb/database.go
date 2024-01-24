package tempdb

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Alviner/drillerfy/tempdb/querier"
	"github.com/google/uuid"
)

type Dialect string
type DBOption func(*DB) error
type Closer func() error

const (
	DialectPostgres Dialect = "postgres"
)

type DB struct {
	templateName string
	querier      querier.Querier
}

func New(opts ...DBOption) (*DB, error) {
	store := new(DB)
	for _, option := range opts {
		if err := option(store); err != nil {
			return nil, fmt.Errorf("failed to apply option: %w", err)
		}
	}
	return store, nil
}

func WithDialect(dialect Dialect) DBOption {
	return func(s *DB) error {
		if dialect == "" {
			return errors.New("dialect must not be empty")
		}
		lookup := map[Dialect]querier.Querier{
			DialectPostgres: new(querier.Postgres),
		}
		querier, ok := lookup[dialect]
		if !ok {
			return fmt.Errorf("unknown dialect: %q", dialect)
		}
		s.querier = querier
		return nil
	}
}

func WithTemplate(templateName string) DBOption {
	return func(s *DB) error {
		s.templateName = templateName
		return nil
	}
}

func (s *DB) Create(ctx context.Context, prefix string, db *sql.DB) (string, Closer, error) {
	databaseName := GenerateName(prefix)

	q := s.querier.CreateDatabase(databaseName, s.templateName)
	if _, err := db.ExecContext(ctx, q); err != nil {
		return "", nil, fmt.Errorf(
			"failed to create database %q from template %q: %w",
			databaseName,
			s.templateName,
			err,
		)
	}
	return databaseName, func() error {
		ctx := context.Background()
		if err := s.disconnect(ctx, databaseName, db); err != nil {
			return err
		}
		return s.Delete(ctx, databaseName, db)
	}, nil
}

func (s *DB) Delete(ctx context.Context, name string, db *sql.DB) error {
	q := s.querier.DeleteDatabase(name)
	if _, err := db.ExecContext(ctx, q); err != nil {
		return fmt.Errorf("failed to delete database %s: %w", name, err)
	}
	return nil
}

func (s *DB) disconnect(ctx context.Context, name string, db *sql.DB) error {
	q := s.querier.DisconnectFromDatabase(name)
	if _, err := db.ExecContext(ctx, q); err != nil {
		return fmt.Errorf("failed to disconnect database %s: %w", name, err)
	}
	return nil
}

func GenerateName(prefix string) string {
	u := uuid.New()
	return fmt.Sprintf("%s-%s", prefix, u)

}
