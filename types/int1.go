// Do not edit. Generated from codegen

package types

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"math"
	"strconv"

	"github.com/pg-uint/pgx-pg-uint128/pgio"

	. "github.com/jackc/pgx/v5/pgtype"

	"lukechampine.com/uint128"

	"github.com/pg-uint/pgx-pg-uint128/int128"
	"go.shabbyrobe.org/num"
)

type Int1 struct {
	Int8  int8
	Valid bool
}

func (n Int1) Int64Value() (Int8, error) {
	return Int8{Int64: int64(n.Int8), Valid: n.Valid}, nil
}

func (n Int1) Uint64Value() (UInt8, error) {
	if n.Int8 < 0 {
		return UInt8{}, fmt.Errorf("Int1 value is less than min UInt8 value")
	}
	return UInt8{Uint64: uint64(n.Int8), Valid: n.Valid}, nil
}

// ScanInt64 implements the Int64Scanner interface.
func (dst *Int1) ScanInt64(n Int8) error {
	if !n.Valid {
		*dst = Int1{}
		return nil
	}

	if n.Int64 < math.MinInt8 {
		return fmt.Errorf("%d is less than minimum value for Int1", n.Int64)
	}
	if n.Int64 > math.MaxInt8 {
		return fmt.Errorf("%d is greater than maximum value for Int1", n.Int64)
	}
	*dst = Int1{Int8: int8(n.Int64), Valid: true}

	return nil
}

// ScanUint64 implements the Uint64Scanner interface.
func (dst *Int1) ScanUint64(n UInt8) error {
	if !n.Valid {
		*dst = Int1{}
		return nil
	}
	if n.Uint64 > math.MaxInt8 {
		return fmt.Errorf("%d is greater than maximum value for Int1", n.Uint64)
	}

	*dst = Int1{Int8: int8(n.Uint64), Valid: true}

	return nil
}

// Scan implements the database/sql Scanner interface.
func (dst *Int1) Scan(src any) error {
	if src == nil {
		*dst = Int1{}
		return nil
	}

	var n int64

	switch src := src.(type) {
	case int64:
		if src < math.MinInt8 {
			return fmt.Errorf("%d is less than minimum value for Int1", n)
		}

		n = int64(src)
	case uint64:
		n = int64(src)
	case string:
		var err error
		n, err = strconv.ParseInt(src, 10, 8)
		if err != nil {
			return err
		}
	case []byte:
		var err error
		n, err = strconv.ParseInt(string(src), 10, 8)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("cannot scan %T", src)
	}

	if n > math.MaxInt8 {
		return fmt.Errorf("%d is greater than maximum value for Int1", n)
	}

	*dst = Int1{Int8: int8(n), Valid: true}

	return nil
}

// Value implements the database/sql/driver Valuer interface.
func (src Int1) Value() (driver.Value, error) {
	if !src.Valid {
		return nil, nil
	}
	return int64(src.Int8), nil
}

func (src Int1) MarshalJSON() ([]byte, error) {
	if !src.Valid {
		return []byte("null"), nil
	}
	return []byte(strconv.FormatInt(int64(src.Int8), 10)), nil
}

func (dst *Int1) UnmarshalJSON(b []byte) error {
	var n *int8
	err := json.Unmarshal(b, &n)
	if err != nil {
		return err
	}

	if n == nil {
		*dst = Int1{}
	} else {
		*dst = Int1{Int8: *n, Valid: true}
	}

	return nil
}

type Int1Codec struct{}

func (Int1Codec) FormatSupported(format int16) bool {
	return format == TextFormatCode || format == BinaryFormatCode
}

func (Int1Codec) PreferredFormat() int16 {
	return BinaryFormatCode
}

func (Int1Codec) PlanEncode(m *Map, oid uint32, format int16, value any) EncodePlan {
	switch format {
	case BinaryFormatCode:
		switch value.(type) {
		case int8:
			return encodePlanInt1CodecBinaryInt8{}
		case Uint64Valuer:
			return encodePlanInt1CodecBinaryInt8Uint64Valuer{}
		case Int64Valuer:
			return encodePlanInt1CodecBinaryInt8Int64Valuer{}
		}
	case TextFormatCode:
		switch value.(type) {
		case int8:
			return encodePlanInt1CodecTextInt8{}
		case Uint64Valuer:
			return encodePlanInt1CodecTextInt8Uint64Valuer{}
		case Int64Valuer:
			return encodePlanInt1CodecTextInt8Int64Valuer{}
		}
	}

	return nil
}

type encodePlanInt1CodecBinaryInt8 struct{}

func (encodePlanInt1CodecBinaryInt8) Encode(value any, buf []byte) (newBuf []byte, err error) {
	n := value.(int8)
	return pgio.AppendInt8(buf, n), nil
}

type encodePlanInt1CodecTextInt8 struct{}

func (encodePlanInt1CodecTextInt8) Encode(value any, buf []byte) (newBuf []byte, err error) {
	n := value.(int8)
	return append(buf, strconv.FormatInt(int64(n), 10)...), nil
}

type encodePlanInt1CodecBinaryInt8Int64Valuer struct{}

func (encodePlanInt1CodecBinaryInt8Int64Valuer) Encode(value any, buf []byte) (newBuf []byte, err error) {
	n, err := value.(Int64Valuer).Int64Value()
	if err != nil {
		return nil, err
	}

	if !n.Valid {
		return nil, nil
	}

	if n.Int64 < 0 {
		return nil, fmt.Errorf("%d is less than minimum value for int1", n.Int64)
	}
	if n.Int64 > math.MaxInt8 {
		return nil, fmt.Errorf("%d is greater than maximum value for int1", n.Int64)
	}

	return pgio.AppendInt8(buf, int8(n.Int64)), nil
}

type encodePlanInt1CodecTextInt8Int64Valuer struct{}

func (encodePlanInt1CodecTextInt8Int64Valuer) Encode(value any, buf []byte) (newBuf []byte, err error) {
	n, err := value.(Int64Valuer).Int64Value()
	if err != nil {
		return nil, err
	}

	if !n.Valid {
		return nil, nil
	}

	if n.Int64 < 0 {
		return nil, fmt.Errorf("%d is less than minimum value for int1", n.Int64)
	}
	if n.Int64 > math.MaxInt8 {
		return nil, fmt.Errorf("%d is greater than maximum value for int1", n.Int64)
	}

	return append(buf, strconv.FormatInt(n.Int64, 10)...), nil
}

type encodePlanInt1CodecBinaryInt8Uint64Valuer struct{}

func (encodePlanInt1CodecBinaryInt8Uint64Valuer) Encode(value any, buf []byte) (newBuf []byte, err error) {
	n, err := value.(Uint64Valuer).Uint64Value()
	if err != nil {
		return nil, err
	}

	if !n.Valid {
		return nil, nil
	}

	if n.Uint64 > math.MaxInt8 {
		return nil, fmt.Errorf("%d is greater than maximum value for int1", n.Uint64)
	}

	return pgio.AppendInt8(buf, int8(n.Uint64)), nil
}

type encodePlanInt1CodecTextInt8Uint64Valuer struct{}

func (encodePlanInt1CodecTextInt8Uint64Valuer) Encode(value any, buf []byte) (newBuf []byte, err error) {
	n, err := value.(Uint64Valuer).Uint64Value()
	if err != nil {
		return nil, err
	}

	if !n.Valid {
		return nil, nil
	}

	if n.Uint64 > math.MaxInt8 {
		return nil, fmt.Errorf("%d is greater than maximum value for int1", n.Uint64)
	}

	return append(buf, strconv.FormatInt(int64(n.Uint64), 10)...), nil
}

func (Int1Codec) PlanScan(m *Map, oid uint32, format int16, target any) ScanPlan {
	switch format {
	case BinaryFormatCode:
		switch target.(type) {
		case *uint8:
			return scanPlanBinaryInt1ToUint8{}
		case *uint16:
			return scanPlanBinaryInt1ToUint16{}
		case *uint32:
			return scanPlanBinaryInt1ToUint32{}
		case *uint64:
			return scanPlanBinaryInt1ToUint64{}
		case *uint128.Uint128:
			return scanPlanBinaryInt1ToUint128{}
		case *uint:
			return scanPlanBinaryInt1ToUint{}
		case *int8:
			return scanPlanBinaryInt1ToInt8{}
		case *int16:
			return scanPlanBinaryInt1ToInt16{}
		case *int32:
			return scanPlanBinaryInt1ToInt32{}
		case *int64:
			return scanPlanBinaryInt1ToInt64{}
		case *num.I128:
			return scanPlanBinaryInt1ToI128{}
		case *int:
			return scanPlanBinaryInt1ToInt{}
		case Int64Scanner:
			return scanPlanBinaryInt1ToInt64Scanner{}
		case Uint64Scanner:
			return scanPlanBinaryInt1ToUint64Scanner{}
		case TextScanner:
			return scanPlanBinaryInt1ToTextScanner{}
		}
	case TextFormatCode:
		switch target.(type) {
		case *uint8:
			return scanPlanTextInt1ToUint8{}
		case *uint16:
			return scanPlanTextInt1ToUint16{}
		case *uint32:
			return scanPlanTextInt1ToUint32{}
		case *uint64:
			return scanPlanTextInt1ToUint64{}
		case *uint128.Uint128:
			return scanPlanTextInt1ToUint128{}
		case *uint:
			return scanPlanTextInt1ToUint{}
		case *int8:
			return scanPlanTextInt1ToInt8{}
		case *int16:
			return scanPlanTextInt1ToInt16{}
		case *int32:
			return scanPlanTextInt1ToInt32{}
		case *int64:
			return scanPlanTextInt1ToInt64{}
		case *num.I128:
			return scanPlanTextInt1ToI128{}
		case *int:
			return scanPlanTextInt1ToInt{}
		case Int64Scanner:
			return scanPlanTextAnyToInt64Scanner{}
		case Uint64Scanner:
			return scanPlanTextAnyToUint64Scanner{}
		}
	}

	return nil
}

func (c Int1Codec) DecodeDatabaseSQLValue(m *Map, oid uint32, format int16, src []byte) (driver.Value, error) {
	if src == nil {
		return nil, nil
	}

	var n uint64
	err := codecScan(c, m, oid, format, src, &n)
	if err != nil {
		return nil, err
	}
	return n, nil
}

func (c Int1Codec) DecodeValue(m *Map, oid uint32, format int16, src []byte) (any, error) {
	if src == nil {
		return nil, nil
	}

	var n int8
	err := codecScan(c, m, oid, format, src, &n)
	if err != nil {
		return nil, err
	}
	return n, nil
}

type scanPlanBinaryInt1ToUint8 struct{}

func (scanPlanBinaryInt1ToUint8) Scan(src []byte, dst any) error {
	if src == nil {
		return fmt.Errorf("cannot scan NULL into %T", dst)
	}

	if len(src) != 1 {
		return fmt.Errorf("invalid length for Int1: %v", len(src))
	}

	p, ok := (dst).(*uint8)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := pgio.ReadInt8(src)
	if n < 0 {
		return fmt.Errorf("%d is less than minimum value for uint8", n)
	}
	*p = uint8(n)

	return nil
}

type scanPlanBinaryInt1ToUint16 struct{}

func (scanPlanBinaryInt1ToUint16) Scan(src []byte, dst any) error {
	if src == nil {
		return fmt.Errorf("cannot scan NULL into %T", dst)
	}

	if len(src) != 1 {
		return fmt.Errorf("invalid length for Int1: %v", len(src))
	}

	p, ok := (dst).(*uint16)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := pgio.ReadInt8(src)
	if n < 0 {
		return fmt.Errorf("%d is less than minimum value for uint16", n)
	}
	*p = uint16(n)

	return nil
}

type scanPlanBinaryInt1ToUint32 struct{}

func (scanPlanBinaryInt1ToUint32) Scan(src []byte, dst any) error {
	if src == nil {
		return fmt.Errorf("cannot scan NULL into %T", dst)
	}

	if len(src) != 1 {
		return fmt.Errorf("invalid length for Int1: %v", len(src))
	}

	p, ok := (dst).(*uint32)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := pgio.ReadInt8(src)
	if n < 0 {
		return fmt.Errorf("%d is less than minimum value for uint32", n)
	}
	*p = uint32(n)

	return nil
}

type scanPlanBinaryInt1ToUint64 struct{}

func (scanPlanBinaryInt1ToUint64) Scan(src []byte, dst any) error {
	if src == nil {
		return fmt.Errorf("cannot scan NULL into %T", dst)
	}

	if len(src) != 1 {
		return fmt.Errorf("invalid length for Int1: %v", len(src))
	}

	p, ok := (dst).(*uint64)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := pgio.ReadInt8(src)
	if n < 0 {
		return fmt.Errorf("%d is less than minimum value for uint64", n)
	}
	*p = uint64(n)

	return nil
}

type scanPlanBinaryInt1ToUint128 struct{}

func (scanPlanBinaryInt1ToUint128) Scan(src []byte, dst any) error {
	if src == nil {
		return fmt.Errorf("cannot scan NULL into %T", dst)
	}

	if len(src) != 1 {
		return fmt.Errorf("invalid length for Int1: %v", len(src))
	}

	p, ok := (dst).(*uint128.Uint128)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := pgio.ReadInt8(src)
	if n < 0 {
		return fmt.Errorf("%d is less than minimum value for Uint128", n)
	}
	*p = uint128.From64(uint64(n))

	return nil
}

type scanPlanBinaryInt1ToUint struct{}

func (scanPlanBinaryInt1ToUint) Scan(src []byte, dst any) error {
	if src == nil {
		return fmt.Errorf("cannot scan NULL into %T", dst)
	}

	if len(src) != 1 {
		return fmt.Errorf("invalid length for Int1: %v", len(src))
	}

	p, ok := (dst).(*uint)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := pgio.ReadInt8(src)
	if n < 0 {
		return fmt.Errorf("%d is less than minimum value for uint", n)
	}
	*p = uint(n)

	return nil
}

type scanPlanBinaryInt1ToInt8 struct{}

func (scanPlanBinaryInt1ToInt8) Scan(src []byte, dst any) error {
	if src == nil {
		return fmt.Errorf("cannot scan NULL into %T", dst)
	}

	if len(src) != 1 {
		return fmt.Errorf("invalid length for Int1: %v", len(src))
	}

	p, ok := (dst).(*int8)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := pgio.ReadInt8(src)

	*p = n

	return nil
}

type scanPlanBinaryInt1ToInt16 struct{}

func (scanPlanBinaryInt1ToInt16) Scan(src []byte, dst any) error {
	if src == nil {
		return fmt.Errorf("cannot scan NULL into %T", dst)
	}

	if len(src) != 1 {
		return fmt.Errorf("invalid length for Int1: %v", len(src))
	}

	p, ok := (dst).(*int16)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := pgio.ReadInt8(src)

	*p = int16(n)

	return nil
}

type scanPlanBinaryInt1ToInt32 struct{}

func (scanPlanBinaryInt1ToInt32) Scan(src []byte, dst any) error {
	if src == nil {
		return fmt.Errorf("cannot scan NULL into %T", dst)
	}

	if len(src) != 1 {
		return fmt.Errorf("invalid length for Int1: %v", len(src))
	}

	p, ok := (dst).(*int32)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := pgio.ReadInt8(src)

	*p = int32(n)

	return nil
}

type scanPlanBinaryInt1ToInt64 struct{}

func (scanPlanBinaryInt1ToInt64) Scan(src []byte, dst any) error {
	if src == nil {
		return fmt.Errorf("cannot scan NULL into %T", dst)
	}

	if len(src) != 1 {
		return fmt.Errorf("invalid length for Int1: %v", len(src))
	}

	p, ok := (dst).(*int64)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := pgio.ReadInt8(src)

	*p = int64(n)

	return nil
}

type scanPlanBinaryInt1ToI128 struct{}

func (scanPlanBinaryInt1ToI128) Scan(src []byte, dst any) error {
	if src == nil {
		return fmt.Errorf("cannot scan NULL into %T", dst)
	}

	if len(src) != 1 {
		return fmt.Errorf("invalid length for Int1: %v", len(src))
	}

	p, ok := (dst).(*num.I128)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := pgio.ReadInt8(src)

	*p = num.I128From64(int64(n))

	return nil
}

type scanPlanBinaryInt1ToInt struct{}

func (scanPlanBinaryInt1ToInt) Scan(src []byte, dst any) error {
	if src == nil {
		return fmt.Errorf("cannot scan NULL into %T", dst)
	}

	if len(src) != 1 {
		return fmt.Errorf("invalid length for Int1: %v", len(src))
	}

	p, ok := (dst).(*int)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := pgio.ReadInt8(src)

	*p = int(n)

	return nil
}

type scanPlanBinaryInt1ToTextScanner struct{}

func (scanPlanBinaryInt1ToTextScanner) Scan(src []byte, dst any) error {
	s, ok := (dst).(TextScanner)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	if src == nil {
		return s.ScanText(Text{})
	}

	if len(src) != 1 {
		return fmt.Errorf("invalid length for int1: %v", len(src))
	}

	n := int64(pgio.ReadInt8(src))

	return s.ScanText(Text{String: strconv.FormatInt(int64(n), 10), Valid: true})
}

type scanPlanBinaryInt1ToInt64Scanner struct{}

func (scanPlanBinaryInt1ToInt64Scanner) Scan(src []byte, dst any) error {
	s, ok := (dst).(Int64Scanner)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	if src == nil {
		return s.ScanInt64(Int8{})
	}

	if len(src) != 1 {
		return fmt.Errorf("invalid length for int1: %v", len(src))
	}

	n := int64(pgio.ReadInt8(src))
	if n > math.MaxInt64 {
		return fmt.Errorf("Int1 value %d is greater than max value for Int8", n)
	}

	return s.ScanInt64(Int8{Int64: int64(n), Valid: true})
}

type scanPlanBinaryInt1ToUint64Scanner struct{}

func (scanPlanBinaryInt1ToUint64Scanner) Scan(src []byte, dst any) error {
	s, ok := (dst).(Uint64Scanner)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	if src == nil {
		return s.ScanUint64(UInt8{})
	}

	if len(src) != 1 {
		return fmt.Errorf("invalid length for int1: %v", len(src))
	}

	n := int64(pgio.ReadInt8(src))

	return s.ScanUint64(UInt8{Uint64: uint64(n), Valid: true})
}

type scanPlanTextInt1ToUint8 struct{}

func (scanPlanTextInt1ToUint8) Scan(src []byte, dst any) error {
	if src == nil {
		return fmt.Errorf("cannot scan NULL into %T", dst)
	}

	p, ok := (dst).(*uint8)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n, err := strconv.ParseUint(string(src), 10, 8)
	if err != nil {
		return err
	}

	*p = uint8(n)
	return nil
}

type scanPlanTextInt1ToUint16 struct{}

func (scanPlanTextInt1ToUint16) Scan(src []byte, dst any) error {
	if src == nil {
		return fmt.Errorf("cannot scan NULL into %T", dst)
	}

	p, ok := (dst).(*uint16)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n, err := strconv.ParseUint(string(src), 10, 16)
	if err != nil {
		return err
	}

	*p = uint16(n)
	return nil
}

type scanPlanTextInt1ToUint32 struct{}

func (scanPlanTextInt1ToUint32) Scan(src []byte, dst any) error {
	if src == nil {
		return fmt.Errorf("cannot scan NULL into %T", dst)
	}

	p, ok := (dst).(*uint32)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n, err := strconv.ParseUint(string(src), 10, 32)
	if err != nil {
		return err
	}

	*p = uint32(n)
	return nil
}

type scanPlanTextInt1ToUint64 struct{}

func (scanPlanTextInt1ToUint64) Scan(src []byte, dst any) error {
	if src == nil {
		return fmt.Errorf("cannot scan NULL into %T", dst)
	}

	p, ok := (dst).(*uint64)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n, err := strconv.ParseUint(string(src), 10, 64)
	if err != nil {
		return err
	}

	*p = uint64(n)
	return nil
}

type scanPlanTextInt1ToUint128 struct{}

func (scanPlanTextInt1ToUint128) Scan(src []byte, dst any) error {
	if src == nil {
		return fmt.Errorf("cannot scan NULL into %T", dst)
	}

	p, ok := (dst).(*uint128.Uint128)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n, err := uint128.FromString(string(src))
	if err != nil {
		return err
	}

	*p = uint128.Uint128(n)
	return nil
}

type scanPlanTextInt1ToUint struct{}

func (scanPlanTextInt1ToUint) Scan(src []byte, dst any) error {
	if src == nil {
		return fmt.Errorf("cannot scan NULL into %T", dst)
	}

	p, ok := (dst).(*uint)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n, err := strconv.ParseUint(string(src), 10, intSize)
	if err != nil {
		return err
	}

	*p = uint(n)
	return nil
}

type scanPlanTextInt1ToInt8 struct{}

func (scanPlanTextInt1ToInt8) Scan(src []byte, dst any) error {
	if src == nil {
		return fmt.Errorf("cannot scan NULL into %T", dst)
	}

	p, ok := (dst).(*int8)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n, err := strconv.ParseInt(string(src), 10, 8)
	if err != nil {
		return err
	}

	*p = int8(n)
	return nil
}

type scanPlanTextInt1ToInt16 struct{}

func (scanPlanTextInt1ToInt16) Scan(src []byte, dst any) error {
	if src == nil {
		return fmt.Errorf("cannot scan NULL into %T", dst)
	}

	p, ok := (dst).(*int16)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n, err := strconv.ParseInt(string(src), 10, 16)
	if err != nil {
		return err
	}

	*p = int16(n)
	return nil
}

type scanPlanTextInt1ToInt32 struct{}

func (scanPlanTextInt1ToInt32) Scan(src []byte, dst any) error {
	if src == nil {
		return fmt.Errorf("cannot scan NULL into %T", dst)
	}

	p, ok := (dst).(*int32)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n, err := strconv.ParseInt(string(src), 10, 32)
	if err != nil {
		return err
	}

	*p = int32(n)
	return nil
}

type scanPlanTextInt1ToInt64 struct{}

func (scanPlanTextInt1ToInt64) Scan(src []byte, dst any) error {
	if src == nil {
		return fmt.Errorf("cannot scan NULL into %T", dst)
	}

	p, ok := (dst).(*int64)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n, err := strconv.ParseInt(string(src), 10, 64)
	if err != nil {
		return err
	}

	*p = int64(n)
	return nil
}

type scanPlanTextInt1ToI128 struct{}

func (scanPlanTextInt1ToI128) Scan(src []byte, dst any) error {
	if src == nil {
		return fmt.Errorf("cannot scan NULL into %T", dst)
	}

	p, ok := (dst).(*num.I128)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n, err := int128.FromString(string(src))
	if err != nil {
		return err
	}

	*p = num.I128(n)
	return nil
}

type scanPlanTextInt1ToInt struct{}

func (scanPlanTextInt1ToInt) Scan(src []byte, dst any) error {
	if src == nil {
		return fmt.Errorf("cannot scan NULL into %T", dst)
	}

	p, ok := (dst).(*int)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n, err := strconv.ParseInt(string(src), 10, intSize)
	if err != nil {
		return err
	}

	*p = int(n)
	return nil
}
