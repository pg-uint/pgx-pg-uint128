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

	"github.com/codercms/pgx-pg-uint128/types"
	"github.com/codercms/pgx-pg-uint128/types/zeronull"
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

		log.Printf("PID=%d Got u16 %d, u32 %d, u64 %d, u128 %s", conn.PgConn().PID(), num16, num32, num64, num128.String())
		log.Printf("PID=%d Got s128 max: %s, s128 min: %s", conn.PgConn().PID(), num128sMax.String(), num128sMin.String())
	}
}

// Correct output should look like this (order of output can be randomized a bit because of parallelism):
//
// Types cache created by conn pid=25380
// Types registered from cache to conn pid=25380
// Types registered from cache to conn pid=24048
// Types registered from cache to conn pid=22676
// Types registered from cache to conn pid=26020
// PID=26020 Got u16 65535, u32 4294967295, u64 18446744073709551615, u128 340282366920938463463374607431768211455
// PID=26020 Got s128 max: 170141183460469231731687303715884105727, s128 min: -170141183460469231731687303715884105728
// PID=25380 Got u16 65535, u32 4294967295, u64 18446744073709551615, u128 340282366920938463463374607431768211455
// PID=24048 Got u16 65535, u32 4294967295, u64 18446744073709551615, u128 340282366920938463463374607431768211455
// PID=24048 Got s128 max: 170141183460469231731687303715884105727, s128 min: -170141183460469231731687303715884105728
// PID=22676 Got u16 65535, u32 4294967295, u64 18446744073709551615, u128 340282366920938463463374607431768211455
// PID=22676 Got s128 max: 170141183460469231731687303715884105727, s128 min: -170141183460469231731687303715884105728
// PID=25380 Got s128 max: 170141183460469231731687303715884105727, s128 min: -170141183460469231731687303715884105728
