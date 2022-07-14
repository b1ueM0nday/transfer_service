package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"time"
)

type (
	Connector interface {
		Connect(address, login, pass string, port uint) (err error)
		Close()
	}
	Database interface {
		InsertLog(dtime time.Time, optype string, data json.RawMessage) (err error)
		InsertReceipt(dtime time.Time, optype string, data json.RawMessage) (err error)
	}
	Repo interface {
		Connector
		Database
	}
	Config struct {
		Login    string `yaml:"login"`
		Password string `yaml:"password"`
		Address  string `yaml:"address"`
		Port     uint   `yaml:"port"`
	}
	Repository struct {
		cfg     *Config
		db      *pgx.Conn
		ctx     context.Context
		queries map[string]*pgconn.StatementDescription
	}
)

func NewRepository(ctx context.Context, cfg *Config) *Repository {
	return &Repository{
		ctx:     ctx,
		queries: make(map[string]*pgconn.StatementDescription, len(queries)),
		cfg:     cfg,
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

func (r *Repository) InsertLog(dtime time.Time, optype string, data json.RawMessage) error {
	res, err := r.db.Exec(context.Background(), r.queries[insertLog].SQL,
		dtime, optype, data)
	if err != nil {
		return err
	}
	if count := res.RowsAffected(); count < 1 {
		return errors.New("no log rows affected")
	}
	return nil
}

func (r *Repository) InsertReceipt(dtime time.Time, optype string, data json.RawMessage) error {
	res, err := r.db.Exec(context.Background(), r.queries[insertReceipt].SQL,
		dtime, optype, data)
	if err != nil {
		return err
	}
	if count := res.RowsAffected(); count < 1 {
		return errors.New("no receipt rows affected")
	}
	return nil
}
