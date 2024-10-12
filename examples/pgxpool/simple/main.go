package main

import (
	"context"
	"log"
	"math"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"go.shabbyrobe.org/num"
	"lukechampine.com/uint128"

	"github.com/codercms/pgx-pg-uint128/types"
	"github.com/codercms/pgx-pg-uint128/types/zeronull"
)

func main() {
	poolCfg, err := pgxpool.ParseConfig(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Cannot create pgxpool config: %v", err)
	}

	poolCfg.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		_, err := types.RegisterAll(ctx, conn)

		// Optionally register zeronull types
		zeronull.Register(conn.TypeMap())

		return err
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), poolCfg)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	defer pool.Close()

	rows, err := pool.Query(
		context.Background(),
		"SELECT $1::uint2, $2::uint4, $3::uint8, $4::uint16, $5::int16, $6::int16",
		uint16(math.MaxUint16),
		uint32(math.MaxUint32),
		uint64(math.MaxUint64),
		uint128.Max,
		num.MaxI128,
		num.MinI128,
	)
	if err != nil {
		log.Fatalf("Cannot exec query: %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var num16 uint16
		var num32 uint32
		var num64 uint64
		var num128 uint128.Uint128
		var num128sMax num.I128
		var num128sMin num.I128

		if err := rows.Scan(&num16, &num32, &num64, &num128, &num128sMax, &num128sMin); err != nil {
			log.Fatalf("Cannot scan row: %v", err)
		}

		log.Printf("Got u16 %d, u32 %d, u64 %d, u128 %s", num16, num32, num64, num128.String())
		log.Printf("Got s128 max: %s, s128 min: %s", num128sMax.String(), num128sMin.String())
	}
}

// Correct output should look like this:
//
// Got u16 65535, u32 4294967295, u64 18446744073709551615, u128 340282366920938463463374607431768211455
// Got s128 max: 170141183460469231731687303715884105727, s128 min: -170141183460469231731687303715884105728
