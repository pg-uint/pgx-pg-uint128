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

type Int16 struct {
	I128  num.I128
	Valid bool
}

func (n Int16) Int64Value() (Int8, error) {
	if n.I128.Cmp64(int64(math.MaxInt64)) > 0 {
		return Int8{}, fmt.Errorf("Int16 value is greater than max Int8 value")
	}
	if n.I128.Cmp64(int64(math.MaxInt64)) < 0 {
		return Int8{}, fmt.Errorf("Int16 value is less than min Int8 value")
	}
	return Int8{Int64: n.I128.AsInt64(), Valid: n.Valid}, nil
}

func (n Int16) Uint64Value() (UInt8, error) {
	if n.I128.Cmp64(0) < 0 {
		return UInt8{}, fmt.Errorf("Int16 value is less than min UInt8 value")
	}
	if n.I128.Cmp(u64MaxInS128) > 0 {
		return UInt8{}, fmt.Errorf("Int16 value is greater than max UInt8 value")
	}
	return UInt8{Uint64: n.I128.AsUint64(), Valid: n.Valid}, nil
}

// ScanInt64 implements the Int64Scanner interface.
func (dst *Int16) ScanInt64(n Int8) error {
	if !n.Valid {
		*dst = Int16{}
		return nil
	}

	*dst = Int16{I128: num.I128From64(n.Int64), Valid: true}

	return nil
}

// ScanUint64 implements the Uint64Scanner interface.
func (dst *Int16) ScanUint64(n UInt8) error {
	if !n.Valid {
		*dst = Int16{}
		return nil
	}

	*dst = Int16{I128: num.I128FromU64(n.Uint64), Valid: true}

	return nil
}

// Scan implements the database/sql Scanner interface.
func (dst *Int16) Scan(src any) error {
	if src == nil {
		*dst = Int16{}
		return nil
	}

	var n num.I128

	switch src := src.(type) {
	case int64:
		n = num.I128From64(src)
	case uint64:
		n = num.I128FromU64(src)
	case num.I128:
		n = src
	case string:
		var err error
		n, err = int128.FromString(src)
		if err != nil {
			return err
		}
	case []byte:
		var err error
		n, err = int128.FromString(string(src))
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("cannot scan %T", src)
	}

	*dst = Int16{I128: n, Valid: true}

	return nil
}

// Value implements the database/sql/driver Valuer interface.
func (src Int16) Value() (driver.Value, error) {
	if !src.Valid {
		return nil, nil
	}
	return src.I128, nil
}

func (src Int16) MarshalJSON() ([]byte, error) {
	if !src.Valid {
		return []byte("null"), nil
	}
	return []byte(`"` + src.I128.String() + `"`), nil
}

func (dst *Int16) UnmarshalJSON(b []byte) error {
	var n *string
	err := json.Unmarshal(b, &n)
	if err != nil {
		return err
	}

	if n == nil {
		*dst = Int16{}
	} else {
		u, err := int128.FromString(*n)
		if err != nil {
			return err
		}

		*dst = Int16{I128: u, Valid: true}
	}

	return nil
}

type Int16Codec struct{}

func (Int16Codec) FormatSupported(format int16) bool {
	return format == TextFormatCode || format == BinaryFormatCode
}

func (Int16Codec) PreferredFormat() int16 {
	return BinaryFormatCode
}

func (Int16Codec) PlanEncode(m *Map, oid uint32, format int16, value any) EncodePlan {
	switch format {
	case BinaryFormatCode:
		switch value.(type) {
		case num.I128:
			return encodePlanInt16CodecBinaryI128{}
		case Uint64Valuer:
			return encodePlanInt16CodecBinaryI128Uint64Valuer{}
		case Int64Valuer:
			return encodePlanInt16CodecBinaryI128Int64Valuer{}
		}
	case TextFormatCode:
		switch value.(type) {
		case num.I128:
			return encodePlanInt16CodecTextI128{}
		case Uint64Valuer:
			return encodePlanInt16CodecTextI128Uint64Valuer{}
		case Int64Valuer:
			return encodePlanInt16CodecTextI128Int64Valuer{}
		}
	}

	return nil
}

type encodePlanInt16CodecBinaryI128 struct{}

func (encodePlanInt16CodecBinaryI128) Encode(value any, buf []byte) (newBuf []byte, err error) {
	n := value.(num.I128)
	return pgio.AppendInt128(buf, n), nil
}

type encodePlanInt16CodecTextI128 struct{}

func (encodePlanInt16CodecTextI128) Encode(value any, buf []byte) (newBuf []byte, err error) {
	n := value.(num.I128)
	return append(buf, n.String()...), nil
}

type encodePlanInt16CodecBinaryI128Int64Valuer struct{}

func (encodePlanInt16CodecBinaryI128Int64Valuer) Encode(value any, buf []byte) (newBuf []byte, err error) {
	n, err := value.(Int64Valuer).Int64Value()
	if err != nil {
		return nil, err
	}

	if !n.Valid {
		return nil, nil
	}

	return pgio.AppendInt128(buf, num.I128From64(n.Int64)), nil
}

type encodePlanInt16CodecTextI128Int64Valuer struct{}

func (encodePlanInt16CodecTextI128Int64Valuer) Encode(value any, buf []byte) (newBuf []byte, err error) {
	n, err := value.(Int64Valuer).Int64Value()
	if err != nil {
		return nil, err
	}

	if !n.Valid {
		return nil, nil
	}

	return append(buf, strconv.FormatInt(n.Int64, 10)...), nil
}

type encodePlanInt16CodecBinaryI128Uint64Valuer struct{}

func (encodePlanInt16CodecBinaryI128Uint64Valuer) Encode(value any, buf []byte) (newBuf []byte, err error) {
	n, err := value.(Uint64Valuer).Uint64Value()
	if err != nil {
		return nil, err
	}

	if !n.Valid {
		return nil, nil
	}

	return pgio.AppendInt128(buf, num.I128FromU64(n.Uint64)), nil
}

type encodePlanInt16CodecTextI128Uint64Valuer struct{}

func (encodePlanInt16CodecTextI128Uint64Valuer) Encode(value any, buf []byte) (newBuf []byte, err error) {
	n, err := value.(Uint64Valuer).Uint64Value()
	if err != nil {
		return nil, err
	}

	if !n.Valid {
		return nil, nil
	}

	return append(buf, strconv.FormatInt(int64(n.Uint64), 10)...), nil
}

func (Int16Codec) PlanScan(m *Map, oid uint32, format int16, target any) ScanPlan {
	switch format {
	case BinaryFormatCode:
		switch target.(type) {
		case *uint8:
			return scanPlanBinaryInt16ToUint8{}
		case *uint16:
			return scanPlanBinaryInt16ToUint16{}
		case *uint32:
			return scanPlanBinaryInt16ToUint32{}
		case *uint64:
			return scanPlanBinaryInt16ToUint64{}
		case *uint128.Uint128:
			return scanPlanBinaryInt16ToUint128{}
		case *uint:
			return scanPlanBinaryInt16ToUint{}
		case *int8:
			return scanPlanBinaryInt16ToInt8{}
		case *int16:
			return scanPlanBinaryInt16ToInt16{}
		case *int32:
			return scanPlanBinaryInt16ToInt32{}
		case *int64:
			return scanPlanBinaryInt16ToInt64{}
		case *num.I128:
			return scanPlanBinaryInt16ToI128{}
		case *int:
			return scanPlanBinaryInt16ToInt{}
		case Int64Scanner:
			return scanPlanBinaryInt16ToInt64Scanner{}
		case Uint64Scanner:
			return scanPlanBinaryInt16ToUint64Scanner{}
		case TextScanner:
			return scanPlanBinaryInt16ToTextScanner{}
		}
	case TextFormatCode:
		switch target.(type) {
		case *uint8:
			return scanPlanTextInt16ToUint8{}
		case *uint16:
			return scanPlanTextInt16ToUint16{}
		case *uint32:
			return scanPlanTextInt16ToUint32{}
		case *uint64:
			return scanPlanTextInt16ToUint64{}
		case *uint128.Uint128:
			return scanPlanTextInt16ToUint128{}
		case *uint:
			return scanPlanTextInt16ToUint{}
		case *int8:
			return scanPlanTextInt16ToInt8{}
		case *int16:
			return scanPlanTextInt16ToInt16{}
		case *int32:
			return scanPlanTextInt16ToInt32{}
		case *int64:
			return scanPlanTextInt16ToInt64{}
		case *num.I128:
			return scanPlanTextInt16ToI128{}
		case *int:
			return scanPlanTextInt16ToInt{}
		case Int64Scanner:
			return scanPlanTextAnyToInt64Scanner{}
		case Uint64Scanner:
			return scanPlanTextAnyToUint64Scanner{}
		}
	}

	return nil
}

func (c Int16Codec) DecodeDatabaseSQLValue(m *Map, oid uint32, format int16, src []byte) (driver.Value, error) {
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

func (c Int16Codec) DecodeValue(m *Map, oid uint32, format int16, src []byte) (any, error) {
	if src == nil {
		return nil, nil
	}

	var n num.I128
	err := codecScan(c, m, oid, format, src, &n)
	if err != nil {
		return nil, err
	}
	return n, nil
}

type scanPlanBinaryInt16ToUint8 struct{}

func (scanPlanBinaryInt16ToUint8) Scan(src []byte, dst any) error {
	if src == nil {
		return fmt.Errorf("cannot scan NULL into %T", dst)
	}

	if len(src) != 16 {
		return fmt.Errorf("invalid length for Int16: %v", len(src))
	}

	p, ok := (dst).(*uint8)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := pgio.ReadInt128(src)
	if n.Cmp64(0) < 0 {
		return fmt.Errorf("%s is less than minimum value for uint8", n.String())
	}
	if n.Cmp(u8MaxInS128) > 0 {
		return fmt.Errorf("%s is greater than maximum value for uint8", n.String())
	}
	*p = uint8(n.AsUint64())

	return nil
}

type scanPlanBinaryInt16ToUint16 struct{}

func (scanPlanBinaryInt16ToUint16) Scan(src []byte, dst any) error {
	if src == nil {
		return fmt.Errorf("cannot scan NULL into %T", dst)
	}

	if len(src) != 16 {
		return fmt.Errorf("invalid length for Int16: %v", len(src))
	}

	p, ok := (dst).(*uint16)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := pgio.ReadInt128(src)
	if n.Cmp64(0) < 0 {
		return fmt.Errorf("%s is less than minimum value for uint16", n.String())
	}
	if n.Cmp(u16MaxInS128) > 0 {
		return fmt.Errorf("%s is greater than maximum value for uint16", n.String())
	}
	*p = uint16(n.AsUint64())

	return nil
}

type scanPlanBinaryInt16ToUint32 struct{}

func (scanPlanBinaryInt16ToUint32) Scan(src []byte, dst any) error {
	if src == nil {
		return fmt.Errorf("cannot scan NULL into %T", dst)
	}

	if len(src) != 16 {
		return fmt.Errorf("invalid length for Int16: %v", len(src))
	}

	p, ok := (dst).(*uint32)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := pgio.ReadInt128(src)
	if n.Cmp64(0) < 0 {
		return fmt.Errorf("%s is less than minimum value for uint32", n.String())
	}
	if n.Cmp(u32MaxInS128) > 0 {
		return fmt.Errorf("%s is greater than maximum value for uint32", n.String())
	}
	*p = uint32(n.AsUint64())

	return nil
}

type scanPlanBinaryInt16ToUint64 struct{}

func (scanPlanBinaryInt16ToUint64) Scan(src []byte, dst any) error {
	if src == nil {
		return fmt.Errorf("cannot scan NULL into %T", dst)
	}

	if len(src) != 16 {
		return fmt.Errorf("invalid length for Int16: %v", len(src))
	}

	p, ok := (dst).(*uint64)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := pgio.ReadInt128(src)
	if n.Cmp64(0) < 0 {
		return fmt.Errorf("%s is less than minimum value for uint64", n.String())
	}
	if n.Cmp(u64MaxInS128) > 0 {
		return fmt.Errorf("%s is greater than maximum value for uint64", n.String())
	}
	*p = uint64(n.AsUint64())

	return nil
}

type scanPlanBinaryInt16ToUint128 struct{}

func (scanPlanBinaryInt16ToUint128) Scan(src []byte, dst any) error {
	if src == nil {
		return fmt.Errorf("cannot scan NULL into %T", dst)
	}

	if len(src) != 16 {
		return fmt.Errorf("invalid length for Int16: %v", len(src))
	}

	p, ok := (dst).(*uint128.Uint128)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := pgio.ReadInt128(src)
	if n.Cmp64(0) < 0 {
		return fmt.Errorf("%s is less than minimum value for Uint128", n.String())
	}
	hi, lo := n.Raw()
	*p = uint128.New(lo, hi)

	return nil
}

type scanPlanBinaryInt16ToUint struct{}

func (scanPlanBinaryInt16ToUint) Scan(src []byte, dst any) error {
	if src == nil {
		return fmt.Errorf("cannot scan NULL into %T", dst)
	}

	if len(src) != 16 {
		return fmt.Errorf("invalid length for Int16: %v", len(src))
	}

	p, ok := (dst).(*uint)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := pgio.ReadInt128(src)
	if n.Cmp64(0) < 0 {
		return fmt.Errorf("%s is less than minimum value for uint", n.String())
	}
	if n.Cmp(uMaxInS128) > 0 {
		return fmt.Errorf("%s is greater than maximum value for uint", n.String())
	}
	*p = uint(n.AsUint64())

	return nil
}

type scanPlanBinaryInt16ToInt8 struct{}

func (scanPlanBinaryInt16ToInt8) Scan(src []byte, dst any) error {
	if src == nil {
		return fmt.Errorf("cannot scan NULL into %T", dst)
	}

	if len(src) != 16 {
		return fmt.Errorf("invalid length for Int16: %v", len(src))
	}

	p, ok := (dst).(*int8)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := pgio.ReadInt128(src)
	if n.Cmp64(int64(math.MinInt8)) < 0 {
		return fmt.Errorf("%s is less than minimum value for int8", n.String())
	}
	if n.Cmp(s8MaxInS128) > 0 {
		return fmt.Errorf("%s is greater than maximum value for int8", n.String())
	}
	*p = int8(n.AsInt64())

	return nil
}

type scanPlanBinaryInt16ToInt16 struct{}

func (scanPlanBinaryInt16ToInt16) Scan(src []byte, dst any) error {
	if src == nil {
		return fmt.Errorf("cannot scan NULL into %T", dst)
	}

	if len(src) != 16 {
		return fmt.Errorf("invalid length for Int16: %v", len(src))
	}

	p, ok := (dst).(*int16)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := pgio.ReadInt128(src)
	if n.Cmp64(int64(math.MinInt16)) < 0 {
		return fmt.Errorf("%s is less than minimum value for int16", n.String())
	}
	if n.Cmp(s16MaxInS128) > 0 {
		return fmt.Errorf("%s is greater than maximum value for int16", n.String())
	}
	*p = int16(n.AsInt64())

	return nil
}

type scanPlanBinaryInt16ToInt32 struct{}

func (scanPlanBinaryInt16ToInt32) Scan(src []byte, dst any) error {
	if src == nil {
		return fmt.Errorf("cannot scan NULL into %T", dst)
	}

	if len(src) != 16 {
		return fmt.Errorf("invalid length for Int16: %v", len(src))
	}

	p, ok := (dst).(*int32)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := pgio.ReadInt128(src)
	if n.Cmp64(int64(math.MinInt32)) < 0 {
		return fmt.Errorf("%s is less than minimum value for int32", n.String())
	}
	if n.Cmp(s32MaxInS128) > 0 {
		return fmt.Errorf("%s is greater than maximum value for int32", n.String())
	}
	*p = int32(n.AsInt64())

	return nil
}

type scanPlanBinaryInt16ToInt64 struct{}

func (scanPlanBinaryInt16ToInt64) Scan(src []byte, dst any) error {
	if src == nil {
		return fmt.Errorf("cannot scan NULL into %T", dst)
	}

	if len(src) != 16 {
		return fmt.Errorf("invalid length for Int16: %v", len(src))
	}

	p, ok := (dst).(*int64)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := pgio.ReadInt128(src)
	if n.Cmp64(int64(math.MinInt64)) < 0 {
		return fmt.Errorf("%s is less than minimum value for int64", n.String())
	}
	if n.Cmp(s64MaxInS128) > 0 {
		return fmt.Errorf("%s is greater than maximum value for int64", n.String())
	}
	*p = int64(n.AsInt64())

	return nil
}

type scanPlanBinaryInt16ToI128 struct{}

func (scanPlanBinaryInt16ToI128) Scan(src []byte, dst any) error {
	if src == nil {
		return fmt.Errorf("cannot scan NULL into %T", dst)
	}

	if len(src) != 16 {
		return fmt.Errorf("invalid length for Int16: %v", len(src))
	}

	p, ok := (dst).(*num.I128)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := pgio.ReadInt128(src)

	*p = n

	return nil
}

type scanPlanBinaryInt16ToInt struct{}

func (scanPlanBinaryInt16ToInt) Scan(src []byte, dst any) error {
	if src == nil {
		return fmt.Errorf("cannot scan NULL into %T", dst)
	}

	if len(src) != 16 {
		return fmt.Errorf("invalid length for Int16: %v", len(src))
	}

	p, ok := (dst).(*int)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := pgio.ReadInt128(src)
	if n.Cmp64(int64(math.MinInt)) < 0 {
		return fmt.Errorf("%s is less than minimum value for int", n.String())
	}
	if n.Cmp(sMaxInS128) > 0 {
		return fmt.Errorf("%s is greater than maximum value for int", n.String())
	}
	*p = int(n.AsInt64())

	return nil
}

type scanPlanBinaryInt16ToTextScanner struct{}

func (scanPlanBinaryInt16ToTextScanner) Scan(src []byte, dst any) error {
	s, ok := (dst).(TextScanner)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	if src == nil {
		return s.ScanText(Text{})
	}

	if len(src) != 16 {
		return fmt.Errorf("invalid length for int16: %v", len(src))
	}

	n := pgio.ReadInt128(src)

	return s.ScanText(Text{String: n.String(), Valid: true})
}

type scanPlanBinaryInt16ToInt64Scanner struct{}

func (scanPlanBinaryInt16ToInt64Scanner) Scan(src []byte, dst any) error {
	s, ok := (dst).(Int64Scanner)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	if src == nil {
		return s.ScanInt64(Int8{})
	}

	if len(src) != 16 {
		return fmt.Errorf("invalid length for int16: %v", len(src))
	}

	n := pgio.ReadInt128(src)
	if n.Cmp64(math.MaxInt64) > 0 {
		return fmt.Errorf("Int16 value %s is greater than max value for Int8", n.String())
	}

	return s.ScanInt64(Int8{Int64: n.AsInt64(), Valid: true})
}

type scanPlanBinaryInt16ToUint64Scanner struct{}

func (scanPlanBinaryInt16ToUint64Scanner) Scan(src []byte, dst any) error {
	s, ok := (dst).(Uint64Scanner)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	if src == nil {
		return s.ScanUint64(UInt8{})
	}

	if len(src) != 16 {
		return fmt.Errorf("invalid length for int16: %v", len(src))
	}

	n := pgio.ReadInt128(src)
	if n.Cmp(u64MaxInS128) > 0 {
		return fmt.Errorf("Int16 value %s is greater than max value for UInt8", n.String())
	}

	return s.ScanUint64(UInt8{Uint64: n.AsUint64(), Valid: true})
}

type scanPlanTextInt16ToUint8 struct{}

func (scanPlanTextInt16ToUint8) Scan(src []byte, dst any) error {
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

type scanPlanTextInt16ToUint16 struct{}

func (scanPlanTextInt16ToUint16) Scan(src []byte, dst any) error {
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

type scanPlanTextInt16ToUint32 struct{}

func (scanPlanTextInt16ToUint32) Scan(src []byte, dst any) error {
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

type scanPlanTextInt16ToUint64 struct{}

func (scanPlanTextInt16ToUint64) Scan(src []byte, dst any) error {
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

type scanPlanTextInt16ToUint128 struct{}

func (scanPlanTextInt16ToUint128) Scan(src []byte, dst any) error {
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

type scanPlanTextInt16ToUint struct{}

func (scanPlanTextInt16ToUint) Scan(src []byte, dst any) error {
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

type scanPlanTextInt16ToInt8 struct{}

func (scanPlanTextInt16ToInt8) Scan(src []byte, dst any) error {
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

type scanPlanTextInt16ToInt16 struct{}

func (scanPlanTextInt16ToInt16) Scan(src []byte, dst any) error {
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

type scanPlanTextInt16ToInt32 struct{}

func (scanPlanTextInt16ToInt32) Scan(src []byte, dst any) error {
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

type scanPlanTextInt16ToInt64 struct{}

func (scanPlanTextInt16ToInt64) Scan(src []byte, dst any) error {
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

type scanPlanTextInt16ToI128 struct{}

func (scanPlanTextInt16ToI128) Scan(src []byte, dst any) error {
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

type scanPlanTextInt16ToInt struct{}

func (scanPlanTextInt16ToInt) Scan(src []byte, dst any) error {
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
