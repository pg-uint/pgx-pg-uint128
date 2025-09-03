// Do not edit. Generated from codegen

package types_test

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/jackc/pgx/v5/pgtype"

	"lukechampine.com/uint128"

	"go.shabbyrobe.org/num"
)

func TestUInt2Binary_ScanUint8(t *testing.T) {
	var dst uint8

	assert.NoError(t,
		typeMap.Scan(UInt2OID, pgtype.BinaryFormatCode, []byte("\x00\xFF"), &dst),
	)

	assert.Equal(t, uint8(255), dst)
}
func TestUInt2Binary_ScanUint8_Overflow(t *testing.T) {
	var dst uint8

	assert.ErrorContains(t,
		typeMap.Scan(UInt2OID, pgtype.BinaryFormatCode, []byte("\xFF\xFF"), &dst),
		`65535 is greater than maximum value for uint8`,
	)
}

func TestUInt2Binary_ScanUint16(t *testing.T) {
	var dst uint16

	assert.NoError(t,
		typeMap.Scan(UInt2OID, pgtype.BinaryFormatCode, []byte("\xFF\xFF"), &dst),
	)

	assert.Equal(t, uint16(65535), dst)
}

func TestUInt2Binary_ScanUint32(t *testing.T) {
	var dst uint32

	assert.NoError(t,
		typeMap.Scan(UInt2OID, pgtype.BinaryFormatCode, []byte("\xFF\xFF"), &dst),
	)

	assert.Equal(t, uint32(65535), dst)
}

func TestUInt2Binary_ScanUint64(t *testing.T) {
	var dst uint64

	assert.NoError(t,
		typeMap.Scan(UInt2OID, pgtype.BinaryFormatCode, []byte("\xFF\xFF"), &dst),
	)

	assert.Equal(t, uint64(65535), dst)
}

func TestUInt2Binary_ScanUint128(t *testing.T) {
	var dst uint128.Uint128

	assert.NoError(t,
		typeMap.Scan(UInt2OID, pgtype.BinaryFormatCode, []byte("\xFF\xFF"), &dst),
	)

	assert.Equal(t, u16MaxInU128, dst)
}

func TestUInt2Binary_ScanUint(t *testing.T) {
	var dst uint

	assert.NoError(t,
		typeMap.Scan(UInt2OID, pgtype.BinaryFormatCode, []byte("\xFF\xFF"), &dst),
	)

	assert.Equal(t, uint(65535), dst)
}

func TestUInt2Binary_ScanInt8(t *testing.T) {
	var dst int8

	assert.NoError(t,
		typeMap.Scan(UInt2OID, pgtype.BinaryFormatCode, []byte("\x00\x7F"), &dst),
	)

	assert.Equal(t, int8(127), dst)
}
func TestUInt2Binary_ScanInt8_Overflow(t *testing.T) {
	var dst int8

	assert.ErrorContains(t,
		typeMap.Scan(UInt2OID, pgtype.BinaryFormatCode, []byte("\xFF\xFF"), &dst),
		`65535 is greater than maximum value for int8`,
	)
}

func TestUInt2Binary_ScanInt16(t *testing.T) {
	var dst int16

	assert.NoError(t,
		typeMap.Scan(UInt2OID, pgtype.BinaryFormatCode, []byte("\x7F\xFF"), &dst),
	)

	assert.Equal(t, int16(32767), dst)
}
func TestUInt2Binary_ScanInt16_Overflow(t *testing.T) {
	var dst int16

	assert.ErrorContains(t,
		typeMap.Scan(UInt2OID, pgtype.BinaryFormatCode, []byte("\xFF\xFF"), &dst),
		`65535 is greater than maximum value for int16`,
	)
}

func TestUInt2Binary_ScanInt32(t *testing.T) {
	var dst int32

	assert.NoError(t,
		typeMap.Scan(UInt2OID, pgtype.BinaryFormatCode, []byte("\xFF\xFF"), &dst),
	)

	assert.Equal(t, int32(65535), dst)
}

func TestUInt2Binary_ScanInt64(t *testing.T) {
	var dst int64

	assert.NoError(t,
		typeMap.Scan(UInt2OID, pgtype.BinaryFormatCode, []byte("\xFF\xFF"), &dst),
	)

	assert.Equal(t, int64(65535), dst)
}

func TestUInt2Binary_ScanI128(t *testing.T) {
	var dst num.I128

	assert.NoError(t,
		typeMap.Scan(UInt2OID, pgtype.BinaryFormatCode, []byte("\xFF\xFF"), &dst),
	)

	assert.Equal(t, u16MaxInS128, dst)
}

func TestUInt2Binary_ScanInt(t *testing.T) {
	var dst int

	assert.NoError(t,
		typeMap.Scan(UInt2OID, pgtype.BinaryFormatCode, []byte("\xFF\xFF"), &dst),
	)

	assert.Equal(t, int(65535), dst)
}

func TestUInt2Text_ScanUint8(t *testing.T) {
	var dst uint8

	assert.NoError(t,
		typeMap.Scan(UInt2OID, pgtype.TextFormatCode, []byte("255"), &dst),
	)

	assert.Equal(t, uint8(255), dst)
}
func TestUInt2Text_ScanUint8_Overflow(t *testing.T) {
	var dst uint8

	assert.ErrorContains(t,
		typeMap.Scan(UInt2OID, pgtype.TextFormatCode, []byte("65535"), &dst),
		`strconv.ParseUint: parsing "65535": value out of range`,
	)
}

func TestUInt2Text_ScanUint16(t *testing.T) {
	var dst uint16

	assert.NoError(t,
		typeMap.Scan(UInt2OID, pgtype.TextFormatCode, []byte("65535"), &dst),
	)

	assert.Equal(t, uint16(65535), dst)
}
func TestUInt2Text_ScanUint16_SelfOverflow(t *testing.T) {
	var dst uint16

	assert.ErrorContains(t,
		typeMap.Scan(UInt2OID, pgtype.TextFormatCode, []byte("65536"), &dst),
		`strconv.ParseUint: parsing "65536": value out of range`,
	)
}

func TestUInt2Text_ScanUint32(t *testing.T) {
	var dst uint32

	assert.NoError(t,
		typeMap.Scan(UInt2OID, pgtype.TextFormatCode, []byte("65535"), &dst),
	)

	assert.Equal(t, uint32(65535), dst)
}

func TestUInt2Text_ScanUint64(t *testing.T) {
	var dst uint64

	assert.NoError(t,
		typeMap.Scan(UInt2OID, pgtype.TextFormatCode, []byte("65535"), &dst),
	)

	assert.Equal(t, uint64(65535), dst)
}

func TestUInt2Text_ScanUint128(t *testing.T) {
	var dst uint128.Uint128

	assert.NoError(t,
		typeMap.Scan(UInt2OID, pgtype.TextFormatCode, []byte("65535"), &dst),
	)

	assert.Equal(t, u16MaxInU128, dst)
}

func TestUInt2Text_ScanUint(t *testing.T) {
	var dst uint

	assert.NoError(t,
		typeMap.Scan(UInt2OID, pgtype.TextFormatCode, []byte("65535"), &dst),
	)

	assert.Equal(t, uint(65535), dst)
}

func TestUInt2Text_ScanInt8(t *testing.T) {
	var dst int8

	assert.NoError(t,
		typeMap.Scan(UInt2OID, pgtype.TextFormatCode, []byte("127"), &dst),
	)

	assert.Equal(t, int8(127), dst)
}
func TestUInt2Text_ScanInt8_Overflow(t *testing.T) {
	var dst int8

	assert.ErrorContains(t,
		typeMap.Scan(UInt2OID, pgtype.TextFormatCode, []byte("65535"), &dst),
		`strconv.ParseInt: parsing "65535": value out of range`,
	)
}

func TestUInt2Text_ScanInt16(t *testing.T) {
	var dst int16

	assert.NoError(t,
		typeMap.Scan(UInt2OID, pgtype.TextFormatCode, []byte("32767"), &dst),
	)

	assert.Equal(t, int16(32767), dst)
}
func TestUInt2Text_ScanInt16_Overflow(t *testing.T) {
	var dst int16

	assert.ErrorContains(t,
		typeMap.Scan(UInt2OID, pgtype.TextFormatCode, []byte("65535"), &dst),
		`strconv.ParseInt: parsing "65535": value out of range`,
	)
}

func TestUInt2Text_ScanInt32(t *testing.T) {
	var dst int32

	assert.NoError(t,
		typeMap.Scan(UInt2OID, pgtype.TextFormatCode, []byte("65535"), &dst),
	)

	assert.Equal(t, int32(65535), dst)
}

func TestUInt2Text_ScanInt64(t *testing.T) {
	var dst int64

	assert.NoError(t,
		typeMap.Scan(UInt2OID, pgtype.TextFormatCode, []byte("65535"), &dst),
	)

	assert.Equal(t, int64(65535), dst)
}

func TestUInt2Text_ScanI128(t *testing.T) {
	var dst num.I128

	assert.NoError(t,
		typeMap.Scan(UInt2OID, pgtype.TextFormatCode, []byte("65535"), &dst),
	)

	assert.Equal(t, u16MaxInS128, dst)
}

func TestUInt2Text_ScanInt(t *testing.T) {
	var dst int

	assert.NoError(t,
		typeMap.Scan(UInt2OID, pgtype.TextFormatCode, []byte("65535"), &dst),
	)

	assert.Equal(t, int(65535), dst)
}
