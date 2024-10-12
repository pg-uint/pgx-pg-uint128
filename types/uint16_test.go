// Do not edit. Generated from codegen

package types_test

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/jackc/pgx/v5/pgtype"

	"lukechampine.com/uint128"

	"go.shabbyrobe.org/num"
)

func TestUInt16Binary_ScanUint16(t *testing.T) {
	var dst uint16

	assert.NoError(t,
		typeMap.Scan(UInt16OID, pgtype.BinaryFormatCode, []byte("\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\xFF\xFF"), &dst),
	)

	assert.Equal(t, uint16(65535), dst)
}
func TestUInt16Binary_ScanUint16_Overflow(t *testing.T) {
	var dst uint16

	assert.ErrorContains(t,
		typeMap.Scan(UInt16OID, pgtype.BinaryFormatCode, []byte("\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF"), &dst),
		`340282366920938463463374607431768211455 is greater than maximum value for uint16`,
	)
}

func TestUInt16Binary_ScanUint32(t *testing.T) {
	var dst uint32

	assert.NoError(t,
		typeMap.Scan(UInt16OID, pgtype.BinaryFormatCode, []byte("\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\xFF\xFF\xFF\xFF"), &dst),
	)

	assert.Equal(t, uint32(4294967295), dst)
}
func TestUInt16Binary_ScanUint32_Overflow(t *testing.T) {
	var dst uint32

	assert.ErrorContains(t,
		typeMap.Scan(UInt16OID, pgtype.BinaryFormatCode, []byte("\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF"), &dst),
		`340282366920938463463374607431768211455 is greater than maximum value for uint32`,
	)
}

func TestUInt16Binary_ScanUint64(t *testing.T) {
	var dst uint64

	assert.NoError(t,
		typeMap.Scan(UInt16OID, pgtype.BinaryFormatCode, []byte("\x00\x00\x00\x00\x00\x00\x00\x00\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF"), &dst),
	)

	assert.Equal(t, uint64(18446744073709551615), dst)
}
func TestUInt16Binary_ScanUint64_Overflow(t *testing.T) {
	var dst uint64

	assert.ErrorContains(t,
		typeMap.Scan(UInt16OID, pgtype.BinaryFormatCode, []byte("\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF"), &dst),
		`340282366920938463463374607431768211455 is greater than maximum value for uint64`,
	)
}

func TestUInt16Binary_ScanUint128(t *testing.T) {
	var dst uint128.Uint128

	assert.NoError(t,
		typeMap.Scan(UInt16OID, pgtype.BinaryFormatCode, []byte("\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF"), &dst),
	)

	assert.Equal(t, u128Max, dst)
}

func TestUInt16Binary_ScanUint(t *testing.T) {
	var dst uint

	assert.NoError(t,
		typeMap.Scan(UInt16OID, pgtype.BinaryFormatCode, []byte("\x00\x00\x00\x00\x00\x00\x00\x00\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF"), &dst),
	)

	assert.Equal(t, uint(18446744073709551615), dst)
}
func TestUInt16Binary_ScanUint_Overflow(t *testing.T) {
	var dst uint

	assert.ErrorContains(t,
		typeMap.Scan(UInt16OID, pgtype.BinaryFormatCode, []byte("\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF"), &dst),
		`340282366920938463463374607431768211455 is greater than maximum value for uint`,
	)
}

func TestUInt16Binary_ScanInt16(t *testing.T) {
	var dst int16

	assert.NoError(t,
		typeMap.Scan(UInt16OID, pgtype.BinaryFormatCode, []byte("\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x7F\xFF"), &dst),
	)

	assert.Equal(t, int16(32767), dst)
}
func TestUInt16Binary_ScanInt16_Overflow(t *testing.T) {
	var dst int16

	assert.ErrorContains(t,
		typeMap.Scan(UInt16OID, pgtype.BinaryFormatCode, []byte("\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF"), &dst),
		`340282366920938463463374607431768211455 is greater than maximum value for int16`,
	)
}

func TestUInt16Binary_ScanInt32(t *testing.T) {
	var dst int32

	assert.NoError(t,
		typeMap.Scan(UInt16OID, pgtype.BinaryFormatCode, []byte("\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x7F\xFF\xFF\xFF"), &dst),
	)

	assert.Equal(t, int32(2147483647), dst)
}
func TestUInt16Binary_ScanInt32_Overflow(t *testing.T) {
	var dst int32

	assert.ErrorContains(t,
		typeMap.Scan(UInt16OID, pgtype.BinaryFormatCode, []byte("\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF"), &dst),
		`340282366920938463463374607431768211455 is greater than maximum value for int32`,
	)
}

func TestUInt16Binary_ScanInt64(t *testing.T) {
	var dst int64

	assert.NoError(t,
		typeMap.Scan(UInt16OID, pgtype.BinaryFormatCode, []byte("\x00\x00\x00\x00\x00\x00\x00\x00\x7F\xFF\xFF\xFF\xFF\xFF\xFF\xFF"), &dst),
	)

	assert.Equal(t, int64(9223372036854775807), dst)
}
func TestUInt16Binary_ScanInt64_Overflow(t *testing.T) {
	var dst int64

	assert.ErrorContains(t,
		typeMap.Scan(UInt16OID, pgtype.BinaryFormatCode, []byte("\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF"), &dst),
		`340282366920938463463374607431768211455 is greater than maximum value for int64`,
	)
}

func TestUInt16Binary_ScanI128(t *testing.T) {
	var dst num.I128

	assert.NoError(t,
		typeMap.Scan(UInt16OID, pgtype.BinaryFormatCode, []byte("\x7F\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF"), &dst),
	)

	assert.Equal(t, s128Max, dst)
}
func TestUInt16Binary_ScanI128_Overflow(t *testing.T) {
	var dst num.I128

	assert.ErrorContains(t,
		typeMap.Scan(UInt16OID, pgtype.BinaryFormatCode, []byte("\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF"), &dst),
		`340282366920938463463374607431768211455 is greater than maximum value for I128`,
	)
}

func TestUInt16Binary_ScanInt(t *testing.T) {
	var dst int

	assert.NoError(t,
		typeMap.Scan(UInt16OID, pgtype.BinaryFormatCode, []byte("\x00\x00\x00\x00\x00\x00\x00\x00\x7F\xFF\xFF\xFF\xFF\xFF\xFF\xFF"), &dst),
	)

	assert.Equal(t, int(9223372036854775807), dst)
}
func TestUInt16Binary_ScanInt_Overflow(t *testing.T) {
	var dst int

	assert.ErrorContains(t,
		typeMap.Scan(UInt16OID, pgtype.BinaryFormatCode, []byte("\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF"), &dst),
		`340282366920938463463374607431768211455 is greater than maximum value for int`,
	)
}

func TestUInt16Text_ScanUint16(t *testing.T) {
	var dst uint16

	assert.NoError(t,
		typeMap.Scan(UInt16OID, pgtype.TextFormatCode, []byte("65535"), &dst),
	)

	assert.Equal(t, uint16(65535), dst)
}
func TestUInt16Text_ScanUint16_Overflow(t *testing.T) {
	var dst uint16

	assert.ErrorContains(t,
		typeMap.Scan(UInt16OID, pgtype.TextFormatCode, []byte("340282366920938463463374607431768211455"), &dst),
		`strconv.ParseUint: parsing "340282366920938463463374607431768211455": value out of range`,
	)
}

func TestUInt16Text_ScanUint32(t *testing.T) {
	var dst uint32

	assert.NoError(t,
		typeMap.Scan(UInt16OID, pgtype.TextFormatCode, []byte("4294967295"), &dst),
	)

	assert.Equal(t, uint32(4294967295), dst)
}
func TestUInt16Text_ScanUint32_Overflow(t *testing.T) {
	var dst uint32

	assert.ErrorContains(t,
		typeMap.Scan(UInt16OID, pgtype.TextFormatCode, []byte("340282366920938463463374607431768211455"), &dst),
		`strconv.ParseUint: parsing "340282366920938463463374607431768211455": value out of range`,
	)
}

func TestUInt16Text_ScanUint64(t *testing.T) {
	var dst uint64

	assert.NoError(t,
		typeMap.Scan(UInt16OID, pgtype.TextFormatCode, []byte("18446744073709551615"), &dst),
	)

	assert.Equal(t, uint64(18446744073709551615), dst)
}
func TestUInt16Text_ScanUint64_Overflow(t *testing.T) {
	var dst uint64

	assert.ErrorContains(t,
		typeMap.Scan(UInt16OID, pgtype.TextFormatCode, []byte("340282366920938463463374607431768211455"), &dst),
		`strconv.ParseUint: parsing "340282366920938463463374607431768211455": value out of range`,
	)
}

func TestUInt16Text_ScanUint128(t *testing.T) {
	var dst uint128.Uint128

	assert.NoError(t,
		typeMap.Scan(UInt16OID, pgtype.TextFormatCode, []byte("340282366920938463463374607431768211455"), &dst),
	)

	assert.Equal(t, u128Max, dst)
}
func TestUInt16Text_ScanUint128_SelfOverflow(t *testing.T) {
	var dst uint128.Uint128

	assert.ErrorContains(t,
		typeMap.Scan(UInt16OID, pgtype.TextFormatCode, []byte("340282366920938463463374607431768211456"), &dst),
		`value overflows Uint128`,
	)
}

func TestUInt16Text_ScanUint(t *testing.T) {
	var dst uint

	assert.NoError(t,
		typeMap.Scan(UInt16OID, pgtype.TextFormatCode, []byte("18446744073709551615"), &dst),
	)

	assert.Equal(t, uint(18446744073709551615), dst)
}
func TestUInt16Text_ScanUint_Overflow(t *testing.T) {
	var dst uint

	assert.ErrorContains(t,
		typeMap.Scan(UInt16OID, pgtype.TextFormatCode, []byte("340282366920938463463374607431768211455"), &dst),
		`strconv.ParseUint: parsing "340282366920938463463374607431768211455": value out of range`,
	)
}

func TestUInt16Text_ScanInt16(t *testing.T) {
	var dst int16

	assert.NoError(t,
		typeMap.Scan(UInt16OID, pgtype.TextFormatCode, []byte("32767"), &dst),
	)

	assert.Equal(t, int16(32767), dst)
}
func TestUInt16Text_ScanInt16_Overflow(t *testing.T) {
	var dst int16

	assert.ErrorContains(t,
		typeMap.Scan(UInt16OID, pgtype.TextFormatCode, []byte("340282366920938463463374607431768211455"), &dst),
		`strconv.ParseInt: parsing "340282366920938463463374607431768211455": value out of range`,
	)
}

func TestUInt16Text_ScanInt32(t *testing.T) {
	var dst int32

	assert.NoError(t,
		typeMap.Scan(UInt16OID, pgtype.TextFormatCode, []byte("2147483647"), &dst),
	)

	assert.Equal(t, int32(2147483647), dst)
}
func TestUInt16Text_ScanInt32_Overflow(t *testing.T) {
	var dst int32

	assert.ErrorContains(t,
		typeMap.Scan(UInt16OID, pgtype.TextFormatCode, []byte("340282366920938463463374607431768211455"), &dst),
		`strconv.ParseInt: parsing "340282366920938463463374607431768211455": value out of range`,
	)
}

func TestUInt16Text_ScanInt64(t *testing.T) {
	var dst int64

	assert.NoError(t,
		typeMap.Scan(UInt16OID, pgtype.TextFormatCode, []byte("9223372036854775807"), &dst),
	)

	assert.Equal(t, int64(9223372036854775807), dst)
}
func TestUInt16Text_ScanInt64_Overflow(t *testing.T) {
	var dst int64

	assert.ErrorContains(t,
		typeMap.Scan(UInt16OID, pgtype.TextFormatCode, []byte("340282366920938463463374607431768211455"), &dst),
		`strconv.ParseInt: parsing "340282366920938463463374607431768211455": value out of range`,
	)
}

func TestUInt16Text_ScanI128(t *testing.T) {
	var dst num.I128

	assert.NoError(t,
		typeMap.Scan(UInt16OID, pgtype.TextFormatCode, []byte("170141183460469231731687303715884105727"), &dst),
	)

	assert.Equal(t, s128Max, dst)
}
func TestUInt16Text_ScanI128_Overflow(t *testing.T) {
	var dst num.I128

	assert.ErrorContains(t,
		typeMap.Scan(UInt16OID, pgtype.TextFormatCode, []byte("340282366920938463463374607431768211455"), &dst),
		`int128.FromString: parsing "340282366920938463463374607431768211455": value out of range`,
	)
}

func TestUInt16Text_ScanInt(t *testing.T) {
	var dst int

	assert.NoError(t,
		typeMap.Scan(UInt16OID, pgtype.TextFormatCode, []byte("9223372036854775807"), &dst),
	)

	assert.Equal(t, int(9223372036854775807), dst)
}
func TestUInt16Text_ScanInt_Overflow(t *testing.T) {
	var dst int

	assert.ErrorContains(t,
		typeMap.Scan(UInt16OID, pgtype.TextFormatCode, []byte("340282366920938463463374607431768211455"), &dst),
		`strconv.ParseInt: parsing "340282366920938463463374607431768211455": value out of range`,
	)
}
