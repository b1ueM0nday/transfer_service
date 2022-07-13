package base

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"time"
)

type Database struct {
	db      *pgx.Conn
	ctx     context.Context
	queries map[string]*pgconn.StatementDescription
}

func NewDatabase(ctx context.Context) *Database {
	return &Database{
		ctx: ctx,
	}
}
func (db *Database) Connect(address, login, pass string, port uint) (err error) {

	db.db, err = pgx.Connect(context.Background(),
		fmt.Sprintf("postgres://%s:%s@%s:%d/postgres", login, pass, address, port))
	if err != nil {
		return err
	}
	for i := range queries {
		db.queries[i], err = db.db.Prepare(context.Background(), i, queries[i])
		if err != nil {
			return err
		}
	}

	return nil
}

func (db *Database) Close() {
	db.Close()
}
func (db *Database) InsertLog(dtime time.Time, mtype string, data json.RawMessage) error {
	_, err := db.db.Exec(context.Background(), db.queries[insertLog].SQL,
		dtime, mtype, data)
	return err
}
func (db *Database) InsertReceipt(dtime time.Time, optype string, data json.RawMessage) error {
	_, err := db.db.Exec(context.Background(), db.queries[insertLog].SQL,
		dtime, optype, data)
	return err
}
