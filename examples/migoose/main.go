package main

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/Alviner/drillerfy/migoose"
	"github.com/pressly/goose/v3"
)

func main() {
	db, err := sql.Open("pgx", "database dns")
	if err != nil {
		log.Fatal(err)
	}
	provider, err := goose.NewProvider(
		goose.DialectPostgres,
		db,
		os.DirFS("migrations"),
	)
	if err != nil {
		log.Fatal(err)
	}
	migrator := migoose.New(provider)

	if err := migrator.Stairway(2 * time.Second); err != nil {
		log.Fatal(err)
	}
}
