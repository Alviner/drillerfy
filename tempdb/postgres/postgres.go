package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"net/url"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type Postgres struct {
	db *sql.DB
}

func (p *Postgres) CreateDatabase(ctx context.Context, name, templateName string) error {
	query := fmt.Sprintf(`CREATE DATABASE "%s"`, name)
	if templateName != "" {
		query = fmt.Sprintf(`CREATE DATABASE "%s" TEMPLATE "%s"`, name, templateName)
	}
	if _, err := p.db.ExecContext(ctx, query); err != nil {
		return fmt.Errorf(
			"failed to create database %q from template %q: %w",
			name,
			templateName,
			err,
		)
	}
	return nil
}

func (p *Postgres) DeleteDatabase(ctx context.Context, name string) error {
	if err := p.DisconnectFromDatabase(ctx, name); err != nil {
		return err
	}

	query := fmt.Sprintf(`DROP DATABASE "%s"`, name)
	if _, err := p.db.ExecContext(ctx, query); err != nil {
		return fmt.Errorf(
			"failed to delete database %q: %w",
			name,
			err,
		)
	}
	return nil
}

func (p *Postgres) DisconnectFromDatabase(ctx context.Context, name string) error {
	query := fmt.Sprintf(
		`SELECT pg_terminate_backend(pg_stat_activity.pid)
          FROM pg_stat_activity
          WHERE pg_stat_activity.datname = '%s'
          AND pg_stat_activity.pid <> pg_backend_pid()`, name)
	if _, err := p.db.ExecContext(ctx, query); err != nil {
		return fmt.Errorf(
			"failed to disconnect from database %q: %w",
			name,
			err,
		)
	}
	return nil
}

func New(url *url.URL) (*Postgres, error) {
	db, err := sql.Open("pgx", url.String())
	if err != nil {
		return nil, fmt.Errorf("cannot connect postgresql: %w", err)
	}
	ins := new(Postgres)
	ins.db = db
	return ins, nil
}
