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

type UInt16 struct {
	Uint128 uint128.Uint128
	Valid   bool
}

func (n UInt16) Int64Value() (Int8, error) {
	if n.Uint128.Cmp64(uint64(math.MaxInt64)) > 0 {
		return Int8{}, fmt.Errorf("UInt16 value is greater than max Int8 value")
	}
	return Int8{Int64: int64(n.Uint128.Lo), Valid: n.Valid}, nil
}

func (n UInt16) Uint64Value() (UInt8, error) {
	return UInt8{Uint64: n.Uint128.Lo, Valid: n.Valid}, nil
}

// ScanInt64 implements the Int64Scanner interface.
func (dst *UInt16) ScanInt64(n Int8) error {
	if !n.Valid {
		*dst = UInt16{}
		return nil
	}

	if n.Int64 < 0 {
		return fmt.Errorf("%d is less than minimum value for UInt16", n.Int64)
	}

	*dst = UInt16{Uint128: uint128.From64(uint64(n.Int64)), Valid: true}

	return nil
}

// ScanUint64 implements the Uint64Scanner interface.
func (dst *UInt16) ScanUint64(n UInt8) error {
	if !n.Valid {
		*dst = UInt16{}
		return nil
	}

	*dst = UInt16{Uint128: uint128.From64(n.Uint64), Valid: true}

	return nil
}

// Scan implements the database/sql Scanner interface.
func (dst *UInt16) Scan(src any) error {
	if src == nil {
		*dst = UInt16{}
		return nil
	}

	var n uint128.Uint128

	switch src := src.(type) {
	case int64:
		if src < 0 {
			return fmt.Errorf("%d is greater than maximum value for UInt16", n)
		}

		n = uint128.From64(uint64(src))
	case uint64:
		n = uint128.From64(src)
	case uint128.Uint128:
		n = src
	case string:
		var err error
		n, err = uint128.FromString(src)
		if err != nil {
			return err
		}
	case []byte:
		var err error
		n, err = uint128.FromString(string(src))
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("cannot scan %T", src)
	}

	*dst = UInt16{Uint128: n, Valid: true}

	return nil
}

// Value implements the database/sql/driver Valuer interface.
func (src UInt16) Value() (driver.Value, error) {
	if !src.Valid {
		return nil, nil
	}
	return src.Uint128, nil
}

func (src UInt16) MarshalJSON() ([]byte, error) {
	if !src.Valid {
		return []byte("null"), nil
	}
	return []byte(`"` + src.Uint128.String() + `"`), nil
}

func (dst *UInt16) UnmarshalJSON(b []byte) error {
	var n *string
	err := json.Unmarshal(b, &n)
	if err != nil {
		return err
	}

	if n == nil {
		*dst = UInt16{}
	} else {
		u, err := uint128.FromString(*n)
		if err != nil {
			return err
		}

		*dst = UInt16{Uint128: u, Valid: true}
	}

	return nil
}

type UInt16Codec struct{}

func (UInt16Codec) FormatSupported(format int16) bool {
	return format == TextFormatCode || format == BinaryFormatCode
}

func (UInt16Codec) PreferredFormat() int16 {
	return BinaryFormatCode
}

func (UInt16Codec) PlanEncode(m *Map, oid uint32, format int16, value any) EncodePlan {
	switch format {
	case BinaryFormatCode:
		switch value.(type) {
		case uint128.Uint128:
			return encodePlanUInt16CodecBinaryUint128{}
		case Uint64Valuer:
			return encodePlanUInt16CodecBinaryUint128Uint64Valuer{}
		case Int64Valuer:
			return encodePlanUInt16CodecBinaryUint128Int64Valuer{}
		}
	case TextFormatCode:
		switch value.(type) {
		case uint128.Uint128:
			return encodePlanUInt16CodecTextUint128{}
		case Uint64Valuer:
			return encodePlanUInt16CodecTextUint128Uint64Valuer{}
		case Int64Valuer:
			return encodePlanUInt16CodecTextUint128Int64Valuer{}
		}
	}

	return nil
}

type encodePlanUInt16CodecBinaryUint128 struct{}

func (encodePlanUInt16CodecBinaryUint128) Encode(value any, buf []byte) (newBuf []byte, err error) {
	n := value.(uint128.Uint128)
	return pgio.AppendUint128(buf, n), nil
}

type encodePlanUInt16CodecTextUint128 struct{}

func (encodePlanUInt16CodecTextUint128) Encode(value any, buf []byte) (newBuf []byte, err error) {
	n := value.(uint128.Uint128)
	return append(buf, n.String()...), nil
}

type encodePlanUInt16CodecBinaryUint128Int64Valuer struct{}

func (encodePlanUInt16CodecBinaryUint128Int64Valuer) Encode(value any, buf []byte) (newBuf []byte, err error) {
	n, err := value.(Int64Valuer).Int64Value()
	if err != nil {
		return nil, err
	}

	if !n.Valid {
		return nil, nil
	}

	if n.Int64 < 0 {
		return nil, fmt.Errorf("%d is less than minimum value for uint16", n.Int64)
	}

	return pgio.AppendUint128(buf, uint128.From64(uint64(n.Int64))), nil
}

type encodePlanUInt16CodecTextUint128Int64Valuer struct{}

func (encodePlanUInt16CodecTextUint128Int64Valuer) Encode(value any, buf []byte) (newBuf []byte, err error) {
	n, err := value.(Int64Valuer).Int64Value()
	if err != nil {
		return nil, err
	}

	if !n.Valid {
		return nil, nil
	}

	if n.Int64 < 0 {
		return nil, fmt.Errorf("%d is less than minimum value for uint16", n.Int64)
	}

	return append(buf, strconv.FormatUint(uint64(n.Int64), 10)...), nil
}

type encodePlanUInt16CodecBinaryUint128Uint64Valuer struct{}

func (encodePlanUInt16CodecBinaryUint128Uint64Valuer) Encode(value any, buf []byte) (newBuf []byte, err error) {
	n, err := value.(Uint64Valuer).Uint64Value()
	if err != nil {
		return nil, err
	}

	if !n.Valid {
		return nil, nil
	}

	return pgio.AppendUint128(buf, uint128.From64(n.Uint64)), nil
}

type encodePlanUInt16CodecTextUint128Uint64Valuer struct{}

func (encodePlanUInt16CodecTextUint128Uint64Valuer) Encode(value any, buf []byte) (newBuf []byte, err error) {
	n, err := value.(Uint64Valuer).Uint64Value()
	if err != nil {
		return nil, err
	}

	if !n.Valid {
		return nil, nil
	}

	return append(buf, strconv.FormatUint(n.Uint64, 10)...), nil
}

func (UInt16Codec) PlanScan(m *Map, oid uint32, format int16, target any) ScanPlan {
	switch format {
	case BinaryFormatCode:
		switch target.(type) {
		case *uint16:
			return scanPlanBinaryUInt16ToUint16{}
		case *uint32:
			return scanPlanBinaryUInt16ToUint32{}
		case *uint64:
			return scanPlanBinaryUInt16ToUint64{}
		case *uint128.Uint128:
			return scanPlanBinaryUInt16ToUint128{}
		case *uint:
			return scanPlanBinaryUInt16ToUint{}
		case *int16:
			return scanPlanBinaryUInt16ToInt16{}
		case *int32:
			return scanPlanBinaryUInt16ToInt32{}
		case *int64:
			return scanPlanBinaryUInt16ToInt64{}
		case *num.I128:
			return scanPlanBinaryUInt16ToI128{}
		case *int:
			return scanPlanBinaryUInt16ToInt{}
		case Int64Scanner:
			return scanPlanBinaryUInt16ToInt64Scanner{}
		case Uint64Scanner:
			return scanPlanBinaryUInt16ToUint64Scanner{}
		case TextScanner:
			return scanPlanBinaryUInt16ToTextScanner{}
		}
	case TextFormatCode:
		switch target.(type) {
		case *uint16:
			return scanPlanTextUInt16ToUint16{}
		case *uint32:
			return scanPlanTextUInt16ToUint32{}
		case *uint64:
			return scanPlanTextUInt16ToUint64{}
		case *uint128.Uint128:
			return scanPlanTextUInt16ToUint128{}
		case *uint:
			return scanPlanTextUInt16ToUint{}
		case *int16:
			return scanPlanTextUInt16ToInt16{}
		case *int32:
			return scanPlanTextUInt16ToInt32{}
		case *int64:
			return scanPlanTextUInt16ToInt64{}
		case *num.I128:
			return scanPlanTextUInt16ToI128{}
		case *int:
			return scanPlanTextUInt16ToInt{}
		case Int64Scanner:
			return scanPlanTextAnyToInt64Scanner{}
		case Uint64Scanner:
			return scanPlanTextAnyToUint64Scanner{}
		}
	}

	return nil
}

func (c UInt16Codec) DecodeDatabaseSQLValue(m *Map, oid uint32, format int16, src []byte) (driver.Value, error) {
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

func (c UInt16Codec) DecodeValue(m *Map, oid uint32, format int16, src []byte) (any, error) {
	if src == nil {
		return nil, nil
	}

	var n uint128.Uint128
	err := codecScan(c, m, oid, format, src, &n)
	if err != nil {
		return nil, err
	}
	return n, nil
}

type scanPlanBinaryUInt16ToUint16 struct{}

func (scanPlanBinaryUInt16ToUint16) Scan(src []byte, dst any) error {
	if src == nil {
		return fmt.Errorf("cannot scan NULL into %T", dst)
	}

	if len(src) != 16 {
		return fmt.Errorf("invalid length for UInt16: %v", len(src))
	}

	p, ok := (dst).(*uint16)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := pgio.ReadUint128(src)
	if n.Cmp(u16MaxInU128) > 0 {
		return fmt.Errorf("%s is greater than maximum value for uint16", n.String())
	}
	*p = uint16(n.Lo)

	return nil
}

type scanPlanBinaryUInt16ToUint32 struct{}

func (scanPlanBinaryUInt16ToUint32) Scan(src []byte, dst any) error {
	if src == nil {
		return fmt.Errorf("cannot scan NULL into %T", dst)
	}

	if len(src) != 16 {
		return fmt.Errorf("invalid length for UInt16: %v", len(src))
	}

	p, ok := (dst).(*uint32)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := pgio.ReadUint128(src)
	if n.Cmp(u32MaxInU128) > 0 {
		return fmt.Errorf("%s is greater than maximum value for uint32", n.String())
	}
	*p = uint32(n.Lo)

	return nil
}

type scanPlanBinaryUInt16ToUint64 struct{}

func (scanPlanBinaryUInt16ToUint64) Scan(src []byte, dst any) error {
	if src == nil {
		return fmt.Errorf("cannot scan NULL into %T", dst)
	}

	if len(src) != 16 {
		return fmt.Errorf("invalid length for UInt16: %v", len(src))
	}

	p, ok := (dst).(*uint64)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := pgio.ReadUint128(src)
	if n.Cmp(u64MaxInU128) > 0 {
		return fmt.Errorf("%s is greater than maximum value for uint64", n.String())
	}
	*p = uint64(n.Lo)

	return nil
}

type scanPlanBinaryUInt16ToUint128 struct{}

func (scanPlanBinaryUInt16ToUint128) Scan(src []byte, dst any) error {
	if src == nil {
		return fmt.Errorf("cannot scan NULL into %T", dst)
	}

	if len(src) != 16 {
		return fmt.Errorf("invalid length for UInt16: %v", len(src))
	}

	p, ok := (dst).(*uint128.Uint128)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := pgio.ReadUint128(src)

	*p = n

	return nil
}

type scanPlanBinaryUInt16ToUint struct{}

func (scanPlanBinaryUInt16ToUint) Scan(src []byte, dst any) error {
	if src == nil {
		return fmt.Errorf("cannot scan NULL into %T", dst)
	}

	if len(src) != 16 {
		return fmt.Errorf("invalid length for UInt16: %v", len(src))
	}

	p, ok := (dst).(*uint)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := pgio.ReadUint128(src)
	if n.Cmp(uMaxInU128) > 0 {
		return fmt.Errorf("%s is greater than maximum value for uint", n.String())
	}
	*p = uint(n.Lo)

	return nil
}

type scanPlanBinaryUInt16ToInt16 struct{}

func (scanPlanBinaryUInt16ToInt16) Scan(src []byte, dst any) error {
	if src == nil {
		return fmt.Errorf("cannot scan NULL into %T", dst)
	}

	if len(src) != 16 {
		return fmt.Errorf("invalid length for UInt16: %v", len(src))
	}

	p, ok := (dst).(*int16)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := pgio.ReadUint128(src)
	if n.Cmp(s16MaxInU128) > 0 {
		return fmt.Errorf("%s is greater than maximum value for int16", n.String())
	}
	*p = int16(n.Lo)

	return nil
}

type scanPlanBinaryUInt16ToInt32 struct{}

func (scanPlanBinaryUInt16ToInt32) Scan(src []byte, dst any) error {
	if src == nil {
		return fmt.Errorf("cannot scan NULL into %T", dst)
	}

	if len(src) != 16 {
		return fmt.Errorf("invalid length for UInt16: %v", len(src))
	}

	p, ok := (dst).(*int32)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := pgio.ReadUint128(src)
	if n.Cmp(s32MaxInU128) > 0 {
		return fmt.Errorf("%s is greater than maximum value for int32", n.String())
	}
	*p = int32(n.Lo)

	return nil
}

type scanPlanBinaryUInt16ToInt64 struct{}

func (scanPlanBinaryUInt16ToInt64) Scan(src []byte, dst any) error {
	if src == nil {
		return fmt.Errorf("cannot scan NULL into %T", dst)
	}

	if len(src) != 16 {
		return fmt.Errorf("invalid length for UInt16: %v", len(src))
	}

	p, ok := (dst).(*int64)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := pgio.ReadUint128(src)
	if n.Cmp(s64MaxInU128) > 0 {
		return fmt.Errorf("%s is greater than maximum value for int64", n.String())
	}
	*p = int64(n.Lo)

	return nil
}

type scanPlanBinaryUInt16ToI128 struct{}

func (scanPlanBinaryUInt16ToI128) Scan(src []byte, dst any) error {
	if src == nil {
		return fmt.Errorf("cannot scan NULL into %T", dst)
	}

	if len(src) != 16 {
		return fmt.Errorf("invalid length for UInt16: %v", len(src))
	}

	p, ok := (dst).(*num.I128)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := pgio.ReadUint128(src)
	if n.Cmp(s128MaxInU128) > 0 {
		return fmt.Errorf("%s is greater than maximum value for I128", n.String())
	}
	*p = num.I128FromRaw(n.Hi, n.Lo)

	return nil
}

type scanPlanBinaryUInt16ToInt struct{}

func (scanPlanBinaryUInt16ToInt) Scan(src []byte, dst any) error {
	if src == nil {
		return fmt.Errorf("cannot scan NULL into %T", dst)
	}

	if len(src) != 16 {
		return fmt.Errorf("invalid length for UInt16: %v", len(src))
	}

	p, ok := (dst).(*int)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := pgio.ReadUint128(src)
	if n.Cmp(sMaxInU128) > 0 {
		return fmt.Errorf("%s is greater than maximum value for int", n.String())
	}
	*p = int(n.Lo)

	return nil
}

type scanPlanBinaryUInt16ToTextScanner struct{}

func (scanPlanBinaryUInt16ToTextScanner) Scan(src []byte, dst any) error {
	s, ok := (dst).(TextScanner)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	if src == nil {
		return s.ScanText(Text{})
	}

	if len(src) != 16 {
		return fmt.Errorf("invalid length for uint16: %v", len(src))
	}

	n := pgio.ReadUint128(src)

	return s.ScanText(Text{String: n.String(), Valid: true})
}

type scanPlanBinaryUInt16ToInt64Scanner struct{}

func (scanPlanBinaryUInt16ToInt64Scanner) Scan(src []byte, dst any) error {
	s, ok := (dst).(Int64Scanner)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	if src == nil {
		return s.ScanInt64(Int8{})
	}

	if len(src) != 16 {
		return fmt.Errorf("invalid length for uint16: %v", len(src))
	}

	n := pgio.ReadUint128(src)
	if n.Cmp64(uint64(math.MaxInt64)) > 0 {
		return fmt.Errorf("UInt16 value %s is greater than max value for Int8", n.String())
	}

	return s.ScanInt64(Int8{Int64: int64(n.Lo), Valid: true})
}

type scanPlanBinaryUInt16ToUint64Scanner struct{}

func (scanPlanBinaryUInt16ToUint64Scanner) Scan(src []byte, dst any) error {
	s, ok := (dst).(Uint64Scanner)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	if src == nil {
		return s.ScanUint64(UInt8{})
	}

	if len(src) != 16 {
		return fmt.Errorf("invalid length for uint16: %v", len(src))
	}

	n := pgio.ReadUint128(src)
	if n.Cmp64(math.MaxUint64) > 0 {
		return fmt.Errorf("UInt16 value %s is greater than max value for UInt8", n.String())
	}

	return s.ScanUint64(UInt8{Uint64: n.Lo, Valid: true})
}

type scanPlanTextUInt16ToUint16 struct{}

func (scanPlanTextUInt16ToUint16) Scan(src []byte, dst any) error {
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

type scanPlanTextUInt16ToUint32 struct{}

func (scanPlanTextUInt16ToUint32) Scan(src []byte, dst any) error {
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

type scanPlanTextUInt16ToUint64 struct{}

func (scanPlanTextUInt16ToUint64) Scan(src []byte, dst any) error {
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

type scanPlanTextUInt16ToUint128 struct{}

func (scanPlanTextUInt16ToUint128) Scan(src []byte, dst any) error {
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

type scanPlanTextUInt16ToUint struct{}

func (scanPlanTextUInt16ToUint) Scan(src []byte, dst any) error {
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

type scanPlanTextUInt16ToInt16 struct{}

func (scanPlanTextUInt16ToInt16) Scan(src []byte, dst any) error {
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

type scanPlanTextUInt16ToInt32 struct{}

func (scanPlanTextUInt16ToInt32) Scan(src []byte, dst any) error {
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

type scanPlanTextUInt16ToInt64 struct{}

func (scanPlanTextUInt16ToInt64) Scan(src []byte, dst any) error {
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

type scanPlanTextUInt16ToI128 struct{}

func (scanPlanTextUInt16ToI128) Scan(src []byte, dst any) error {
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

type scanPlanTextUInt16ToInt struct{}

func (scanPlanTextUInt16ToInt) Scan(src []byte, dst any) error {
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
