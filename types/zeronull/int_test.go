package zeronull_test

import (
	"context"
	"go.shabbyrobe.org/num"
	"lukechampine.com/uint128"
	"testing"

	"github.com/jackc/pgx/v5/pgxtest"

	"github.com/pg-uint/pgx-pg-uint128/types/zeronull"
)

func isExpectedEq(a any) func(any) bool {
	return func(v any) bool {
		return a == v
	}
}

func TestUInt2Transcode(t *testing.T) {
	pgxtest.RunValueRoundTripTests(context.Background(), t, defaultConnTestRunner, nil, "uint2", []pgxtest.ValueRoundTripTest{
		{
			(zeronull.UInt2)(1),
			new(zeronull.UInt2),
			isExpectedEq((zeronull.UInt2)(1)),
		},
		{
			nil,
			new(zeronull.UInt2),
			isExpectedEq((zeronull.UInt2)(0)),
		},
		{
			(zeronull.UInt2)(0),
			new(any),
			isExpectedEq(nil),
		},
	})
}

func TestUInt4Transcode(t *testing.T) {
	pgxtest.RunValueRoundTripTests(context.Background(), t, defaultConnTestRunner, nil, "uint4", []pgxtest.ValueRoundTripTest{
		{
			(zeronull.UInt4)(1),
			new(zeronull.UInt4),
			isExpectedEq((zeronull.UInt4)(1)),
		},
		{
			nil,
			new(zeronull.UInt4),
			isExpectedEq((zeronull.UInt4)(0)),
		},
		{
			(zeronull.UInt4)(0),
			new(any),
			isExpectedEq(nil),
		},
	})
}

func TestUInt8Transcode(t *testing.T) {
	pgxtest.RunValueRoundTripTests(context.Background(), t, defaultConnTestRunner, nil, "uint8", []pgxtest.ValueRoundTripTest{
		{
			(zeronull.UInt8)(1),
			new(zeronull.UInt8),
			isExpectedEq((zeronull.UInt8)(1)),
		},
		{
			nil,
			new(zeronull.UInt8),
			isExpectedEq((zeronull.UInt8)(0)),
		},
		{
			(zeronull.UInt8)(0),
			new(any),
			isExpectedEq(nil),
		},
	})
}

func TestUInt16Transcode(t *testing.T) {
	pgxtest.RunValueRoundTripTests(context.Background(), t, defaultConnTestRunner, nil, "uint16", []pgxtest.ValueRoundTripTest{
		{
			(zeronull.UInt16)(uint128.From64(1)),
			new(zeronull.UInt16),
			isExpectedEq((zeronull.UInt16)(uint128.From64(1))),
		},
		{
			nil,
			new(zeronull.UInt16),
			isExpectedEq(zeronull.UInt16{}),
		},
		{
			zeronull.UInt16{},
			new(any),
			isExpectedEq(nil),
		},
	})
}

func TestInt16Transcode(t *testing.T) {
	pgxtest.RunValueRoundTripTests(context.Background(), t, defaultConnTestRunner, nil, "int16", []pgxtest.ValueRoundTripTest{
		{
			(zeronull.Int16)(num.I128From64(1)),
			new(zeronull.Int16),
			isExpectedEq((zeronull.Int16)(num.I128From64(1))),
		},
		{
			nil,
			new(zeronull.Int16),
			isExpectedEq(zeronull.Int16{}),
		},
		{
			zeronull.Int16{},
			new(any),
			isExpectedEq(nil),
		},
	})
}
