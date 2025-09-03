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

type UInt4 struct {
	Uint32 uint32
	Valid  bool
}

func (n UInt4) Int64Value() (Int8, error) {
	return Int8{Int64: int64(n.Uint32), Valid: n.Valid}, nil
}

func (n UInt4) Uint64Value() (UInt8, error) {
	return UInt8{Uint64: uint64(n.Uint32), Valid: n.Valid}, nil
}

// ScanInt64 implements the Int64Scanner interface.
func (dst *UInt4) ScanInt64(n Int8) error {
	if !n.Valid {
		*dst = UInt4{}
		return nil
	}

	if n.Int64 < 0 {
		return fmt.Errorf("%d is less than minimum value for UInt4", n.Int64)
	}
	if n.Int64 > math.MaxUint32 {
		return fmt.Errorf("%d is greater than maximum value for UInt4", n.Int64)
	}
	*dst = UInt4{Uint32: uint32(n.Int64), Valid: true}

	return nil
}

// ScanUint64 implements the Uint64Scanner interface.
func (dst *UInt4) ScanUint64(n UInt8) error {
	if !n.Valid {
		*dst = UInt4{}
		return nil
	}
	if n.Uint64 > math.MaxUint32 {
		return fmt.Errorf("%d is greater than maximum value for UInt4", n.Uint64)
	}

	*dst = UInt4{Uint32: uint32(n.Uint64), Valid: true}

	return nil
}

// Scan implements the database/sql Scanner interface.
func (dst *UInt4) Scan(src any) error {
	if src == nil {
		*dst = UInt4{}
		return nil
	}

	var n uint64

	switch src := src.(type) {
	case int64:
		if src < 0 {
			return fmt.Errorf("%d is less than minimum value for UInt4", n)
		}

		n = uint64(src)
	case uint64:
		n = src
	case string:
		var err error
		n, err = strconv.ParseUint(src, 10, 32)
		if err != nil {
			return err
		}
	case []byte:
		var err error
		n, err = strconv.ParseUint(string(src), 10, 32)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("cannot scan %T", src)
	}

	if n > math.MaxUint32 {
		return fmt.Errorf("%d is greater than maximum value for UInt4", n)
	}

	*dst = UInt4{Uint32: uint32(n), Valid: true}

	return nil
}

// Value implements the database/sql/driver Valuer interface.
func (src UInt4) Value() (driver.Value, error) {
	if !src.Valid {
		return nil, nil
	}
	return uint64(src.Uint32), nil
}

func (src UInt4) MarshalJSON() ([]byte, error) {
	if !src.Valid {
		return []byte("null"), nil
	}
	return []byte(strconv.FormatUint(uint64(src.Uint32), 10)), nil
}

func (dst *UInt4) UnmarshalJSON(b []byte) error {
	var n *uint32
	err := json.Unmarshal(b, &n)
	if err != nil {
		return err
	}

	if n == nil {
		*dst = UInt4{}
	} else {
		*dst = UInt4{Uint32: *n, Valid: true}
	}

	return nil
}

type UInt4Codec struct{}

func (UInt4Codec) FormatSupported(format int16) bool {
	return format == TextFormatCode || format == BinaryFormatCode
}

func (UInt4Codec) PreferredFormat() int16 {
	return BinaryFormatCode
}

func (UInt4Codec) PlanEncode(m *Map, oid uint32, format int16, value any) EncodePlan {
	switch format {
	case BinaryFormatCode:
		switch value.(type) {
		case uint32:
			return encodePlanUInt4CodecBinaryUint32{}
		case Uint64Valuer:
			return encodePlanUInt4CodecBinaryUint32Uint64Valuer{}
		case Int64Valuer:
			return encodePlanUInt4CodecBinaryUint32Int64Valuer{}
		}
	case TextFormatCode:
		switch value.(type) {
		case uint32:
			return encodePlanUInt4CodecTextUint32{}
		case Uint64Valuer:
			return encodePlanUInt4CodecTextUint32Uint64Valuer{}
		case Int64Valuer:
			return encodePlanUInt4CodecTextUint32Int64Valuer{}
		}
	}

	return nil
}

type encodePlanUInt4CodecBinaryUint32 struct{}

func (encodePlanUInt4CodecBinaryUint32) Encode(value any, buf []byte) (newBuf []byte, err error) {
	n := value.(uint32)
	return pgio.AppendUint32(buf, n), nil
}

type encodePlanUInt4CodecTextUint32 struct{}

func (encodePlanUInt4CodecTextUint32) Encode(value any, buf []byte) (newBuf []byte, err error) {
	n := value.(uint32)
	return append(buf, strconv.FormatUint(uint64(n), 10)...), nil
}

type encodePlanUInt4CodecBinaryUint32Int64Valuer struct{}

func (encodePlanUInt4CodecBinaryUint32Int64Valuer) Encode(value any, buf []byte) (newBuf []byte, err error) {
	n, err := value.(Int64Valuer).Int64Value()
	if err != nil {
		return nil, err
	}

	if !n.Valid {
		return nil, nil
	}

	if n.Int64 < 0 {
		return nil, fmt.Errorf("%d is less than minimum value for uint4", n.Int64)
	}
	if n.Int64 > math.MaxUint32 {
		return nil, fmt.Errorf("%d is greater than maximum value for uint4", n.Int64)
	}

	return pgio.AppendUint32(buf, uint32(n.Int64)), nil
}

type encodePlanUInt4CodecTextUint32Int64Valuer struct{}

func (encodePlanUInt4CodecTextUint32Int64Valuer) Encode(value any, buf []byte) (newBuf []byte, err error) {
	n, err := value.(Int64Valuer).Int64Value()
	if err != nil {
		return nil, err
	}

	if !n.Valid {
		return nil, nil
	}

	if n.Int64 < 0 {
		return nil, fmt.Errorf("%d is less than minimum value for uint4", n.Int64)
	}
	if n.Int64 > math.MaxUint32 {
		return nil, fmt.Errorf("%d is greater than maximum value for uint4", n.Int64)
	}

	return append(buf, strconv.FormatUint(uint64(n.Int64), 10)...), nil
}

type encodePlanUInt4CodecBinaryUint32Uint64Valuer struct{}

func (encodePlanUInt4CodecBinaryUint32Uint64Valuer) Encode(value any, buf []byte) (newBuf []byte, err error) {
	n, err := value.(Uint64Valuer).Uint64Value()
	if err != nil {
		return nil, err
	}

	if !n.Valid {
		return nil, nil
	}

	if n.Uint64 > math.MaxUint32 {
		return nil, fmt.Errorf("%d is greater than maximum value for uint4", n.Uint64)
	}

	return pgio.AppendUint32(buf, uint32(n.Uint64)), nil
}

type encodePlanUInt4CodecTextUint32Uint64Valuer struct{}

func (encodePlanUInt4CodecTextUint32Uint64Valuer) Encode(value any, buf []byte) (newBuf []byte, err error) {
	n, err := value.(Uint64Valuer).Uint64Value()
	if err != nil {
		return nil, err
	}

	if !n.Valid {
		return nil, nil
	}

	if n.Uint64 > math.MaxUint32 {
		return nil, fmt.Errorf("%d is greater than maximum value for uint4", n.Uint64)
	}

	return append(buf, strconv.FormatUint(uint64(n.Uint64), 10)...), nil
}

func (UInt4Codec) PlanScan(m *Map, oid uint32, format int16, target any) ScanPlan {
	switch format {
	case BinaryFormatCode:
		switch target.(type) {
		case *uint8:
			return scanPlanBinaryUInt4ToUint8{}
		case *uint16:
			return scanPlanBinaryUInt4ToUint16{}
		case *uint32:
			return scanPlanBinaryUInt4ToUint32{}
		case *uint64:
			return scanPlanBinaryUInt4ToUint64{}
		case *uint128.Uint128:
			return scanPlanBinaryUInt4ToUint128{}
		case *uint:
			return scanPlanBinaryUInt4ToUint{}
		case *int8:
			return scanPlanBinaryUInt4ToInt8{}
		case *int16:
			return scanPlanBinaryUInt4ToInt16{}
		case *int32:
			return scanPlanBinaryUInt4ToInt32{}
		case *int64:
			return scanPlanBinaryUInt4ToInt64{}
		case *num.I128:
			return scanPlanBinaryUInt4ToI128{}
		case *int:
			return scanPlanBinaryUInt4ToInt{}
		case Int64Scanner:
			return scanPlanBinaryUInt4ToInt64Scanner{}
		case Uint64Scanner:
			return scanPlanBinaryUInt4ToUint64Scanner{}
		case TextScanner:
			return scanPlanBinaryUInt4ToTextScanner{}
		}
	case TextFormatCode:
		switch target.(type) {
		case *uint8:
			return scanPlanTextUInt4ToUint8{}
		case *uint16:
			return scanPlanTextUInt4ToUint16{}
		case *uint32:
			return scanPlanTextUInt4ToUint32{}
		case *uint64:
			return scanPlanTextUInt4ToUint64{}
		case *uint128.Uint128:
			return scanPlanTextUInt4ToUint128{}
		case *uint:
			return scanPlanTextUInt4ToUint{}
		case *int8:
			return scanPlanTextUInt4ToInt8{}
		case *int16:
			return scanPlanTextUInt4ToInt16{}
		case *int32:
			return scanPlanTextUInt4ToInt32{}
		case *int64:
			return scanPlanTextUInt4ToInt64{}
		case *num.I128:
			return scanPlanTextUInt4ToI128{}
		case *int:
			return scanPlanTextUInt4ToInt{}
		case Int64Scanner:
			return scanPlanTextAnyToInt64Scanner{}
		case Uint64Scanner:
			return scanPlanTextAnyToUint64Scanner{}
		}
	}

	return nil
}

func (c UInt4Codec) DecodeDatabaseSQLValue(m *Map, oid uint32, format int16, src []byte) (driver.Value, error) {
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

func (c UInt4Codec) DecodeValue(m *Map, oid uint32, format int16, src []byte) (any, error) {
	if src == nil {
		return nil, nil
	}

	var n uint32
	err := codecScan(c, m, oid, format, src, &n)
	if err != nil {
		return nil, err
	}
	return n, nil
}

type scanPlanBinaryUInt4ToUint8 struct{}

func (scanPlanBinaryUInt4ToUint8) Scan(src []byte, dst any) error {
	if src == nil {
		return fmt.Errorf("cannot scan NULL into %T", dst)
	}

	if len(src) != 4 {
		return fmt.Errorf("invalid length for UInt4: %v", len(src))
	}

	p, ok := (dst).(*uint8)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := pgio.ReadUint32(src)
	if n > uint32(math.MaxUint8) {
		return fmt.Errorf("%d is greater than maximum value for uint8", n)
	}
	*p = uint8(n)

	return nil
}

type scanPlanBinaryUInt4ToUint16 struct{}

func (scanPlanBinaryUInt4ToUint16) Scan(src []byte, dst any) error {
	if src == nil {
		return fmt.Errorf("cannot scan NULL into %T", dst)
	}

	if len(src) != 4 {
		return fmt.Errorf("invalid length for UInt4: %v", len(src))
	}

	p, ok := (dst).(*uint16)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := pgio.ReadUint32(src)
	if n > uint32(math.MaxUint16) {
		return fmt.Errorf("%d is greater than maximum value for uint16", n)
	}
	*p = uint16(n)

	return nil
}

type scanPlanBinaryUInt4ToUint32 struct{}

func (scanPlanBinaryUInt4ToUint32) Scan(src []byte, dst any) error {
	if src == nil {
		return fmt.Errorf("cannot scan NULL into %T", dst)
	}

	if len(src) != 4 {
		return fmt.Errorf("invalid length for UInt4: %v", len(src))
	}

	p, ok := (dst).(*uint32)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := pgio.ReadUint32(src)

	*p = n

	return nil
}

type scanPlanBinaryUInt4ToUint64 struct{}

func (scanPlanBinaryUInt4ToUint64) Scan(src []byte, dst any) error {
	if src == nil {
		return fmt.Errorf("cannot scan NULL into %T", dst)
	}

	if len(src) != 4 {
		return fmt.Errorf("invalid length for UInt4: %v", len(src))
	}

	p, ok := (dst).(*uint64)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := pgio.ReadUint32(src)

	*p = uint64(n)

	return nil
}

type scanPlanBinaryUInt4ToUint128 struct{}

func (scanPlanBinaryUInt4ToUint128) Scan(src []byte, dst any) error {
	if src == nil {
		return fmt.Errorf("cannot scan NULL into %T", dst)
	}

	if len(src) != 4 {
		return fmt.Errorf("invalid length for UInt4: %v", len(src))
	}

	p, ok := (dst).(*uint128.Uint128)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := pgio.ReadUint32(src)

	*p = uint128.From64(uint64(n))

	return nil
}

type scanPlanBinaryUInt4ToUint struct{}

func (scanPlanBinaryUInt4ToUint) Scan(src []byte, dst any) error {
	if src == nil {
		return fmt.Errorf("cannot scan NULL into %T", dst)
	}

	if len(src) != 4 {
		return fmt.Errorf("invalid length for UInt4: %v", len(src))
	}

	p, ok := (dst).(*uint)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := pgio.ReadUint32(src)

	*p = uint(n)

	return nil
}

type scanPlanBinaryUInt4ToInt8 struct{}

func (scanPlanBinaryUInt4ToInt8) Scan(src []byte, dst any) error {
	if src == nil {
		return fmt.Errorf("cannot scan NULL into %T", dst)
	}

	if len(src) != 4 {
		return fmt.Errorf("invalid length for UInt4: %v", len(src))
	}

	p, ok := (dst).(*int8)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := pgio.ReadUint32(src)
	if n > uint32(math.MaxInt8) {
		return fmt.Errorf("%d is greater than maximum value for int8", n)
	}
	*p = int8(n)

	return nil
}

type scanPlanBinaryUInt4ToInt16 struct{}

func (scanPlanBinaryUInt4ToInt16) Scan(src []byte, dst any) error {
	if src == nil {
		return fmt.Errorf("cannot scan NULL into %T", dst)
	}

	if len(src) != 4 {
		return fmt.Errorf("invalid length for UInt4: %v", len(src))
	}

	p, ok := (dst).(*int16)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := pgio.ReadUint32(src)
	if n > uint32(math.MaxInt16) {
		return fmt.Errorf("%d is greater than maximum value for int16", n)
	}
	*p = int16(n)

	return nil
}

type scanPlanBinaryUInt4ToInt32 struct{}

func (scanPlanBinaryUInt4ToInt32) Scan(src []byte, dst any) error {
	if src == nil {
		return fmt.Errorf("cannot scan NULL into %T", dst)
	}

	if len(src) != 4 {
		return fmt.Errorf("invalid length for UInt4: %v", len(src))
	}

	p, ok := (dst).(*int32)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := pgio.ReadUint32(src)
	if n > uint32(math.MaxInt32) {
		return fmt.Errorf("%d is greater than maximum value for int32", n)
	}
	*p = int32(n)

	return nil
}

type scanPlanBinaryUInt4ToInt64 struct{}

func (scanPlanBinaryUInt4ToInt64) Scan(src []byte, dst any) error {
	if src == nil {
		return fmt.Errorf("cannot scan NULL into %T", dst)
	}

	if len(src) != 4 {
		return fmt.Errorf("invalid length for UInt4: %v", len(src))
	}

	p, ok := (dst).(*int64)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := pgio.ReadUint32(src)

	*p = int64(n)

	return nil
}

type scanPlanBinaryUInt4ToI128 struct{}

func (scanPlanBinaryUInt4ToI128) Scan(src []byte, dst any) error {
	if src == nil {
		return fmt.Errorf("cannot scan NULL into %T", dst)
	}

	if len(src) != 4 {
		return fmt.Errorf("invalid length for UInt4: %v", len(src))
	}

	p, ok := (dst).(*num.I128)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := pgio.ReadUint32(src)

	*p = num.I128FromU64(uint64(n))

	return nil
}

type scanPlanBinaryUInt4ToInt struct{}

func (scanPlanBinaryUInt4ToInt) Scan(src []byte, dst any) error {
	if src == nil {
		return fmt.Errorf("cannot scan NULL into %T", dst)
	}

	if len(src) != 4 {
		return fmt.Errorf("invalid length for UInt4: %v", len(src))
	}

	p, ok := (dst).(*int)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := pgio.ReadUint32(src)
	var maxNum int = math.MaxInt
	if intSize == 32 && n > uint32(maxNum) {
		return fmt.Errorf("%d is greater than maximum value for int", n)
	}
	*p = int(n)

	return nil
}

type scanPlanBinaryUInt4ToTextScanner struct{}

func (scanPlanBinaryUInt4ToTextScanner) Scan(src []byte, dst any) error {
	s, ok := (dst).(TextScanner)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	if src == nil {
		return s.ScanText(Text{})
	}

	if len(src) != 4 {
		return fmt.Errorf("invalid length for uint4: %v", len(src))
	}

	n := uint64(pgio.ReadUint32(src))

	return s.ScanText(Text{String: strconv.FormatUint(uint64(n), 10), Valid: true})
}

type scanPlanBinaryUInt4ToInt64Scanner struct{}

func (scanPlanBinaryUInt4ToInt64Scanner) Scan(src []byte, dst any) error {
	s, ok := (dst).(Int64Scanner)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	if src == nil {
		return s.ScanInt64(Int8{})
	}

	if len(src) != 4 {
		return fmt.Errorf("invalid length for uint4: %v", len(src))
	}

	n := uint64(pgio.ReadUint32(src))

	return s.ScanInt64(Int8{Int64: int64(n), Valid: true})
}

type scanPlanBinaryUInt4ToUint64Scanner struct{}

func (scanPlanBinaryUInt4ToUint64Scanner) Scan(src []byte, dst any) error {
	s, ok := (dst).(Uint64Scanner)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	if src == nil {
		return s.ScanUint64(UInt8{})
	}

	if len(src) != 4 {
		return fmt.Errorf("invalid length for uint4: %v", len(src))
	}

	n := uint64(pgio.ReadUint32(src))

	return s.ScanUint64(UInt8{Uint64: n, Valid: true})
}

type scanPlanTextUInt4ToUint8 struct{}

func (scanPlanTextUInt4ToUint8) Scan(src []byte, dst any) error {
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

type scanPlanTextUInt4ToUint16 struct{}

func (scanPlanTextUInt4ToUint16) Scan(src []byte, dst any) error {
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

type scanPlanTextUInt4ToUint32 struct{}

func (scanPlanTextUInt4ToUint32) Scan(src []byte, dst any) error {
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

type scanPlanTextUInt4ToUint64 struct{}

func (scanPlanTextUInt4ToUint64) Scan(src []byte, dst any) error {
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

type scanPlanTextUInt4ToUint128 struct{}

func (scanPlanTextUInt4ToUint128) Scan(src []byte, dst any) error {
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

type scanPlanTextUInt4ToUint struct{}

func (scanPlanTextUInt4ToUint) Scan(src []byte, dst any) error {
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

type scanPlanTextUInt4ToInt8 struct{}

func (scanPlanTextUInt4ToInt8) Scan(src []byte, dst any) error {
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

type scanPlanTextUInt4ToInt16 struct{}

func (scanPlanTextUInt4ToInt16) Scan(src []byte, dst any) error {
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

type scanPlanTextUInt4ToInt32 struct{}

func (scanPlanTextUInt4ToInt32) Scan(src []byte, dst any) error {
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

type scanPlanTextUInt4ToInt64 struct{}

func (scanPlanTextUInt4ToInt64) Scan(src []byte, dst any) error {
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

type scanPlanTextUInt4ToI128 struct{}

func (scanPlanTextUInt4ToI128) Scan(src []byte, dst any) error {
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

type scanPlanTextUInt4ToInt struct{}

func (scanPlanTextUInt4ToInt) Scan(src []byte, dst any) error {
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
