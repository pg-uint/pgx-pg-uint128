package pgio

import (
	"encoding/binary"

	"go.shabbyrobe.org/num"
	"lukechampine.com/uint128"
)

func ReadUint8(buf []byte) uint8 {
	_ = buf[0] // bounds check hint to compiler; see golang.org/issue/14808

	return buf[0]
}

func ReadUint16(buf []byte) uint16 {
	return binary.BigEndian.Uint16(buf)
}

func ReadUint32(buf []byte) uint32 {
	return binary.BigEndian.Uint32(buf)
}

func ReadUint64(buf []byte) uint64 {
	return binary.BigEndian.Uint64(buf)
}

func ReadUint128(buf []byte) uint128.Uint128 {
	_ = buf[15] // bounds check hint to compiler; see golang.org/issue/14808

	high := binary.BigEndian.Uint64(buf[0:8])
	low := binary.BigEndian.Uint64(buf[8:16])

	return uint128.New(low, high)
}

func ReadInt8(buf []byte) int8 {
	_ = buf[0] // bounds check hint to compiler; see golang.org/issue/14808

	return int8(buf[0])
}

func ReadInt128(buf []byte) num.I128 {
	_ = buf[15] // bounds check hint to compiler; see golang.org/issue/14808

	high := binary.BigEndian.Uint64(buf[0:8])
	low := binary.BigEndian.Uint64(buf[8:16])

	return num.I128FromRaw(high, low)
}
