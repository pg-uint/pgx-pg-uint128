package types_test

import (
	. "github.com/codercms/pgx-pg-uint128/types"
	"github.com/jackc/pgx/v5/pgtype"
	"go.shabbyrobe.org/num"
	"lukechampine.com/uint128"
	"math"
)

var typeMap = pgtype.NewMap()

const intSize = 32 << (^uint(0) >> 63)

const UInt2OID = 787898
const UInt4OID = 787899
const UInt8OID = 787900
const UInt16OID = 787901

const Int16OID = 787902

func registerTestUint2(tMap *pgtype.Map) {
	tMap.RegisterType(&pgtype.Type{Name: "uint2", OID: UInt2OID, Codec: UInt2Codec{}})
}

func registerTestUint4(tMap *pgtype.Map) {
	tMap.RegisterType(&pgtype.Type{Name: "uint4", OID: UInt4OID, Codec: UInt4Codec{}})
}

func registerTestUint8(tMap *pgtype.Map) {
	tMap.RegisterType(&pgtype.Type{Name: "uint8", OID: UInt8OID, Codec: UInt8Codec{}})
}

func registerTestUint16(tMap *pgtype.Map) {
	tMap.RegisterType(&pgtype.Type{Name: "uint16", OID: UInt16OID, Codec: UInt16Codec{}})
}

func registerTestInt16(tMap *pgtype.Map) {
	tMap.RegisterType(&pgtype.Type{Name: "int16", OID: Int16OID, Codec: Int16Codec{}})
}

func registerTestTypes(tMap *pgtype.Map) {
	registerTestUint2(tMap)
	registerTestUint4(tMap)
	registerTestUint8(tMap)
	registerTestUint16(tMap)

	registerTestInt16(tMap)
}

func init() {
	registerTestTypes(typeMap)
}

// Unsigned integer types max values in [uint128.Uint128]
var (
	u8MaxInU128  = uint128.From64(math.MaxUint8)
	u16MaxInU128 = uint128.From64(math.MaxUint16)
	u32MaxInU128 = uint128.From64(math.MaxUint32)
	u64MaxInU128 = uint128.From64(math.MaxUint64)
	u128Max      = uint128.Max

	uMaxInU128 = uint128.From64(math.MaxUint)
)

// Signed integer types max values in [uint128.Uint128]
var (
	s8MaxInU128   = uint128.From64(math.MaxInt8)
	s16MaxInU128  = uint128.From64(math.MaxInt16)
	s32MaxInU128  = uint128.From64(math.MaxInt32)
	s64MaxInU128  = uint128.From64(math.MaxInt64)
	s128MaxInU128 = uint128.New(math.MaxUint64, math.MaxInt64)
	sMaxInU128    = uint128.From64(math.MaxInt)
)

// Unsigned integer types max values in [num.I128]
var (
	u8MaxInS128  = num.I128FromU64(math.MaxUint8)
	u16MaxInS128 = num.I128FromU64(math.MaxUint16)
	u32MaxInS128 = num.I128FromU64(math.MaxUint32)
	u64MaxInS128 = num.I128FromU64(math.MaxUint64)

	uMaxInS128 = num.I128FromU64(math.MaxUint)
)

// Signed integer types max values in [num.I128]
var (
	s8MaxInS128  = num.I128From64(math.MaxInt8)
	s16MaxInS128 = num.I128From64(math.MaxInt16)
	s32MaxInS128 = num.I128From64(math.MaxInt32)
	s64MaxInS128 = num.I128From64(math.MaxInt64)

	s128Max = num.MaxI128

	sMaxInS128 = num.I128From64(math.MaxInt)
)
