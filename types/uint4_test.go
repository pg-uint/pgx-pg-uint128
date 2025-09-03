// Do not edit. Generated from codegen

package types_test

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/jackc/pgx/v5/pgtype"

	"lukechampine.com/uint128"

	"go.shabbyrobe.org/num"
)

func TestUInt4Binary_ScanUint8(t *testing.T) {
	var dst uint8

	assert.NoError(t,
		typeMap.Scan(UInt4OID, pgtype.BinaryFormatCode, []byte("\x00\x00\x00\xFF"), &dst),
	)

	assert.Equal(t, uint8(255), dst)
}
func TestUInt4Binary_ScanUint8_Overflow(t *testing.T) {
	var dst uint8

	assert.ErrorContains(t,
		typeMap.Scan(UInt4OID, pgtype.BinaryFormatCode, []byte("\xFF\xFF\xFF\xFF"), &dst),
		`4294967295 is greater than maximum value for uint8`,
	)
}

func TestUInt4Binary_ScanUint16(t *testing.T) {
	var dst uint16

	assert.NoError(t,
		typeMap.Scan(UInt4OID, pgtype.BinaryFormatCode, []byte("\x00\x00\xFF\xFF"), &dst),
	)

	assert.Equal(t, uint16(65535), dst)
}
func TestUInt4Binary_ScanUint16_Overflow(t *testing.T) {
	var dst uint16

	assert.ErrorContains(t,
		typeMap.Scan(UInt4OID, pgtype.BinaryFormatCode, []byte("\xFF\xFF\xFF\xFF"), &dst),
		`4294967295 is greater than maximum value for uint16`,
	)
}

func TestUInt4Binary_ScanUint32(t *testing.T) {
	var dst uint32

	assert.NoError(t,
		typeMap.Scan(UInt4OID, pgtype.BinaryFormatCode, []byte("\xFF\xFF\xFF\xFF"), &dst),
	)

	assert.Equal(t, uint32(4294967295), dst)
}

func TestUInt4Binary_ScanUint64(t *testing.T) {
	var dst uint64

	assert.NoError(t,
		typeMap.Scan(UInt4OID, pgtype.BinaryFormatCode, []byte("\xFF\xFF\xFF\xFF"), &dst),
	)

	assert.Equal(t, uint64(4294967295), dst)
}

func TestUInt4Binary_ScanUint128(t *testing.T) {
	var dst uint128.Uint128

	assert.NoError(t,
		typeMap.Scan(UInt4OID, pgtype.BinaryFormatCode, []byte("\xFF\xFF\xFF\xFF"), &dst),
	)

	assert.Equal(t, u32MaxInU128, dst)
}

func TestUInt4Binary_ScanUint(t *testing.T) {
	var dst uint

	assert.NoError(t,
		typeMap.Scan(UInt4OID, pgtype.BinaryFormatCode, []byte("\xFF\xFF\xFF\xFF"), &dst),
	)

	assert.Equal(t, uint(4294967295), dst)
}

func TestUInt4Binary_ScanInt8(t *testing.T) {
	var dst int8

	assert.NoError(t,
		typeMap.Scan(UInt4OID, pgtype.BinaryFormatCode, []byte("\x00\x00\x00\x7F"), &dst),
	)

	assert.Equal(t, int8(127), dst)
}
func TestUInt4Binary_ScanInt8_Overflow(t *testing.T) {
	var dst int8

	assert.ErrorContains(t,
		typeMap.Scan(UInt4OID, pgtype.BinaryFormatCode, []byte("\xFF\xFF\xFF\xFF"), &dst),
		`4294967295 is greater than maximum value for int8`,
	)
}

func TestUInt4Binary_ScanInt16(t *testing.T) {
	var dst int16

	assert.NoError(t,
		typeMap.Scan(UInt4OID, pgtype.BinaryFormatCode, []byte("\x00\x00\x7F\xFF"), &dst),
	)

	assert.Equal(t, int16(32767), dst)
}
func TestUInt4Binary_ScanInt16_Overflow(t *testing.T) {
	var dst int16

	assert.ErrorContains(t,
		typeMap.Scan(UInt4OID, pgtype.BinaryFormatCode, []byte("\xFF\xFF\xFF\xFF"), &dst),
		`4294967295 is greater than maximum value for int16`,
	)
}

func TestUInt4Binary_ScanInt32(t *testing.T) {
	var dst int32

	assert.NoError(t,
		typeMap.Scan(UInt4OID, pgtype.BinaryFormatCode, []byte("\x7F\xFF\xFF\xFF"), &dst),
	)

	assert.Equal(t, int32(2147483647), dst)
}
func TestUInt4Binary_ScanInt32_Overflow(t *testing.T) {
	var dst int32

	assert.ErrorContains(t,
		typeMap.Scan(UInt4OID, pgtype.BinaryFormatCode, []byte("\xFF\xFF\xFF\xFF"), &dst),
		`4294967295 is greater than maximum value for int32`,
	)
}

func TestUInt4Binary_ScanInt64(t *testing.T) {
	var dst int64

	assert.NoError(t,
		typeMap.Scan(UInt4OID, pgtype.BinaryFormatCode, []byte("\xFF\xFF\xFF\xFF"), &dst),
	)

	assert.Equal(t, int64(4294967295), dst)
}

func TestUInt4Binary_ScanI128(t *testing.T) {
	var dst num.I128

	assert.NoError(t,
		typeMap.Scan(UInt4OID, pgtype.BinaryFormatCode, []byte("\xFF\xFF\xFF\xFF"), &dst),
	)

	assert.Equal(t, u32MaxInS128, dst)
}

func TestUInt4Binary_ScanInt(t *testing.T) {
	var dst int

	assert.NoError(t,
		typeMap.Scan(UInt4OID, pgtype.BinaryFormatCode, []byte("\xFF\xFF\xFF\xFF"), &dst),
	)

	assert.Equal(t, int(4294967295), dst)
}
func TestUInt4Binary_ScanInt_Overflow(t *testing.T) {
	var dst int

	if intSize == 32 {
		assert.ErrorContains(t,
			typeMap.Scan(UInt4OID, pgtype.BinaryFormatCode, []byte("\xFF\xFF\xFF\xFF"), &dst),
			`4294967295 is greater than maximum value for int`,
		)
	}
}

func TestUInt4Text_ScanUint8(t *testing.T) {
	var dst uint8

	assert.NoError(t,
		typeMap.Scan(UInt4OID, pgtype.TextFormatCode, []byte("255"), &dst),
	)

	assert.Equal(t, uint8(255), dst)
}
func TestUInt4Text_ScanUint8_Overflow(t *testing.T) {
	var dst uint8

	assert.ErrorContains(t,
		typeMap.Scan(UInt4OID, pgtype.TextFormatCode, []byte("4294967295"), &dst),
		`strconv.ParseUint: parsing "4294967295": value out of range`,
	)
}

func TestUInt4Text_ScanUint16(t *testing.T) {
	var dst uint16

	assert.NoError(t,
		typeMap.Scan(UInt4OID, pgtype.TextFormatCode, []byte("65535"), &dst),
	)

	assert.Equal(t, uint16(65535), dst)
}
func TestUInt4Text_ScanUint16_Overflow(t *testing.T) {
	var dst uint16

	assert.ErrorContains(t,
		typeMap.Scan(UInt4OID, pgtype.TextFormatCode, []byte("4294967295"), &dst),
		`strconv.ParseUint: parsing "4294967295": value out of range`,
	)
}

func TestUInt4Text_ScanUint32(t *testing.T) {
	var dst uint32

	assert.NoError(t,
		typeMap.Scan(UInt4OID, pgtype.TextFormatCode, []byte("4294967295"), &dst),
	)

	assert.Equal(t, uint32(4294967295), dst)
}
func TestUInt4Text_ScanUint32_SelfOverflow(t *testing.T) {
	var dst uint32

	assert.ErrorContains(t,
		typeMap.Scan(UInt4OID, pgtype.TextFormatCode, []byte("4294967296"), &dst),
		`strconv.ParseUint: parsing "4294967296": value out of range`,
	)
}

func TestUInt4Text_ScanUint64(t *testing.T) {
	var dst uint64

	assert.NoError(t,
		typeMap.Scan(UInt4OID, pgtype.TextFormatCode, []byte("4294967295"), &dst),
	)

	assert.Equal(t, uint64(4294967295), dst)
}

func TestUInt4Text_ScanUint128(t *testing.T) {
	var dst uint128.Uint128

	assert.NoError(t,
		typeMap.Scan(UInt4OID, pgtype.TextFormatCode, []byte("4294967295"), &dst),
	)

	assert.Equal(t, u32MaxInU128, dst)
}

func TestUInt4Text_ScanUint(t *testing.T) {
	var dst uint

	assert.NoError(t,
		typeMap.Scan(UInt4OID, pgtype.TextFormatCode, []byte("4294967295"), &dst),
	)

	assert.Equal(t, uint(4294967295), dst)
}

func TestUInt4Text_ScanInt8(t *testing.T) {
	var dst int8

	assert.NoError(t,
		typeMap.Scan(UInt4OID, pgtype.TextFormatCode, []byte("127"), &dst),
	)

	assert.Equal(t, int8(127), dst)
}
func TestUInt4Text_ScanInt8_Overflow(t *testing.T) {
	var dst int8

	assert.ErrorContains(t,
		typeMap.Scan(UInt4OID, pgtype.TextFormatCode, []byte("4294967295"), &dst),
		`strconv.ParseInt: parsing "4294967295": value out of range`,
	)
}

func TestUInt4Text_ScanInt16(t *testing.T) {
	var dst int16

	assert.NoError(t,
		typeMap.Scan(UInt4OID, pgtype.TextFormatCode, []byte("32767"), &dst),
	)

	assert.Equal(t, int16(32767), dst)
}
func TestUInt4Text_ScanInt16_Overflow(t *testing.T) {
	var dst int16

	assert.ErrorContains(t,
		typeMap.Scan(UInt4OID, pgtype.TextFormatCode, []byte("4294967295"), &dst),
		`strconv.ParseInt: parsing "4294967295": value out of range`,
	)
}

func TestUInt4Text_ScanInt32(t *testing.T) {
	var dst int32

	assert.NoError(t,
		typeMap.Scan(UInt4OID, pgtype.TextFormatCode, []byte("2147483647"), &dst),
	)

	assert.Equal(t, int32(2147483647), dst)
}
func TestUInt4Text_ScanInt32_Overflow(t *testing.T) {
	var dst int32

	assert.ErrorContains(t,
		typeMap.Scan(UInt4OID, pgtype.TextFormatCode, []byte("4294967295"), &dst),
		`strconv.ParseInt: parsing "4294967295": value out of range`,
	)
}

func TestUInt4Text_ScanInt64(t *testing.T) {
	var dst int64

	assert.NoError(t,
		typeMap.Scan(UInt4OID, pgtype.TextFormatCode, []byte("4294967295"), &dst),
	)

	assert.Equal(t, int64(4294967295), dst)
}

func TestUInt4Text_ScanI128(t *testing.T) {
	var dst num.I128

	assert.NoError(t,
		typeMap.Scan(UInt4OID, pgtype.TextFormatCode, []byte("4294967295"), &dst),
	)

	assert.Equal(t, u32MaxInS128, dst)
}

func TestUInt4Text_ScanInt(t *testing.T) {
	var dst int

	assert.NoError(t,
		typeMap.Scan(UInt4OID, pgtype.TextFormatCode, []byte("4294967295"), &dst),
	)

	assert.Equal(t, int(4294967295), dst)
}
func TestUInt4Text_ScanInt_Overflow(t *testing.T) {
	var dst int

	if intSize == 32 {
		assert.ErrorContains(t,
			typeMap.Scan(UInt4OID, pgtype.TextFormatCode, []byte("4294967295"), &dst),
			`strconv.ParseInt: parsing "4294967295": value out of range`,
		)
	}
}
