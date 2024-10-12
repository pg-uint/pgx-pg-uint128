<?php

declare(strict_types=1);

require_once __DIR__ . '/Type.php';
require_once __DIR__ . '/TypeConfig.php';

const UINT8 = new Type(
    name: 'UInt1',
    pgName: 'uint1',
    goType: 'uint8',
    bitSize: 8,
    isUnsigned: true,
    maxValConst: 'math.MaxUint8',
    minValConst: '0',
    maxVal: '255',
    oid: 0,

);

const UINT16 = new Type(
    name: 'UInt2',
    pgName: 'uint2',
    goType: 'uint16',
    bitSize: 16,
    isUnsigned: true,
    maxValConst: 'math.MaxUint16',
    minValConst: '0',
    maxVal: '65535',
    oid: 787800,
);

const UINT32 = new Type(
    name: 'UInt4',
    pgName: 'uint4',
    goType: 'uint32',
    bitSize: 32,
    isUnsigned: true,
    maxValConst: 'math.MaxUint32',
    minValConst: '0',
    maxVal: '4294967295',
    oid: 787801,
);

const UINT64 = new Type(
    name: 'UInt8',
    pgName: 'uint8',
    goType: 'uint64',
    bitSize: 64,
    isUnsigned: true,
    maxValConst: 'math.MaxUint64',
    minValConst: '0',
    maxVal: '18446744073709551615',
    oid: 787802,
);

const UINT128 = new Uint128(
    name: 'UInt16',
    pgName: 'uint16',
    goType: 'Uint128',
    bitSize: 128,
    isUnsigned: true,
    maxValConst: '',
    minValConst: '0',
    maxVal: '340282366920938463463374607431768211455',
    oid: 787803,
);

const UINT = new Type(
    name: '',
    pgName: '',
    goType: 'uint',
    bitSize: 0,
    isUnsigned: true,
    maxValConst: 'math.MaxUint',
    minValConst: '0',
    maxVal: '18446744073709551615',
    oid: 0,
);

const INT8 = new Type(
    name: 'Int1',
    pgName: 'int1',
    goType: 'int8',
    bitSize: 8,
    isUnsigned: false,
    maxValConst: 'math.MaxInt8',
    minValConst: 'math.MinInt8',
    maxVal: '127',
    oid: 0,
);

const INT16 = new Type(
    name: 'Int2',
    pgName: 'int2',
    goType: 'int16',
    bitSize: 16,
    isUnsigned: false,
    maxValConst: 'math.MaxInt16',
    minValConst: 'math.MinInt16',
    maxVal: '32767',
    oid: 0,
);

const INT32 = new Type(
    name: 'Int4',
    pgName: 'int4',
    goType: 'int32',
    bitSize: 32,
    isUnsigned: false,
    maxValConst: 'math.MaxInt32',
    minValConst: 'math.MinInt32',
    maxVal: '2147483647',
    oid: 0,
);

const INT64 = new Type(
    name: 'Int8',
    pgName: 'int8',
    goType: 'int64',
    bitSize: 64,
    isUnsigned: false,
    maxValConst: 'math.MaxInt64',
    minValConst: 'math.MinInt64',
    maxVal: '9223372036854775807',
    oid: 0,
);

const INT128 = new Int128(
    name: 'Int16',
    pgName: 'int16',
    goType: 'I128',
    bitSize: 128,
    isUnsigned: false,
    maxValConst: 'Int128Max',
    minValConst: 'Int128Min',
    maxVal: '170141183460469231731687303715884105727',
    oid: 0,
);

const INT = new Type(
    name: '',
    pgName: '',
    goType: 'int',
    bitSize: 0,
    isUnsigned: false,
    maxValConst: 'math.MaxInt',
    minValConst: 'math.MinInt',
    maxVal: '9223372036854775807',
    oid: 0,
);

const TEXT_SCANNER = new Type(
    name: '',
    pgName: '',
    goType: 'TextScanner',
    bitSize: 0,
    isUnsigned: false,
    maxValConst: '',
    minValConst: '',
    maxVal: '',
    oid: 0,
);

const INT64_SCANNER = new Type(
    name: '',
    pgName: '',
    goType: 'Int64Scanner',
    bitSize: 0,
    isUnsigned: false,
    maxValConst: '',
    minValConst: '',
    maxVal: '',
    oid: 0,
);

const UINT64_SCANNER = new Type(
    name: '',
    pgName: '',
    goType: 'Uint64Scanner',
    bitSize: 0,
    isUnsigned: false,
    maxValConst: '',
    minValConst: '',
    maxVal: '',
    oid: 0,
);

const SCAN_TYPES = [
    UINT16,
    UINT32,
    UINT64,
    UINT128,
    UINT,
    INT16,
    INT32,
    INT64,
    INT128,
    INT,
];

/**
 * @var array<TypeConfig>
 */
const CONFIGURED_TYPES = [
    new TypeConfig(type: UINT16, scanTypes: SCAN_TYPES),
    new TypeConfig(type: UINT32, scanTypes: SCAN_TYPES),
    new TypeConfig(type: UINT64, scanTypes: SCAN_TYPES),
    new TypeConfig(type: UINT128, scanTypes: SCAN_TYPES),

    new TypeConfig(type: INT128, scanTypes: SCAN_TYPES)
];
