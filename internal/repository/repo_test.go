package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"regexp"
	"testing"
	"time"
)

func TestRepository_InsertLog(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	message, _ := json.Marshal("some message")
	tests := []struct {
		name      string
		message   []byte
		operation string
		timestamp time.Time
		wantErr   bool
	}{
		{
			name:      "empty message",
			message:   nil,
			operation: "op",
			timestamp: time.Now(),
			wantErr:   true,
		},

		{
			name:      "empty operation",
			message:   message,
			timestamp: time.Now(),
			wantErr:   false,
		},
		{
			name:      "default timestamp",
			message:   message,
			operation: "",
			wantErr:   false,
		},

		{
			name:      "valid message",
			message:   message,
			operation: "op",
			timestamp: time.Now(),
			wantErr:   false,
		},
	}

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	r := Repository{
		cfg: nil,
		db:  sqlxDB,
		ctx: context.Background(),
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantErr {
				mock.ExpectExec(regexp.QuoteMeta(queries[insertLog])).WithArgs(
					tt.timestamp, tt.operation, tt.message).WillReturnError(fmt.Errorf(""))

			} else {
				mock.ExpectExec(regexp.QuoteMeta(queries[insertLog])).WithArgs(
					tt.timestamp, tt.operation, tt.message).WillReturnResult(sqlmock.NewResult(1, 1))

			}

			if err = r.InsertLog(tt.timestamp, tt.operation, tt.message); (err != nil) != tt.wantErr {
				t.Fatalf("InsertLog: %s \n expected error = %t", tt.name, tt.wantErr)
			}
		})
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestRepository_InsertReciept(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	message, _ := json.Marshal("some message")
	tests := []struct {
		name      string
		message   []byte
		operation string
		timestamp time.Time
		wantErr   bool
	}{
		{
			name:      "empty message",
			message:   nil,
			operation: "op",
			timestamp: time.Now(),
			wantErr:   true,
		},

		{
			name:      "empty operation",
			message:   message,
			timestamp: time.Now(),
			wantErr:   false,
		},
		{
			name:      "default timestamp",
			message:   message,
			operation: "",
			wantErr:   false,
		},

		{
			name:      "valid message",
			message:   message,
			operation: "op",
			timestamp: time.Now(),
			wantErr:   false,
		},
	}

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	r := Repository{
		cfg: nil,
		db:  sqlxDB,
		ctx: context.Background(),
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantErr {
				mock.ExpectExec(regexp.QuoteMeta(queries[insertReceipt])).WithArgs(
					tt.timestamp, tt.operation, tt.message).WillReturnError(fmt.Errorf(""))

			} else {
				mock.ExpectExec(regexp.QuoteMeta(queries[insertReceipt])).WithArgs(
					tt.timestamp, tt.operation, tt.message).WillReturnResult(sqlmock.NewResult(1, 1))

			}

			if err = r.InsertReceipt(tt.timestamp, tt.operation, tt.message); (err != nil) != tt.wantErr {
				t.Fatalf("InsertLog: %s \n expected error = %t", tt.name, tt.wantErr)
			}
		})
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestRepository_Close(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	r := Repository{
		cfg: nil,
		db:  sqlxDB,
		ctx: context.Background(),
	}
	r.Close()
	if err = db.Ping(); err == nil {
		t.Fatalf("database was not closed")
	}
}
