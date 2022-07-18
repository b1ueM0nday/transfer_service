package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"log"
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
		cfg *Config
		db  *sqlx.DB
		ctx context.Context
	}
)

var DefaultConfig = Config{
	Login:    "postgres",
	Password: "",
	Address:  "localhost",
	Port:     5432,
}

func NewRepository(ctx context.Context, cfg *Config) *Repository {
	return &Repository{
		ctx: ctx,
		cfg: cfg,
	}
}
func (r *Repository) Connect(address, login, pass string, port uint) (err error) {

	r.db, err = sqlx.Open("pgx", fmt.Sprintf("postgres://%s:%s@%s:%d/postgres", login, pass, address, port))
	if err != nil {
		log.Fatal(err)
	}

	r.db.DB.SetMaxOpenConns(5) //??

	for i := range queries {
		_, err = r.db.PrepareContext(r.ctx, queries[i])
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *Repository) Close() {
	r.db.Close()
}

func (r *Repository) InsertLog(dtime time.Time, optype string, data json.RawMessage) error {
	res, err := r.db.Exec(queries[insertLog], dtime, optype, data)

	if err != nil {
		return err
	}
	if count, err := res.RowsAffected(); count < 1 || err != nil {
		return errors.New("no log rows affected")
	}
	return nil
}

func (r *Repository) InsertReceipt(dtime time.Time, optype string, data json.RawMessage) error {
	res, err := r.db.Exec(queries[insertReceipt], dtime, optype, data)
	if err != nil {
		return err
	}
	if count, err := res.RowsAffected(); count < 1 || err != nil {
		return errors.New("no receipt rows affected")
	}
	return nil
}
