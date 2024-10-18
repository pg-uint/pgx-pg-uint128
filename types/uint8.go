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

type UInt8 struct {
	Uint64 uint64
	Valid  bool
}

func (n UInt8) Int64Value() (Int8, error) {
	if n.Uint64 > math.MaxInt64 {
		return Int8{}, fmt.Errorf("UInt8 value is greater than max Int8 value")
	}
	return Int8{Int64: int64(n.Uint64), Valid: n.Valid}, nil
}

func (n UInt8) Uint64Value() (UInt8, error) {
	return UInt8{Uint64: uint64(n.Uint64), Valid: n.Valid}, nil
}

// ScanInt64 implements the Int64Scanner interface.
func (dst *UInt8) ScanInt64(n Int8) error {
	if !n.Valid {
		*dst = UInt8{}
		return nil
	}

	if n.Int64 < 0 {
		return fmt.Errorf("%d is less than minimum value for UInt8", n.Int64)
	}

	*dst = UInt8{Uint64: uint64(n.Int64), Valid: true}

	return nil
}

// ScanUint64 implements the Uint64Scanner interface.
func (dst *UInt8) ScanUint64(n UInt8) error {
	if !n.Valid {
		*dst = UInt8{}
		return nil
	}

	*dst = UInt8{Uint64: uint64(n.Uint64), Valid: true}

	return nil
}

// Scan implements the database/sql Scanner interface.
func (dst *UInt8) Scan(src any) error {
	if src == nil {
		*dst = UInt8{}
		return nil
	}

	var n uint64

	switch src := src.(type) {
	case int64:
		if src < 0 {
			return fmt.Errorf("%d is less than minimum value for UInt8", n)
		}

		n = uint64(src)
	case uint64:
		n = src
	case string:
		var err error
		n, err = strconv.ParseUint(src, 10, 64)
		if err != nil {
			return err
		}
	case []byte:
		var err error
		n, err = strconv.ParseUint(string(src), 10, 64)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("cannot scan %T", src)
	}

	*dst = UInt8{Uint64: uint64(n), Valid: true}

	return nil
}

// Value implements the database/sql/driver Valuer interface.
func (src UInt8) Value() (driver.Value, error) {
	if !src.Valid {
		return nil, nil
	}
	return uint64(src.Uint64), nil
}

func (src UInt8) MarshalJSON() ([]byte, error) {
	if !src.Valid {
		return []byte("null"), nil
	}
	return []byte(strconv.FormatUint(uint64(src.Uint64), 10)), nil
}

func (dst *UInt8) UnmarshalJSON(b []byte) error {
	var n *uint64
	err := json.Unmarshal(b, &n)
	if err != nil {
		return err
	}

	if n == nil {
		*dst = UInt8{}
	} else {
		*dst = UInt8{Uint64: *n, Valid: true}
	}

	return nil
}

type UInt8Codec struct{}

func (UInt8Codec) FormatSupported(format int16) bool {
	return format == TextFormatCode || format == BinaryFormatCode
}

func (UInt8Codec) PreferredFormat() int16 {
	return BinaryFormatCode
}

func (UInt8Codec) PlanEncode(m *Map, oid uint32, format int16, value any) EncodePlan {
	switch format {
	case BinaryFormatCode:
		switch value.(type) {
		case uint64:
			return encodePlanUInt8CodecBinaryUint64{}
		case Uint64Valuer:
			return encodePlanUInt8CodecBinaryUint64Uint64Valuer{}
		case Int64Valuer:
			return encodePlanUInt8CodecBinaryUint64Int64Valuer{}
		}
	case TextFormatCode:
		switch value.(type) {
		case uint64:
			return encodePlanUInt8CodecTextUint64{}
		case Uint64Valuer:
			return encodePlanUInt8CodecTextUint64Uint64Valuer{}
		case Int64Valuer:
			return encodePlanUInt8CodecTextUint64Int64Valuer{}
		}
	}

	return nil
}

type encodePlanUInt8CodecBinaryUint64 struct{}

func (encodePlanUInt8CodecBinaryUint64) Encode(value any, buf []byte) (newBuf []byte, err error) {
	n := value.(uint64)
	return pgio.AppendUint64(buf, n), nil
}

type encodePlanUInt8CodecTextUint64 struct{}

func (encodePlanUInt8CodecTextUint64) Encode(value any, buf []byte) (newBuf []byte, err error) {
	n := value.(uint64)
	return append(buf, strconv.FormatUint(uint64(n), 10)...), nil
}

type encodePlanUInt8CodecBinaryUint64Int64Valuer struct{}

func (encodePlanUInt8CodecBinaryUint64Int64Valuer) Encode(value any, buf []byte) (newBuf []byte, err error) {
	n, err := value.(Int64Valuer).Int64Value()
	if err != nil {
		return nil, err
	}

	if !n.Valid {
		return nil, nil
	}

	if n.Int64 < 0 {
		return nil, fmt.Errorf("%d is less than minimum value for uint8", n.Int64)
	}

	return pgio.AppendUint64(buf, uint64(n.Int64)), nil
}

type encodePlanUInt8CodecTextUint64Int64Valuer struct{}

func (encodePlanUInt8CodecTextUint64Int64Valuer) Encode(value any, buf []byte) (newBuf []byte, err error) {
	n, err := value.(Int64Valuer).Int64Value()
	if err != nil {
		return nil, err
	}

	if !n.Valid {
		return nil, nil
	}

	if n.Int64 < 0 {
		return nil, fmt.Errorf("%d is less than minimum value for uint8", n.Int64)
	}

	return append(buf, strconv.FormatUint(uint64(n.Int64), 10)...), nil
}

type encodePlanUInt8CodecBinaryUint64Uint64Valuer struct{}

func (encodePlanUInt8CodecBinaryUint64Uint64Valuer) Encode(value any, buf []byte) (newBuf []byte, err error) {
	n, err := value.(Uint64Valuer).Uint64Value()
	if err != nil {
		return nil, err
	}

	if !n.Valid {
		return nil, nil
	}

	return pgio.AppendUint64(buf, uint64(n.Uint64)), nil
}

type encodePlanUInt8CodecTextUint64Uint64Valuer struct{}

func (encodePlanUInt8CodecTextUint64Uint64Valuer) Encode(value any, buf []byte) (newBuf []byte, err error) {
	n, err := value.(Uint64Valuer).Uint64Value()
	if err != nil {
		return nil, err
	}

	if !n.Valid {
		return nil, nil
	}

	return append(buf, strconv.FormatUint(uint64(n.Uint64), 10)...), nil
}

func (UInt8Codec) PlanScan(m *Map, oid uint32, format int16, target any) ScanPlan {
	switch format {
	case BinaryFormatCode:
		switch target.(type) {
		case *uint16:
			return scanPlanBinaryUInt8ToUint16{}
		case *uint32:
			return scanPlanBinaryUInt8ToUint32{}
		case *uint64:
			return scanPlanBinaryUInt8ToUint64{}
		case *uint128.Uint128:
			return scanPlanBinaryUInt8ToUint128{}
		case *uint:
			return scanPlanBinaryUInt8ToUint{}
		case *int16:
			return scanPlanBinaryUInt8ToInt16{}
		case *int32:
			return scanPlanBinaryUInt8ToInt32{}
		case *int64:
			return scanPlanBinaryUInt8ToInt64{}
		case *num.I128:
			return scanPlanBinaryUInt8ToI128{}
		case *int:
			return scanPlanBinaryUInt8ToInt{}
		case Int64Scanner:
			return scanPlanBinaryUInt8ToInt64Scanner{}
		case Uint64Scanner:
			return scanPlanBinaryUInt8ToUint64Scanner{}
		case TextScanner:
			return scanPlanBinaryUInt8ToTextScanner{}
		}
	case TextFormatCode:
		switch target.(type) {
		case *uint16:
			return scanPlanTextUInt8ToUint16{}
		case *uint32:
			return scanPlanTextUInt8ToUint32{}
		case *uint64:
			return scanPlanTextUInt8ToUint64{}
		case *uint128.Uint128:
			return scanPlanTextUInt8ToUint128{}
		case *uint:
			return scanPlanTextUInt8ToUint{}
		case *int16:
			return scanPlanTextUInt8ToInt16{}
		case *int32:
			return scanPlanTextUInt8ToInt32{}
		case *int64:
			return scanPlanTextUInt8ToInt64{}
		case *num.I128:
			return scanPlanTextUInt8ToI128{}
		case *int:
			return scanPlanTextUInt8ToInt{}
		case Int64Scanner:
			return scanPlanTextAnyToInt64Scanner{}
		case Uint64Scanner:
			return scanPlanTextAnyToUint64Scanner{}
		}
	}

	return nil
}

func (c UInt8Codec) DecodeDatabaseSQLValue(m *Map, oid uint32, format int16, src []byte) (driver.Value, error) {
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

func (c UInt8Codec) DecodeValue(m *Map, oid uint32, format int16, src []byte) (any, error) {
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

type scanPlanBinaryUInt8ToUint16 struct{}

func (scanPlanBinaryUInt8ToUint16) Scan(src []byte, dst any) error {
	if src == nil {
		return fmt.Errorf("cannot scan NULL into %T", dst)
	}

	if len(src) != 8 {
		return fmt.Errorf("invalid length for UInt8: %v", len(src))
	}

	p, ok := (dst).(*uint16)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := pgio.ReadUint64(src)
	if n > uint64(math.MaxUint16) {
		return fmt.Errorf("%d is greater than maximum value for uint16", n)
	}
	*p = uint16(n)

	return nil
}

type scanPlanBinaryUInt8ToUint32 struct{}

func (scanPlanBinaryUInt8ToUint32) Scan(src []byte, dst any) error {
	if src == nil {
		return fmt.Errorf("cannot scan NULL into %T", dst)
	}

	if len(src) != 8 {
		return fmt.Errorf("invalid length for UInt8: %v", len(src))
	}

	p, ok := (dst).(*uint32)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := pgio.ReadUint64(src)
	if n > uint64(math.MaxUint32) {
		return fmt.Errorf("%d is greater than maximum value for uint32", n)
	}
	*p = uint32(n)

	return nil
}

type scanPlanBinaryUInt8ToUint64 struct{}

func (scanPlanBinaryUInt8ToUint64) Scan(src []byte, dst any) error {
	if src == nil {
		return fmt.Errorf("cannot scan NULL into %T", dst)
	}

	if len(src) != 8 {
		return fmt.Errorf("invalid length for UInt8: %v", len(src))
	}

	p, ok := (dst).(*uint64)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := pgio.ReadUint64(src)

	*p = n

	return nil
}

type scanPlanBinaryUInt8ToUint128 struct{}

func (scanPlanBinaryUInt8ToUint128) Scan(src []byte, dst any) error {
	if src == nil {
		return fmt.Errorf("cannot scan NULL into %T", dst)
	}

	if len(src) != 8 {
		return fmt.Errorf("invalid length for UInt8: %v", len(src))
	}

	p, ok := (dst).(*uint128.Uint128)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := pgio.ReadUint64(src)

	*p = uint128.From64(uint64(n))

	return nil
}

type scanPlanBinaryUInt8ToUint struct{}

func (scanPlanBinaryUInt8ToUint) Scan(src []byte, dst any) error {
	if src == nil {
		return fmt.Errorf("cannot scan NULL into %T", dst)
	}

	if len(src) != 8 {
		return fmt.Errorf("invalid length for UInt8: %v", len(src))
	}

	p, ok := (dst).(*uint)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := pgio.ReadUint64(src)
	if intSize == 32 && n > uint64(math.MaxUint) {
		return fmt.Errorf("%d is greater than maximum value for uint", n)
	}
	*p = uint(n)

	return nil
}

type scanPlanBinaryUInt8ToInt16 struct{}

func (scanPlanBinaryUInt8ToInt16) Scan(src []byte, dst any) error {
	if src == nil {
		return fmt.Errorf("cannot scan NULL into %T", dst)
	}

	if len(src) != 8 {
		return fmt.Errorf("invalid length for UInt8: %v", len(src))
	}

	p, ok := (dst).(*int16)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := pgio.ReadUint64(src)
	if n > uint64(math.MaxInt16) {
		return fmt.Errorf("%d is greater than maximum value for int16", n)
	}
	*p = int16(n)

	return nil
}

type scanPlanBinaryUInt8ToInt32 struct{}

func (scanPlanBinaryUInt8ToInt32) Scan(src []byte, dst any) error {
	if src == nil {
		return fmt.Errorf("cannot scan NULL into %T", dst)
	}

	if len(src) != 8 {
		return fmt.Errorf("invalid length for UInt8: %v", len(src))
	}

	p, ok := (dst).(*int32)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := pgio.ReadUint64(src)
	if n > uint64(math.MaxInt32) {
		return fmt.Errorf("%d is greater than maximum value for int32", n)
	}
	*p = int32(n)

	return nil
}

type scanPlanBinaryUInt8ToInt64 struct{}

func (scanPlanBinaryUInt8ToInt64) Scan(src []byte, dst any) error {
	if src == nil {
		return fmt.Errorf("cannot scan NULL into %T", dst)
	}

	if len(src) != 8 {
		return fmt.Errorf("invalid length for UInt8: %v", len(src))
	}

	p, ok := (dst).(*int64)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := pgio.ReadUint64(src)
	if n > uint64(math.MaxInt64) {
		return fmt.Errorf("%d is greater than maximum value for int64", n)
	}
	*p = int64(n)

	return nil
}

type scanPlanBinaryUInt8ToI128 struct{}

func (scanPlanBinaryUInt8ToI128) Scan(src []byte, dst any) error {
	if src == nil {
		return fmt.Errorf("cannot scan NULL into %T", dst)
	}

	if len(src) != 8 {
		return fmt.Errorf("invalid length for UInt8: %v", len(src))
	}

	p, ok := (dst).(*num.I128)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := pgio.ReadUint64(src)

	*p = num.I128FromU64(uint64(n))

	return nil
}

type scanPlanBinaryUInt8ToInt struct{}

func (scanPlanBinaryUInt8ToInt) Scan(src []byte, dst any) error {
	if src == nil {
		return fmt.Errorf("cannot scan NULL into %T", dst)
	}

	if len(src) != 8 {
		return fmt.Errorf("invalid length for UInt8: %v", len(src))
	}

	p, ok := (dst).(*int)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := pgio.ReadUint64(src)
	var maxNum int = math.MaxInt
	if n > uint64(maxNum) {
		return fmt.Errorf("%d is greater than maximum value for int", n)
	}
	*p = int(n)

	return nil
}

type scanPlanBinaryUInt8ToTextScanner struct{}

func (scanPlanBinaryUInt8ToTextScanner) Scan(src []byte, dst any) error {
	s, ok := (dst).(TextScanner)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	if src == nil {
		return s.ScanText(Text{})
	}

	if len(src) != 8 {
		return fmt.Errorf("invalid length for uint8: %v", len(src))
	}

	n := uint64(pgio.ReadUint64(src))

	return s.ScanText(Text{String: strconv.FormatUint(uint64(n), 10), Valid: true})
}

type scanPlanBinaryUInt8ToInt64Scanner struct{}

func (scanPlanBinaryUInt8ToInt64Scanner) Scan(src []byte, dst any) error {
	s, ok := (dst).(Int64Scanner)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	if src == nil {
		return s.ScanInt64(Int8{})
	}

	if len(src) != 8 {
		return fmt.Errorf("invalid length for uint8: %v", len(src))
	}

	n := uint64(pgio.ReadUint64(src))
	if n > math.MaxInt64 {
		return fmt.Errorf("UInt8 value %d is greater than max value for Int8", n)
	}

	return s.ScanInt64(Int8{Int64: int64(n), Valid: true})
}

type scanPlanBinaryUInt8ToUint64Scanner struct{}

func (scanPlanBinaryUInt8ToUint64Scanner) Scan(src []byte, dst any) error {
	s, ok := (dst).(Uint64Scanner)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	if src == nil {
		return s.ScanUint64(UInt8{})
	}

	if len(src) != 8 {
		return fmt.Errorf("invalid length for uint8: %v", len(src))
	}

	n := uint64(pgio.ReadUint64(src))

	return s.ScanUint64(UInt8{Uint64: n, Valid: true})
}

type scanPlanTextUInt8ToUint16 struct{}

func (scanPlanTextUInt8ToUint16) Scan(src []byte, dst any) error {
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

type scanPlanTextUInt8ToUint32 struct{}

func (scanPlanTextUInt8ToUint32) Scan(src []byte, dst any) error {
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

type scanPlanTextUInt8ToUint64 struct{}

func (scanPlanTextUInt8ToUint64) Scan(src []byte, dst any) error {
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

type scanPlanTextUInt8ToUint128 struct{}

func (scanPlanTextUInt8ToUint128) Scan(src []byte, dst any) error {
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

type scanPlanTextUInt8ToUint struct{}

func (scanPlanTextUInt8ToUint) Scan(src []byte, dst any) error {
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

type scanPlanTextUInt8ToInt16 struct{}

func (scanPlanTextUInt8ToInt16) Scan(src []byte, dst any) error {
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

type scanPlanTextUInt8ToInt32 struct{}

func (scanPlanTextUInt8ToInt32) Scan(src []byte, dst any) error {
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

type scanPlanTextUInt8ToInt64 struct{}

func (scanPlanTextUInt8ToInt64) Scan(src []byte, dst any) error {
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

type scanPlanTextUInt8ToI128 struct{}

func (scanPlanTextUInt8ToI128) Scan(src []byte, dst any) error {
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

type scanPlanTextUInt8ToInt struct{}

func (scanPlanTextUInt8ToInt) Scan(src []byte, dst any) error {
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
