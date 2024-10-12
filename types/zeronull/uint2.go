// Do not edit. Generated from codegen

package zeronull

import (
	"database/sql/driver"
	"fmt"
	"math"

	"github.com/codercms/pgx-pg-uint128/types"
)

type UInt2 uint16

func (UInt2) SkipUnderlyingTypePlan() {}

// ScanInt64 implements the Int64Scanner interface.
func (dst *UInt2) ScanInt64(n int64, valid bool) error {
	if !valid {
		*dst = 0
		return nil
	}

	if n < 0 {
		return fmt.Errorf("%d is less than minimum value for UInt2", n)
	}
	if n > math.MaxUint16 {
		return fmt.Errorf("%d is greater than maximum value for UInt2", n)
	}
	*dst = UInt2(n)

	return nil
}

// ScanUint64 implements the Uint64Scanner interface.
func (dst *UInt2) ScanUint64(n uint64, valid bool) error {
	if !valid {
		*dst = 0
		return nil
	}

	if n > math.MaxUint16 {
		return fmt.Errorf("%d is greater than maximum value for UInt2", n)
	}

	*dst = UInt2(n)

	return nil
}

// Scan implements the database/sql Scanner interface.
func (dst *UInt2) Scan(src any) error {
	if src == nil {
		*dst = 0
		return nil
	}

	var nullable types.UInt2
	err := nullable.Scan(src)
	if err != nil {
		return err
	}

	*dst = UInt2(nullable.Uint16)

	return nil
}

// Value implements the database/sql/driver Valuer interface.
func (src UInt2) Value() (driver.Value, error) {
	if src == 0 {
		return nil, nil
	}
	return uint64(src), nil
}
