<?php

declare(strict_types=1);

require_once __DIR__ . '/Type.php';
require_once __DIR__ . '/types.php';

function genType(TypeConfig $typeConfig): string
{
    $type = $typeConfig->type;
//    if (!$type->isUnsigned) {
//        throw new InvalidArgumentException("Only unsigned types supported");
//    }

    $structProp = $type->getStructPropName();

    $wideFmtType = $type->isUnsigned ? 'uint64' : 'int64';

    $maxCheck = '';
    if ($type->bitSize < 64) {
        $maxCheck = <<<GO
	if n.Int64 > {$type->maxValConst} {
		return fmt.Errorf("%d is greater than maximum value for {$type->name}", n.Int64)
	}
GO;
    }

    $int64ValFunc = genTypeInt64ValueFunc($type);
    $uint64ValFunc = genTypeUInt64ValueFunc($type);

    $scanUint64MaxCheck = '';
    if (UINT64->canOverflow($type)) {
        $scanUint64MaxCheck = <<<GO
	if n.Uint64 > {$type->maxValConst} {
		return fmt.Errorf("%d is greater than maximum value for {$type->name}", n.Uint64)
	}

GO;
    }

    $scanMaxCheck = '';
    if (($type->isUnsigned ? UINT64 : INT64)->canOverflow($type)) {
        $scanMaxCheck = <<<GO
	if n > {$type->maxValConst} {
		return fmt.Errorf("%d is greater than maximum value for {$type->name}", n)
	}

GO;

    }

    $uint64Case = $type->isUnsigned ? 'n = src' : 'n = int64(src)';

    $buf = <<<GO
type $type->name struct {
	$structProp $type->goType
	Valid bool
}

$int64ValFunc

$uint64ValFunc

// ScanInt64 implements the Int64Scanner interface.
func (dst *{$type->name}) ScanInt64(n Int8) error {
	if !n.Valid {
		*dst = {$type->name}{}
		return nil
	}

	if n.Int64 < {$type->minValConst} {
		return fmt.Errorf("%d is less than minimum value for {$type->name}", n.Int64)
	}
$maxCheck
	*dst = {$type->name}{{$structProp}: {$type->goType}(n.Int64), Valid: true}

	return nil
}

// ScanUint64 implements the Uint64Scanner interface.
func (dst *{$type->name}) ScanUint64(n UInt8) error {
	if !n.Valid {
		*dst = {$type->name}{}
		return nil
	}
	$scanUint64MaxCheck
	*dst = {$type->name}{{$structProp}: {$type->goType}(n.Uint64), Valid: true}

	return nil
}

// Scan implements the database/sql Scanner interface.
func (dst *{$type->name}) Scan(src any) error {
	if src == nil {
		*dst = {$type->name}{}
		return nil
	}

	var n $wideFmtType

	switch src := src.(type) {
	case int64:
        if src < {$type->minValConst} {
		    return fmt.Errorf("%d is less than minimum value for {$type->name}", n)
	    }

		n = $wideFmtType(src)
    case uint64:
        $uint64Case
	case string:
		var err error
		n, err = {$type->getParseIntFunctionCall("src")}
		if err != nil {
			return err
		}
	case []byte:
		var err error
		n, err = {$type->getParseIntFunctionCall("string(src)")}
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("cannot scan %T", src)
	}

	$scanMaxCheck
	*dst = {$type->name}{{$structProp}: {$type->goType}(n), Valid: true}

	return nil
}

// Value implements the database/sql/driver Valuer interface.
func (src {$type->name}) Value() (driver.Value, error) {
	if !src.Valid {
		return nil, nil
	}
	return $wideFmtType(src.{$structProp}), nil
}

func (src {$type->name}) MarshalJSON() ([]byte, error) {
	if !src.Valid {
		return []byte("null"), nil
	}
	return []byte({$type->getFormatIntFunctionName()}($wideFmtType(src.{$structProp}), 10)), nil
}

func (dst *{$type->name}) UnmarshalJSON(b []byte) error {
	var n *{$type->goType}
	err := json.Unmarshal(b, &n)
	if err != nil {
		return err
	}

	if n == nil {
		*dst = {$type->name}{}
	} else {
		*dst = {$type->name}{{$structProp}: *n, Valid: true}
	}

	return nil
}
GO;

    $buf .= "\n\n";
    $buf .= genCodec($typeConfig) . "\n";

    return $buf;
}

function genTypeU128(TypeConfig $typeConfig): string
{
    $type = $typeConfig->type;
    if ($type !== UINT128) {
        throw new InvalidArgumentException("Only uint128 type is supported");
    }

    $structProp = $type->getStructPropName();

    $int64ValFunc = genTypeInt64ValueFunc($type);
    $uint64ValFunc = genTypeUInt64ValueFunc($type);

    $buf = <<<GO
type $type->name struct {
	$structProp {$type->getFullGoTypeName()}
	Valid bool
}

$int64ValFunc

$uint64ValFunc

// ScanInt64 implements the Int64Scanner interface.
func (dst *{$type->name}) ScanInt64(n Int8) error {
	if !n.Valid {
		*dst = {$type->name}{}
		return nil
	}

	if n.Int64 < {$type->minValConst} {
		return fmt.Errorf("%d is less than minimum value for {$type->name}", n.Int64)
	}

	*dst = {$type->name}{{$structProp}: uint128.From64(uint64(n.Int64)), Valid: true}

	return nil
}

// ScanUint64 implements the Uint64Scanner interface.
func (dst *{$type->name}) ScanUint64(n UInt8) error {
	if !n.Valid {
		*dst = {$type->name}{}
		return nil
	}

	*dst = {$type->name}{{$structProp}: uint128.From64(n.Uint64), Valid: true}

	return nil
}

// Scan implements the database/sql Scanner interface.
func (dst *{$type->name}) Scan(src any) error {
	if src == nil {
		*dst = {$type->name}{}
		return nil
	}

	var n {$type->getFullGoTypeName()}

	switch src := src.(type) {
	case int64:
        if src < {$type->minValConst} {
		    return fmt.Errorf("%d is greater than maximum value for {$type->name}", n)
	    }

		n = uint128.From64(uint64(src))
    case uint64:
        n = uint128.From64(src)
    case {$type->getFullGoTypeName()}:
        n = src
	case string:
		var err error
		n, err = {$type->getParseIntFunctionCall("src")}
		if err != nil {
			return err
		}
	case []byte:
		var err error
		n, err = {$type->getParseIntFunctionCall("string(src)")}
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("cannot scan %T", src)
	}

	*dst = {$type->name}{{$structProp}: n, Valid: true}

	return nil
}

// Value implements the database/sql/driver Valuer interface.
func (src {$type->name}) Value() (driver.Value, error) {
	if !src.Valid {
		return nil, nil
	}
	return src.{$structProp}, nil
}

func (src {$type->name}) MarshalJSON() ([]byte, error) {
	if !src.Valid {
		return []byte("null"), nil
	}
	return []byte(`"` + src.{$structProp}.String() + `"`), nil
}

func (dst *{$type->name}) UnmarshalJSON(b []byte) error {
	var n *string
	err := json.Unmarshal(b, &n)
	if err != nil {
		return err
	}

	if n == nil {
		*dst = {$type->name}{}
	} else {
	    u, err := uint128.FromString(*n)
	    if err != nil {
	        return err
	    }
	    
		*dst = {$type->name}{{$structProp}: u, Valid: true}
	}

	return nil
}
GO;

    $buf .= "\n\n";
    $buf .= genCodec($typeConfig) . "\n";

    return $buf;
}

function genTypeS128(TypeConfig $typeConfig): string
{
    $type = $typeConfig->type;
    if ($type !== INT128) {
        throw new InvalidArgumentException("Only int128 type is supported");
    }

    $structProp = $type->getStructPropName();

    $int64ValFunc = genTypeInt64ValueFunc($type);
    $uint64ValFunc = genTypeUInt64ValueFunc($type);

    $buf = <<<GO
type $type->name struct {
	$structProp {$type->getFullGoTypeName()}
	Valid bool
}

$int64ValFunc

$uint64ValFunc

// ScanInt64 implements the Int64Scanner interface.
func (dst *{$type->name}) ScanInt64(n Int8) error {
	if !n.Valid {
		*dst = {$type->name}{}
		return nil
	}

	*dst = {$type->name}{{$structProp}: num.I128From64(n.Int64), Valid: true}

	return nil
}

// ScanUint64 implements the Uint64Scanner interface.
func (dst *{$type->name}) ScanUint64(n UInt8) error {
	if !n.Valid {
		*dst = {$type->name}{}
		return nil
	}

	*dst = {$type->name}{{$structProp}: num.I128FromU64(n.Uint64), Valid: true}

	return nil
}

// Scan implements the database/sql Scanner interface.
func (dst *{$type->name}) Scan(src any) error {
	if src == nil {
		*dst = {$type->name}{}
		return nil
	}

	var n {$type->getFullGoTypeName()}

	switch src := src.(type) {
	case int64:
		n = num.I128From64(src)
    case uint64:
        n = num.I128FromU64(src)
    case {$type->getFullGoTypeName()}:
        n = src
	case string:
		var err error
		n, err = {$type->getParseIntFunctionCall("src")}
		if err != nil {
			return err
		}
	case []byte:
		var err error
		n, err = {$type->getParseIntFunctionCall("string(src)")}
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("cannot scan %T", src)
	}

	*dst = {$type->name}{{$structProp}: n, Valid: true}

	return nil
}

// Value implements the database/sql/driver Valuer interface.
func (src {$type->name}) Value() (driver.Value, error) {
	if !src.Valid {
		return nil, nil
	}
	return src.{$structProp}, nil
}

func (src {$type->name}) MarshalJSON() ([]byte, error) {
	if !src.Valid {
		return []byte("null"), nil
	}
	return []byte(`"` + src.{$structProp}.String() + `"`), nil
}

func (dst *{$type->name}) UnmarshalJSON(b []byte) error {
	var n *string
	err := json.Unmarshal(b, &n)
	if err != nil {
		return err
	}

	if n == nil {
		*dst = {$type->name}{}
	} else {
	    u, err := int128.FromString(*n)
	    if err != nil {
	        return err
	    }
	    
		*dst = {$type->name}{{$structProp}: u, Valid: true}
	}

	return nil
}
GO;

    $buf .= "\n\n";
    $buf .= genCodec($typeConfig) . "\n";

    return $buf;
}

function genZeroNullType(TypeConfig $typeConfig): string
{
    $type = $typeConfig->type;
    $structProp = $type->getStructPropName();

    $wideFmtType = $type->isUnsigned ? 'uint64' : 'int64';

    $maxCheck = '';
    if ($type->bitSize < 64) {
        $maxCheck = <<<GO
	if n > {$type->maxValConst} {
		return fmt.Errorf("%d is greater than maximum value for {$type->name}", n)
	}
GO;
    }

    $scanUint64MaxCheck = '';
    if (UINT64->canOverflow($type)) {
        $scanUint64MaxCheck = <<<GO
	if n > {$type->maxValConst} {
		return fmt.Errorf("%d is greater than maximum value for {$type->name}", n)
	}

GO;

    }

    return <<<GO
type {$type->name} {$type->getFullGoTypeName()}

func ($type->name) SkipUnderlyingTypePlan() {}

// ScanInt64 implements the Int64Scanner interface.
func (dst *$type->name) ScanInt64(n int64, valid bool) error {
	if !valid {
		*dst = 0
		return nil
	}

	if n < {$type->minValConst} {
		return fmt.Errorf("%d is less than minimum value for {$type->name}", n)
	}
$maxCheck
	*dst = $type->name(n)

	return nil
}

// ScanUint64 implements the Uint64Scanner interface.
func (dst *{$type->name}) ScanUint64(n uint64, valid bool) error {
	if !valid {
		*dst = 0
		return nil
	}

	$scanUint64MaxCheck
	*dst = $type->name(n)

	return nil
}

// Scan implements the database/sql Scanner interface.
func (dst *$type->name) Scan(src any) error {
	if src == nil {
		*dst = 0
		return nil
	}

	var nullable types.{$type->name}
	err := nullable.Scan(src)
	if err != nil {
		return err
	}

	*dst = $type->name(nullable.$structProp)

	return nil
}

// Value implements the database/sql/driver Valuer interface.
func (src $type->name) Value() (driver.Value, error) {
	if src == 0 {
		return nil, nil
	}
	return $wideFmtType(src), nil
}
GO;
}

function genZeroNullUint128Type(TypeConfig $typeConfig): string
{
    if ($typeConfig->type !== UINT128) {
        throw new InvalidArgumentException("Only uint128 is supported");
    }

    $type = $typeConfig->type;
    $structProp = $type->getStructPropName();

        return <<<GO
type {$type->name} {$type->getFullGoTypeName()}

func ($type->name) SkipUnderlyingTypePlan() {}

// ScanInt64 implements the Int64Scanner interface.
func (dst *$type->name) ScanInt64(n int64, valid bool) error {
	if !valid {
		*dst = $type->name{}
		return nil
	}

	if n < 0 {
		return fmt.Errorf("%d is less than minimum value for {$type->name}", n)
	}

	*dst = $type->name(uint128.From64(uint64(n)))

	return nil
}

// ScanUint64 implements the Uint64Scanner interface.
func (dst *{$type->name}) ScanUint64(n uint64, valid bool) error {
	if !valid {
		*dst = $type->name{}
		return nil
	}

	*dst = $type->name(uint128.From64(n))

	return nil
}

// Scan implements the database/sql Scanner interface.
func (dst *$type->name) Scan(src any) error {
	if src == nil {
		*dst = $type->name{}
		return nil
	}

	var nullable types.{$type->name}
	err := nullable.Scan(src)
	if err != nil {
		return err
	}

	*dst = $type->name(nullable.$structProp)

	return nil
}

// Value implements the database/sql/driver Valuer interface.
func (src $type->name) Value() (driver.Value, error) {
	if {$type->getFullGoTypeName()}(src).IsZero() {
		return nil, nil
	}
	return {$type->getFullGoTypeName()}(src), nil
}
GO;
}

function genZeroNullInt128Type(TypeConfig $typeConfig): string
{
    if ($typeConfig->type !== INT128) {
        throw new InvalidArgumentException("Only uint128 is supported");
    }

    $type = $typeConfig->type;
    $structProp = $type->getStructPropName();

        return <<<GO
type {$type->name} {$type->getFullGoTypeName()}

func ($type->name) SkipUnderlyingTypePlan() {}

// ScanInt64 implements the Int64Scanner interface.
func (dst *$type->name) ScanInt64(n int64, valid bool) error {
	if !valid {
		*dst = $type->name{}
		return nil
	}

	*dst = $type->name(num.I128From64(n))

	return nil
}

// ScanUint64 implements the Uint64Scanner interface.
func (dst *{$type->name}) ScanUint64(n uint64, valid bool) error {
	if !valid {
		*dst = $type->name{}
		return nil
	}

	*dst = $type->name(num.I128FromU64(n))

	return nil
}

// Scan implements the database/sql Scanner interface.
func (dst *$type->name) Scan(src any) error {
	if src == nil {
		*dst = $type->name{}
		return nil
	}

	var nullable types.{$type->name}
	err := nullable.Scan(src)
	if err != nil {
		return err
	}

	*dst = $type->name(nullable.$structProp)

	return nil
}

// Value implements the database/sql/driver Valuer interface.
func (src $type->name) Value() (driver.Value, error) {
	if {$type->getFullGoTypeName()}(src).IsZero() {
		return nil, nil
	}
	return {$type->getFullGoTypeName()}(src), nil
}
GO;
}

function genTypeInt64ValueFunc(Type $type): string
{
    $structProp = $type->getStructPropName();

    $funcDef = "func (n {$type->name}) Int64Value() (Int8, error) {\n";

    $canOverflow = $type->canOverflow(INT64);
    $canUnderflow = $type->canUnderflow(INT64);

    if ($canOverflow) {
        $overflowCheck = match ($type) {
            UINT128 => "if n.{$structProp}.Cmp64(uint64(math.MaxInt64)) > 0",
            INT128 => "if n.{$structProp}.Cmp64(int64(math.MaxInt64)) > 0",

            default => "if n.{$structProp} > math.MaxInt64"
        };

        $funcDef .= <<<GO
	$overflowCheck {
	    return Int8{}, fmt.Errorf("$type->name value is greater than max Int8 value")
	}

GO;
    }

    if ($canUnderflow) {
        $underflowCheck = match ($type) {
            INT128 => "if n.{$structProp}.Cmp64(int64(math.MaxInt64)) < 0",
        };

        $funcDef .= <<<GO
	$underflowCheck {
	    return Int8{}, fmt.Errorf("$type->name value is less than min Int8 value")
	}

GO;
    }

    $funcDef .= "	" . match ($type) {
            UINT128 => "return Int8{Int64: int64(n.{$structProp}.Lo), Valid: n.Valid}, nil",
            INT128 => "return Int8{Int64: n.{$structProp}.AsInt64(), Valid: n.Valid}, nil",

            default => "return Int8{Int64: int64(n.{$structProp}), Valid: n.Valid}, nil",
        } . "\n";
    $funcDef .= "}\n";

    return $funcDef;
}

function genTypeUInt64ValueFunc(Type $type): string
{
    $structProp = $type->getStructPropName();

    $canOverflow = $type->canOverflow(UINT64);
    $canUnderflow = $type->canUnderflow(UINT64);

    $funcDef = "func (n {$type->name}) Uint64Value() (UInt8, error) {\n";

    if ($canUnderflow) {
        // Actually underflow is only possible for signed integer
        if ($type->isUnsigned) {
            throw new InvalidArgumentException("Wrong underflow detection for unsigned type {$type->name}");
        }

        $cmpCond = match ($type) {
            INT128 => "if n.{$structProp}.Cmp64(0) < 0",
            default => "if n.{$structProp} < 0",
        };

        $funcDef .= <<<GO
	$cmpCond {
	    return UInt8{}, fmt.Errorf("$type->name value is less than min UInt8 value")
	}

GO;
    }

    if ($canOverflow) {
        // Actually only 128 bit types can overflow uint64
        $cmpCond = match ($type) {
            UINT128 => "if n.{$structProp}.Cmp64(math.MaxUint64) > 0",
            INT128 => "if n.{$structProp}.Cmp(u64MaxInS128) > 0",
        };

        $funcDef .= <<<GO
	$cmpCond {
	    return UInt8{}, fmt.Errorf("$type->name value is greater than max UInt8 value")
	}

GO;
    }

    $retVal = match ($type) {
        UINT128 => "return UInt8{Uint64: n.{$structProp}.Lo, Valid: n.Valid}, nil",
        INT128 => "return UInt8{Uint64: n.{$structProp}.AsUint64(), Valid: n.Valid}, nil",

        default => "return UInt8{Uint64: uint64(n.{$structProp}), Valid: n.Valid}, nil"
    };

    $funcDef .= "    $retVal\n";
    $funcDef .= "}\n";

    return $funcDef;
}

function genCodecBinaryEncode(Type $type): string
{
    $encodeName = $type->getBinaryEncodeCodecName();

    return <<<GO
type $encodeName struct{}

func ($encodeName) Encode(value any, buf []byte) (newBuf []byte, err error) {
	n := value.({$type->getFullGoTypeName()})
	return {$type->getPGIoWriteFuncName()}(buf, n), nil
}
GO;
}

function getCodecBinaryEncodeToFuncName(Type $type, Type $to): string
{
    $goType = ucfirst($to->goType);

    return "encodePlan{$type->name}CodecBinary{$goType}";
}

function genCodecEncodeToInt64Valuer(Type $type, bool $isBinary): string
{
    $encodeName = ($isBinary ? $type->getBinaryEncodeCodecName() : $type->getTextEncodeCodecName()) . "Int64Valuer";

    $minCheck = '';
    $maxCheck = '';
    $canUnderflow = INT64->canUnderflow($type);
    $canOverflow = INT64->canOverflow($type);

    if ($canUnderflow) {
        $minCheck = <<<GO
    if n.Int64 < 0 {
        return nil, fmt.Errorf("%d is less than minimum value for $type->pgName", n.Int64)
    }
GO;
    }

    if ($canOverflow) {
        $maxCheck = <<<GO
	if n.Int64 > $type->maxValConst {
		return nil, fmt.Errorf("%d is greater than maximum value for $type->pgName", n.Int64)
	}
GO;
    }

    $checks = implode("\n", array_filter([$minCheck, $maxCheck]));

    if ($isBinary) {
        $retVal = match ($type) {
            INT128 => "return {$type->getPGIoWriteFuncName()}(buf, num.I128From64(n.Int64)), nil",
            UINT128 => "return {$type->getPGIoWriteFuncName()}(buf, uint128.From64(uint64(n.Int64))), nil",
            default => "return {$type->getPGIoWriteFuncName()}(buf, $type->goType(n.Int64)), nil",
        };
    } else {
        $fmtUint64Func = UINT64->getFormatIntFunctionCall("uint64(n.Int64)");
        $fmtInt64Func = INT64->getFormatIntFunctionCall("n.Int64");

        $retVal = match ($type) {
            INT128 => "return append(buf, $fmtInt64Func...), nil",
            UINT128 => "return append(buf, $fmtUint64Func...), nil",
            INT8 => "return append(buf, {$type->getFormatIntFunctionCall("n.Int64")}...), nil",
            default => "return append(buf, {$type->getFormatIntFunctionCall("uint64(n.Int64)")}...), nil"
        };
    }

    $go = <<<GO
type $encodeName struct{}

func ($encodeName) Encode(value any, buf []byte) (newBuf []byte, err error) {
	n, err := value.(Int64Valuer).Int64Value()
	if err != nil {
		return nil, err
	}

	if !n.Valid {
		return nil, nil
	}

$checks

	$retVal
}
GO;


    return $go;
}

function genCodecEncodeToUInt64Valuer(Type $type, bool $isBinary): string
{
//    if (!$type->isUnsigned) {
//        throw new InvalidArgumentException("Only unsigned types supported");
//    }

    $encodeName = ($isBinary ? $type->getBinaryEncodeCodecName() : $type->getTextEncodeCodecName()) . "Uint64Valuer";

    $minCheck = '';
    $maxCheck = '';
    $canUnderflow = UINT64->canUnderflow($type);
    $canOverflow = UINT64->canOverflow($type);

    if ($canUnderflow) {
        $minCheck = <<<GO
    if n.Uint64 < 0 {
        return nil, fmt.Errorf("%d is less than minimum value for $type->pgName", n.Uint64)
    }
GO;
    }

    if ($canOverflow) {
        $maxCheck = <<<GO
	if n.Uint64 > $type->maxValConst {
		return nil, fmt.Errorf("%d is greater than maximum value for $type->pgName", n.Uint64)
	}
GO;
    }

    $checks = implode("\n", array_filter([$minCheck, $maxCheck]));

    if ($isBinary) {
        $retVal = match ($type) {
            UINT128 => "return {$type->getPGIoWriteFuncName()}(buf, uint128.From64(n.Uint64)), nil",
            INT128 => "return {$type->getPGIoWriteFuncName()}(buf, num.I128FromU64(n.Uint64)), nil",

            default => "return {$type->getPGIoWriteFuncName()}(buf, $type->goType(n.Uint64)), nil",
        };
    } else {
        $uint64FmtFunc = UINT64->getFormatIntFunctionCall("n.Uint64");
        $int64FmtFunc = INT64->getFormatIntFunctionCall("int64(n.Uint64)");

        $retVal = match ($type) {
            UINT128 => "return append(buf, $uint64FmtFunc...), nil",
            INT128 => "return append(buf, $int64FmtFunc...), nil",
            INT8 => "return append(buf, {$type->getFormatIntFunctionCall("int64(n.Uint64)")}...), nil",

            default => "return append(buf, {$type->getFormatIntFunctionCall("uint64(n.Uint64)")}...), nil",
        };
    }

    $go = <<<GO
type $encodeName struct{}

func ($encodeName) Encode(value any, buf []byte) (newBuf []byte, err error) {
	n, err := value.(Uint64Valuer).Uint64Value()
	if err != nil {
		return nil, err
	}

	if !n.Valid {
		return nil, nil
	}

$checks

	$retVal
}
GO;


    return $go;
}

function genCodecTextEncode(Type $type): string
{
    $encodeName = $type->getTextEncodeCodecName();

    if ($type->isUnsigned) {
        $fmtFunc = $type->getFormatIntFunctionCall("uint64(n)");
    } else {
        $fmtFunc = $type->getFormatIntFunctionCall("int64(n)");
    }

    if ($type === UINT128 || $type === INT128) {
        $fmtFunc = $type->getFormatIntFunctionCall("n");
    }

    return <<<GO
type $encodeName struct{}

func ($encodeName) Encode(value any, buf []byte) (newBuf []byte, err error) {
	n := value.({$type->getFullGoTypeName()})
	return append(buf, $fmtFunc...), nil
}
GO;
}

function getCodecBinaryScanToName(Type $type, Type $to): string
{
    $t = ucfirst($to->goType);

    return "scanPlanBinary{$type->name}To{$t}";
}

function getCodecTextScanToName(Type $type, Type $to): string
{
    $t = ucfirst($to->goType);

    return "scanPlanText{$type->name}To{$t}";
}

function genCodecBinaryScanTo(Type $type, Type $to): string
{
    $scanName = getCodecBinaryScanToName($type, $to);

    $minCheck = '';
    $maxCheck = '';
    $canUnderflow = $type->canUnderflow($to);
    $canOverflow = $type->canOverflow($to);

    if ($canUnderflow) {
        if ($type === INT128) {
            if ($to->isUnsigned) {
                $minCheck = <<<GO
    if n.Cmp64(0) < 0 {
        return fmt.Errorf("%s is less than minimum value for $to->goType", n.String())
    }
GO;
            } else {
                $minCheck = <<<GO
    if n.Cmp64(int64($to->minValConst)) < 0 {
        return fmt.Errorf("%s is less than minimum value for $to->goType", n.String())
    }
GO;
            }
        } else {
            $minCheck = <<<GO
    if n < $to->minValConst {
        return fmt.Errorf("%d is less than minimum value for $to->goType", n)
    }
GO;
        }
    }

    if ($canOverflow) {
        $maxCond = "if n > $type->goType($to->maxValConst)";

        // Special case for variable bit int overflow
        if ($to === UINT) {
            if ($type->isUnsigned) {
                if ($type->bitSize > 32) {
                    $maxCond = "if intSize == 32 && n > $type->goType($to->maxValConst)";
                }
            }
        }
        if ($to === INT) {
            if ($type->isUnsigned) {
                if ($type->bitSize == 32) {
                    $maxCond = "var maxNum $to->goType = $to->maxValConst\n";
                    $maxCond .= "if intSize == 32 && n > $type->goType(maxNum)";
                } else {
                    if ($type->bitSize == 64) {
                        $maxCond = "var maxNum $to->goType = $to->maxValConst\n";
                        $maxCond .= "if n > $type->goType(maxNum)";
                    }
                }
            }
        }

        // Special case for uint128 type
        if ($type === UINT128) {
            $maxCheck = <<<GO
    if n.Cmp({$type->getMaxValConstForType($to)}) > 0 {
        return fmt.Errorf("%s is greater than maximum value for $to->goType", n.String())
    }
GO;
        } elseif ($type === INT128) {
            $maxCheck = <<<GO
    if n.Cmp({$type->getMaxValConstForType($to)}) > 0 {
        return fmt.Errorf("%s is greater than maximum value for $to->goType", n.String())
    }
GO;
        } else {
            $maxCheck = <<<GO
	$maxCond {
		return fmt.Errorf("%d is greater than maximum value for $to->goType", n)
	}
GO;
        }
    }

    $checks = implode("\n", array_filter([$minCheck, $maxCheck]));

    $constructPrefix = '';
    if ($type === INT128 && $to === UINT128) {
        $constructPrefix = "hi, lo := n.Raw()\n";
    }

    $construct = match (true) {
        // Special case INT128 => UINT128
        $type === INT128 && $to === UINT128 => "uint128.New(lo, hi)",

        // Special case UINT128 => INT128
        $type === UINT128 && $to === INT128 => "num.I128FromRaw(n.Hi, n.Lo)",

        // UINT128 to unknown type
        $type === UINT128 && $to !== UINT128 => "$to->goType(n.Lo)",
        // Unknown type to UINT128
        $type !== UINT128 && $to === UINT128 => "uint128.From64(uint64(n))",

        $type === INT128 && $to !== INT128 => $to->isUnsigned ? "$to->goType(n.AsUint64())" : "$to->goType(n.AsInt64())",
        $type !== INT128 && $to === INT128 => $type->isUnsigned ? "num.I128FromU64(uint64(n))" : "num.I128From64(int64(n))",

        $type === $to => "n",

        default => "{$to->getFullGoTypeName()}(n)"
    };

    return <<<GO
type $scanName struct{}

func ($scanName) Scan(src []byte, dst any) error {
	if src == nil {
		return fmt.Errorf("cannot scan NULL into %T", dst)
	}

	if len(src) != $type->byteSize {
		return fmt.Errorf("invalid length for $type->name: %v", len(src))
	}

	p, ok := (dst).(*{$to->getFullGoTypeName()})
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := {$type->getPGIoReadFuncName()}(src)
$checks
	{$constructPrefix}*p = $construct

	return nil
}
GO;
}

function genCodecBinaryScanToTextScanner(Type $type): string
{
    $scanName = getCodecBinaryScanToName($type, TEXT_SCANNER);

    $wideType = $type->isUnsigned ? 'uint64' : 'int64';

    $to = TEXT_SCANNER;

    $ending = <<<GO
	n := $wideType({$type->getPGIoReadFuncName()}(src))

	return s.ScanText(Text{String: {$type->getFormatIntFunctionCall("$wideType(n)")}, Valid: true})
GO;

    if ($type === UINT128 || $type === INT128) {
        $ending = <<<GO
	n := {$type->getPGIoReadFuncName()}(src)

	return s.ScanText(Text{String: n.String(), Valid: true})
GO;
    }

    return <<<GO
type $scanName struct{}

func ($scanName) Scan(src []byte, dst any) error {
	s, ok := (dst).($to->goType)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	if src == nil {
		return s.ScanText(Text{})
	}

	if len(src) != $type->byteSize {
		return fmt.Errorf("invalid length for $type->pgName: %v", len(src))
	}

$ending
}
GO;
}

function genCodecBinaryScanToInt64Scanner(Type $type): string
{
    $scanName = getCodecBinaryScanToName($type, INT64_SCANNER);

    $wideType = $type->isUnsigned ? 'uint64' : 'int64';
    $maxConst = INT64->maxValConst;

    $to = INT64_SCANNER;

    $overflowCheck = <<<GO
	if n > $maxConst {
	    return fmt.Errorf("$type->name value %d is greater than max value for Int8", n)
	}
GO;

    if (!$type->canOverflow(INT64)) {
        $overflowCheck = '';
    }


    $ending = <<<GO
	n := $wideType({$type->getPGIoReadFuncName()}(src))
{$overflowCheck}

	return s.ScanInt64(Int8{Int64: int64(n), Valid: true})
GO;

    if ($type === UINT128) {
        $ending = <<<GO
	n := {$type->getPGIoReadFuncName()}(src)
	if n.Cmp64(uint64($maxConst)) > 0 {
	    return fmt.Errorf("$type->name value %s is greater than max value for Int8", n.String())
	}

	return s.ScanInt64(Int8{Int64: int64(n.Lo), Valid: true})
GO;
    }

    if ($type === INT128) {
        $ending = <<<GO
	n := {$type->getPGIoReadFuncName()}(src)
	if n.Cmp64(math.MaxInt64) > 0 {
	    return fmt.Errorf("$type->name value %s is greater than max value for Int8", n.String())
	}

	return s.ScanInt64(Int8{Int64: n.AsInt64(), Valid: true})
GO;
    }

    return <<<GO
type $scanName struct{}

func ($scanName) Scan(src []byte, dst any) error {
	s, ok := (dst).($to->goType)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	if src == nil {
		return s.ScanInt64(Int8{})
	}

	if len(src) != $type->byteSize {
		return fmt.Errorf("invalid length for {$type->pgName}: %v", len(src))
	}

$ending
}
GO;
}

function genCodecBinaryScanToUInt64Scanner(Type $type): string
{
    $scanName = getCodecBinaryScanToName($type, UINT64_SCANNER);

    $wideType = $type->isUnsigned ? 'uint64' : 'int64';

    $to = UINT64_SCANNER;

    $ending = <<<GO
	n := $wideType({$type->getPGIoReadFuncName()}(src))

	return s.ScanUint64(UInt8{Uint64: n, Valid: true})
GO;

    if ($type === INT8) {
        $ending = <<<GO
	n := $wideType({$type->getPGIoReadFuncName()}(src))

	return s.ScanUint64(UInt8{Uint64: uint64(n), Valid: true})
GO;
    }

    if ($type === UINT128) {
        $ending = <<<GO
	n := {$type->getPGIoReadFuncName()}(src)
	if n.Cmp64(math.MaxUint64) > 0 {
	    return fmt.Errorf("$type->name value %s is greater than max value for UInt8", n.String())
	}

	return s.ScanUint64(UInt8{Uint64: n.Lo, Valid: true})
GO;
    }

    if ($type === INT128) {
        $ending = <<<GO
	n := {$type->getPGIoReadFuncName()}(src)
	if n.Cmp(u64MaxInS128) > 0 {
	    return fmt.Errorf("$type->name value %s is greater than max value for UInt8", n.String())
	}

	return s.ScanUint64(UInt8{Uint64: n.AsUint64(), Valid: true})
GO;
    }


    return <<<GO
type $scanName struct{}

func ($scanName) Scan(src []byte, dst any) error {
	s, ok := (dst).($to->goType)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	if src == nil {
		return s.ScanUint64(UInt8{})
	}

	if len(src) != $type->byteSize {
		return fmt.Errorf("invalid length for {$type->pgName}: %v", len(src))
	}

$ending
}
GO;
}


function genCodecTextScanTo(Type $type, Type $to): string
{
    $scanName = getCodecTextScanToName($type, $to);

    return <<<GO
type $scanName struct{}

func ($scanName) Scan(src []byte, dst any) error {
	if src == nil {
		return fmt.Errorf("cannot scan NULL into %T", dst)
	}

	p, ok := (dst).(*{$to->getFullGoTypeName()})
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n, err := {$to->getParseIntFunctionCall("string(src)")}
	if err != nil {
		return err
	}

	*p = {$to->getFullGoTypeName()}(n)
	return nil
}
GO;
}

function genCodec(TypeConfig $typeConfig): string
{
    $type = $typeConfig->type;
    $codecName = "{$type->name}Codec";

    $buf = <<<GO
type {$codecName} struct{}

func ({$codecName}) FormatSupported(format int16) bool {
	return format == TextFormatCode || format == BinaryFormatCode
}

func ({$codecName}) PreferredFormat() int16 {
	return BinaryFormatCode
}

func ({$codecName}) PlanEncode(m *Map, oid uint32, format int16, value any) EncodePlan {
	switch format {
	case BinaryFormatCode:
		switch value.(type) {
		case {$type->getFullGoTypeName()}:
			return {$type->getBinaryEncodeCodecName()}{}
		case Uint64Valuer:
			return {$type->getBinaryEncodeCodecName()}Uint64Valuer{}
		case Int64Valuer:
			return {$type->getBinaryEncodeCodecName()}Int64Valuer{}
		}
	case TextFormatCode:
		switch value.(type) {
		case {$type->getFullGoTypeName()}:
			return {$type->getTextEncodeCodecName()}{}
		case Uint64Valuer:
			return {$type->getTextEncodeCodecName()}Uint64Valuer{}
		case Int64Valuer:
			return {$type->getTextEncodeCodecName()}Int64Valuer{}
		}
	}

	return nil
}
GO;

    $buf .= "\n";

    $buf .= genCodecBinaryEncode($type) . "\n";
    $buf .= genCodecTextEncode($type) . "\n";
    $buf .= "\n";

    $buf .= genCodecEncodeToInt64Valuer($type, isBinary: true) . "\n";
    $buf .= genCodecEncodeToInt64Valuer($type, isBinary: false) . "\n";
    $buf .= "\n";

    $buf .= genCodecEncodeToUInt64Valuer($type, isBinary: true) . "\n";
    $buf .= genCodecEncodeToUInt64Valuer($type, isBinary: false) . "\n";
    $buf .= "\n";

    $binScanParts = [];
    $txtScanParts = [];

    foreach ($typeConfig->scanTypes as $scanType) {
        $binScanName = getCodecBinaryScanToName($type, $scanType);
        $txtScanName = getCodecTextScanToName($type, $scanType);

        $binScanParts[] = <<<GO
		case *{$scanType->getFullGoTypeName()}:
			return $binScanName{}
GO;

        $txtScanParts[] = <<<GO
		case *{$scanType->getFullGoTypeName()}:
			return $txtScanName{}
GO;
    }

    $binScansSwitchBody = implode("\n", $binScanParts);
    $txtScansSwitchBody = implode("\n", $txtScanParts);

    $binToTextScannerName = getCodecBinaryScanToName($type, TEXT_SCANNER);
    $binToInt64ScannerName = getCodecBinaryScanToName($type, INT64_SCANNER);
    $binToUInt64ScannerName = getCodecBinaryScanToName($type, UINT64_SCANNER);

    $buf .= <<<GO
func ({$codecName}) PlanScan(m *Map, oid uint32, format int16, target any) ScanPlan {
	switch format {
	case BinaryFormatCode:
		switch target.(type) {
$binScansSwitchBody
		case Int64Scanner:
			return $binToInt64ScannerName{}
		case Uint64Scanner:
			return $binToUInt64ScannerName{}
		case TextScanner:
			return $binToTextScannerName{}
		}
	case TextFormatCode:
		switch target.(type) {
$txtScansSwitchBody
		case Int64Scanner:
		    return scanPlanTextAnyToInt64Scanner{}
		case Uint64Scanner:
		    return scanPlanTextAnyToUint64Scanner{}
		}
	}

	return nil
}

func (c {$codecName}) DecodeDatabaseSQLValue(m *Map, oid uint32, format int16, src []byte) (driver.Value, error) {
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

func (c {$codecName}) DecodeValue(m *Map, oid uint32, format int16, src []byte) (any, error) {
	if src == nil {
		return nil, nil
	}

	var n {$type->getFullGoTypeName()}
	err := codecScan(c, m, oid, format, src, &n)
	if err != nil {
		return nil, err
	}
	return n, nil
}
GO;

    $buf .= "\n\n";

    foreach ($typeConfig->scanTypes as $scanType) {
        $buf .= genCodecBinaryScanTo($type, $scanType) . "\n\n";
    }

    $buf .= genCodecBinaryScanToTextScanner($type) . "\n\n";
    $buf .= genCodecBinaryScanToInt64Scanner($type) . "\n\n";
    $buf .= genCodecBinaryScanToUInt64Scanner($type) . "\n\n";

    $buf .= "\n\n";

    foreach ($typeConfig->scanTypes as $scanType) {
        $buf .= genCodecTextScanTo($type, $scanType) . "\n\n";
    }

    $buf .= "\n\n";

    return $buf;
}

function genTypeScanTest(TypeConfig $typeConfig, Type $scanType, bool $isBinary): string
{
    $type = $typeConfig->type;
    $test = "";

    // Source type could be greater than scan target
    $canOverflow = $type->canOverflow($scanType);
    $canUnderflow = $type->canUnderflow($scanType);

    if ($type->isUnsigned && $canOverflow && $canUnderflow) {
        throw new InvalidArgumentException(
            "[bug] Both overflow and underflow detected for {$type->name} and {$scanType->name}"
        );
    }

    $scanTypName = $scanType->getStructPropName();

    $maxValue = getMaxValNoOverflowForType($type, $scanType);

    if ($isBinary) {
        $maxBin = getMaxValBytesNoOverflowForType($type, $scanType);
    } else {
        $maxBin = $maxValue;
    }

    if ($scanType === UINT128) {
        if ($type === UINT128) {
            $maxValue = "u128Max";
        } else {
            $prefix = $type->isUnsigned ? "u" : "s";
            $maxValue = "{$prefix}{$type->bitSize}MaxInU128";
        }
    } else {
        if ($scanType === INT128) {
            if ($type === INT128 || $type === UINT128) {
                $maxValue = "s128Max";
            } else {
                $prefix = $type->isUnsigned ? "u" : "s";
                $maxValue = "{$prefix}{$type->bitSize}MaxInS128";
            }
        } else {
            $maxValue = "$scanType->goType($maxValue)";
        }
    }

    $testFnPrefix = $isBinary ? 'Binary' : 'Text';
    $format = $isBinary ? "pgtype.BinaryFormatCode" : "pgtype.TextFormatCode";

    $fn = <<<GO
func Test{$type->name}{$testFnPrefix}_Scan$scanTypName(t *testing.T) {
    var dst {$scanType->getFullGoTypeName()}

    assert.NoError(t,
		typeMap.Scan({$type->name}OID, $format, []byte("$maxBin"), &dst),
	)

	assert.Equal(t, $maxValue, dst)
GO;

    $test .= $fn;
    $test .= "\n}\n";

    // Check self underflow
    if (!$type->isUnsigned && $type === $scanType && !$isBinary) {
        $errMsg = "{$scanType->getParseIntFunctionName()}: parsing \"{$type->getMinUnderflowVal()}\": value out of range";

        $minBin = $type->getMinUnderflowVal();

        $test .= <<<GO
func Test{$type->name}{$testFnPrefix}_Scan{$scanTypName}_SelfUnderflow(t *testing.T) {
	var dst {$scanType->getFullGoTypeName()}

	assert.ErrorContains(t,
		typeMap.Scan({$type->name}OID, $format, []byte("$minBin"), &dst),
		`$errMsg`,
	)
}
GO;

        $test .= "\n";
    }

    // Check self overflow
    if ($type === $scanType && !$isBinary) {
        $errMsg = "{$scanType->getParseIntFunctionName()}: parsing \"{$type->getMaxValOverflow()}\": value out of range";
        if ($type === UINT128) {
            $errMsg = "value overflows Uint128";
        }

        $maxBin = $type->getMaxValOverflow();

        $test .= <<<GO
func Test{$type->name}{$testFnPrefix}_Scan{$scanTypName}_SelfOverflow(t *testing.T) {
	var dst {$scanType->getFullGoTypeName()}

	assert.ErrorContains(t,
		typeMap.Scan({$type->name}OID, $format, []byte("$maxBin"), &dst),
		`$errMsg`,
	)
}
GO;

        $test .= "\n";
    }

    // Test overflow
    if ($canOverflow) {
        $errMsg = $isBinary
            ? "$type->maxVal is greater than maximum value for $scanType->goType"
            : "{$scanType->getParseIntFunctionName()}: parsing \"{$type->getMaxVal()}\": value out of range";

        $maxBin = $isBinary ? $type->getMaxValBytes() : $type->getMaxVal();


        if ($typeConfig->type->isUnsigned && $typeConfig->type->bitSize === 32 && $scanType === INT) {
                $test .= <<<GO
func Test{$type->name}{$testFnPrefix}_Scan{$scanTypName}_Overflow(t *testing.T) {
	var dst {$scanType->getFullGoTypeName()}

	if intSize == 32 {
		assert.ErrorContains(t,
			typeMap.Scan({$type->name}OID, $format, []byte("$maxBin"), &dst),
			`$errMsg`,
		)
	}
}
GO;
        } else if ($typeConfig->type->isUnsigned && $typeConfig->type->bitSize === 64 && $scanType === UINT) {
            $test .= <<<GO
func Test{$type->name}{$testFnPrefix}_Scan{$scanTypName}_Overflow(t *testing.T) {
	var dst {$scanType->getFullGoTypeName()}

	if intSize == 32 {
		assert.ErrorContains(t,
			typeMap.Scan({$type->name}OID, $format, []byte("$maxBin"), &dst),
			`$errMsg`,
		)
	}
}
GO;
        } else {
            $test .= <<<GO
func Test{$type->name}{$testFnPrefix}_Scan{$scanTypName}_Overflow(t *testing.T) {
	var dst {$scanType->getFullGoTypeName()}

	assert.ErrorContains(t,
		typeMap.Scan({$type->name}OID, $format, []byte("$maxBin"), &dst),
		`$errMsg`,
	)
}
GO;
        }

        $test .= "\n";
    }

    // Test underflow
    if ($canUnderflow) {
        if ($scanType->isUnsigned) {
            $errMsg = $isBinary
                ? "{$type->getMinVal()} is less than minimum value for $scanType->goType"
                : "{$scanType->getParseIntFunctionName()}: parsing \"{$type->getMinVal()}\": invalid syntax";

            if (!$isBinary && $scanType === UINT128) {
                $errMsg = "value cannot be negative";
            }
        } else {
            $errMsg = $isBinary
                ? "{$type->getMinVal()} is less than minimum value for $scanType->goType"
                : "{$scanType->getParseIntFunctionName()}: parsing \"{$type->getMinVal()}\": value out of range";
        }

        $minBin = $isBinary ? $type->getMinValBytes() : $type->getMinVal();

        $test .= <<<GO
func Test{$type->name}{$testFnPrefix}_Scan{$scanTypName}_Underflow(t *testing.T) {
	var dst {$scanType->getFullGoTypeName()}

	assert.ErrorContains(t,
		typeMap.Scan({$type->name}OID, $format, []byte("$minBin"), &dst),
		`$errMsg`,
	)
}
GO;

        $test .= "\n";
    }

    $test .= "\n";

    return $test;
}

function genTypeTests(TypeConfig $typeConfig): string
{
    $includes = [
        "testing",
        "github.com/stretchr/testify/assert",
        "",
        "github.com/jackc/pgx/v5/pgtype",
        "",
        "lukechampine.com/uint128",
        "",
        "go.shabbyrobe.org/num",
    ];

    $header = <<<GO
// Do not edit. Generated from codegen

package types_test

GO;

    $header .= genImportsSection($includes) . "\n";

    $test = '';

    // Binary scans
    foreach ($typeConfig->scanTypes as $scanType) {
        $test .= genTypeScanTest($typeConfig, $scanType, true);
    }

    // Text scans
    foreach ($typeConfig->scanTypes as $scanType) {
        $test .= genTypeScanTest($typeConfig, $scanType, false);
    }

    return $header . "\n" . $test;
}

function genImportsSection(array $includes): string
{
    if ($includes === []) {
        return '';
    }

    $begin = "import (\n";
    $end = "\n)\n";

    $imports = array_map(static function (string $include, string|int $alias): string {
        if ($include === "") {
            return "\n";
        }

        if (is_string($alias)) {
            return "{$alias} \"{$include}\"";
        }

        return "    \"$include\"";
    }, $includes, array_keys($includes));

    return $begin . implode("\n", $imports) . $end;
}

const TYPE_INCLUDES = [
    "database/sql/driver",
    "encoding/json",
    "fmt",
    "math",
    "strconv",
    "",
    "github.com/pg-uint/pgx-pg-uint128/pgio",
    "",
    "." => "github.com/jackc/pgx/v5/pgtype",
    "",
    "lukechampine.com/uint128",
    "",
    "go.shabbyrobe.org/num",
    "github.com/pg-uint/pgx-pg-uint128/int128",
];

const TYPE_ZERONULL_INCLUDES = [
    "database/sql/driver",
    "fmt",
    "math",
    "",
    "github.com/pg-uint/pgx-pg-uint128/types"
];

$header = <<<GO
// Do not edit. Generated from codegen

package types

GO;

$zeroNullHeader = <<<GO
// Do not edit. Generated from codegen

package zeronull

GO;


@mkdir("types");
@mkdir("types/zeronull");

/** @var TypeConfig $type */
foreach (CONFIGURED_TYPES as $type) {
    $buf = '';

    $typeIncludes = TYPE_INCLUDES;

    if ($type->type === UINT128) {
        $buf .= genTypeU128($type);
    } elseif ($type->type === INT128) {
        $buf .= genTypeS128($type);
    } else {
        $buf .= genType($type);
    }

    $fileHeader = $header;
    $fileHeader .= genImportsSection($typeIncludes);

    $buf = $fileHeader . "\n" . $buf;

    file_put_contents("types/{$type->type->getFilename()}", $buf);

    $buf = '';
    $buf .= genTypeTests($type);

    file_put_contents("types/{$type->type->pgName}_test.go", $buf);
}

/** @var TypeConfig $type */
foreach (CONFIGURED_TYPES as $type) {
    $buf = '';

    $typeIncludes = TYPE_ZERONULL_INCLUDES;

    if ($type->type === UINT128) {
        unset($typeIncludes[array_search("math", $typeIncludes, true)]);
        $typeIncludes[] = "lukechampine.com/uint128";

        $buf .= genZeroNullUint128Type($type);
    } elseif ($type->type === INT128) {
        unset($typeIncludes[array_search("math", $typeIncludes, true)]);
        unset($typeIncludes[array_search("fmt", $typeIncludes, true)]);
        $typeIncludes[] = "go.shabbyrobe.org/num";

        $buf .= genZeroNullInt128Type($type);
    } else {
        if ($type->type === UINT64) {
            unset($typeIncludes[array_search("math", $typeIncludes, true)]);
        }

        $buf .= genZeroNullType($type);
    }

    $fileHeader = $zeroNullHeader;
    $fileHeader .= genImportsSection($typeIncludes);

    $buf = $fileHeader . "\n" . $buf;

    file_put_contents("types/zeronull/{$type->type->getFilename()}", $buf);

//    $buf = '';
//    $buf .= genTypeTests($type);
//
//    file_put_contents("types/{$type->type->pgName}_test.go", $buf);
}

$fmtPath = realpath(__DIR__ . '/../types');
shell_exec("go fmt $fmtPath");

$fmtPath = realpath(__DIR__ . '/../types/zeronull');
shell_exec("go fmt $fmtPath");
