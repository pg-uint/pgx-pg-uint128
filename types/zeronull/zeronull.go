package zeronull

import (
	"github.com/jackc/pgx/v5/pgtype"
)

// Register registers the zeronull types so they can be used in query exec modes that do not know the server OIDs.
func Register(m *pgtype.Map) {
	m.RegisterDefaultPgType(UInt2(0), "uint2")
	m.RegisterDefaultPgType(UInt4(0), "uint4")
	m.RegisterDefaultPgType(UInt8(0), "uint8")

	// 128-bit types
	m.RegisterDefaultPgType(UInt16{}, "uint16")
	m.RegisterDefaultPgType(Int16{}, "int16")
}
