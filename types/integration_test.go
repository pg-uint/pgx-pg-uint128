package types_test

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxtest"
	"github.com/stretchr/testify/require"
	"go.shabbyrobe.org/num"
	"lukechampine.com/uint128"
	"math"
	"os"
	"testing"
	"time"

	. "github.com/pg-uint/pgx-pg-uint128/types"
)

var defaultConnTestRunner pgxtest.ConnTestRunner

func init() {
	defaultConnTestRunner = pgxtest.DefaultConnTestRunner()
	defaultConnTestRunner.CreateConfig = func(ctx context.Context, t testing.TB) *pgx.ConnConfig {
		config, err := pgx.ParseConfig(os.Getenv("PGX_TEST_DATABASE"))
		require.NoError(t, err)
		return config
	}

	defaultConnTestRunner.AfterConnect = func(ctx context.Context, t testing.TB, conn *pgx.Conn) {
		_, err := RegisterAll(ctx, conn)
		require.NoError(t, err)
	}
}

func isExpectedEq(a any) func(any) bool {
	return func(v any) bool {
		return a == v
	}
}

func TestUInt2_Integration(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pgxtest.RunValueRoundTripTests(context.Background(), t, defaultConnTestRunner, nil, "uint2", []pgxtest.ValueRoundTripTest{
		{int8(1), new(uint16), isExpectedEq(uint16(1))},
		{int16(1), new(uint16), isExpectedEq(uint16(1))},
		{int32(1), new(uint16), isExpectedEq(uint16(1))},
		{int64(1), new(uint16), isExpectedEq(uint16(1))},
		{uint8(1), new(uint16), isExpectedEq(uint16(1))},
		{uint16(1), new(uint16), isExpectedEq(uint16(1))},
		{uint32(1), new(uint16), isExpectedEq(uint16(1))},
		{uint64(1), new(uint16), isExpectedEq(uint16(1))},
		{int(1), new(uint16), isExpectedEq(uint16(1))},
		{uint(1), new(uint16), isExpectedEq(uint16(1))},
		{pgtype.Int2{Int16: 1, Valid: true}, new(uint16), isExpectedEq(uint16(1))},
		{UInt2{Uint16: 1, Valid: true}, new(uint16), isExpectedEq(uint16(1))},
		{int32(1), new(UInt2), isExpectedEq(UInt2{Uint16: 1, Valid: true})},
		{1, new(UInt2), isExpectedEq(UInt2{Uint16: 1, Valid: true})},
		{"1", new(string), isExpectedEq("1")},
		{UInt2{}, new(UInt2), isExpectedEq(UInt2{})},
		{nil, new(*uint16), isExpectedEq((*uint16)(nil))},

		// Signed types target
		{uint16(1), new(int8), isExpectedEq(int8(1))},
		{uint16(1), new(int16), isExpectedEq(int16(1))},
		{uint16(1), new(int32), isExpectedEq(int32(1))},
		{uint16(1), new(int64), isExpectedEq(int64(1))},
		{uint16(1), new(int), isExpectedEq(int(1))},

		// Unsigned types target
		{uint16(1), new(uint8), isExpectedEq(uint8(1))},
		{uint16(1), new(uint16), isExpectedEq(uint16(1))},
		{uint16(1), new(uint32), isExpectedEq(uint32(1))},
		{uint16(1), new(uint64), isExpectedEq(uint64(1))},
		{uint16(1), new(uint), isExpectedEq(uint(1))},
		{uint16(1), new(uint128.Uint128), isExpectedEq(uint128.From64(1))},
		//{uint16(1), new(UInt2), isExpectedEq(UInt2{Uint16: 1, Valid: true})},
		//{uint16(1), new(UInt4), isExpectedEq(UInt4{Uint32: 1, Valid: true})},
		//{uint16(1), new(UInt8), isExpectedEq(UInt8{Uint64: 1, Valid: true})},
		//{uint16(1), new(UInt16), isExpectedEq(UInt16{Uint128: uint128.From64(1), Valid: true})},
	})

	pgxtest.RunWithQueryExecModes(ctx, t, defaultConnTestRunner, pgxtest.AllQueryExecModes, func(ctx context.Context, t testing.TB, conn *pgx.Conn) {
		var h *uint16
		var expectedNil *uint16
		var expectedEmpty uint16
		var expectedMax uint16 = math.MaxUint16

		err := conn.QueryRow(ctx, `select cast(null as uint2)`).Scan(&h)
		if err != nil {
			t.Fatal(err)
		}

		require.Equalf(t, expectedNil, h, "plain conn.Scan failed expectedNil=%#v actual=%#v", expectedNil, h)

		err = conn.QueryRow(ctx, `select cast(0 as uint2)`).Scan(&h)
		if err != nil {
			t.Fatal(err)
		}

		require.NotNil(t, h)
		require.Equalf(t, expectedEmpty, *h, "plain conn.Scan failed expectedEmpty=%#v actual=%#v\", expectedEmpty, h", expectedEmpty, h)

		err = conn.QueryRow(ctx, `select 65535::uint2`).Scan(&h)
		if err != nil {
			t.Fatal(err)
		}

		require.NotNil(t, h)
		require.Equalf(t, expectedMax, *h, "plain conn.Scan failed expectedMax=%#v actual=%#v\", expectedMax, h", expectedMax, h)
	})
}

func TestUInt4_Integration(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pgxtest.RunValueRoundTripTests(context.Background(), t, defaultConnTestRunner, nil, "uint4", []pgxtest.ValueRoundTripTest{
		{int8(1), new(uint32), isExpectedEq(uint32(1))},
		{int16(1), new(uint32), isExpectedEq(uint32(1))},
		{int32(1), new(uint32), isExpectedEq(uint32(1))},
		{int64(1), new(uint32), isExpectedEq(uint32(1))},
		{uint8(1), new(uint32), isExpectedEq(uint32(1))},
		{uint16(1), new(uint32), isExpectedEq(uint32(1))},
		{uint32(1), new(uint32), isExpectedEq(uint32(1))},
		{uint64(1), new(uint32), isExpectedEq(uint32(1))},
		{int(1), new(uint32), isExpectedEq(uint32(1))},
		{uint(1), new(uint32), isExpectedEq(uint32(1))},
		{pgtype.Int4{Int32: 1, Valid: true}, new(uint32), isExpectedEq(uint32(1))},
		{UInt4{Uint32: 1, Valid: true}, new(uint32), isExpectedEq(uint32(1))},
		{int32(1), new(UInt4), isExpectedEq(UInt4{Uint32: 1, Valid: true})},
		{1, new(UInt4), isExpectedEq(UInt4{Uint32: 1, Valid: true})},
		{"1", new(string), isExpectedEq("1")},
		{UInt4{}, new(UInt4), isExpectedEq(UInt4{})},
		{nil, new(*uint32), isExpectedEq((*uint32)(nil))},

		// Signed types target
		{uint32(1), new(int8), isExpectedEq(int8(1))},
		{uint32(1), new(int16), isExpectedEq(int16(1))},
		{uint32(1), new(int32), isExpectedEq(int32(1))},
		{uint32(1), new(int64), isExpectedEq(int64(1))},
		{uint32(1), new(int), isExpectedEq(int(1))},

		// Unsigned types target
		{uint32(1), new(uint8), isExpectedEq(uint8(1))},
		{uint32(1), new(uint16), isExpectedEq(uint16(1))},
		{uint32(1), new(uint32), isExpectedEq(uint32(1))},
		{uint32(1), new(uint64), isExpectedEq(uint64(1))},
		{uint32(1), new(uint), isExpectedEq(uint(1))},
		{uint32(1), new(uint128.Uint128), isExpectedEq(uint128.From64(1))},
	})

	pgxtest.RunWithQueryExecModes(ctx, t, defaultConnTestRunner, pgxtest.AllQueryExecModes, func(ctx context.Context, t testing.TB, conn *pgx.Conn) {
		var h *uint32
		var expectedNil *uint32
		var expectedEmpty uint32
		var expectedMax uint32 = math.MaxUint32

		err := conn.QueryRow(ctx, `select cast(null as uint4)`).Scan(&h)
		if err != nil {
			t.Fatal(err)
		}

		require.Equalf(t, expectedNil, h, "plain conn.Scan failed expectedNil=%#v actual=%#v", expectedNil, h)

		err = conn.QueryRow(ctx, `select cast(0 as uint4)`).Scan(&h)
		if err != nil {
			t.Fatal(err)
		}

		require.NotNil(t, h)
		require.Equalf(t, expectedEmpty, *h, "plain conn.Scan failed expectedEmpty=%#v actual=%#v\", expectedEmpty, h", expectedEmpty, h)

		err = conn.QueryRow(ctx, `select 4294967295::uint4`).Scan(&h)
		if err != nil {
			t.Fatal(err)
		}

		require.NotNil(t, h)
		require.Equalf(t, expectedMax, *h, "plain conn.Scan failed expectedMax=%#v actual=%#v\", expectedMax, h", expectedMax, h)
	})
}

func TestUInt8_Integration(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pgxtest.RunValueRoundTripTests(context.Background(), t, defaultConnTestRunner, nil, "uint8", []pgxtest.ValueRoundTripTest{
		{int8(1), new(uint64), isExpectedEq(uint64(1))},
		{int16(1), new(uint64), isExpectedEq(uint64(1))},
		{int32(1), new(uint64), isExpectedEq(uint64(1))},
		{int64(1), new(uint64), isExpectedEq(uint64(1))},
		{uint8(1), new(uint64), isExpectedEq(uint64(1))},
		{uint16(1), new(uint64), isExpectedEq(uint64(1))},
		{uint32(1), new(uint64), isExpectedEq(uint64(1))},
		{uint64(1), new(uint64), isExpectedEq(uint64(1))},
		{int(1), new(uint64), isExpectedEq(uint64(1))},
		{uint(1), new(uint64), isExpectedEq(uint64(1))},
		{pgtype.Int8{Int64: 1, Valid: true}, new(uint64), isExpectedEq(uint64(1))},
		{UInt8{Uint64: 1, Valid: true}, new(uint64), isExpectedEq(uint64(1))},
		{int32(1), new(UInt8), isExpectedEq(UInt8{Uint64: 1, Valid: true})},
		{1, new(UInt8), isExpectedEq(UInt8{Uint64: 1, Valid: true})},
		{"1", new(string), isExpectedEq("1")},
		{UInt8{}, new(UInt8), isExpectedEq(UInt8{})},
		{nil, new(*uint64), isExpectedEq((*uint64)(nil))},

		// Signed types target
		{uint64(1), new(int8), isExpectedEq(int8(1))},
		{uint64(1), new(int16), isExpectedEq(int16(1))},
		{uint64(1), new(int32), isExpectedEq(int32(1))},
		{uint64(1), new(int64), isExpectedEq(int64(1))},
		{uint64(1), new(int), isExpectedEq(int(1))},

		// Unsigned types target
		{uint64(1), new(uint8), isExpectedEq(uint8(1))},
		{uint64(1), new(uint16), isExpectedEq(uint16(1))},
		{uint64(1), new(uint32), isExpectedEq(uint32(1))},
		{uint64(1), new(uint64), isExpectedEq(uint64(1))},
		{uint64(1), new(uint), isExpectedEq(uint(1))},
		{uint64(1), new(uint128.Uint128), isExpectedEq(uint128.From64(1))},
	})

	pgxtest.RunWithQueryExecModes(ctx, t, defaultConnTestRunner, pgxtest.AllQueryExecModes, func(ctx context.Context, t testing.TB, conn *pgx.Conn) {
		var h *uint64
		var expectedNil *uint64
		var expectedEmpty uint64
		var expectedMax uint64 = math.MaxUint64

		err := conn.QueryRow(ctx, `select cast(null as uint8)`).Scan(&h)
		if err != nil {
			t.Fatal(err)
		}

		require.Equalf(t, expectedNil, h, "plain conn.Scan failed expectedNil=%#v actual=%#v", expectedNil, h)

		err = conn.QueryRow(ctx, `select cast(0 as uint8)`).Scan(&h)
		if err != nil {
			t.Fatal(err)
		}

		require.NotNil(t, h)
		require.Equalf(t, expectedEmpty, *h, "plain conn.Scan failed expectedEmpty=%#v actual=%#v\", expectedEmpty, h", expectedEmpty, h)

		err = conn.QueryRow(ctx, `select 18446744073709551615::uint8`).Scan(&h)
		if err != nil {
			t.Fatal(err)
		}

		require.NotNil(t, h)
		require.Equalf(t, expectedMax, *h, "plain conn.Scan failed expectedMax=%#v actual=%#v\", expectedMax, h", expectedMax, h)
	})
}

func TestUInt16_Integration(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var uint128_1 = uint128.From64(1)

	pgxtest.RunValueRoundTripTests(context.Background(), t, defaultConnTestRunner, nil, "uint16", []pgxtest.ValueRoundTripTest{
		{int8(1), new(uint128.Uint128), isExpectedEq(uint128_1)},
		{int16(1), new(uint128.Uint128), isExpectedEq(uint128_1)},
		{int32(1), new(uint128.Uint128), isExpectedEq(uint128_1)},
		{int64(1), new(uint128.Uint128), isExpectedEq(uint128_1)},
		{uint8(1), new(uint128.Uint128), isExpectedEq(uint128_1)},
		{uint16(1), new(uint128.Uint128), isExpectedEq(uint128_1)},
		{uint32(1), new(uint128.Uint128), isExpectedEq(uint128_1)},
		{uint64(1), new(uint128.Uint128), isExpectedEq(uint128_1)},
		{int(1), new(uint128.Uint128), isExpectedEq(uint128_1)},
		{uint(1), new(uint128.Uint128), isExpectedEq(uint128_1)},
		{pgtype.Int8{Int64: 1, Valid: true}, new(uint128.Uint128), isExpectedEq(uint128_1)},
		{UInt8{Uint64: 1, Valid: true}, new(uint128.Uint128), isExpectedEq(uint128_1)},
		{int32(1), new(UInt16), isExpectedEq(UInt16{Uint128: uint128.From64(1), Valid: true})},
		{1, new(UInt16), isExpectedEq(UInt16{Uint128: uint128.From64(1), Valid: true})},
		{"1", new(string), isExpectedEq("1")},
		{UInt16{}, new(UInt16), isExpectedEq(UInt16{})},
		{nil, new(*uint128.Uint128), isExpectedEq((*uint128.Uint128)(nil))},

		// Signed types target
		{uint128_1, new(int8), isExpectedEq(int8(1))},
		{uint128_1, new(int16), isExpectedEq(int16(1))},
		{uint128_1, new(int32), isExpectedEq(int32(1))},
		{uint128_1, new(int64), isExpectedEq(int64(1))},
		{uint128_1, new(int), isExpectedEq(int(1))},

		// Unsigned types target
		{uint128_1, new(uint8), isExpectedEq(uint8(1))},
		{uint128_1, new(uint16), isExpectedEq(uint16(1))},
		{uint128_1, new(uint32), isExpectedEq(uint32(1))},
		{uint128_1, new(uint64), isExpectedEq(uint64(1))},
		{uint128_1, new(uint), isExpectedEq(uint(1))},
		{uint128_1, new(uint128.Uint128), isExpectedEq(uint128.From64(1))},
	})

	pgxtest.RunWithQueryExecModes(ctx, t, defaultConnTestRunner, pgxtest.AllQueryExecModes, func(ctx context.Context, t testing.TB, conn *pgx.Conn) {
		var h *uint128.Uint128
		var expectedNil *uint128.Uint128
		var expectedEmpty uint128.Uint128
		var expectedMax uint128.Uint128 = uint128.Max

		err := conn.QueryRow(ctx, `select cast(null as uint16)`).Scan(&h)
		if err != nil {
			t.Fatal(err)
		}

		require.Equalf(t, expectedNil, h, "plain conn.Scan failed expectedNil=%#v actual=%#v", expectedNil, h)

		err = conn.QueryRow(ctx, `select cast(0 as uint16)`).Scan(&h)
		if err != nil {
			t.Fatal(err)
		}

		require.NotNil(t, h)
		require.Equalf(t, expectedEmpty, *h, "plain conn.Scan failed expectedEmpty=%#v actual=%#v\", expectedEmpty, h", expectedEmpty, h)

		err = conn.QueryRow(ctx, `select 340282366920938463463374607431768211455::uint16`).Scan(&h)
		if err != nil {
			t.Fatal(err)
		}

		require.NotNil(t, h)
		require.Equalf(t, expectedMax, *h, "plain conn.Scan failed expectedMax=%#v actual=%#v\", expectedMax, h", expectedMax, h)
	})
}

func TestInt16_Integration(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var int128_1 = num.I128From64(1)

	pgxtest.RunValueRoundTripTests(context.Background(), t, defaultConnTestRunner, nil, "int16", []pgxtest.ValueRoundTripTest{
		{int8(1), new(num.I128), isExpectedEq(int128_1)},
		{int16(1), new(num.I128), isExpectedEq(int128_1)},
		{int32(1), new(num.I128), isExpectedEq(int128_1)},
		{int64(1), new(num.I128), isExpectedEq(int128_1)},
		{uint8(1), new(num.I128), isExpectedEq(int128_1)},
		{uint16(1), new(num.I128), isExpectedEq(int128_1)},
		{uint32(1), new(num.I128), isExpectedEq(int128_1)},
		{uint64(1), new(num.I128), isExpectedEq(int128_1)},
		{int(1), new(num.I128), isExpectedEq(int128_1)},
		{uint(1), new(num.I128), isExpectedEq(int128_1)},
		{pgtype.Int8{Int64: 1, Valid: true}, new(num.I128), isExpectedEq(int128_1)},
		{UInt8{Uint64: 1, Valid: true}, new(num.I128), isExpectedEq(int128_1)},
		{int32(1), new(Int16), isExpectedEq(Int16{I128: num.I128From64(1), Valid: true})},
		{1, new(Int16), isExpectedEq(Int16{I128: num.I128From64(1), Valid: true})},
		{"1", new(string), isExpectedEq("1")},
		{Int16{}, new(Int16), isExpectedEq(Int16{})},
		{nil, new(*num.I128), isExpectedEq((*num.I128)(nil))},

		// Signed types target
		{int128_1, new(int8), isExpectedEq(int8(1))},
		{int128_1, new(int16), isExpectedEq(int16(1))},
		{int128_1, new(int32), isExpectedEq(int32(1))},
		{int128_1, new(int64), isExpectedEq(int64(1))},
		{int128_1, new(int), isExpectedEq(int(1))},

		// Unsigned types target
		{int128_1, new(uint8), isExpectedEq(uint8(1))},
		{int128_1, new(uint16), isExpectedEq(uint16(1))},
		{int128_1, new(uint32), isExpectedEq(uint32(1))},
		{int128_1, new(uint64), isExpectedEq(uint64(1))},
		{int128_1, new(uint), isExpectedEq(uint(1))},
		{int128_1, new(uint128.Uint128), isExpectedEq(uint128.From64(1))},
	})

	pgxtest.RunWithQueryExecModes(ctx, t, defaultConnTestRunner, pgxtest.AllQueryExecModes, func(ctx context.Context, t testing.TB, conn *pgx.Conn) {
		var h *num.I128
		var expectedNil *num.I128
		var expectedEmpty num.I128
		var expectedMax num.I128 = s128Max

		err := conn.QueryRow(ctx, `select cast(null as int16)`).Scan(&h)
		if err != nil {
			t.Fatal(err)
		}

		require.Equalf(t, expectedNil, h, "plain conn.Scan failed expectedNil=%#v actual=%#v", expectedNil, h)

		err = conn.QueryRow(ctx, `select cast(0 as int16)`).Scan(&h)
		if err != nil {
			t.Fatal(err)
		}

		require.NotNil(t, h)
		require.Equalf(t, expectedEmpty, *h, "plain conn.Scan failed expectedEmpty=%#v actual=%#v\", expectedEmpty, h", expectedEmpty, h)

		err = conn.QueryRow(ctx, `select 170141183460469231731687303715884105727::int16`).Scan(&h)
		if err != nil {
			t.Fatal(err)
		}

		require.NotNil(t, h)
		require.Equalf(t, expectedMax, *h, "plain conn.Scan failed expectedMax=%#v actual=%#v\", expectedMax, h", expectedMax, h)
	})
}
