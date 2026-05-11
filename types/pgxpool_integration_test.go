package types_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/require"

	"github.com/pg-uint/pgx-pg-uint128/v2/types"
)

func TestRegisterAll_WithPGXPoolAfterConnect_Integration(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	dsn := os.Getenv("PGX_TEST_DATABASE")
	cfg, err := pgxpool.ParseConfig(dsn)
	require.NoError(t, err)

	cfg.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		_, err := types.RegisterAll(ctx, conn)
		return err
	}

	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	require.NoError(t, err)
	defer pool.Close()

	var got uint64
	err = pool.QueryRow(ctx, `select 1::uint8`).Scan(&got)
	require.NoError(t, err)
	require.Equal(t, uint64(1), got)
}
