// Do not edit. Generated from codegen

package types_test

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/jackc/pgx/v5/pgtype"

	"lukechampine.com/uint128"

	"go.shabbyrobe.org/num"
)

func TestInt1Binary_ScanUint8(t *testing.T) {
	var dst uint8

	assert.NoError(t,
		typeMap.Scan(Int1OID, pgtype.BinaryFormatCode, []byte("\x7F"), &dst),
	)

	assert.Equal(t, uint8(127), dst)
}
func TestInt1Binary_ScanUint8_Underflow(t *testing.T) {
	var dst uint8

	assert.ErrorContains(t,
		typeMap.Scan(Int1OID, pgtype.BinaryFormatCode, []byte("\x80"), &dst),
		`-128 is less than minimum value for uint8`,
	)
}

func TestInt1Binary_ScanUint16(t *testing.T) {
	var dst uint16

	assert.NoError(t,
		typeMap.Scan(Int1OID, pgtype.BinaryFormatCode, []byte("\x7F"), &dst),
	)

	assert.Equal(t, uint16(127), dst)
}
func TestInt1Binary_ScanUint16_Underflow(t *testing.T) {
	var dst uint16

	assert.ErrorContains(t,
		typeMap.Scan(Int1OID, pgtype.BinaryFormatCode, []byte("\x80"), &dst),
		`-128 is less than minimum value for uint16`,
	)
}

func TestInt1Binary_ScanUint32(t *testing.T) {
	var dst uint32

	assert.NoError(t,
		typeMap.Scan(Int1OID, pgtype.BinaryFormatCode, []byte("\x7F"), &dst),
	)

	assert.Equal(t, uint32(127), dst)
}
func TestInt1Binary_ScanUint32_Underflow(t *testing.T) {
	var dst uint32

	assert.ErrorContains(t,
		typeMap.Scan(Int1OID, pgtype.BinaryFormatCode, []byte("\x80"), &dst),
		`-128 is less than minimum value for uint32`,
	)
}

func TestInt1Binary_ScanUint64(t *testing.T) {
	var dst uint64

	assert.NoError(t,
		typeMap.Scan(Int1OID, pgtype.BinaryFormatCode, []byte("\x7F"), &dst),
	)

	assert.Equal(t, uint64(127), dst)
}
func TestInt1Binary_ScanUint64_Underflow(t *testing.T) {
	var dst uint64

	assert.ErrorContains(t,
		typeMap.Scan(Int1OID, pgtype.BinaryFormatCode, []byte("\x80"), &dst),
		`-128 is less than minimum value for uint64`,
	)
}

func TestInt1Binary_ScanUint128(t *testing.T) {
	var dst uint128.Uint128

	assert.NoError(t,
		typeMap.Scan(Int1OID, pgtype.BinaryFormatCode, []byte("\x7F"), &dst),
	)

	assert.Equal(t, s8MaxInU128, dst)
}
func TestInt1Binary_ScanUint128_Underflow(t *testing.T) {
	var dst uint128.Uint128

	assert.ErrorContains(t,
		typeMap.Scan(Int1OID, pgtype.BinaryFormatCode, []byte("\x80"), &dst),
		`-128 is less than minimum value for Uint128`,
	)
}

func TestInt1Binary_ScanUint(t *testing.T) {
	var dst uint

	assert.NoError(t,
		typeMap.Scan(Int1OID, pgtype.BinaryFormatCode, []byte("\x7F"), &dst),
	)

	assert.Equal(t, uint(127), dst)
}
func TestInt1Binary_ScanUint_Underflow(t *testing.T) {
	var dst uint

	assert.ErrorContains(t,
		typeMap.Scan(Int1OID, pgtype.BinaryFormatCode, []byte("\x80"), &dst),
		`-128 is less than minimum value for uint`,
	)
}

func TestInt1Binary_ScanInt8(t *testing.T) {
	var dst int8

	assert.NoError(t,
		typeMap.Scan(Int1OID, pgtype.BinaryFormatCode, []byte("\x7F"), &dst),
	)

	assert.Equal(t, int8(127), dst)
}

func TestInt1Binary_ScanInt16(t *testing.T) {
	var dst int16

	assert.NoError(t,
		typeMap.Scan(Int1OID, pgtype.BinaryFormatCode, []byte("\x7F"), &dst),
	)

	assert.Equal(t, int16(127), dst)
}

func TestInt1Binary_ScanInt32(t *testing.T) {
	var dst int32

	assert.NoError(t,
		typeMap.Scan(Int1OID, pgtype.BinaryFormatCode, []byte("\x7F"), &dst),
	)

	assert.Equal(t, int32(127), dst)
}

func TestInt1Binary_ScanInt64(t *testing.T) {
	var dst int64

	assert.NoError(t,
		typeMap.Scan(Int1OID, pgtype.BinaryFormatCode, []byte("\x7F"), &dst),
	)

	assert.Equal(t, int64(127), dst)
}

func TestInt1Binary_ScanI128(t *testing.T) {
	var dst num.I128

	assert.NoError(t,
		typeMap.Scan(Int1OID, pgtype.BinaryFormatCode, []byte("\x7F"), &dst),
	)

	assert.Equal(t, s8MaxInS128, dst)
}

func TestInt1Binary_ScanInt(t *testing.T) {
	var dst int

	assert.NoError(t,
		typeMap.Scan(Int1OID, pgtype.BinaryFormatCode, []byte("\x7F"), &dst),
	)

	assert.Equal(t, int(127), dst)
}

func TestInt1Text_ScanUint8(t *testing.T) {
	var dst uint8

	assert.NoError(t,
		typeMap.Scan(Int1OID, pgtype.TextFormatCode, []byte("127"), &dst),
	)

	assert.Equal(t, uint8(127), dst)
}
func TestInt1Text_ScanUint8_Underflow(t *testing.T) {
	var dst uint8

	assert.ErrorContains(t,
		typeMap.Scan(Int1OID, pgtype.TextFormatCode, []byte("-128"), &dst),
		`strconv.ParseUint: parsing "-128": invalid syntax`,
	)
}

func TestInt1Text_ScanUint16(t *testing.T) {
	var dst uint16

	assert.NoError(t,
		typeMap.Scan(Int1OID, pgtype.TextFormatCode, []byte("127"), &dst),
	)

	assert.Equal(t, uint16(127), dst)
}
func TestInt1Text_ScanUint16_Underflow(t *testing.T) {
	var dst uint16

	assert.ErrorContains(t,
		typeMap.Scan(Int1OID, pgtype.TextFormatCode, []byte("-128"), &dst),
		`strconv.ParseUint: parsing "-128": invalid syntax`,
	)
}

func TestInt1Text_ScanUint32(t *testing.T) {
	var dst uint32

	assert.NoError(t,
		typeMap.Scan(Int1OID, pgtype.TextFormatCode, []byte("127"), &dst),
	)

	assert.Equal(t, uint32(127), dst)
}
func TestInt1Text_ScanUint32_Underflow(t *testing.T) {
	var dst uint32

	assert.ErrorContains(t,
		typeMap.Scan(Int1OID, pgtype.TextFormatCode, []byte("-128"), &dst),
		`strconv.ParseUint: parsing "-128": invalid syntax`,
	)
}

func TestInt1Text_ScanUint64(t *testing.T) {
	var dst uint64

	assert.NoError(t,
		typeMap.Scan(Int1OID, pgtype.TextFormatCode, []byte("127"), &dst),
	)

	assert.Equal(t, uint64(127), dst)
}
func TestInt1Text_ScanUint64_Underflow(t *testing.T) {
	var dst uint64

	assert.ErrorContains(t,
		typeMap.Scan(Int1OID, pgtype.TextFormatCode, []byte("-128"), &dst),
		`strconv.ParseUint: parsing "-128": invalid syntax`,
	)
}

func TestInt1Text_ScanUint128(t *testing.T) {
	var dst uint128.Uint128

	assert.NoError(t,
		typeMap.Scan(Int1OID, pgtype.TextFormatCode, []byte("127"), &dst),
	)

	assert.Equal(t, s8MaxInU128, dst)
}
func TestInt1Text_ScanUint128_Underflow(t *testing.T) {
	var dst uint128.Uint128

	assert.ErrorContains(t,
		typeMap.Scan(Int1OID, pgtype.TextFormatCode, []byte("-128"), &dst),
		`value cannot be negative`,
	)
}

func TestInt1Text_ScanUint(t *testing.T) {
	var dst uint

	assert.NoError(t,
		typeMap.Scan(Int1OID, pgtype.TextFormatCode, []byte("127"), &dst),
	)

	assert.Equal(t, uint(127), dst)
}
func TestInt1Text_ScanUint_Underflow(t *testing.T) {
	var dst uint

	assert.ErrorContains(t,
		typeMap.Scan(Int1OID, pgtype.TextFormatCode, []byte("-128"), &dst),
		`strconv.ParseUint: parsing "-128": invalid syntax`,
	)
}

func TestInt1Text_ScanInt8(t *testing.T) {
	var dst int8

	assert.NoError(t,
		typeMap.Scan(Int1OID, pgtype.TextFormatCode, []byte("127"), &dst),
	)

	assert.Equal(t, int8(127), dst)
}
func TestInt1Text_ScanInt8_SelfUnderflow(t *testing.T) {
	var dst int8

	assert.ErrorContains(t,
		typeMap.Scan(Int1OID, pgtype.TextFormatCode, []byte("-129"), &dst),
		`strconv.ParseInt: parsing "-129": value out of range`,
	)
}
func TestInt1Text_ScanInt8_SelfOverflow(t *testing.T) {
	var dst int8

	assert.ErrorContains(t,
		typeMap.Scan(Int1OID, pgtype.TextFormatCode, []byte("128"), &dst),
		`strconv.ParseInt: parsing "128": value out of range`,
	)
}

func TestInt1Text_ScanInt16(t *testing.T) {
	var dst int16

	assert.NoError(t,
		typeMap.Scan(Int1OID, pgtype.TextFormatCode, []byte("127"), &dst),
	)

	assert.Equal(t, int16(127), dst)
}

func TestInt1Text_ScanInt32(t *testing.T) {
	var dst int32

	assert.NoError(t,
		typeMap.Scan(Int1OID, pgtype.TextFormatCode, []byte("127"), &dst),
	)

	assert.Equal(t, int32(127), dst)
}

func TestInt1Text_ScanInt64(t *testing.T) {
	var dst int64

	assert.NoError(t,
		typeMap.Scan(Int1OID, pgtype.TextFormatCode, []byte("127"), &dst),
	)

	assert.Equal(t, int64(127), dst)
}

func TestInt1Text_ScanI128(t *testing.T) {
	var dst num.I128

	assert.NoError(t,
		typeMap.Scan(Int1OID, pgtype.TextFormatCode, []byte("127"), &dst),
	)

	assert.Equal(t, s8MaxInS128, dst)
}

func TestInt1Text_ScanInt(t *testing.T) {
	var dst int

	assert.NoError(t,
		typeMap.Scan(Int1OID, pgtype.TextFormatCode, []byte("127"), &dst),
	)

	assert.Equal(t, int(127), dst)
}
