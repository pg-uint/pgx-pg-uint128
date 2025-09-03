package main

import (
	"context"
	"log"
	"math"
	"os"
	"sync"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"

	"go.shabbyrobe.org/num"
	"lukechampine.com/uint128"

	"github.com/pg-uint/pgx-pg-uint128/types"
	"github.com/pg-uint/pgx-pg-uint128/types/zeronull"
)

func main() {
	poolCfg, err := pgxpool.ParseConfig(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Cannot create pgxpool config: %v", err)
	}

	// Cache types information instead of fetching their info everytime
	var typesCacheOnce sync.Once
	var typesCache []*pgtype.Type

	poolCfg.MaxConns = 4
	poolCfg.MinConns = 4
	poolCfg.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		var cacheErr error

		typesCacheOnce.Do(func() {
			pgTypes, regErr := types.RegisterAll(ctx, conn)
			if regErr != nil {
				cacheErr = regErr
			}

			typesCache = pgTypes

			log.Printf("Types cache created by conn pid=%d", conn.PgConn().PID())
		})

		if cacheErr != nil {
			log.Fatalf("Unable to register types: %v", cacheErr)
		}

		// RegisterTypes is available from pgx 5.7+
		//conn.TypeMap().RegisterTypes(typesCache)

		// In older versions of pgx you can use:
		for _, t := range typesCache {
			conn.TypeMap().RegisterType(t)
		}

		types.RegisterDefaultPgTypeVariants(conn.TypeMap())

		// Optionally register zeronull types
		zeronull.Register(conn.TypeMap())

		log.Printf("Types registered from cache to conn pid=%d", conn.PgConn().PID())

		return nil
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), poolCfg)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	defer pool.Close()

	// And now lets try to test that the types cache is applied correctly to all connections
	conns := make([]*pgxpool.Conn, 0, poolCfg.MaxConns)
	for i := 0; i < int(poolCfg.MaxConns); i++ {
		conn, err := pool.Acquire(context.Background())
		if err != nil {
			log.Fatalf("Unable to acquire connection: %v", err)
		}

		conns = append(conns, conn)
	}

	var wg sync.WaitGroup
	wg.Add(len(conns))

	for _, conn := range conns {
		go func(conn *pgxpool.Conn) {
			defer func() {
				conn.Release()
				wg.Done()
			}()

			testConn(conn.Conn())
		}(conn)
	}

	wg.Wait()
}

func testConn(conn *pgx.Conn) {
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

		log.Printf("PID=%d Got u8 %d, u16 %d, u32 %d, u64 %d, u128 %s", conn.PgConn().PID(), num8, num16, num32, num64, num128.String())
		log.Printf("PID=%d Got s8 max: %d, s8 min: %d", conn.PgConn().PID(), num8sMax, num8sMin)
		log.Printf("PID=%d Got s128 max: %s, s128 min: %s", conn.PgConn().PID(), num128sMax.String(), num128sMin.String())
	}
}

// Correct output should look like this (order of output can be randomized a bit because of parallelism):
//
// Types cache created by conn pid=45733
// Types registered from cache to conn pid=45733
// Types registered from cache to conn pid=45734
// Types registered from cache to conn pid=45732
// Types registered from cache to conn pid=45731
// PID=45731 Got u8 255, u16 65535, u32 4294967295, u64 18446744073709551615, u128 340282366920938463463374607431768211455
// PID=45731 Got s8 max: 127, s8 min: -128
// PID=45731 Got s128 max: 170141183460469231731687303715884105727, s128 min: -170141183460469231731687303715884105728
// PID=45733 Got u8 255, u16 65535, u32 4294967295, u64 18446744073709551615, u128 340282366920938463463374607431768211455
// PID=45733 Got s8 max: 127, s8 min: -128
// PID=45733 Got s128 max: 170141183460469231731687303715884105727, s128 min: -170141183460469231731687303715884105728
// PID=45732 Got u8 255, u16 65535, u32 4294967295, u64 18446744073709551615, u128 340282366920938463463374607431768211455
// PID=45732 Got s8 max: 127, s8 min: -128
// PID=45732 Got s128 max: 170141183460469231731687303715884105727, s128 min: -170141183460469231731687303715884105728
// PID=45734 Got u8 255, u16 65535, u32 4294967295, u64 18446744073709551615, u128 340282366920938463463374607431768211455
// PID=45734 Got s8 max: 127, s8 min: -128
// PID=45734 Got s128 max: 170141183460469231731687303715884105727, s128 min: -170141183460469231731687303715884105728
