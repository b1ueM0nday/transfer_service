package base

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"time"
)

type Logger interface {
	InsertData(dtime time.Time, optype string, data json.RawMessage, isLog bool) (err error)
	Close()
}
type Repository struct {
	db      *pgx.Conn
	ctx     context.Context
	queries map[string]*pgconn.StatementDescription
	Logger
}

func NewRepository(ctx context.Context) *Repository {
	return &Repository{
		ctx: ctx,
	}
}
func (r *Repository) Connect(address, login, pass string, port uint) (err error) {

	r.db, err = pgx.Connect(context.Background(),
		fmt.Sprintf("postgres://%s:%s@%s:%d/postgres", login, pass, address, port))
	if err != nil {
		return err
	}
	for i := range queries {
		r.queries[i], err = r.db.Prepare(context.Background(), i, queries[i])
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *Repository) Close() {
	r.Close()
}
func (r *Repository) InsertData(dtime time.Time, optype string, data json.RawMessage, isLog bool) (err error) {
	if isLog {
		_, err = r.db.Exec(context.Background(), r.queries[insertLog].SQL,
			dtime, optype, data)
	} else {
		_, err = r.db.Exec(context.Background(), r.queries[insertReceipt].SQL,
			dtime, optype, data)
	}

	return err
}
