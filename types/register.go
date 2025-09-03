package types

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	. "github.com/jackc/pgx/v5/pgtype"
	"regexp"
	"strconv"
	"strings"
)

// RegisterAll registers both integer types and their associated variants.
//
// This function invokes RegisterTypes to register the main integer types
// (uint2, uint4, uint8, uint16 and int16) and subsequently registers their variants
// (array, range, multirange) in the TypeMap of [pgx.Conn] by calling
// RegisterDefaultPgTypeVariants.
//
// It returns a slice of pointers to the registered Type structs. If an error
// occurs during the registration process, it returns nil along with the error.
//
// The returned types can be used to register them on new connections without
// issuing actual database queries. For example, you can use the following code:
//
//		conn.TypeMap().RegisterTypes(regTypes)
//	 types.RegisterDefaultPgTypeVariants(conn.TypeMap())
//
// Note that the uniqueness of type OIDs is guaranteed only within a single server
// and a single database on that server. Consequently, OIDs may conflict across
// different databases or servers.
func RegisterAll(ctx context.Context, conn *pgx.Conn) ([]*Type, error) {
	types, err := RegisterTypes(ctx, conn)
	if err != nil {
		return nil, err
	}

	RegisterDefaultPgTypeVariants(conn.TypeMap())

	return types, nil
}

// RegisterTypes registers all integer types (uint2, uint4, uint8, uint16 and int16) at once.
//
// It returns a slice of pointers to the registered Type structs corresponding to these types.
//
// The returned types can be used to register them on new connections without
// issuing actual database queries. For example, you can use the following code:
//
//	conn.TypeMap().RegisterTypes(types)
//
// Note that the uniqueness of type OIDs is guaranteed only within a single server
// and a single database on that server. Consequently, OIDs may conflict across
// different databases or servers.
//
// If there is an error while loading the types, it will be returned, along with a nil slice.
func RegisterTypes(ctx context.Context, conn *pgx.Conn) ([]*Type, error) {
	types, err := loadTypes(ctx, conn)
	if err != nil {
		return nil, err
	}

	return types, nil
}

// buildLoadDerivedTypesSQL generates the correct query for retrieving type information.
//
// Actually this function is a clone of PGX function but with support for base types and a bit simplified.
//
//	pgVersion: the major version of the PostgreSQL server
func buildLoadDerivedTypesSQL(pgVersion int64) string {
	supportsMultirange := (pgVersion >= 14)

	const sql = `
WITH types AS (
    SELECT
        typname AS _typname,
        typname::regtype::oid AS _oid
    FROM unnest(ARRAY['uint1', 'uint2', 'uint4', 'uint8', 'uint16', 'int1', 'int16']) AS typname
)
SELECT
    oid,
    typname,
    typtype,
    typnamespace::regnamespace AS typnsp,
    typarray AS typarray_oid,
    (
        SELECT typname
        FROM pg_type pg_type_arr
        WHERE pg_type_arr.oid = pg_type.typarray
    ) AS typ_array,
    rtype.*
FROM pg_type
JOIN types ON pg_type.oid = types._oid
LEFT JOIN LATERAL (
    SELECT
        rngtypid AS range_type_oid,
        pg_rtype.typname AS range_type,
        {mrr_typeid} AS multi_range_type_oid,
        {mrr_typename} AS multi_range_type
    FROM pg_range
    JOIN pg_type pg_rtype ON pg_rtype.oid = pg_range.rngtypid
	{mrr_join}
    WHERE rngsubtype = pg_type.oid
    LIMIT 1
) rtype ON TRUE
`

	q := sql

	if supportsMultirange {
		q = strings.Replace(q, "{mrr_typeid}", "rngmultitypid", 1)
		q = strings.Replace(q, "{mrr_typename}", "pg_mrtype.typname", 1)
		q = strings.Replace(q, "{mrr_join}", "LEFT JOIN pg_type pg_mrtype ON pg_mrtype.oid = pg_range.rngmultitypid", 1)
	} else {
		q = strings.Replace(q, "{mrr_typeid}", "0::oid", 1)
		q = strings.Replace(q, "{mrr_typename}", "''", 1)
		q = strings.Replace(q, "{mrr_join}", "", 1)
	}

	return q
}

// loadTypes performs a single (complex) query, returning all the required
// information to register the named types, as well as any other types directly
// or indirectly required to complete the registration.
// The result of this call can be passed into RegisterTypes to complete the process.
func loadTypes(ctx context.Context, c *pgx.Conn) ([]*Type, error) {
	tMap := c.TypeMap()

	// Disregard server version errors. This will result in
	// the SQL not support recent structures such as multirange
	serverVersion, _ := serverVersion(c)
	sql := buildLoadDerivedTypesSQL(serverVersion)

	rows, err := c.Query(ctx, sql, pgx.QueryExecModeSimpleProtocol)
	if err != nil {
		return nil, fmt.Errorf("While generating load types query: %w", err)
	}
	defer rows.Close()

	result := make([]*Type, 0, 100)

	for rows.Next() {
		var typtype, typnsp string
		var typoid, typarroid, typrangeoid, typmultirangeoid uint32
		var typname, typarrname, typrangename, typmultirangename string

		err = rows.Scan(&typoid, &typname, &typtype, &typnsp, &typarroid, &typarrname, &typrangeoid, &typrangename, &typmultirangeoid, &typmultirangename)
		if err != nil {
			return nil, fmt.Errorf("While scanning type information: %w", err)
		}

		if typtype != "b" {
			return nil, fmt.Errorf("Mailformed typtype for %s (%s), expected to have %s", typname, typtype, "b")
		}

		typ := &Type{
			Codec: nil,
			Name:  typname,
			OID:   typoid,
		}

		switch typname {
		case Uint1TypName:
			typ.Codec = &UInt1Codec{}
		case Uint2TypName:
			typ.Codec = &UInt2Codec{}
		case Uint4TypName:
			typ.Codec = &UInt4Codec{}
		case Uint8TypName:
			typ.Codec = &UInt8Codec{}
		case Uint16TypName:
			typ.Codec = &UInt16Codec{}
		case Int1TypName:
			typ.Codec = &Int1Codec{}
		case Int16TypName:
			typ.Codec = &Int16Codec{}
		default:
			return nil, fmt.Errorf("Unknown type: %s", typname)
		}

		typArr := &Type{
			Codec: &ArrayCodec{ElementType: typ},
			Name:  typarrname,
			OID:   typarroid,
		}

		typRange := &Type{
			Codec: &RangeCodec{ElementType: typ},
			Name:  typrangename,
			OID:   typrangeoid,
		}

		var typMultiRange *Type
		if typmultirangeoid > 0 {
			typMultiRange = &Type{
				Codec: &MultirangeCodec{ElementType: typ},
				Name:  typmultirangename,
				OID:   typmultirangeoid,
			}
		}

		tMap.RegisterType(typ)
		tMap.RegisterType(typArr)
		tMap.RegisterType(typRange)

		if typMultiRange != nil {
			tMap.RegisterType(typMultiRange)
		}

		result = append(result, typ, typArr, typRange)
		if typMultiRange != nil {
			result = append(result, typMultiRange)
		}

		if typnsp != "" {
			nspTyp := &Type{Name: typnsp + "." + typ.Name, OID: typ.OID, Codec: typ.Codec}
			nspTypArr := &Type{Name: typnsp + "." + typArr.Name, OID: typArr.OID, Codec: typArr.Codec}
			nspTypRange := &Type{Name: typnsp + "." + typRange.Name, OID: typRange.OID, Codec: typRange.Codec}

			var nspTypMultiRange *Type
			if typMultiRange != nil {
				nspTypMultiRange = &Type{Name: typnsp + "." + typMultiRange.Name, OID: typMultiRange.OID, Codec: typMultiRange.Codec}
			}

			tMap.RegisterType(nspTyp)
			tMap.RegisterType(nspTypArr)
			tMap.RegisterType(nspTypRange)

			if nspTypMultiRange != nil {
				tMap.RegisterType(nspTypMultiRange)
			}

			result = append(result, nspTyp, nspTypArr, nspTypRange)
			if nspTypMultiRange != nil {
				result = append(result, nspTypMultiRange)
			}
		}
	}

	return result, nil
}

// serverVersion copy of unexported function from PGX.
//
// returns the postgresql server version.
func serverVersion(c *pgx.Conn) (int64, error) {
	serverVersionStr := c.PgConn().ParameterStatus("server_version")
	serverVersionStr = regexp.MustCompile(`^[0-9]+`).FindString(serverVersionStr)
	// if not PostgreSQL do nothing
	if serverVersionStr == "" {
		return 0, fmt.Errorf("Cannot identify server version in %q", serverVersionStr)
	}

	version, err := strconv.ParseInt(serverVersionStr, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("postgres version parsing failed: %w", err)
	}
	return version, nil
}
