//go:build !nopgxregisterdefaulttypes

package types

import (
	"go.shabbyrobe.org/num"
	"lukechampine.com/uint128"
)
import . "github.com/jackc/pgx/v5/pgtype"

func RegisterDefaultPgTypeVariants(tMap *Map) {
	registerDefaultPgTypeVariants[uint16](tMap, Uint2TypName)
	registerDefaultPgTypeVariants[uint32](tMap, Uint4TypName)
	registerDefaultPgTypeVariants[uint64](tMap, Uint8TypName)
	registerDefaultPgTypeVariants[uint128.Uint128](tMap, Uint16TypName)
	registerDefaultPgTypeVariants[num.I128](tMap, Int16TypName)

	registerDefaultPgTypeVariants[Range[uint16]](tMap, Uint2TypName+"range")
	registerDefaultPgTypeVariants[Range[uint32]](tMap, Uint4TypName+"range")
	registerDefaultPgTypeVariants[Range[uint64]](tMap, Uint8TypName+"range")
	registerDefaultPgTypeVariants[Range[uint128.Uint128]](tMap, Uint16TypName+"range")
	registerDefaultPgTypeVariants[Range[num.I128]](tMap, Int16TypName+"range")

	registerDefaultPgTypeVariants[Multirange[Range[uint16]]](tMap, Uint2TypName+"multirange")
	registerDefaultPgTypeVariants[Multirange[Range[uint32]]](tMap, Uint4TypName+"multirange")
	registerDefaultPgTypeVariants[Multirange[Range[uint64]]](tMap, Uint8TypName+"multirange")
	registerDefaultPgTypeVariants[Multirange[Range[uint128.Uint128]]](tMap, Uint16TypName+"multirange")
	registerDefaultPgTypeVariants[Multirange[Range[num.I128]]](tMap, Int16TypName+"multirange")
}

// registerDefaultPgTypeVariants copy of unexported function from PGX
func registerDefaultPgTypeVariants[T any](m *Map, name string) {
	arrayName := "_" + name

	var value T
	m.RegisterDefaultPgType(value, name)  // T
	m.RegisterDefaultPgType(&value, name) // *T

	var sliceT []T
	m.RegisterDefaultPgType(sliceT, arrayName)  // []T
	m.RegisterDefaultPgType(&sliceT, arrayName) // *[]T

	var slicePtrT []*T
	m.RegisterDefaultPgType(slicePtrT, arrayName)  // []*T
	m.RegisterDefaultPgType(&slicePtrT, arrayName) // *[]*T

	var arrayOfT Array[T]
	m.RegisterDefaultPgType(arrayOfT, arrayName)  // Array[T]
	m.RegisterDefaultPgType(&arrayOfT, arrayName) // *Array[T]

	var arrayOfPtrT Array[*T]
	m.RegisterDefaultPgType(arrayOfPtrT, arrayName)  // Array[*T]
	m.RegisterDefaultPgType(&arrayOfPtrT, arrayName) // *Array[*T]

	var flatArrayOfT FlatArray[T]
	m.RegisterDefaultPgType(flatArrayOfT, arrayName)  // FlatArray[T]
	m.RegisterDefaultPgType(&flatArrayOfT, arrayName) // *FlatArray[T]

	var flatArrayOfPtrT FlatArray[*T]
	m.RegisterDefaultPgType(flatArrayOfPtrT, arrayName)  // FlatArray[*T]
	m.RegisterDefaultPgType(&flatArrayOfPtrT, arrayName) // *FlatArray[*T]
}
