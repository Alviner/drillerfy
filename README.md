# Drillerfy: Go Database Testing Simplified

## Overview

Drillerfy is a Go package designed to simplify database testing.
It provides a streamlined approach for setting up and tearing down databases,
making it easier to test Go applications that interact with databases.

## Features

- Easy setup and teardown of databases.
- Easy migrations stairway tests.
- Support for multiple database engines.

## Installation

```(bash)
go get github.com/Alviner/drillerfy
```

## Usage

### Database Module

Provides functionality to easily create and drop databases.
This is particularly useful in testing environments where you need to set up a fresh database instance for each test run and clean it up afterward.

```(go)
package main

import (
    "context"
    "database/sql"
    "log"

    _ "github.com/jackc/pgx/v5/stdlib"
    "github.com/Alviner/drillerfy/database"

)

func main() {
    conn, err := sql.Open("pgx", "database dns")
    if err != nil {
        log.Fatal(err)
    }
    database, err := database.New(database.WithDialect(database.DialectPostgres))
    if err != nil {
        log.Fatal(err)
    }

    dbName, closer, err := database.Create(ctx, "test", conn)
    if err != nil {
        log.Fatal(err)
    }
    defer closer()
    // ... some useful staff with created db
}

```

### Migrations Module

Provides functionality to easily run stairway tests for migrations via goose Provider.
This module simplifies the process of applying and reverting database schema changes,
which is essential in maintaining consistent database states for testing.

```(go)
package main

import (
    "context"
    "database/sql"
    "log"
    "time"

    _ "github.com/jackc/pgx/v5/stdlib"

    "github.com/pressly/goose/v3"
    "github.com/Alviner/drillerfy/migrations"

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
    migrator := migrations.New(provider)

    if err := migrations.Stairway(2 * time.Second); err != nil {
        log.Fatal(err)
    }
    // ... some useful staff with created db
}

```

## Contributing

Contributions to Drillerfy are welcome.
Please read the contributing guidelines in the repository
for instructions on how to submit pull requests, report issues, and suggest enhancements.

## License

Drillerfy is released under the MIT License.
See the LICENSE file in the repository for full license text.

## Authors

Drillerfy was created and is maintained by [Alviner](https://github.com/Alviner).
Contributions from the community are appreciated.
