// Do not edit. Generated from codegen

package types_test

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/jackc/pgx/v5/pgtype"

	"lukechampine.com/uint128"

	"go.shabbyrobe.org/num"
)

func TestInt16Binary_ScanUint8(t *testing.T) {
	var dst uint8

	assert.NoError(t,
		typeMap.Scan(Int16OID, pgtype.BinaryFormatCode, []byte("\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\xFF"), &dst),
	)

	assert.Equal(t, uint8(255), dst)
}
func TestInt16Binary_ScanUint8_Overflow(t *testing.T) {
	var dst uint8

	assert.ErrorContains(t,
		typeMap.Scan(Int16OID, pgtype.BinaryFormatCode, []byte("\x7F\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF"), &dst),
		`170141183460469231731687303715884105727 is greater than maximum value for uint8`,
	)
}
func TestInt16Binary_ScanUint8_Underflow(t *testing.T) {
	var dst uint8

	assert.ErrorContains(t,
		typeMap.Scan(Int16OID, pgtype.BinaryFormatCode, []byte("\x80\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), &dst),
		`-170141183460469231731687303715884105728 is less than minimum value for uint8`,
	)
}

func TestInt16Binary_ScanUint16(t *testing.T) {
	var dst uint16

	assert.NoError(t,
		typeMap.Scan(Int16OID, pgtype.BinaryFormatCode, []byte("\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\xFF\xFF"), &dst),
	)

	assert.Equal(t, uint16(65535), dst)
}
func TestInt16Binary_ScanUint16_Overflow(t *testing.T) {
	var dst uint16

	assert.ErrorContains(t,
		typeMap.Scan(Int16OID, pgtype.BinaryFormatCode, []byte("\x7F\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF"), &dst),
		`170141183460469231731687303715884105727 is greater than maximum value for uint16`,
	)
}
func TestInt16Binary_ScanUint16_Underflow(t *testing.T) {
	var dst uint16

	assert.ErrorContains(t,
		typeMap.Scan(Int16OID, pgtype.BinaryFormatCode, []byte("\x80\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), &dst),
		`-170141183460469231731687303715884105728 is less than minimum value for uint16`,
	)
}

func TestInt16Binary_ScanUint32(t *testing.T) {
	var dst uint32

	assert.NoError(t,
		typeMap.Scan(Int16OID, pgtype.BinaryFormatCode, []byte("\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\xFF\xFF\xFF\xFF"), &dst),
	)

	assert.Equal(t, uint32(4294967295), dst)
}
func TestInt16Binary_ScanUint32_Overflow(t *testing.T) {
	var dst uint32

	assert.ErrorContains(t,
		typeMap.Scan(Int16OID, pgtype.BinaryFormatCode, []byte("\x7F\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF"), &dst),
		`170141183460469231731687303715884105727 is greater than maximum value for uint32`,
	)
}
func TestInt16Binary_ScanUint32_Underflow(t *testing.T) {
	var dst uint32

	assert.ErrorContains(t,
		typeMap.Scan(Int16OID, pgtype.BinaryFormatCode, []byte("\x80\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), &dst),
		`-170141183460469231731687303715884105728 is less than minimum value for uint32`,
	)
}

func TestInt16Binary_ScanUint64(t *testing.T) {
	var dst uint64

	assert.NoError(t,
		typeMap.Scan(Int16OID, pgtype.BinaryFormatCode, []byte("\x00\x00\x00\x00\x00\x00\x00\x00\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF"), &dst),
	)

	assert.Equal(t, uint64(18446744073709551615), dst)
}
func TestInt16Binary_ScanUint64_Overflow(t *testing.T) {
	var dst uint64

	assert.ErrorContains(t,
		typeMap.Scan(Int16OID, pgtype.BinaryFormatCode, []byte("\x7F\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF"), &dst),
		`170141183460469231731687303715884105727 is greater than maximum value for uint64`,
	)
}
func TestInt16Binary_ScanUint64_Underflow(t *testing.T) {
	var dst uint64

	assert.ErrorContains(t,
		typeMap.Scan(Int16OID, pgtype.BinaryFormatCode, []byte("\x80\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), &dst),
		`-170141183460469231731687303715884105728 is less than minimum value for uint64`,
	)
}

func TestInt16Binary_ScanUint128(t *testing.T) {
	var dst uint128.Uint128

	assert.NoError(t,
		typeMap.Scan(Int16OID, pgtype.BinaryFormatCode, []byte("\x7F\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF"), &dst),
	)

	assert.Equal(t, s128MaxInU128, dst)
}
func TestInt16Binary_ScanUint128_Underflow(t *testing.T) {
	var dst uint128.Uint128

	assert.ErrorContains(t,
		typeMap.Scan(Int16OID, pgtype.BinaryFormatCode, []byte("\x80\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), &dst),
		`-170141183460469231731687303715884105728 is less than minimum value for Uint128`,
	)
}

func TestInt16Binary_ScanUint(t *testing.T) {
	var dst uint

	assert.NoError(t,
		typeMap.Scan(Int16OID, pgtype.BinaryFormatCode, []byte("\x00\x00\x00\x00\x00\x00\x00\x00\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF"), &dst),
	)

	assert.Equal(t, uint(18446744073709551615), dst)
}
func TestInt16Binary_ScanUint_Overflow(t *testing.T) {
	var dst uint

	assert.ErrorContains(t,
		typeMap.Scan(Int16OID, pgtype.BinaryFormatCode, []byte("\x7F\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF"), &dst),
		`170141183460469231731687303715884105727 is greater than maximum value for uint`,
	)
}
func TestInt16Binary_ScanUint_Underflow(t *testing.T) {
	var dst uint

	assert.ErrorContains(t,
		typeMap.Scan(Int16OID, pgtype.BinaryFormatCode, []byte("\x80\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), &dst),
		`-170141183460469231731687303715884105728 is less than minimum value for uint`,
	)
}

func TestInt16Binary_ScanInt8(t *testing.T) {
	var dst int8

	assert.NoError(t,
		typeMap.Scan(Int16OID, pgtype.BinaryFormatCode, []byte("\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x7F"), &dst),
	)

	assert.Equal(t, int8(127), dst)
}
func TestInt16Binary_ScanInt8_Overflow(t *testing.T) {
	var dst int8

	assert.ErrorContains(t,
		typeMap.Scan(Int16OID, pgtype.BinaryFormatCode, []byte("\x7F\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF"), &dst),
		`170141183460469231731687303715884105727 is greater than maximum value for int8`,
	)
}
func TestInt16Binary_ScanInt8_Underflow(t *testing.T) {
	var dst int8

	assert.ErrorContains(t,
		typeMap.Scan(Int16OID, pgtype.BinaryFormatCode, []byte("\x80\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), &dst),
		`-170141183460469231731687303715884105728 is less than minimum value for int8`,
	)
}

func TestInt16Binary_ScanInt16(t *testing.T) {
	var dst int16

	assert.NoError(t,
		typeMap.Scan(Int16OID, pgtype.BinaryFormatCode, []byte("\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x7F\xFF"), &dst),
	)

	assert.Equal(t, int16(32767), dst)
}
func TestInt16Binary_ScanInt16_Overflow(t *testing.T) {
	var dst int16

	assert.ErrorContains(t,
		typeMap.Scan(Int16OID, pgtype.BinaryFormatCode, []byte("\x7F\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF"), &dst),
		`170141183460469231731687303715884105727 is greater than maximum value for int16`,
	)
}
func TestInt16Binary_ScanInt16_Underflow(t *testing.T) {
	var dst int16

	assert.ErrorContains(t,
		typeMap.Scan(Int16OID, pgtype.BinaryFormatCode, []byte("\x80\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), &dst),
		`-170141183460469231731687303715884105728 is less than minimum value for int16`,
	)
}

func TestInt16Binary_ScanInt32(t *testing.T) {
	var dst int32

	assert.NoError(t,
		typeMap.Scan(Int16OID, pgtype.BinaryFormatCode, []byte("\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x7F\xFF\xFF\xFF"), &dst),
	)

	assert.Equal(t, int32(2147483647), dst)
}
func TestInt16Binary_ScanInt32_Overflow(t *testing.T) {
	var dst int32

	assert.ErrorContains(t,
		typeMap.Scan(Int16OID, pgtype.BinaryFormatCode, []byte("\x7F\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF"), &dst),
		`170141183460469231731687303715884105727 is greater than maximum value for int32`,
	)
}
func TestInt16Binary_ScanInt32_Underflow(t *testing.T) {
	var dst int32

	assert.ErrorContains(t,
		typeMap.Scan(Int16OID, pgtype.BinaryFormatCode, []byte("\x80\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), &dst),
		`-170141183460469231731687303715884105728 is less than minimum value for int32`,
	)
}

func TestInt16Binary_ScanInt64(t *testing.T) {
	var dst int64

	assert.NoError(t,
		typeMap.Scan(Int16OID, pgtype.BinaryFormatCode, []byte("\x00\x00\x00\x00\x00\x00\x00\x00\x7F\xFF\xFF\xFF\xFF\xFF\xFF\xFF"), &dst),
	)

	assert.Equal(t, int64(9223372036854775807), dst)
}
func TestInt16Binary_ScanInt64_Overflow(t *testing.T) {
	var dst int64

	assert.ErrorContains(t,
		typeMap.Scan(Int16OID, pgtype.BinaryFormatCode, []byte("\x7F\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF"), &dst),
		`170141183460469231731687303715884105727 is greater than maximum value for int64`,
	)
}
func TestInt16Binary_ScanInt64_Underflow(t *testing.T) {
	var dst int64

	assert.ErrorContains(t,
		typeMap.Scan(Int16OID, pgtype.BinaryFormatCode, []byte("\x80\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), &dst),
		`-170141183460469231731687303715884105728 is less than minimum value for int64`,
	)
}

func TestInt16Binary_ScanI128(t *testing.T) {
	var dst num.I128

	assert.NoError(t,
		typeMap.Scan(Int16OID, pgtype.BinaryFormatCode, []byte("\x7F\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF"), &dst),
	)

	assert.Equal(t, s128Max, dst)
}

func TestInt16Binary_ScanInt(t *testing.T) {
	var dst int

	assert.NoError(t,
		typeMap.Scan(Int16OID, pgtype.BinaryFormatCode, []byte("\x00\x00\x00\x00\x00\x00\x00\x00\x7F\xFF\xFF\xFF\xFF\xFF\xFF\xFF"), &dst),
	)

	assert.Equal(t, int(9223372036854775807), dst)
}
func TestInt16Binary_ScanInt_Overflow(t *testing.T) {
	var dst int

	assert.ErrorContains(t,
		typeMap.Scan(Int16OID, pgtype.BinaryFormatCode, []byte("\x7F\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF"), &dst),
		`170141183460469231731687303715884105727 is greater than maximum value for int`,
	)
}
func TestInt16Binary_ScanInt_Underflow(t *testing.T) {
	var dst int

	assert.ErrorContains(t,
		typeMap.Scan(Int16OID, pgtype.BinaryFormatCode, []byte("\x80\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"), &dst),
		`-170141183460469231731687303715884105728 is less than minimum value for int`,
	)
}

func TestInt16Text_ScanUint8(t *testing.T) {
	var dst uint8

	assert.NoError(t,
		typeMap.Scan(Int16OID, pgtype.TextFormatCode, []byte("255"), &dst),
	)

	assert.Equal(t, uint8(255), dst)
}
func TestInt16Text_ScanUint8_Overflow(t *testing.T) {
	var dst uint8

	assert.ErrorContains(t,
		typeMap.Scan(Int16OID, pgtype.TextFormatCode, []byte("170141183460469231731687303715884105727"), &dst),
		`strconv.ParseUint: parsing "170141183460469231731687303715884105727": value out of range`,
	)
}
func TestInt16Text_ScanUint8_Underflow(t *testing.T) {
	var dst uint8

	assert.ErrorContains(t,
		typeMap.Scan(Int16OID, pgtype.TextFormatCode, []byte("-170141183460469231731687303715884105728"), &dst),
		`strconv.ParseUint: parsing "-170141183460469231731687303715884105728": invalid syntax`,
	)
}

func TestInt16Text_ScanUint16(t *testing.T) {
	var dst uint16

	assert.NoError(t,
		typeMap.Scan(Int16OID, pgtype.TextFormatCode, []byte("65535"), &dst),
	)

	assert.Equal(t, uint16(65535), dst)
}
func TestInt16Text_ScanUint16_Overflow(t *testing.T) {
	var dst uint16

	assert.ErrorContains(t,
		typeMap.Scan(Int16OID, pgtype.TextFormatCode, []byte("170141183460469231731687303715884105727"), &dst),
		`strconv.ParseUint: parsing "170141183460469231731687303715884105727": value out of range`,
	)
}
func TestInt16Text_ScanUint16_Underflow(t *testing.T) {
	var dst uint16

	assert.ErrorContains(t,
		typeMap.Scan(Int16OID, pgtype.TextFormatCode, []byte("-170141183460469231731687303715884105728"), &dst),
		`strconv.ParseUint: parsing "-170141183460469231731687303715884105728": invalid syntax`,
	)
}

func TestInt16Text_ScanUint32(t *testing.T) {
	var dst uint32

	assert.NoError(t,
		typeMap.Scan(Int16OID, pgtype.TextFormatCode, []byte("4294967295"), &dst),
	)

	assert.Equal(t, uint32(4294967295), dst)
}
func TestInt16Text_ScanUint32_Overflow(t *testing.T) {
	var dst uint32

	assert.ErrorContains(t,
		typeMap.Scan(Int16OID, pgtype.TextFormatCode, []byte("170141183460469231731687303715884105727"), &dst),
		`strconv.ParseUint: parsing "170141183460469231731687303715884105727": value out of range`,
	)
}
func TestInt16Text_ScanUint32_Underflow(t *testing.T) {
	var dst uint32

	assert.ErrorContains(t,
		typeMap.Scan(Int16OID, pgtype.TextFormatCode, []byte("-170141183460469231731687303715884105728"), &dst),
		`strconv.ParseUint: parsing "-170141183460469231731687303715884105728": invalid syntax`,
	)
}

func TestInt16Text_ScanUint64(t *testing.T) {
	var dst uint64

	assert.NoError(t,
		typeMap.Scan(Int16OID, pgtype.TextFormatCode, []byte("18446744073709551615"), &dst),
	)

	assert.Equal(t, uint64(18446744073709551615), dst)
}
func TestInt16Text_ScanUint64_Overflow(t *testing.T) {
	var dst uint64

	assert.ErrorContains(t,
		typeMap.Scan(Int16OID, pgtype.TextFormatCode, []byte("170141183460469231731687303715884105727"), &dst),
		`strconv.ParseUint: parsing "170141183460469231731687303715884105727": value out of range`,
	)
}
func TestInt16Text_ScanUint64_Underflow(t *testing.T) {
	var dst uint64

	assert.ErrorContains(t,
		typeMap.Scan(Int16OID, pgtype.TextFormatCode, []byte("-170141183460469231731687303715884105728"), &dst),
		`strconv.ParseUint: parsing "-170141183460469231731687303715884105728": invalid syntax`,
	)
}

func TestInt16Text_ScanUint128(t *testing.T) {
	var dst uint128.Uint128

	assert.NoError(t,
		typeMap.Scan(Int16OID, pgtype.TextFormatCode, []byte("170141183460469231731687303715884105727"), &dst),
	)

	assert.Equal(t, s128MaxInU128, dst)
}
func TestInt16Text_ScanUint128_Underflow(t *testing.T) {
	var dst uint128.Uint128

	assert.ErrorContains(t,
		typeMap.Scan(Int16OID, pgtype.TextFormatCode, []byte("-170141183460469231731687303715884105728"), &dst),
		`value cannot be negative`,
	)
}

func TestInt16Text_ScanUint(t *testing.T) {
	var dst uint

	assert.NoError(t,
		typeMap.Scan(Int16OID, pgtype.TextFormatCode, []byte("18446744073709551615"), &dst),
	)

	assert.Equal(t, uint(18446744073709551615), dst)
}
func TestInt16Text_ScanUint_Overflow(t *testing.T) {
	var dst uint

	assert.ErrorContains(t,
		typeMap.Scan(Int16OID, pgtype.TextFormatCode, []byte("170141183460469231731687303715884105727"), &dst),
		`strconv.ParseUint: parsing "170141183460469231731687303715884105727": value out of range`,
	)
}
func TestInt16Text_ScanUint_Underflow(t *testing.T) {
	var dst uint

	assert.ErrorContains(t,
		typeMap.Scan(Int16OID, pgtype.TextFormatCode, []byte("-170141183460469231731687303715884105728"), &dst),
		`strconv.ParseUint: parsing "-170141183460469231731687303715884105728": invalid syntax`,
	)
}

func TestInt16Text_ScanInt8(t *testing.T) {
	var dst int8

	assert.NoError(t,
		typeMap.Scan(Int16OID, pgtype.TextFormatCode, []byte("127"), &dst),
	)

	assert.Equal(t, int8(127), dst)
}
func TestInt16Text_ScanInt8_Overflow(t *testing.T) {
	var dst int8

	assert.ErrorContains(t,
		typeMap.Scan(Int16OID, pgtype.TextFormatCode, []byte("170141183460469231731687303715884105727"), &dst),
		`strconv.ParseInt: parsing "170141183460469231731687303715884105727": value out of range`,
	)
}
func TestInt16Text_ScanInt8_Underflow(t *testing.T) {
	var dst int8

	assert.ErrorContains(t,
		typeMap.Scan(Int16OID, pgtype.TextFormatCode, []byte("-170141183460469231731687303715884105728"), &dst),
		`strconv.ParseInt: parsing "-170141183460469231731687303715884105728": value out of range`,
	)
}

func TestInt16Text_ScanInt16(t *testing.T) {
	var dst int16

	assert.NoError(t,
		typeMap.Scan(Int16OID, pgtype.TextFormatCode, []byte("32767"), &dst),
	)

	assert.Equal(t, int16(32767), dst)
}
func TestInt16Text_ScanInt16_Overflow(t *testing.T) {
	var dst int16

	assert.ErrorContains(t,
		typeMap.Scan(Int16OID, pgtype.TextFormatCode, []byte("170141183460469231731687303715884105727"), &dst),
		`strconv.ParseInt: parsing "170141183460469231731687303715884105727": value out of range`,
	)
}
func TestInt16Text_ScanInt16_Underflow(t *testing.T) {
	var dst int16

	assert.ErrorContains(t,
		typeMap.Scan(Int16OID, pgtype.TextFormatCode, []byte("-170141183460469231731687303715884105728"), &dst),
		`strconv.ParseInt: parsing "-170141183460469231731687303715884105728": value out of range`,
	)
}

func TestInt16Text_ScanInt32(t *testing.T) {
	var dst int32

	assert.NoError(t,
		typeMap.Scan(Int16OID, pgtype.TextFormatCode, []byte("2147483647"), &dst),
	)

	assert.Equal(t, int32(2147483647), dst)
}
func TestInt16Text_ScanInt32_Overflow(t *testing.T) {
	var dst int32

	assert.ErrorContains(t,
		typeMap.Scan(Int16OID, pgtype.TextFormatCode, []byte("170141183460469231731687303715884105727"), &dst),
		`strconv.ParseInt: parsing "170141183460469231731687303715884105727": value out of range`,
	)
}
func TestInt16Text_ScanInt32_Underflow(t *testing.T) {
	var dst int32

	assert.ErrorContains(t,
		typeMap.Scan(Int16OID, pgtype.TextFormatCode, []byte("-170141183460469231731687303715884105728"), &dst),
		`strconv.ParseInt: parsing "-170141183460469231731687303715884105728": value out of range`,
	)
}

func TestInt16Text_ScanInt64(t *testing.T) {
	var dst int64

	assert.NoError(t,
		typeMap.Scan(Int16OID, pgtype.TextFormatCode, []byte("9223372036854775807"), &dst),
	)

	assert.Equal(t, int64(9223372036854775807), dst)
}
func TestInt16Text_ScanInt64_Overflow(t *testing.T) {
	var dst int64

	assert.ErrorContains(t,
		typeMap.Scan(Int16OID, pgtype.TextFormatCode, []byte("170141183460469231731687303715884105727"), &dst),
		`strconv.ParseInt: parsing "170141183460469231731687303715884105727": value out of range`,
	)
}
func TestInt16Text_ScanInt64_Underflow(t *testing.T) {
	var dst int64

	assert.ErrorContains(t,
		typeMap.Scan(Int16OID, pgtype.TextFormatCode, []byte("-170141183460469231731687303715884105728"), &dst),
		`strconv.ParseInt: parsing "-170141183460469231731687303715884105728": value out of range`,
	)
}

func TestInt16Text_ScanI128(t *testing.T) {
	var dst num.I128

	assert.NoError(t,
		typeMap.Scan(Int16OID, pgtype.TextFormatCode, []byte("170141183460469231731687303715884105727"), &dst),
	)

	assert.Equal(t, s128Max, dst)
}
func TestInt16Text_ScanI128_SelfUnderflow(t *testing.T) {
	var dst num.I128

	assert.ErrorContains(t,
		typeMap.Scan(Int16OID, pgtype.TextFormatCode, []byte("-170141183460469231731687303715884105729"), &dst),
		`int128.FromString: parsing "-170141183460469231731687303715884105729": value out of range`,
	)
}
func TestInt16Text_ScanI128_SelfOverflow(t *testing.T) {
	var dst num.I128

	assert.ErrorContains(t,
		typeMap.Scan(Int16OID, pgtype.TextFormatCode, []byte("170141183460469231731687303715884105728"), &dst),
		`int128.FromString: parsing "170141183460469231731687303715884105728": value out of range`,
	)
}

func TestInt16Text_ScanInt(t *testing.T) {
	var dst int

	assert.NoError(t,
		typeMap.Scan(Int16OID, pgtype.TextFormatCode, []byte("9223372036854775807"), &dst),
	)

	assert.Equal(t, int(9223372036854775807), dst)
}
func TestInt16Text_ScanInt_Overflow(t *testing.T) {
	var dst int

	assert.ErrorContains(t,
		typeMap.Scan(Int16OID, pgtype.TextFormatCode, []byte("170141183460469231731687303715884105727"), &dst),
		`strconv.ParseInt: parsing "170141183460469231731687303715884105727": value out of range`,
	)
}
func TestInt16Text_ScanInt_Underflow(t *testing.T) {
	var dst int

	assert.ErrorContains(t,
		typeMap.Scan(Int16OID, pgtype.TextFormatCode, []byte("-170141183460469231731687303715884105728"), &dst),
		`strconv.ParseInt: parsing "-170141183460469231731687303715884105728": value out of range`,
	)
}
