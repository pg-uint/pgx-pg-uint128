package types

import "fmt"
import . "github.com/jackc/pgx/v5/pgtype"

func codecScan(codec Codec, m *Map, oid uint32, format int16, src []byte, dst any) error {
	scanPlan := codec.PlanScan(m, oid, format, dst)
	if scanPlan == nil {
		return fmt.Errorf("PlanScan did not find a plan")
	}
	return scanPlan.Scan(src, dst)
}
