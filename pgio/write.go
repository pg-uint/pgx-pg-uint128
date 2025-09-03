package pgio

import (
	"encoding/binary"

	"go.shabbyrobe.org/num"
	"lukechampine.com/uint128"
)

func AppendUint8(buf []byte, n uint8) []byte {
	wp := len(buf)
	buf = append(buf, 0)
	buf[wp] = n
	return buf
}

func AppendUint16(buf []byte, n uint16) []byte {
	wp := len(buf)
	buf = append(buf, 0, 0)
	binary.BigEndian.PutUint16(buf[wp:], n)
	return buf
}

func AppendUint32(buf []byte, n uint32) []byte {
	wp := len(buf)
	buf = append(buf, 0, 0, 0, 0)
	binary.BigEndian.PutUint32(buf[wp:], n)
	return buf
}

func AppendUint64(buf []byte, n uint64) []byte {
	wp := len(buf)
	buf = append(buf, 0, 0, 0, 0, 0, 0, 0, 0)
	binary.BigEndian.PutUint64(buf[wp:], n)
	return buf
}

func AppendUint128(buf []byte, n uint128.Uint128) []byte {
	wp := len(buf)
	buf = append(buf, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0)

	binary.BigEndian.PutUint64(buf[wp:], n.Hi)

	wp += 8
	binary.BigEndian.PutUint64(buf[wp:], n.Lo)

	return buf
}

func AppendInt8(buf []byte, n int8) []byte {
	wp := len(buf)
	buf = append(buf, 0)
	buf[wp] = uint8(n)
	return buf
}

func AppendInt128(buf []byte, n num.I128) []byte {
	wp := len(buf)
	buf = append(buf, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0)

	high, low := n.Raw()

	binary.BigEndian.PutUint64(buf[wp:], high)

	wp += 8
	binary.BigEndian.PutUint64(buf[wp:], low)

	return buf
}
