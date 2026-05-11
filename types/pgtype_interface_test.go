package types_test

import (
	"testing"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testUint32Valuer struct {
	v pgtype.Uint32
}

func (t testUint32Valuer) Uint32Value() (pgtype.Uint32, error) {
	return t.v, nil
}

type testUint64Valuer struct {
	v pgtype.Uint64
}

func (t testUint64Valuer) Uint64Value() (pgtype.Uint64, error) {
	return t.v, nil
}

func TestUInt4Codec_PGTypeUint32Valuer_EncodeBinaryAndText(t *testing.T) {
	bin, err := typeMap.Encode(UInt4OID, pgtype.BinaryFormatCode, testUint32Valuer{
		v: pgtype.Uint32{Uint32: 12345, Valid: true},
	}, nil)
	require.NoError(t, err)
	assert.Equal(t, []byte{0x00, 0x00, 0x30, 0x39}, bin)

	txt, err := typeMap.Encode(UInt4OID, pgtype.TextFormatCode, testUint32Valuer{
		v: pgtype.Uint32{Uint32: 12345, Valid: true},
	}, nil)
	require.NoError(t, err)
	assert.Equal(t, []byte("12345"), txt)
}

func TestUInt8Codec_PGTypeUint64Valuer_EncodeBinaryAndText(t *testing.T) {
	bin, err := typeMap.Encode(UInt8OID, pgtype.BinaryFormatCode, testUint64Valuer{
		v: pgtype.Uint64{Uint64: 12345, Valid: true},
	}, nil)
	require.NoError(t, err)
	assert.Equal(t, []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x30, 0x39}, bin)

	txt, err := typeMap.Encode(UInt8OID, pgtype.TextFormatCode, testUint64Valuer{
		v: pgtype.Uint64{Uint64: 12345, Valid: true},
	}, nil)
	require.NoError(t, err)
	assert.Equal(t, []byte("12345"), txt)
}

func TestUInt4Codec_PGTypeUint32Scanner_ScanBinaryAndText(t *testing.T) {
	var dst pgtype.Uint32

	err := typeMap.Scan(UInt4OID, pgtype.BinaryFormatCode, []byte{0x00, 0x00, 0x30, 0x39}, &dst)
	require.NoError(t, err)
	assert.Equal(t, pgtype.Uint32{Uint32: 12345, Valid: true}, dst)

	dst = pgtype.Uint32{}
	err = typeMap.Scan(UInt4OID, pgtype.TextFormatCode, []byte("12345"), &dst)
	require.NoError(t, err)
	assert.Equal(t, pgtype.Uint32{Uint32: 12345, Valid: true}, dst)
}

func TestUInt8Codec_PGTypeUint64Scanner_ScanBinaryAndText(t *testing.T) {
	var dst pgtype.Uint64

	err := typeMap.Scan(UInt8OID, pgtype.BinaryFormatCode, []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x30, 0x39}, &dst)
	require.NoError(t, err)
	assert.Equal(t, pgtype.Uint64{Uint64: 12345, Valid: true}, dst)

	dst = pgtype.Uint64{}
	err = typeMap.Scan(UInt8OID, pgtype.TextFormatCode, []byte("12345"), &dst)
	require.NoError(t, err)
	assert.Equal(t, pgtype.Uint64{Uint64: 12345, Valid: true}, dst)
}
