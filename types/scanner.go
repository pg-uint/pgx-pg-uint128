package types

import (
	"github.com/jackc/pgx/v5/pgtype"
	"strconv"
)

type Uint64Scanner interface {
	ScanUint64(UInt8) error
}

type Uint64Valuer interface {
	Uint64Value() (UInt8, error)
}

type scanPlanTextAnyToInt64Scanner struct{}

func (scanPlanTextAnyToInt64Scanner) Scan(src []byte, dst any) error {
	s, ok := (dst).(pgtype.Int64Scanner)
	if !ok {
		return pgtype.ErrScanTargetTypeChanged
	}

	if src == nil {
		return s.ScanInt64(pgtype.Int8{})
	}

	n, err := strconv.ParseInt(string(src), 10, 64)
	if err != nil {
		return err
	}

	err = s.ScanInt64(pgtype.Int8{Int64: n, Valid: true})
	if err != nil {
		return err
	}

	return nil
}

type scanPlanTextAnyToUint64Scanner struct{}

func (scanPlanTextAnyToUint64Scanner) Scan(src []byte, dst any) error {
	s, ok := (dst).(Uint64Scanner)
	if !ok {
		return pgtype.ErrScanTargetTypeChanged
	}

	if src == nil {
		return s.ScanUint64(UInt8{})
	}

	n, err := strconv.ParseUint(string(src), 10, 64)
	if err != nil {
		return err
	}

	err = s.ScanUint64(UInt8{Uint64: n, Valid: true})
	if err != nil {
		return err
	}

	return nil
}
