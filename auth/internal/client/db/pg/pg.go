package pg

import (
	"context"
	"github.com/Nol1feee/CLI-chat/auth/internal/client/db"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

var _ db.Client = (*pg)(nil)

type pg struct {
	pgx *pgxpool.Pool
}

func NewPGClient(ctx context.Context, dsn string) (db.Client, error) {
	con, _ := pgxpool.Connect(ctx, dsn)

	return &pg{pgx: con}, con.Ping(ctx)
}

func (p *pg) BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error) {
	return p.pgx.BeginTx(ctx, txOptions)
}

func (p *pg) Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error) {

	return p.pgx.Exec(ctx, sql, arguments...)
}

func (p *pg) QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row {

	return p.pgx.QueryRow(ctx, sql, args...)
}

func (p *pg) Close() {
	p.pgx.Close()
}
