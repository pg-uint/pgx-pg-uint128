package main

import (
	"context"
	"log"
	"math"
	"os"

	"github.com/jackc/pgx/v5"

	"go.shabbyrobe.org/num"
	"lukechampine.com/uint128"

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

	rows, err := conn.Query(
		context.Background(),
		"SELECT $1::uint1, $2::uint2, $3::uint4, $4::uint8, $5::uint16, $6::int1, $7::int1, $8::int16, $9::int16",
		uint8(math.MaxUint8),
		uint16(math.MaxUint16),
		uint32(math.MaxUint32),
		uint64(math.MaxUint64),
		uint128.Max,
		int8(math.MaxInt8),
		int8(math.MinInt8),
		num.MaxI128,
		num.MinI128,
	)
	if err != nil {
		log.Fatalf("Cannot exec query: %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var num8 uint8
		var num16 uint16
		var num32 uint32
		var num64 uint64
		var num128 uint128.Uint128
		var num8sMax int8
		var num8sMin int8
		var num128sMax num.I128
		var num128sMin num.I128

		if err := rows.Scan(&num8, &num16, &num32, &num64, &num128, &num8sMax, &num8sMin, &num128sMax, &num128sMin); err != nil {
			log.Fatalf("Cannot scan row: %v", err)
		}

		log.Printf("Got u8 %d, u16 %d, u32 %d, u64 %d, u128 %s", num8, num16, num32, num64, num128.String())
		log.Printf("Got s8 max: %d, s8 min: %d", num8sMax, num8sMin)
		log.Printf("Got s128 max: %s, s128 min: %s", num128sMax.String(), num128sMin.String())
	}
}

// Correct output should look like this:
//
// Got u8 255, u16 65535, u32 4294967295, u64 18446744073709551615, u128 340282366920938463463374607431768211455
// Got s8 max: 127, s8 min: -128
// Got s128 max: 170141183460469231731687303715884105727, s128 min: -170141183460469231731687303715884105728
