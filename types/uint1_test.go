// Do not edit. Generated from codegen

package types_test

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/jackc/pgx/v5/pgtype"

	"lukechampine.com/uint128"

	"go.shabbyrobe.org/num"
)

func TestUInt1Binary_ScanUint8(t *testing.T) {
	var dst uint8

	assert.NoError(t,
		typeMap.Scan(UInt1OID, pgtype.BinaryFormatCode, []byte("\xFF"), &dst),
	)

	assert.Equal(t, uint8(255), dst)
}

func TestUInt1Binary_ScanUint16(t *testing.T) {
	var dst uint16

	assert.NoError(t,
		typeMap.Scan(UInt1OID, pgtype.BinaryFormatCode, []byte("\xFF"), &dst),
	)

	assert.Equal(t, uint16(255), dst)
}

func TestUInt1Binary_ScanUint32(t *testing.T) {
	var dst uint32

	assert.NoError(t,
		typeMap.Scan(UInt1OID, pgtype.BinaryFormatCode, []byte("\xFF"), &dst),
	)

	assert.Equal(t, uint32(255), dst)
}

func TestUInt1Binary_ScanUint64(t *testing.T) {
	var dst uint64

	assert.NoError(t,
		typeMap.Scan(UInt1OID, pgtype.BinaryFormatCode, []byte("\xFF"), &dst),
	)

	assert.Equal(t, uint64(255), dst)
}

func TestUInt1Binary_ScanUint128(t *testing.T) {
	var dst uint128.Uint128

	assert.NoError(t,
		typeMap.Scan(UInt1OID, pgtype.BinaryFormatCode, []byte("\xFF"), &dst),
	)

	assert.Equal(t, u8MaxInU128, dst)
}

func TestUInt1Binary_ScanUint(t *testing.T) {
	var dst uint

	assert.NoError(t,
		typeMap.Scan(UInt1OID, pgtype.BinaryFormatCode, []byte("\xFF"), &dst),
	)

	assert.Equal(t, uint(255), dst)
}

func TestUInt1Binary_ScanInt8(t *testing.T) {
	var dst int8

	assert.NoError(t,
		typeMap.Scan(UInt1OID, pgtype.BinaryFormatCode, []byte("\x7F"), &dst),
	)

	assert.Equal(t, int8(127), dst)
}
func TestUInt1Binary_ScanInt8_Overflow(t *testing.T) {
	var dst int8

	assert.ErrorContains(t,
		typeMap.Scan(UInt1OID, pgtype.BinaryFormatCode, []byte("\xFF"), &dst),
		`255 is greater than maximum value for int8`,
	)
}

func TestUInt1Binary_ScanInt16(t *testing.T) {
	var dst int16

	assert.NoError(t,
		typeMap.Scan(UInt1OID, pgtype.BinaryFormatCode, []byte("\xFF"), &dst),
	)

	assert.Equal(t, int16(255), dst)
}

func TestUInt1Binary_ScanInt32(t *testing.T) {
	var dst int32

	assert.NoError(t,
		typeMap.Scan(UInt1OID, pgtype.BinaryFormatCode, []byte("\xFF"), &dst),
	)

	assert.Equal(t, int32(255), dst)
}

func TestUInt1Binary_ScanInt64(t *testing.T) {
	var dst int64

	assert.NoError(t,
		typeMap.Scan(UInt1OID, pgtype.BinaryFormatCode, []byte("\xFF"), &dst),
	)

	assert.Equal(t, int64(255), dst)
}

func TestUInt1Binary_ScanI128(t *testing.T) {
	var dst num.I128

	assert.NoError(t,
		typeMap.Scan(UInt1OID, pgtype.BinaryFormatCode, []byte("\xFF"), &dst),
	)

	assert.Equal(t, u8MaxInS128, dst)
}

func TestUInt1Binary_ScanInt(t *testing.T) {
	var dst int

	assert.NoError(t,
		typeMap.Scan(UInt1OID, pgtype.BinaryFormatCode, []byte("\xFF"), &dst),
	)

	assert.Equal(t, int(255), dst)
}

func TestUInt1Text_ScanUint8(t *testing.T) {
	var dst uint8

	assert.NoError(t,
		typeMap.Scan(UInt1OID, pgtype.TextFormatCode, []byte("255"), &dst),
	)

	assert.Equal(t, uint8(255), dst)
}
func TestUInt1Text_ScanUint8_SelfOverflow(t *testing.T) {
	var dst uint8

	assert.ErrorContains(t,
		typeMap.Scan(UInt1OID, pgtype.TextFormatCode, []byte("256"), &dst),
		`strconv.ParseUint: parsing "256": value out of range`,
	)
}

func TestUInt1Text_ScanUint16(t *testing.T) {
	var dst uint16

	assert.NoError(t,
		typeMap.Scan(UInt1OID, pgtype.TextFormatCode, []byte("255"), &dst),
	)

	assert.Equal(t, uint16(255), dst)
}

func TestUInt1Text_ScanUint32(t *testing.T) {
	var dst uint32

	assert.NoError(t,
		typeMap.Scan(UInt1OID, pgtype.TextFormatCode, []byte("255"), &dst),
	)

	assert.Equal(t, uint32(255), dst)
}

func TestUInt1Text_ScanUint64(t *testing.T) {
	var dst uint64

	assert.NoError(t,
		typeMap.Scan(UInt1OID, pgtype.TextFormatCode, []byte("255"), &dst),
	)

	assert.Equal(t, uint64(255), dst)
}

func TestUInt1Text_ScanUint128(t *testing.T) {
	var dst uint128.Uint128

	assert.NoError(t,
		typeMap.Scan(UInt1OID, pgtype.TextFormatCode, []byte("255"), &dst),
	)

	assert.Equal(t, u8MaxInU128, dst)
}

func TestUInt1Text_ScanUint(t *testing.T) {
	var dst uint

	assert.NoError(t,
		typeMap.Scan(UInt1OID, pgtype.TextFormatCode, []byte("255"), &dst),
	)

	assert.Equal(t, uint(255), dst)
}

func TestUInt1Text_ScanInt8(t *testing.T) {
	var dst int8

	assert.NoError(t,
		typeMap.Scan(UInt1OID, pgtype.TextFormatCode, []byte("127"), &dst),
	)

	assert.Equal(t, int8(127), dst)
}
func TestUInt1Text_ScanInt8_Overflow(t *testing.T) {
	var dst int8

	assert.ErrorContains(t,
		typeMap.Scan(UInt1OID, pgtype.TextFormatCode, []byte("255"), &dst),
		`strconv.ParseInt: parsing "255": value out of range`,
	)
}

func TestUInt1Text_ScanInt16(t *testing.T) {
	var dst int16

	assert.NoError(t,
		typeMap.Scan(UInt1OID, pgtype.TextFormatCode, []byte("255"), &dst),
	)

	assert.Equal(t, int16(255), dst)
}

func TestUInt1Text_ScanInt32(t *testing.T) {
	var dst int32

	assert.NoError(t,
		typeMap.Scan(UInt1OID, pgtype.TextFormatCode, []byte("255"), &dst),
	)

	assert.Equal(t, int32(255), dst)
}

func TestUInt1Text_ScanInt64(t *testing.T) {
	var dst int64

	assert.NoError(t,
		typeMap.Scan(UInt1OID, pgtype.TextFormatCode, []byte("255"), &dst),
	)

	assert.Equal(t, int64(255), dst)
}

func TestUInt1Text_ScanI128(t *testing.T) {
	var dst num.I128

	assert.NoError(t,
		typeMap.Scan(UInt1OID, pgtype.TextFormatCode, []byte("255"), &dst),
	)

	assert.Equal(t, u8MaxInS128, dst)
}

func TestUInt1Text_ScanInt(t *testing.T) {
	var dst int

	assert.NoError(t,
		typeMap.Scan(UInt1OID, pgtype.TextFormatCode, []byte("255"), &dst),
	)

	assert.Equal(t, int(255), dst)
}
