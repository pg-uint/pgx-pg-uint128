//go:build nopgxregisterdefaulttypes

package types

import . "github.com/jackc/pgx/v5/pgtype"

func RegisterDefaultPgTypeVariants(tMap *Map) {
}
