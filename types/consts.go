package types

import (
	"go.shabbyrobe.org/num"
	"lukechampine.com/uint128"
	"math"
)

const intSize = 32 << (^uint(0) >> 63)

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

const (
	Uint2TypName  = "uint2"
	Uint4TypName  = "uint4"
	Uint8TypName  = "uint8"
	Uint16TypName = "uint16"

	Int16TypName = "int16"
)
