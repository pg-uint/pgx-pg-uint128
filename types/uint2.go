// Do not edit. Generated from codegen

package types

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"math"
	"strconv"

	"github.com/codercms/pgx-pg-uint128/pgio"

	. "github.com/jackc/pgx/v5/pgtype"

	"lukechampine.com/uint128"

	"github.com/codercms/pgx-pg-uint128/int128"
	"go.shabbyrobe.org/num"
)

type UInt2 struct {
	Uint16 uint16
	Valid  bool
}

func (n UInt2) Int64Value() (Int8, error) {
	return Int8{Int64: int64(n.Uint16), Valid: n.Valid}, nil
}

func (n UInt2) Uint64Value() (UInt8, error) {
	return UInt8{Uint64: uint64(n.Uint16), Valid: n.Valid}, nil
}

// ScanInt64 implements the Int64Scanner interface.
func (dst *UInt2) ScanInt64(n Int8) error {
	if !n.Valid {
		*dst = UInt2{}
		return nil
	}

	if n.Int64 < 0 {
		return fmt.Errorf("%d is less than minimum value for UInt2", n.Int64)
	}
	if n.Int64 > math.MaxUint16 {
		return fmt.Errorf("%d is greater than maximum value for UInt2", n.Int64)
	}
	*dst = UInt2{Uint16: uint16(n.Int64), Valid: true}

	return nil
}

// ScanUint64 implements the Uint64Scanner interface.
func (dst *UInt2) ScanUint64(n UInt8) error {
	if !n.Valid {
		*dst = UInt2{}
		return nil
	}
	if n.Uint64 > math.MaxUint16 {
		return fmt.Errorf("%d is greater than maximum value for UInt2", n.Uint64)
	}

	*dst = UInt2{Uint16: uint16(n.Uint64), Valid: true}

	return nil
}

// Scan implements the database/sql Scanner interface.
func (dst *UInt2) Scan(src any) error {
	if src == nil {
		*dst = UInt2{}
		return nil
	}

	var n uint64

	switch src := src.(type) {
	case int64:
		if src < 0 {
			return fmt.Errorf("%d is less than minimum value for UInt2", n)
		}

		n = uint64(src)
	case uint64:
		n = src
	case string:
		var err error
		n, err = strconv.ParseUint(src, 10, 16)
		if err != nil {
			return err
		}
	case []byte:
		var err error
		n, err = strconv.ParseUint(string(src), 10, 16)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("cannot scan %T", src)
	}

	if n > math.MaxUint16 {
		return fmt.Errorf("%d is greater than maximum value for UInt2", n)
	}

	*dst = UInt2{Uint16: uint16(n), Valid: true}

	return nil
}

// Value implements the database/sql/driver Valuer interface.
func (src UInt2) Value() (driver.Value, error) {
	if !src.Valid {
		return nil, nil
	}
	return uint64(src.Uint16), nil
}

func (src UInt2) MarshalJSON() ([]byte, error) {
	if !src.Valid {
		return []byte("null"), nil
	}
	return []byte(strconv.FormatUint(uint64(src.Uint16), 10)), nil
}

func (dst *UInt2) UnmarshalJSON(b []byte) error {
	var n *uint16
	err := json.Unmarshal(b, &n)
	if err != nil {
		return err
	}

	if n == nil {
		*dst = UInt2{}
	} else {
		*dst = UInt2{Uint16: *n, Valid: true}
	}

	return nil
}

type UInt2Codec struct{}

func (UInt2Codec) FormatSupported(format int16) bool {
	return format == TextFormatCode || format == BinaryFormatCode
}

func (UInt2Codec) PreferredFormat() int16 {
	return BinaryFormatCode
}

func (UInt2Codec) PlanEncode(m *Map, oid uint32, format int16, value any) EncodePlan {
	switch format {
	case BinaryFormatCode:
		switch value.(type) {
		case uint16:
			return encodePlanUInt2CodecBinaryUint16{}
		case Uint64Valuer:
			return encodePlanUInt2CodecBinaryUint16Uint64Valuer{}
		case Int64Valuer:
			return encodePlanUInt2CodecBinaryUint16Int64Valuer{}
		}
	case TextFormatCode:
		switch value.(type) {
		case uint16:
			return encodePlanUInt2CodecTextUint16{}
		case Uint64Valuer:
			return encodePlanUInt2CodecTextUint16Uint64Valuer{}
		case Int64Valuer:
			return encodePlanUInt2CodecTextUint16Int64Valuer{}
		}
	}

	return nil
}

type encodePlanUInt2CodecBinaryUint16 struct{}

func (encodePlanUInt2CodecBinaryUint16) Encode(value any, buf []byte) (newBuf []byte, err error) {
	n := value.(uint16)
	return pgio.AppendUint16(buf, n), nil
}

type encodePlanUInt2CodecTextUint16 struct{}

func (encodePlanUInt2CodecTextUint16) Encode(value any, buf []byte) (newBuf []byte, err error) {
	n := value.(uint16)
	return append(buf, strconv.FormatUint(uint64(n), 10)...), nil
}

type encodePlanUInt2CodecBinaryUint16Int64Valuer struct{}

func (encodePlanUInt2CodecBinaryUint16Int64Valuer) Encode(value any, buf []byte) (newBuf []byte, err error) {
	n, err := value.(Int64Valuer).Int64Value()
	if err != nil {
		return nil, err
	}

	if !n.Valid {
		return nil, nil
	}

	if n.Int64 < 0 {
		return nil, fmt.Errorf("%d is less than minimum value for uint2", n.Int64)
	}
	if n.Int64 > math.MaxUint16 {
		return nil, fmt.Errorf("%d is greater than maximum value for uint2", n.Int64)
	}

	return pgio.AppendUint16(buf, uint16(n.Int64)), nil
}

type encodePlanUInt2CodecTextUint16Int64Valuer struct{}

func (encodePlanUInt2CodecTextUint16Int64Valuer) Encode(value any, buf []byte) (newBuf []byte, err error) {
	n, err := value.(Int64Valuer).Int64Value()
	if err != nil {
		return nil, err
	}

	if !n.Valid {
		return nil, nil
	}

	if n.Int64 < 0 {
		return nil, fmt.Errorf("%d is less than minimum value for uint2", n.Int64)
	}
	if n.Int64 > math.MaxUint16 {
		return nil, fmt.Errorf("%d is greater than maximum value for uint2", n.Int64)
	}

	return append(buf, strconv.FormatUint(uint64(n.Int64), 10)...), nil
}

type encodePlanUInt2CodecBinaryUint16Uint64Valuer struct{}

func (encodePlanUInt2CodecBinaryUint16Uint64Valuer) Encode(value any, buf []byte) (newBuf []byte, err error) {
	n, err := value.(Uint64Valuer).Uint64Value()
	if err != nil {
		return nil, err
	}

	if !n.Valid {
		return nil, nil
	}

	if n.Uint64 > math.MaxUint16 {
		return nil, fmt.Errorf("%d is greater than maximum value for uint2", n.Uint64)
	}

	return pgio.AppendUint16(buf, uint16(n.Uint64)), nil
}

type encodePlanUInt2CodecTextUint16Uint64Valuer struct{}

func (encodePlanUInt2CodecTextUint16Uint64Valuer) Encode(value any, buf []byte) (newBuf []byte, err error) {
	n, err := value.(Uint64Valuer).Uint64Value()
	if err != nil {
		return nil, err
	}

	if !n.Valid {
		return nil, nil
	}

	if n.Uint64 > math.MaxUint16 {
		return nil, fmt.Errorf("%d is greater than maximum value for uint2", n.Uint64)
	}

	return append(buf, strconv.FormatUint(uint64(n.Uint64), 10)...), nil
}

func (UInt2Codec) PlanScan(m *Map, oid uint32, format int16, target any) ScanPlan {
	switch format {
	case BinaryFormatCode:
		switch target.(type) {
		case *uint16:
			return scanPlanBinaryUInt2ToUint16{}
		case *uint32:
			return scanPlanBinaryUInt2ToUint32{}
		case *uint64:
			return scanPlanBinaryUInt2ToUint64{}
		case *uint128.Uint128:
			return scanPlanBinaryUInt2ToUint128{}
		case *uint:
			return scanPlanBinaryUInt2ToUint{}
		case *int16:
			return scanPlanBinaryUInt2ToInt16{}
		case *int32:
			return scanPlanBinaryUInt2ToInt32{}
		case *int64:
			return scanPlanBinaryUInt2ToInt64{}
		case *num.I128:
			return scanPlanBinaryUInt2ToI128{}
		case *int:
			return scanPlanBinaryUInt2ToInt{}
		case Int64Scanner:
			return scanPlanBinaryUInt2ToInt64Scanner{}
		case Uint64Scanner:
			return scanPlanBinaryUInt2ToUint64Scanner{}
		case TextScanner:
			return scanPlanBinaryUInt2ToTextScanner{}
		}
	case TextFormatCode:
		switch target.(type) {
		case *uint16:
			return scanPlanTextUInt2ToUint16{}
		case *uint32:
			return scanPlanTextUInt2ToUint32{}
		case *uint64:
			return scanPlanTextUInt2ToUint64{}
		case *uint128.Uint128:
			return scanPlanTextUInt2ToUint128{}
		case *uint:
			return scanPlanTextUInt2ToUint{}
		case *int16:
			return scanPlanTextUInt2ToInt16{}
		case *int32:
			return scanPlanTextUInt2ToInt32{}
		case *int64:
			return scanPlanTextUInt2ToInt64{}
		case *num.I128:
			return scanPlanTextUInt2ToI128{}
		case *int:
			return scanPlanTextUInt2ToInt{}
		case Int64Scanner:
			return scanPlanTextAnyToInt64Scanner{}
		case Uint64Scanner:
			return scanPlanTextAnyToUint64Scanner{}
		}
	}

	return nil
}

func (c UInt2Codec) DecodeDatabaseSQLValue(m *Map, oid uint32, format int16, src []byte) (driver.Value, error) {
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

func (c UInt2Codec) DecodeValue(m *Map, oid uint32, format int16, src []byte) (any, error) {
	if src == nil {
		return nil, nil
	}

	var n uint16
	err := codecScan(c, m, oid, format, src, &n)
	if err != nil {
		return nil, err
	}
	return n, nil
}

type scanPlanBinaryUInt2ToUint16 struct{}

func (scanPlanBinaryUInt2ToUint16) Scan(src []byte, dst any) error {
	if src == nil {
		return fmt.Errorf("cannot scan NULL into %T", dst)
	}

	if len(src) != 2 {
		return fmt.Errorf("invalid length for UInt2: %v", len(src))
	}

	p, ok := (dst).(*uint16)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := pgio.ReadUint16(src)

	*p = n

	return nil
}

type scanPlanBinaryUInt2ToUint32 struct{}

func (scanPlanBinaryUInt2ToUint32) Scan(src []byte, dst any) error {
	if src == nil {
		return fmt.Errorf("cannot scan NULL into %T", dst)
	}

	if len(src) != 2 {
		return fmt.Errorf("invalid length for UInt2: %v", len(src))
	}

	p, ok := (dst).(*uint32)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := pgio.ReadUint16(src)

	*p = uint32(n)

	return nil
}

type scanPlanBinaryUInt2ToUint64 struct{}

func (scanPlanBinaryUInt2ToUint64) Scan(src []byte, dst any) error {
	if src == nil {
		return fmt.Errorf("cannot scan NULL into %T", dst)
	}

	if len(src) != 2 {
		return fmt.Errorf("invalid length for UInt2: %v", len(src))
	}

	p, ok := (dst).(*uint64)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := pgio.ReadUint16(src)

	*p = uint64(n)

	return nil
}

type scanPlanBinaryUInt2ToUint128 struct{}

func (scanPlanBinaryUInt2ToUint128) Scan(src []byte, dst any) error {
	if src == nil {
		return fmt.Errorf("cannot scan NULL into %T", dst)
	}

	if len(src) != 2 {
		return fmt.Errorf("invalid length for UInt2: %v", len(src))
	}

	p, ok := (dst).(*uint128.Uint128)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := pgio.ReadUint16(src)

	*p = uint128.From64(uint64(n))

	return nil
}

type scanPlanBinaryUInt2ToUint struct{}

func (scanPlanBinaryUInt2ToUint) Scan(src []byte, dst any) error {
	if src == nil {
		return fmt.Errorf("cannot scan NULL into %T", dst)
	}

	if len(src) != 2 {
		return fmt.Errorf("invalid length for UInt2: %v", len(src))
	}

	p, ok := (dst).(*uint)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := pgio.ReadUint16(src)

	*p = uint(n)

	return nil
}

type scanPlanBinaryUInt2ToInt16 struct{}

func (scanPlanBinaryUInt2ToInt16) Scan(src []byte, dst any) error {
	if src == nil {
		return fmt.Errorf("cannot scan NULL into %T", dst)
	}

	if len(src) != 2 {
		return fmt.Errorf("invalid length for UInt2: %v", len(src))
	}

	p, ok := (dst).(*int16)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := pgio.ReadUint16(src)
	if n > uint16(math.MaxInt16) {
		return fmt.Errorf("%d is greater than maximum value for int16", n)
	}
	*p = int16(n)

	return nil
}

type scanPlanBinaryUInt2ToInt32 struct{}

func (scanPlanBinaryUInt2ToInt32) Scan(src []byte, dst any) error {
	if src == nil {
		return fmt.Errorf("cannot scan NULL into %T", dst)
	}

	if len(src) != 2 {
		return fmt.Errorf("invalid length for UInt2: %v", len(src))
	}

	p, ok := (dst).(*int32)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := pgio.ReadUint16(src)

	*p = int32(n)

	return nil
}

type scanPlanBinaryUInt2ToInt64 struct{}

func (scanPlanBinaryUInt2ToInt64) Scan(src []byte, dst any) error {
	if src == nil {
		return fmt.Errorf("cannot scan NULL into %T", dst)
	}

	if len(src) != 2 {
		return fmt.Errorf("invalid length for UInt2: %v", len(src))
	}

	p, ok := (dst).(*int64)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := pgio.ReadUint16(src)

	*p = int64(n)

	return nil
}

type scanPlanBinaryUInt2ToI128 struct{}

func (scanPlanBinaryUInt2ToI128) Scan(src []byte, dst any) error {
	if src == nil {
		return fmt.Errorf("cannot scan NULL into %T", dst)
	}

	if len(src) != 2 {
		return fmt.Errorf("invalid length for UInt2: %v", len(src))
	}

	p, ok := (dst).(*num.I128)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := pgio.ReadUint16(src)

	*p = num.I128FromU64(uint64(n))

	return nil
}

type scanPlanBinaryUInt2ToInt struct{}

func (scanPlanBinaryUInt2ToInt) Scan(src []byte, dst any) error {
	if src == nil {
		return fmt.Errorf("cannot scan NULL into %T", dst)
	}

	if len(src) != 2 {
		return fmt.Errorf("invalid length for UInt2: %v", len(src))
	}

	p, ok := (dst).(*int)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := pgio.ReadUint16(src)

	*p = int(n)

	return nil
}

type scanPlanBinaryUInt2ToTextScanner struct{}

func (scanPlanBinaryUInt2ToTextScanner) Scan(src []byte, dst any) error {
	s, ok := (dst).(TextScanner)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	if src == nil {
		return s.ScanText(Text{})
	}

	if len(src) != 2 {
		return fmt.Errorf("invalid length for uint2: %v", len(src))
	}

	n := uint64(pgio.ReadUint16(src))

	return s.ScanText(Text{String: strconv.FormatUint(uint64(n), 10), Valid: true})
}

type scanPlanBinaryUInt2ToInt64Scanner struct{}

func (scanPlanBinaryUInt2ToInt64Scanner) Scan(src []byte, dst any) error {
	s, ok := (dst).(Int64Scanner)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	if src == nil {
		return s.ScanInt64(Int8{})
	}

	if len(src) != 2 {
		return fmt.Errorf("invalid length for uint2: %v", len(src))
	}

	n := uint64(pgio.ReadUint16(src))
	if n > math.MaxInt64 {
		return fmt.Errorf("UInt2 value %d is greater than max value for Int8", n)
	}

	return s.ScanInt64(Int8{Int64: int64(n), Valid: true})
}

type scanPlanBinaryUInt2ToUint64Scanner struct{}

func (scanPlanBinaryUInt2ToUint64Scanner) Scan(src []byte, dst any) error {
	s, ok := (dst).(Uint64Scanner)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	if src == nil {
		return s.ScanUint64(UInt8{})
	}

	if len(src) != 2 {
		return fmt.Errorf("invalid length for uint2: %v", len(src))
	}

	n := uint64(pgio.ReadUint16(src))

	return s.ScanUint64(UInt8{Uint64: n, Valid: true})
}

type scanPlanTextUInt2ToUint16 struct{}

func (scanPlanTextUInt2ToUint16) Scan(src []byte, dst any) error {
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

type scanPlanTextUInt2ToUint32 struct{}

func (scanPlanTextUInt2ToUint32) Scan(src []byte, dst any) error {
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

type scanPlanTextUInt2ToUint64 struct{}

func (scanPlanTextUInt2ToUint64) Scan(src []byte, dst any) error {
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

type scanPlanTextUInt2ToUint128 struct{}

func (scanPlanTextUInt2ToUint128) Scan(src []byte, dst any) error {
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

type scanPlanTextUInt2ToUint struct{}

func (scanPlanTextUInt2ToUint) Scan(src []byte, dst any) error {
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

type scanPlanTextUInt2ToInt16 struct{}

func (scanPlanTextUInt2ToInt16) Scan(src []byte, dst any) error {
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

type scanPlanTextUInt2ToInt32 struct{}

func (scanPlanTextUInt2ToInt32) Scan(src []byte, dst any) error {
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

type scanPlanTextUInt2ToInt64 struct{}

func (scanPlanTextUInt2ToInt64) Scan(src []byte, dst any) error {
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

type scanPlanTextUInt2ToI128 struct{}

func (scanPlanTextUInt2ToI128) Scan(src []byte, dst any) error {
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

type scanPlanTextUInt2ToInt struct{}

func (scanPlanTextUInt2ToInt) Scan(src []byte, dst any) error {
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
