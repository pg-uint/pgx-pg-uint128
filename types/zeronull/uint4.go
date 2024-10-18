// Do not edit. Generated from codegen

package zeronull

import (
	"database/sql/driver"
	"fmt"
	"math"

	"github.com/pg-uint/pgx-pg-uint128/types"
)

type UInt4 uint32

func (UInt4) SkipUnderlyingTypePlan() {}

// ScanInt64 implements the Int64Scanner interface.
func (dst *UInt4) ScanInt64(n int64, valid bool) error {
	if !valid {
		*dst = 0
		return nil
	}

	if n < 0 {
		return fmt.Errorf("%d is less than minimum value for UInt4", n)
	}
	if n > math.MaxUint32 {
		return fmt.Errorf("%d is greater than maximum value for UInt4", n)
	}
	*dst = UInt4(n)

	return nil
}

// ScanUint64 implements the Uint64Scanner interface.
func (dst *UInt4) ScanUint64(n uint64, valid bool) error {
	if !valid {
		*dst = 0
		return nil
	}

	if n > math.MaxUint32 {
		return fmt.Errorf("%d is greater than maximum value for UInt4", n)
	}

	*dst = UInt4(n)

	return nil
}

// Scan implements the database/sql Scanner interface.
func (dst *UInt4) Scan(src any) error {
	if src == nil {
		*dst = 0
		return nil
	}

	var nullable types.UInt4
	err := nullable.Scan(src)
	if err != nil {
		return err
	}

	*dst = UInt4(nullable.Uint32)

	return nil
}

// Value implements the database/sql/driver Valuer interface.
func (src UInt4) Value() (driver.Value, error) {
	if src == 0 {
		return nil, nil
	}
	return uint64(src), nil
}
