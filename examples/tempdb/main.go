package main

import (
	"context"
	"log"
	"time"

	"github.com/Alviner/drillerfy/tempdb"
)

const (
	PG_DNS = "postgresql://pguser:pgpass@localhost:5432/pgdb"
)

func main() {
	db, err := tempdb.New(PG_DNS)
	if err != nil {
		log.Fatal(err)
	}
	ctx, done := context.WithTimeout(context.Background(), 2*time.Second)
	defer done()

	if err := db.CreateDatabase(ctx, "example", ""); err != nil {
		log.Fatal(err)
	}
	defer db.DeleteDatabase(ctx, "example")
	// .. useful staff

}
