// Do not edit. Generated from codegen

package zeronull

import (
	"database/sql/driver"
	"fmt"

	"github.com/pg-uint/pgx-pg-uint128/types"
)

type UInt8 uint64

func (UInt8) SkipUnderlyingTypePlan() {}

// ScanInt64 implements the Int64Scanner interface.
func (dst *UInt8) ScanInt64(n int64, valid bool) error {
	if !valid {
		*dst = 0
		return nil
	}

	if n < 0 {
		return fmt.Errorf("%d is less than minimum value for UInt8", n)
	}

	*dst = UInt8(n)

	return nil
}

// ScanUint64 implements the Uint64Scanner interface.
func (dst *UInt8) ScanUint64(n uint64, valid bool) error {
	if !valid {
		*dst = 0
		return nil
	}

	*dst = UInt8(n)

	return nil
}

// Scan implements the database/sql Scanner interface.
func (dst *UInt8) Scan(src any) error {
	if src == nil {
		*dst = 0
		return nil
	}

	var nullable types.UInt8
	err := nullable.Scan(src)
	if err != nil {
		return err
	}

	*dst = UInt8(nullable.Uint64)

	return nil
}

// Value implements the database/sql/driver Valuer interface.
func (src UInt8) Value() (driver.Value, error) {
	if src == 0 {
		return nil, nil
	}
	return uint64(src), nil
}
