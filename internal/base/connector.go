package base

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jackc/pgx/v4"
	"os"
	"time"
)

type Database struct {
	db  *pgx.Conn
	ctx context.Context
}

func NewDatabase(ctx context.Context) *Database {
	return &Database{
		ctx: ctx,
	}
}
func (db *Database) Connect() {
	var err error
	db.db, err = pgx.Connect(context.Background(), "postgres://postgres:postgrespw@localhost:55000/postgres")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
}

func (db *Database) InsertLog(dtime time.Time, mtype string, data json.RawMessage) error {
	_, err := db.db.Exec(context.Background(), "insert into public.logs (date, message_type, message) values ($1,$2,$3)",
		dtime, mtype, data)
	return err
}
func (db *Database) InsertReceipt(dtime time.Time, optype string, data json.RawMessage) error {
	_, err := db.db.Exec(context.Background(), "insert into public.receipts (date, op_type, receipt) values ($1,$2,$3)",
		dtime, optype, data)
	return err
}
