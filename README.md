# pgx-pg-uint128

**pgx-pg-uint128** adds support for the [pg-uint128](https://github.com/codercms/pg-uint128) PostgreSQL extension in
the Go [pgx](https://github.com/jackc/pgx) driver.

## Requirements

- PostgreSQL 12+ with the [pg-uint128](https://github.com/codercms/pg-uint128) extension installed
- `pgx` driver version 5.6.0+
- `Go` version 1.21+

## Features

- **New Datatypes for pgx**:
    - `uint2` (maps to Go's `uint16`)
    - `uint4` (maps to Go's `uint32`)
    - `uint8` (maps to Go's `uint64`)
    - `uint16` (maps to `uint128` via [this package](https://pkg.go.dev/lukechampine.com/uint128) to emulate 128-bit
      unsigned integers in Go)
    - `int16` (maps to `int128` via [this package](https://pkg.go.dev/go.shabbyrobe.org/num) to emulate 128-bit
      signed integers in Go)

- **Derivative Type Support**:
    - Arrays
    - Ranges
    - Multi-ranges

- **Efficient Encoding/Decoding**:
    - Full support for both binary and text protocols

- **Flexible Scanning**:
    - `zeronull` types support
    - Automatically scans new types into Go's standard integer types.
    - Includes safeguards for overflow and underflow, returning errors when necessary

## Installation

To add `pgx-pg-uint128` to your Go project, run:

```sh
go get github.com/pg-uint/pgx-pg-uint128
```

## Usage

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"

	"github.com/pg-uint/pgx-pg-uint128/types"
	"github.com/pg-uint/pgx-pg-uint128/types/zeronull"
)

func main() {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	defer conn.Close(context.Background())

	if _, err := types.RegisterAll(context.Background(), conn); err != nil {
		log.Fatalf("Unable to register types: %v", err)
	}

	// Optionally register zeronull types
	zeronull.Register(conn.TypeMap())

	// Do regular work with connection
}
```

- See this example for pgx: [example](examples/pgx/main.go)
- See these examples for pgxpool:
    - [Simple pgxpool usage](examples/pgxpool/simple/main.go)
    - [Caching types between connections](examples/pgxpool/cache/main.go)
