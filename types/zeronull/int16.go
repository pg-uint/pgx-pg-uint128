// Do not edit. Generated from codegen

package zeronull

import (
	"database/sql/driver"

	"github.com/pg-uint/pgx-pg-uint128/types"
	"go.shabbyrobe.org/num"
)

type Int16 num.I128

func (Int16) SkipUnderlyingTypePlan() {}

// ScanInt64 implements the Int64Scanner interface.
func (dst *Int16) ScanInt64(n int64, valid bool) error {
	if !valid {
		*dst = Int16{}
		return nil
	}

	*dst = Int16(num.I128From64(n))

	return nil
}

// ScanUint64 implements the Uint64Scanner interface.
func (dst *Int16) ScanUint64(n uint64, valid bool) error {
	if !valid {
		*dst = Int16{}
		return nil
	}

	*dst = Int16(num.I128FromU64(n))

	return nil
}

// Scan implements the database/sql Scanner interface.
func (dst *Int16) Scan(src any) error {
	if src == nil {
		*dst = Int16{}
		return nil
	}

	var nullable types.Int16
	err := nullable.Scan(src)
	if err != nil {
		return err
	}

	*dst = Int16(nullable.I128)

	return nil
}

// Value implements the database/sql/driver Valuer interface.
func (src Int16) Value() (driver.Value, error) {
	if num.I128(src).IsZero() {
		return nil, nil
	}
	return num.I128(src), nil
}
