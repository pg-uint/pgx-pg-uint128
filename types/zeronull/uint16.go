// Do not edit. Generated from codegen

package zeronull

import (
	"database/sql/driver"
	"fmt"

	"github.com/pg-uint/pgx-pg-uint128/types"
	"lukechampine.com/uint128"
)

type UInt16 uint128.Uint128

func (UInt16) SkipUnderlyingTypePlan() {}

// ScanInt64 implements the Int64Scanner interface.
func (dst *UInt16) ScanInt64(n int64, valid bool) error {
	if !valid {
		*dst = UInt16{}
		return nil
	}

	if n < 0 {
		return fmt.Errorf("%d is less than minimum value for UInt16", n)
	}

	*dst = UInt16(uint128.From64(uint64(n)))

	return nil
}

// ScanUint64 implements the Uint64Scanner interface.
func (dst *UInt16) ScanUint64(n uint64, valid bool) error {
	if !valid {
		*dst = UInt16{}
		return nil
	}

	*dst = UInt16(uint128.From64(n))

	return nil
}

// Scan implements the database/sql Scanner interface.
func (dst *UInt16) Scan(src any) error {
	if src == nil {
		*dst = UInt16{}
		return nil
	}

	var nullable types.UInt16
	err := nullable.Scan(src)
	if err != nil {
		return err
	}

	*dst = UInt16(nullable.Uint128)

	return nil
}

// Value implements the database/sql/driver Valuer interface.
func (src UInt16) Value() (driver.Value, error) {
	if uint128.Uint128(src).IsZero() {
		return nil, nil
	}
	return uint128.Uint128(src), nil
}
