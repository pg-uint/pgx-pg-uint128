<?php

declare(strict_types=1);

class Type
{
    public readonly int $byteSize;

    public function __construct(
        public string $name,
        public string $pgName,

        public string $goType,

        public int $bitSize,
        public bool $isUnsigned,

        public string $maxValConst,
        public string $minValConst,

        public string $maxVal,

        public int $oid,
    ) {
        $this->byteSize = $bitSize / 8;
    }

    public function getStructPropName(): string
    {
        return ucfirst($this->goType);
    }

    public function getFilename(): string
    {
        return "{$this->pgName}.go";
    }

    public function getBinaryEncodeCodecName(): string
    {
        $goType = ucfirst($this->goType);

        return "encodePlan{$this->name}CodecBinary{$goType}";
    }

    public function getTextEncodeCodecName(): string
    {
        $goType = ucfirst($this->goType);

        return "encodePlan{$this->name}CodecText{$goType}";
    }

    public function getPGIoReadFuncName(): string
    {
        if ($this->isUnsigned) {
            return "pgio.ReadUint{$this->bitSize}";
        }

        throw new InvalidArgumentException("signed types is not supported");
    }

    public function getPGIoWriteFuncName(): string
    {
        if ($this->isUnsigned) {
            return "pgio.AppendUint{$this->bitSize}";
        }

        throw new InvalidArgumentException("signed types is not supported");
    }

    public function getParseIntFunctionName(): string
    {
        return match ($this->isUnsigned) {
            true => "strconv.ParseUint",
            false => "strconv.ParseInt",
        };
    }

    public function getParseIntFunctionCall(string $funcArg): string
    {
        $func = $this->getParseIntFunctionName();

        $bitSize = $this->bitSize > 0 ? (string)($this->bitSize) : "intSize";

        return "$func($funcArg, 10, $bitSize)";
    }

    public function getFormatIntFunctionName(): string
    {
        return match ($this->isUnsigned) {
            true => "strconv.FormatUint",
            false => "strconv.FormatInt",
        };
    }

    public function getFormatIntFunctionCall(string $funcArg): string
    {
        $func = $this->getFormatIntFunctionName();

        return "$func($funcArg, 10)";
    }

    public function getMaxVal(): string
    {
        if ($this->isUnsigned) {
            return match ($this->bitSize) {
                8 => '255',
                16 => '65535',
                32 => '4294967295',
                64 => '18446744073709551615',
                128 => '340282366920938463463374607431768211455',
            };
        }

        return match ($this->bitSize) {
            8 => '127',
            16 => '32767',
            32 => '2147483647',
            64 => '9223372036854775807',
            128 => '170141183460469231731687303715884105727',
        };
    }

    public function getMaxValOverflow(): string
    {
        if ($this->isUnsigned) {
            return match ($this->bitSize) {
                8 => '256',
                16 => '65536',
                32 => '4294967296',
                64 => '18446744073709551616',
                128 => '340282366920938463463374607431768211456',
            };
        }

        return match ($this->bitSize) {
            8 => '128',
            16 => '32768',
            32 => '2147483648',
            64 => '9223372036854775808',
            128 => '170141183460469231731687303715884105728',
        };
    }

    public function getMinVal(): string
    {
        if ($this->isUnsigned) {
            return '0';
        }

        return match ($this->bitSize) {
            8 => '-128',
            16 => '-32768',
            32 => '-2147483648',
            64 => '-9223372036854775808',
            128 => '-170141183460469231731687303715884105728',
        };
    }

    public function getMinValBytes(): string
    {
        if ($this->isUnsigned) {
            return str_repeat("\\x00", $this->byteSize);
        }

        return "\\x80" . str_repeat("\\x00", $this->byteSize - 1);
    }

    public function getMinUnderflowVal(): string
    {
        if ($this->isUnsigned) {
            return '0';
        }

        return match ($this->bitSize) {
            8 => '-129',
            16 => '-32769',
            32 => '-2147483649',
            64 => '-9223372036854775809',
            128 => '-170141183460469231731687303715884105729',
        };
    }

    public function getMinUnderflowValBytes(): string
    {
        if ($this->isUnsigned) {
            throw new InvalidArgumentException("Unsigned integer types cannot underflow");
        }

        return "\\x81" . str_repeat("\\x00", $this->byteSize - 1);
    }

    public function getMaxValBytes(): string
    {
        if ($this->isUnsigned) {
            return str_repeat("\\xFF", $this->byteSize);
        }

        return "\\x7F" . str_repeat("\\xFF", $this->byteSize - 1);
    }

    public function canOverflow(Type $type): bool
    {
        if ($this->isUnsigned && $type === UINT) {
            return $this->canOverflow(UINT64);
        }
        if ($this->isUnsigned && $type === INT) {
            return $this->canOverflow(INT64);
        }

        // Unsigned types can overflow signed types only when precA >= precB
        if ($this->isUnsigned && !$type->isUnsigned) {
            return $this->bitSize >= $type->bitSize;
        }

        // Unsigned types can overflow each other only when precA > precB
        // Signed types can overflow unsigned types only when precA > precB
        // Signed types can overflow each other only when precA > precB
        return $this->bitSize > $type->bitSize;
    }

    public function canUnderflow(Type $type): bool
    {
        // Min val for unsigned types is 0, so no underflow is possible
        if ($this->isUnsigned && $type->isUnsigned) {
            return false;
        }

        // Min val for unsigned types is 0, so it never can underflow signed
        if ($this->isUnsigned && !$type->isUnsigned) {
            return false;
        }

        // Signed type can underflow any unsigned type
        if (!$this->isUnsigned && $type->isUnsigned) {
            return true;
        }

        // Signed type can underflow another signed type only when precA > precB
        return $this->bitSize > $type->bitSize;
    }

    public function canOverflowMachineDependedType(Type $type): bool
    {
        // UINT can be overflowed only when both signed/unsigned bit size is larger than 32 bit
        if ($type === UINT) {
            return $this->bitSize > 32;
        }
        if ($type === INT) {
            // Unsigned can overflow signed even with the same bit size
            if ($this->isUnsigned) {
                return $this->bitSize >= 32;
            }

            return $this->bitSize > 32;
        }

        throw new InvalidArgumentException("Cannot check for not machine depended types");
    }

    public function getGoPackage(): string
    {
        return '';
    }

    public function getFullGoTypeName(): string
    {
        $pkg = $this->getGoPackage();

        if ($pkg === '') {
            return $this->goType;
        }

        return "{$pkg}.{$this->goType}";
    }
}

function getMaxValNoOverflowForType(Type $src, Type $dst): string
{
    // UINT <=> UINT
    if ($src->isUnsigned && $dst->isUnsigned) {
        if ($dst === UINT) {
            return getMaxValNoOverflowForType($src, UINT64);
        }

        if ($src->bitSize > $dst->bitSize) {
            return $dst->maxVal;
        }

        return $src->maxVal;
    }

    // INT <=> INT
    if (!$src->isUnsigned && !$dst->isUnsigned) {
        if ($dst === INT) {
            return getMaxValNoOverflowForType($src, INT64);
        }

        if ($src->bitSize > $dst->bitSize) {
            return $dst->maxVal;
        }

        return $src->maxVal;
    }

    // UINT <=> INT
    if ($src->isUnsigned && !$dst->isUnsigned) {
        if ($dst === INT) {
            return getMaxValNoOverflowForType($src, INT64);
        }

        // uint can overflow signed int even with same bit size
        if ($src->bitSize >= $dst->bitSize) {
            return $dst->maxVal;
        }

        if ($src->bitSize < $dst->bitSize) {
            // Choose corresponding datatype
            $cutType = match ($src->bitSize) {
                16 => UINT16,
                32 => UINT32,
                64 => UINT64,
                128 => UINT128,
            };

            return $cutType->maxVal;
        }

        return $dst->maxVal;
    }

    // INT <=> UINT
    if (!$src->isUnsigned && $dst->isUnsigned) {
        if ($dst === UINT) {
            return getMaxValNoOverflowForType($src, UINT64);
        }

        // Signed int can overflow UINT only when it has larger bit size
        if ($src->bitSize > $dst->bitSize) {
            return $dst->maxVal;
        }

        if ($src->bitSize < $dst->bitSize) {
            // Choose corresponding datatype
            $cutType = match ($src->bitSize) {
                16 => UINT16,
                32 => UINT32,
                64 => UINT64,
                128 => UINT128,
            };

            return $cutType->maxVal;
        }

        return $src->maxVal;
    }

    throw new InvalidArgumentException("Signed to unsigned is not supported");
}

function getMaxValBytesNoOverflowForType(Type $src, Type $dst): string
{
    // UINT <=> UINT
    if ($src->isUnsigned && $dst->isUnsigned) {
        if ($dst === UINT) {
            return getMaxValBytesNoOverflowForType($src, UINT64);
        }

        if ($src->bitSize > $dst->bitSize) {
            return str_repeat("\\x00", $src->byteSize - $dst->byteSize) . $dst->getMaxValBytes();
        }

        return $src->getMaxValBytes();
    }

    // INT <=> INT
    if (!$src->isUnsigned && !$dst->isUnsigned) {
        if ($dst === INT) {
            return getMaxValBytesNoOverflowForType($src, INT64);
        }

        if ($src->bitSize > $dst->bitSize) {
            return str_repeat("\\x00", $src->byteSize - $dst->byteSize) . $dst->getMaxValBytes();
        }

        return $src->getMaxValBytes();
    }

    // UINT <=> INT
    if ($src->isUnsigned && !$dst->isUnsigned) {
        if ($dst === INT) {
            return getMaxValBytesNoOverflowForType($src, INT64);
        }

        // Unsigned can overflow signed even with same bti size
        if ($src->bitSize >= $dst->bitSize) {
            return str_repeat("\\x00", $src->byteSize - $dst->byteSize) . $dst->getMaxValBytes();
        }

        if ($src->bitSize < $dst->bitSize) {
            // Choose corresponding datatype
            $cutType = match ($src->bitSize) {
                16 => UINT16,
                32 => UINT32,
                64 => UINT64,
                128 => UINT128,
            };

            return $cutType->getMaxValBytes();
        }

        return $dst->getMaxValBytes();
    }

    // INT <=> UINT
    if (!$src->isUnsigned && $dst->isUnsigned) {
        if ($dst === UINT) {
            return getMaxValBytesNoOverflowForType($src, UINT64);
        }

        if ($src->bitSize > $dst->bitSize) {
            return str_repeat("\\x00", $src->byteSize - $dst->byteSize) . $dst->getMaxValBytes();
        }

        if ($src->bitSize < $dst->bitSize) {
            // Choose corresponding datatype
            $cutType = match ($src->bitSize) {
                16 => INT16,
                32 => INT32,
                64 => INT64,
                128 => INT128,
            };

            return $cutType->getMaxValBytes();
        }

        return $src->getMaxValBytes();
    }

    throw new InvalidArgumentException("Signed to unsigned is not supported");
}

class Uint128 extends Type
{
    public function getGoPackage(): string
    {
        return "uint128";
    }

    public function getFormatIntFunctionCall(string $funcArg): string
    {
        return "$funcArg.String()";
    }

    public function getParseIntFunctionName(): string
    {
        return match ($this->isUnsigned) {
            true => "uint128.FromString",
            false => throw new InvalidArgumentException("Only unsigned uint128 is supported"),
        };
    }

    public function getParseIntFunctionCall(string $funcArg): string
    {
        $func = $this->getParseIntFunctionName();

        return "$func($funcArg)";
    }

    public function getMaxValConstForType(Type $type): string
    {
        $prefix = $type->isUnsigned ? 'u' : 's';
        $bitSize = $type->bitSize === 0 ? '' : (string)$type->bitSize;

        return "{$prefix}{$bitSize}MaxInU128";
    }
}

class Int128 extends Type
{
    public function getGoPackage(): string
    {
        return "num";
    }

    public function getFormatIntFunctionCall(string $funcArg): string
    {
        return "$funcArg.String()";
    }

    public function getParseIntFunctionName(): string
    {
        return match (!$this->isUnsigned) {
            true => "int128.FromString",
            false => throw new InvalidArgumentException("Only signed int128 is supported"),
        };
    }

    public function getParseIntFunctionCall(string $funcArg): string
    {
        $func = $this->getParseIntFunctionName();

        return "$func($funcArg)";
    }

    public function getPGIoReadFuncName(): string
    {
        return "pgio.ReadInt{$this->bitSize}";
    }

    public function getPGIoWriteFuncName(): string
    {
        return "pgio.AppendInt{$this->bitSize}";
    }

    public function getMaxValConstForType(Type $type): string
    {
        $prefix = $type->isUnsigned ? 'u' : 's';
        $bitSize = $type->bitSize === 0 ? '' : (string)$type->bitSize;

        return "{$prefix}{$bitSize}MaxInS128";
    }
}
