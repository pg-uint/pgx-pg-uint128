// Do not edit. Generated from codegen

package types_test

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/jackc/pgx/v5/pgtype"

	"lukechampine.com/uint128"

	"go.shabbyrobe.org/num"
)

func TestUInt8Binary_ScanUint16(t *testing.T) {
	var dst uint16

	assert.NoError(t,
		typeMap.Scan(UInt8OID, pgtype.BinaryFormatCode, []byte("\x00\x00\x00\x00\x00\x00\xFF\xFF"), &dst),
	)

	assert.Equal(t, uint16(65535), dst)
}
func TestUInt8Binary_ScanUint16_Overflow(t *testing.T) {
	var dst uint16

	assert.ErrorContains(t,
		typeMap.Scan(UInt8OID, pgtype.BinaryFormatCode, []byte("\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF"), &dst),
		`18446744073709551615 is greater than maximum value for uint16`,
	)
}

func TestUInt8Binary_ScanUint32(t *testing.T) {
	var dst uint32

	assert.NoError(t,
		typeMap.Scan(UInt8OID, pgtype.BinaryFormatCode, []byte("\x00\x00\x00\x00\xFF\xFF\xFF\xFF"), &dst),
	)

	assert.Equal(t, uint32(4294967295), dst)
}
func TestUInt8Binary_ScanUint32_Overflow(t *testing.T) {
	var dst uint32

	assert.ErrorContains(t,
		typeMap.Scan(UInt8OID, pgtype.BinaryFormatCode, []byte("\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF"), &dst),
		`18446744073709551615 is greater than maximum value for uint32`,
	)
}

func TestUInt8Binary_ScanUint64(t *testing.T) {
	var dst uint64

	assert.NoError(t,
		typeMap.Scan(UInt8OID, pgtype.BinaryFormatCode, []byte("\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF"), &dst),
	)

	assert.Equal(t, uint64(18446744073709551615), dst)
}

func TestUInt8Binary_ScanUint128(t *testing.T) {
	var dst uint128.Uint128

	assert.NoError(t,
		typeMap.Scan(UInt8OID, pgtype.BinaryFormatCode, []byte("\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF"), &dst),
	)

	assert.Equal(t, u64MaxInU128, dst)
}

func TestUInt8Binary_ScanUint(t *testing.T) {
	var dst uint

	assert.NoError(t,
		typeMap.Scan(UInt8OID, pgtype.BinaryFormatCode, []byte("\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF"), &dst),
	)

	assert.Equal(t, uint(18446744073709551615), dst)
}

func TestUInt8Binary_ScanInt16(t *testing.T) {
	var dst int16

	assert.NoError(t,
		typeMap.Scan(UInt8OID, pgtype.BinaryFormatCode, []byte("\x00\x00\x00\x00\x00\x00\x7F\xFF"), &dst),
	)

	assert.Equal(t, int16(32767), dst)
}
func TestUInt8Binary_ScanInt16_Overflow(t *testing.T) {
	var dst int16

	assert.ErrorContains(t,
		typeMap.Scan(UInt8OID, pgtype.BinaryFormatCode, []byte("\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF"), &dst),
		`18446744073709551615 is greater than maximum value for int16`,
	)
}

func TestUInt8Binary_ScanInt32(t *testing.T) {
	var dst int32

	assert.NoError(t,
		typeMap.Scan(UInt8OID, pgtype.BinaryFormatCode, []byte("\x00\x00\x00\x00\x7F\xFF\xFF\xFF"), &dst),
	)

	assert.Equal(t, int32(2147483647), dst)
}
func TestUInt8Binary_ScanInt32_Overflow(t *testing.T) {
	var dst int32

	assert.ErrorContains(t,
		typeMap.Scan(UInt8OID, pgtype.BinaryFormatCode, []byte("\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF"), &dst),
		`18446744073709551615 is greater than maximum value for int32`,
	)
}

func TestUInt8Binary_ScanInt64(t *testing.T) {
	var dst int64

	assert.NoError(t,
		typeMap.Scan(UInt8OID, pgtype.BinaryFormatCode, []byte("\x7F\xFF\xFF\xFF\xFF\xFF\xFF\xFF"), &dst),
	)

	assert.Equal(t, int64(9223372036854775807), dst)
}
func TestUInt8Binary_ScanInt64_Overflow(t *testing.T) {
	var dst int64

	assert.ErrorContains(t,
		typeMap.Scan(UInt8OID, pgtype.BinaryFormatCode, []byte("\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF"), &dst),
		`18446744073709551615 is greater than maximum value for int64`,
	)
}

func TestUInt8Binary_ScanI128(t *testing.T) {
	var dst num.I128

	assert.NoError(t,
		typeMap.Scan(UInt8OID, pgtype.BinaryFormatCode, []byte("\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF"), &dst),
	)

	assert.Equal(t, u64MaxInS128, dst)
}

func TestUInt8Binary_ScanInt(t *testing.T) {
	var dst int

	assert.NoError(t,
		typeMap.Scan(UInt8OID, pgtype.BinaryFormatCode, []byte("\x7F\xFF\xFF\xFF\xFF\xFF\xFF\xFF"), &dst),
	)

	assert.Equal(t, int(9223372036854775807), dst)
}
func TestUInt8Binary_ScanInt_Overflow(t *testing.T) {
	var dst int

	assert.ErrorContains(t,
		typeMap.Scan(UInt8OID, pgtype.BinaryFormatCode, []byte("\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF"), &dst),
		`18446744073709551615 is greater than maximum value for int`,
	)
}

func TestUInt8Text_ScanUint16(t *testing.T) {
	var dst uint16

	assert.NoError(t,
		typeMap.Scan(UInt8OID, pgtype.TextFormatCode, []byte("65535"), &dst),
	)

	assert.Equal(t, uint16(65535), dst)
}
func TestUInt8Text_ScanUint16_Overflow(t *testing.T) {
	var dst uint16

	assert.ErrorContains(t,
		typeMap.Scan(UInt8OID, pgtype.TextFormatCode, []byte("18446744073709551615"), &dst),
		`strconv.ParseUint: parsing "18446744073709551615": value out of range`,
	)
}

func TestUInt8Text_ScanUint32(t *testing.T) {
	var dst uint32

	assert.NoError(t,
		typeMap.Scan(UInt8OID, pgtype.TextFormatCode, []byte("4294967295"), &dst),
	)

	assert.Equal(t, uint32(4294967295), dst)
}
func TestUInt8Text_ScanUint32_Overflow(t *testing.T) {
	var dst uint32

	assert.ErrorContains(t,
		typeMap.Scan(UInt8OID, pgtype.TextFormatCode, []byte("18446744073709551615"), &dst),
		`strconv.ParseUint: parsing "18446744073709551615": value out of range`,
	)
}

func TestUInt8Text_ScanUint64(t *testing.T) {
	var dst uint64

	assert.NoError(t,
		typeMap.Scan(UInt8OID, pgtype.TextFormatCode, []byte("18446744073709551615"), &dst),
	)

	assert.Equal(t, uint64(18446744073709551615), dst)
}
func TestUInt8Text_ScanUint64_SelfOverflow(t *testing.T) {
	var dst uint64

	assert.ErrorContains(t,
		typeMap.Scan(UInt8OID, pgtype.TextFormatCode, []byte("18446744073709551616"), &dst),
		`strconv.ParseUint: parsing "18446744073709551616": value out of range`,
	)
}

func TestUInt8Text_ScanUint128(t *testing.T) {
	var dst uint128.Uint128

	assert.NoError(t,
		typeMap.Scan(UInt8OID, pgtype.TextFormatCode, []byte("18446744073709551615"), &dst),
	)

	assert.Equal(t, u64MaxInU128, dst)
}

func TestUInt8Text_ScanUint(t *testing.T) {
	var dst uint

	assert.NoError(t,
		typeMap.Scan(UInt8OID, pgtype.TextFormatCode, []byte("18446744073709551615"), &dst),
	)

	assert.Equal(t, uint(18446744073709551615), dst)
}

func TestUInt8Text_ScanInt16(t *testing.T) {
	var dst int16

	assert.NoError(t,
		typeMap.Scan(UInt8OID, pgtype.TextFormatCode, []byte("32767"), &dst),
	)

	assert.Equal(t, int16(32767), dst)
}
func TestUInt8Text_ScanInt16_Overflow(t *testing.T) {
	var dst int16

	assert.ErrorContains(t,
		typeMap.Scan(UInt8OID, pgtype.TextFormatCode, []byte("18446744073709551615"), &dst),
		`strconv.ParseInt: parsing "18446744073709551615": value out of range`,
	)
}

func TestUInt8Text_ScanInt32(t *testing.T) {
	var dst int32

	assert.NoError(t,
		typeMap.Scan(UInt8OID, pgtype.TextFormatCode, []byte("2147483647"), &dst),
	)

	assert.Equal(t, int32(2147483647), dst)
}
func TestUInt8Text_ScanInt32_Overflow(t *testing.T) {
	var dst int32

	assert.ErrorContains(t,
		typeMap.Scan(UInt8OID, pgtype.TextFormatCode, []byte("18446744073709551615"), &dst),
		`strconv.ParseInt: parsing "18446744073709551615": value out of range`,
	)
}

func TestUInt8Text_ScanInt64(t *testing.T) {
	var dst int64

	assert.NoError(t,
		typeMap.Scan(UInt8OID, pgtype.TextFormatCode, []byte("9223372036854775807"), &dst),
	)

	assert.Equal(t, int64(9223372036854775807), dst)
}
func TestUInt8Text_ScanInt64_Overflow(t *testing.T) {
	var dst int64

	assert.ErrorContains(t,
		typeMap.Scan(UInt8OID, pgtype.TextFormatCode, []byte("18446744073709551615"), &dst),
		`strconv.ParseInt: parsing "18446744073709551615": value out of range`,
	)
}

func TestUInt8Text_ScanI128(t *testing.T) {
	var dst num.I128

	assert.NoError(t,
		typeMap.Scan(UInt8OID, pgtype.TextFormatCode, []byte("18446744073709551615"), &dst),
	)

	assert.Equal(t, u64MaxInS128, dst)
}

func TestUInt8Text_ScanInt(t *testing.T) {
	var dst int

	assert.NoError(t,
		typeMap.Scan(UInt8OID, pgtype.TextFormatCode, []byte("9223372036854775807"), &dst),
	)

	assert.Equal(t, int(9223372036854775807), dst)
}
func TestUInt8Text_ScanInt_Overflow(t *testing.T) {
	var dst int

	assert.ErrorContains(t,
		typeMap.Scan(UInt8OID, pgtype.TextFormatCode, []byte("18446744073709551615"), &dst),
		`strconv.ParseInt: parsing "18446744073709551615": value out of range`,
	)
}
