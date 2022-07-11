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

func (db *Database) InsertLog(mtype string, data json.RawMessage) error {
	err := db.db.QueryRow(context.Background(), "insert into public.logs (date, message_type, message) values ($1,$2,$3)", time.Now(), mtype, data)
	if err != nil {
		return fmt.Errorf("QueryRow failed: %v\n", err)
	}
	return nil
}
