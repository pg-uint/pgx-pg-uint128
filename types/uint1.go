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

type UInt1 struct {
	Uint8 uint8
	Valid bool
}

func (n UInt1) Int64Value() (Int8, error) {
	return Int8{Int64: int64(n.Uint8), Valid: n.Valid}, nil
}

func (n UInt1) Uint64Value() (UInt8, error) {
	return UInt8{Uint64: uint64(n.Uint8), Valid: n.Valid}, nil
}

// ScanInt64 implements the Int64Scanner interface.
func (dst *UInt1) ScanInt64(n Int8) error {
	if !n.Valid {
		*dst = UInt1{}
		return nil
	}

	if n.Int64 < 0 {
		return fmt.Errorf("%d is less than minimum value for UInt1", n.Int64)
	}
	if n.Int64 > math.MaxUint8 {
		return fmt.Errorf("%d is greater than maximum value for UInt1", n.Int64)
	}
	*dst = UInt1{Uint8: uint8(n.Int64), Valid: true}

	return nil
}

// ScanUint64 implements the Uint64Scanner interface.
func (dst *UInt1) ScanUint64(n UInt8) error {
	if !n.Valid {
		*dst = UInt1{}
		return nil
	}
	if n.Uint64 > math.MaxUint8 {
		return fmt.Errorf("%d is greater than maximum value for UInt1", n.Uint64)
	}

	*dst = UInt1{Uint8: uint8(n.Uint64), Valid: true}

	return nil
}

// Scan implements the database/sql Scanner interface.
func (dst *UInt1) Scan(src any) error {
	if src == nil {
		*dst = UInt1{}
		return nil
	}

	var n uint64

	switch src := src.(type) {
	case int64:
		if src < 0 {
			return fmt.Errorf("%d is less than minimum value for UInt1", n)
		}

		n = uint64(src)
	case uint64:
		n = src
	case string:
		var err error
		n, err = strconv.ParseUint(src, 10, 8)
		if err != nil {
			return err
		}
	case []byte:
		var err error
		n, err = strconv.ParseUint(string(src), 10, 8)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("cannot scan %T", src)
	}

	if n > math.MaxUint8 {
		return fmt.Errorf("%d is greater than maximum value for UInt1", n)
	}

	*dst = UInt1{Uint8: uint8(n), Valid: true}

	return nil
}

// Value implements the database/sql/driver Valuer interface.
func (src UInt1) Value() (driver.Value, error) {
	if !src.Valid {
		return nil, nil
	}
	return uint64(src.Uint8), nil
}

func (src UInt1) MarshalJSON() ([]byte, error) {
	if !src.Valid {
		return []byte("null"), nil
	}
	return []byte(strconv.FormatUint(uint64(src.Uint8), 10)), nil
}

func (dst *UInt1) UnmarshalJSON(b []byte) error {
	var n *uint8
	err := json.Unmarshal(b, &n)
	if err != nil {
		return err
	}

	if n == nil {
		*dst = UInt1{}
	} else {
		*dst = UInt1{Uint8: *n, Valid: true}
	}

	return nil
}

type UInt1Codec struct{}

func (UInt1Codec) FormatSupported(format int16) bool {
	return format == TextFormatCode || format == BinaryFormatCode
}

func (UInt1Codec) PreferredFormat() int16 {
	return BinaryFormatCode
}

func (UInt1Codec) PlanEncode(m *Map, oid uint32, format int16, value any) EncodePlan {
	switch format {
	case BinaryFormatCode:
		switch value.(type) {
		case uint8:
			return encodePlanUInt1CodecBinaryUint8{}
		case Uint64Valuer:
			return encodePlanUInt1CodecBinaryUint8Uint64Valuer{}
		case Int64Valuer:
			return encodePlanUInt1CodecBinaryUint8Int64Valuer{}
		}
	case TextFormatCode:
		switch value.(type) {
		case uint8:
			return encodePlanUInt1CodecTextUint8{}
		case Uint64Valuer:
			return encodePlanUInt1CodecTextUint8Uint64Valuer{}
		case Int64Valuer:
			return encodePlanUInt1CodecTextUint8Int64Valuer{}
		}
	}

	return nil
}

type encodePlanUInt1CodecBinaryUint8 struct{}

func (encodePlanUInt1CodecBinaryUint8) Encode(value any, buf []byte) (newBuf []byte, err error) {
	n := value.(uint8)
	return pgio.AppendUint8(buf, n), nil
}

type encodePlanUInt1CodecTextUint8 struct{}

func (encodePlanUInt1CodecTextUint8) Encode(value any, buf []byte) (newBuf []byte, err error) {
	n := value.(uint8)
	return append(buf, strconv.FormatUint(uint64(n), 10)...), nil
}

type encodePlanUInt1CodecBinaryUint8Int64Valuer struct{}

func (encodePlanUInt1CodecBinaryUint8Int64Valuer) Encode(value any, buf []byte) (newBuf []byte, err error) {
	n, err := value.(Int64Valuer).Int64Value()
	if err != nil {
		return nil, err
	}

	if !n.Valid {
		return nil, nil
	}

	if n.Int64 < 0 {
		return nil, fmt.Errorf("%d is less than minimum value for uint1", n.Int64)
	}
	if n.Int64 > math.MaxUint8 {
		return nil, fmt.Errorf("%d is greater than maximum value for uint1", n.Int64)
	}

	return pgio.AppendUint8(buf, uint8(n.Int64)), nil
}

type encodePlanUInt1CodecTextUint8Int64Valuer struct{}

func (encodePlanUInt1CodecTextUint8Int64Valuer) Encode(value any, buf []byte) (newBuf []byte, err error) {
	n, err := value.(Int64Valuer).Int64Value()
	if err != nil {
		return nil, err
	}

	if !n.Valid {
		return nil, nil
	}

	if n.Int64 < 0 {
		return nil, fmt.Errorf("%d is less than minimum value for uint1", n.Int64)
	}
	if n.Int64 > math.MaxUint8 {
		return nil, fmt.Errorf("%d is greater than maximum value for uint1", n.Int64)
	}

	return append(buf, strconv.FormatUint(uint64(n.Int64), 10)...), nil
}

type encodePlanUInt1CodecBinaryUint8Uint64Valuer struct{}

func (encodePlanUInt1CodecBinaryUint8Uint64Valuer) Encode(value any, buf []byte) (newBuf []byte, err error) {
	n, err := value.(Uint64Valuer).Uint64Value()
	if err != nil {
		return nil, err
	}

	if !n.Valid {
		return nil, nil
	}

	if n.Uint64 > math.MaxUint8 {
		return nil, fmt.Errorf("%d is greater than maximum value for uint1", n.Uint64)
	}

	return pgio.AppendUint8(buf, uint8(n.Uint64)), nil
}

type encodePlanUInt1CodecTextUint8Uint64Valuer struct{}

func (encodePlanUInt1CodecTextUint8Uint64Valuer) Encode(value any, buf []byte) (newBuf []byte, err error) {
	n, err := value.(Uint64Valuer).Uint64Value()
	if err != nil {
		return nil, err
	}

	if !n.Valid {
		return nil, nil
	}

	if n.Uint64 > math.MaxUint8 {
		return nil, fmt.Errorf("%d is greater than maximum value for uint1", n.Uint64)
	}

	return append(buf, strconv.FormatUint(uint64(n.Uint64), 10)...), nil
}

func (UInt1Codec) PlanScan(m *Map, oid uint32, format int16, target any) ScanPlan {
	switch format {
	case BinaryFormatCode:
		switch target.(type) {
		case *uint8:
			return scanPlanBinaryUInt1ToUint8{}
		case *uint16:
			return scanPlanBinaryUInt1ToUint16{}
		case *uint32:
			return scanPlanBinaryUInt1ToUint32{}
		case *uint64:
			return scanPlanBinaryUInt1ToUint64{}
		case *uint128.Uint128:
			return scanPlanBinaryUInt1ToUint128{}
		case *uint:
			return scanPlanBinaryUInt1ToUint{}
		case *int8:
			return scanPlanBinaryUInt1ToInt8{}
		case *int16:
			return scanPlanBinaryUInt1ToInt16{}
		case *int32:
			return scanPlanBinaryUInt1ToInt32{}
		case *int64:
			return scanPlanBinaryUInt1ToInt64{}
		case *num.I128:
			return scanPlanBinaryUInt1ToI128{}
		case *int:
			return scanPlanBinaryUInt1ToInt{}
		case Int64Scanner:
			return scanPlanBinaryUInt1ToInt64Scanner{}
		case Uint64Scanner:
			return scanPlanBinaryUInt1ToUint64Scanner{}
		case TextScanner:
			return scanPlanBinaryUInt1ToTextScanner{}
		}
	case TextFormatCode:
		switch target.(type) {
		case *uint8:
			return scanPlanTextUInt1ToUint8{}
		case *uint16:
			return scanPlanTextUInt1ToUint16{}
		case *uint32:
			return scanPlanTextUInt1ToUint32{}
		case *uint64:
			return scanPlanTextUInt1ToUint64{}
		case *uint128.Uint128:
			return scanPlanTextUInt1ToUint128{}
		case *uint:
			return scanPlanTextUInt1ToUint{}
		case *int8:
			return scanPlanTextUInt1ToInt8{}
		case *int16:
			return scanPlanTextUInt1ToInt16{}
		case *int32:
			return scanPlanTextUInt1ToInt32{}
		case *int64:
			return scanPlanTextUInt1ToInt64{}
		case *num.I128:
			return scanPlanTextUInt1ToI128{}
		case *int:
			return scanPlanTextUInt1ToInt{}
		case Int64Scanner:
			return scanPlanTextAnyToInt64Scanner{}
		case Uint64Scanner:
			return scanPlanTextAnyToUint64Scanner{}
		}
	}

	return nil
}

func (c UInt1Codec) DecodeDatabaseSQLValue(m *Map, oid uint32, format int16, src []byte) (driver.Value, error) {
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

func (c UInt1Codec) DecodeValue(m *Map, oid uint32, format int16, src []byte) (any, error) {
	if src == nil {
		return nil, nil
	}

	var n uint8
	err := codecScan(c, m, oid, format, src, &n)
	if err != nil {
		return nil, err
	}
	return n, nil
}

type scanPlanBinaryUInt1ToUint8 struct{}

func (scanPlanBinaryUInt1ToUint8) Scan(src []byte, dst any) error {
	if src == nil {
		return fmt.Errorf("cannot scan NULL into %T", dst)
	}

	if len(src) != 1 {
		return fmt.Errorf("invalid length for UInt1: %v", len(src))
	}

	p, ok := (dst).(*uint8)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := pgio.ReadUint8(src)

	*p = n

	return nil
}

type scanPlanBinaryUInt1ToUint16 struct{}

func (scanPlanBinaryUInt1ToUint16) Scan(src []byte, dst any) error {
	if src == nil {
		return fmt.Errorf("cannot scan NULL into %T", dst)
	}

	if len(src) != 1 {
		return fmt.Errorf("invalid length for UInt1: %v", len(src))
	}

	p, ok := (dst).(*uint16)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := pgio.ReadUint8(src)

	*p = uint16(n)

	return nil
}

type scanPlanBinaryUInt1ToUint32 struct{}

func (scanPlanBinaryUInt1ToUint32) Scan(src []byte, dst any) error {
	if src == nil {
		return fmt.Errorf("cannot scan NULL into %T", dst)
	}

	if len(src) != 1 {
		return fmt.Errorf("invalid length for UInt1: %v", len(src))
	}

	p, ok := (dst).(*uint32)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := pgio.ReadUint8(src)

	*p = uint32(n)

	return nil
}

type scanPlanBinaryUInt1ToUint64 struct{}

func (scanPlanBinaryUInt1ToUint64) Scan(src []byte, dst any) error {
	if src == nil {
		return fmt.Errorf("cannot scan NULL into %T", dst)
	}

	if len(src) != 1 {
		return fmt.Errorf("invalid length for UInt1: %v", len(src))
	}

	p, ok := (dst).(*uint64)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := pgio.ReadUint8(src)

	*p = uint64(n)

	return nil
}

type scanPlanBinaryUInt1ToUint128 struct{}

func (scanPlanBinaryUInt1ToUint128) Scan(src []byte, dst any) error {
	if src == nil {
		return fmt.Errorf("cannot scan NULL into %T", dst)
	}

	if len(src) != 1 {
		return fmt.Errorf("invalid length for UInt1: %v", len(src))
	}

	p, ok := (dst).(*uint128.Uint128)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := pgio.ReadUint8(src)

	*p = uint128.From64(uint64(n))

	return nil
}

type scanPlanBinaryUInt1ToUint struct{}

func (scanPlanBinaryUInt1ToUint) Scan(src []byte, dst any) error {
	if src == nil {
		return fmt.Errorf("cannot scan NULL into %T", dst)
	}

	if len(src) != 1 {
		return fmt.Errorf("invalid length for UInt1: %v", len(src))
	}

	p, ok := (dst).(*uint)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := pgio.ReadUint8(src)

	*p = uint(n)

	return nil
}

type scanPlanBinaryUInt1ToInt8 struct{}

func (scanPlanBinaryUInt1ToInt8) Scan(src []byte, dst any) error {
	if src == nil {
		return fmt.Errorf("cannot scan NULL into %T", dst)
	}

	if len(src) != 1 {
		return fmt.Errorf("invalid length for UInt1: %v", len(src))
	}

	p, ok := (dst).(*int8)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := pgio.ReadUint8(src)
	if n > uint8(math.MaxInt8) {
		return fmt.Errorf("%d is greater than maximum value for int8", n)
	}
	*p = int8(n)

	return nil
}

type scanPlanBinaryUInt1ToInt16 struct{}

func (scanPlanBinaryUInt1ToInt16) Scan(src []byte, dst any) error {
	if src == nil {
		return fmt.Errorf("cannot scan NULL into %T", dst)
	}

	if len(src) != 1 {
		return fmt.Errorf("invalid length for UInt1: %v", len(src))
	}

	p, ok := (dst).(*int16)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := pgio.ReadUint8(src)

	*p = int16(n)

	return nil
}

type scanPlanBinaryUInt1ToInt32 struct{}

func (scanPlanBinaryUInt1ToInt32) Scan(src []byte, dst any) error {
	if src == nil {
		return fmt.Errorf("cannot scan NULL into %T", dst)
	}

	if len(src) != 1 {
		return fmt.Errorf("invalid length for UInt1: %v", len(src))
	}

	p, ok := (dst).(*int32)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := pgio.ReadUint8(src)

	*p = int32(n)

	return nil
}

type scanPlanBinaryUInt1ToInt64 struct{}

func (scanPlanBinaryUInt1ToInt64) Scan(src []byte, dst any) error {
	if src == nil {
		return fmt.Errorf("cannot scan NULL into %T", dst)
	}

	if len(src) != 1 {
		return fmt.Errorf("invalid length for UInt1: %v", len(src))
	}

	p, ok := (dst).(*int64)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := pgio.ReadUint8(src)

	*p = int64(n)

	return nil
}

type scanPlanBinaryUInt1ToI128 struct{}

func (scanPlanBinaryUInt1ToI128) Scan(src []byte, dst any) error {
	if src == nil {
		return fmt.Errorf("cannot scan NULL into %T", dst)
	}

	if len(src) != 1 {
		return fmt.Errorf("invalid length for UInt1: %v", len(src))
	}

	p, ok := (dst).(*num.I128)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := pgio.ReadUint8(src)

	*p = num.I128FromU64(uint64(n))

	return nil
}

type scanPlanBinaryUInt1ToInt struct{}

func (scanPlanBinaryUInt1ToInt) Scan(src []byte, dst any) error {
	if src == nil {
		return fmt.Errorf("cannot scan NULL into %T", dst)
	}

	if len(src) != 1 {
		return fmt.Errorf("invalid length for UInt1: %v", len(src))
	}

	p, ok := (dst).(*int)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := pgio.ReadUint8(src)

	*p = int(n)

	return nil
}

type scanPlanBinaryUInt1ToTextScanner struct{}

func (scanPlanBinaryUInt1ToTextScanner) Scan(src []byte, dst any) error {
	s, ok := (dst).(TextScanner)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	if src == nil {
		return s.ScanText(Text{})
	}

	if len(src) != 1 {
		return fmt.Errorf("invalid length for uint1: %v", len(src))
	}

	n := uint64(pgio.ReadUint8(src))

	return s.ScanText(Text{String: strconv.FormatUint(uint64(n), 10), Valid: true})
}

type scanPlanBinaryUInt1ToInt64Scanner struct{}

func (scanPlanBinaryUInt1ToInt64Scanner) Scan(src []byte, dst any) error {
	s, ok := (dst).(Int64Scanner)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	if src == nil {
		return s.ScanInt64(Int8{})
	}

	if len(src) != 1 {
		return fmt.Errorf("invalid length for uint1: %v", len(src))
	}

	n := uint64(pgio.ReadUint8(src))
	if n > math.MaxInt64 {
		return fmt.Errorf("UInt1 value %d is greater than max value for Int8", n)
	}

	return s.ScanInt64(Int8{Int64: int64(n), Valid: true})
}

type scanPlanBinaryUInt1ToUint64Scanner struct{}

func (scanPlanBinaryUInt1ToUint64Scanner) Scan(src []byte, dst any) error {
	s, ok := (dst).(Uint64Scanner)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	if src == nil {
		return s.ScanUint64(UInt8{})
	}

	if len(src) != 1 {
		return fmt.Errorf("invalid length for uint1: %v", len(src))
	}

	n := uint64(pgio.ReadUint8(src))

	return s.ScanUint64(UInt8{Uint64: n, Valid: true})
}

type scanPlanTextUInt1ToUint8 struct{}

func (scanPlanTextUInt1ToUint8) Scan(src []byte, dst any) error {
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

type scanPlanTextUInt1ToUint16 struct{}

func (scanPlanTextUInt1ToUint16) Scan(src []byte, dst any) error {
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

type scanPlanTextUInt1ToUint32 struct{}

func (scanPlanTextUInt1ToUint32) Scan(src []byte, dst any) error {
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

type scanPlanTextUInt1ToUint64 struct{}

func (scanPlanTextUInt1ToUint64) Scan(src []byte, dst any) error {
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

type scanPlanTextUInt1ToUint128 struct{}

func (scanPlanTextUInt1ToUint128) Scan(src []byte, dst any) error {
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

type scanPlanTextUInt1ToUint struct{}

func (scanPlanTextUInt1ToUint) Scan(src []byte, dst any) error {
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

type scanPlanTextUInt1ToInt8 struct{}

func (scanPlanTextUInt1ToInt8) Scan(src []byte, dst any) error {
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

type scanPlanTextUInt1ToInt16 struct{}

func (scanPlanTextUInt1ToInt16) Scan(src []byte, dst any) error {
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

type scanPlanTextUInt1ToInt32 struct{}

func (scanPlanTextUInt1ToInt32) Scan(src []byte, dst any) error {
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

type scanPlanTextUInt1ToInt64 struct{}

func (scanPlanTextUInt1ToInt64) Scan(src []byte, dst any) error {
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

type scanPlanTextUInt1ToI128 struct{}

func (scanPlanTextUInt1ToI128) Scan(src []byte, dst any) error {
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

type scanPlanTextUInt1ToInt struct{}

func (scanPlanTextUInt1ToInt) Scan(src []byte, dst any) error {
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
