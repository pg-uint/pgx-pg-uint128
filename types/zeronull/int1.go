// Do not edit. Generated from codegen

package zeronull

import (
	"database/sql/driver"
	"fmt"
	"math"

	"github.com/pg-uint/pgx-pg-uint128/types"
)

type Int1 int8

func (Int1) SkipUnderlyingTypePlan() {}

// ScanInt64 implements the Int64Scanner interface.
func (dst *Int1) ScanInt64(n int64, valid bool) error {
	if !valid {
		*dst = 0
		return nil
	}

	if n < math.MinInt8 {
		return fmt.Errorf("%d is less than minimum value for Int1", n)
	}
	if n > math.MaxInt8 {
		return fmt.Errorf("%d is greater than maximum value for Int1", n)
	}
	*dst = Int1(n)

	return nil
}

// ScanUint64 implements the Uint64Scanner interface.
func (dst *Int1) ScanUint64(n uint64, valid bool) error {
	if !valid {
		*dst = 0
		return nil
	}

	if n > math.MaxInt8 {
		return fmt.Errorf("%d is greater than maximum value for Int1", n)
	}

	*dst = Int1(n)

	return nil
}

// Scan implements the database/sql Scanner interface.
func (dst *Int1) Scan(src any) error {
	if src == nil {
		*dst = 0
		return nil
	}

	var nullable types.Int1
	err := nullable.Scan(src)
	if err != nil {
		return err
	}

	*dst = Int1(nullable.Int8)

	return nil
}

// Value implements the database/sql/driver Valuer interface.
func (src Int1) Value() (driver.Value, error) {
	if src == 0 {
		return nil, nil
	}
	return int64(src), nil
}
