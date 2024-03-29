# Drillerfy: Go Database Testing Simplified

[![CI](https://github.com/Alviner/drillerfy/actions/workflows/ci.yml/badge.svg)](https://github.com/Alviner/drillerfy/actions/workflows/ci.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/Alviner/drillerfy.svg)](https://pkg.go.dev/github.com/Alviner/drillerfy)
[![Go Report Card](https://goreportcard.com/badge/github.com/Alviner/drillerfy)](https://goreportcard.com/report/github.com/Alviner/drillerfy)

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

Inspired by [sqlalchemy-utils](https://github.com/kvesteri/sqlalchemy-utils)

[Example](examples/tempdb/main.go)

### Migrations Module

Provides functionality to easily run stairway tests for migrations via goose Provider.
This module simplifies the process of applying and reverting database schema changes,
which is essential in maintaining consistent database states for testing.

[Example](examples/migoose/main.go)

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

## Repo activity

![Alt](https://repobeats.axiom.co/api/embed/480a84c4a38844d8a0a6524f1a5665d3706d96d2.svg "Repobeats analytics image")
