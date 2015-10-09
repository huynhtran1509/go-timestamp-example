package db

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"

	// Include sqlite3 driver
	_ "github.com/mattn/go-sqlite3"
)

//go:generate go-bindata -pkg db -o sql.go sql/...
//go:generate rootx-gen -mode code -psql=false -readType "(db *DB)" -writeType "(tx *TX)" -dir sql -pkg db -o db_gen.go

// DB wraps sqlx.DB
type DB struct {
	*sqlx.DB
}

// TX wraps sqlx.TX
type TX struct {
	*sqlx.Tx
}

// NewDB creates a new DB
func NewDB(path string) (*DB, error) {
	db, err := sqlx.Connect("sqlite3", path)
	if err != nil {
		return nil, err
	}
	return &DB{db}, nil
}

// SetUp creates the database
func (db *DB) SetUp() error {
	instruction, err := Asset("sql/create_timestamps.sql")
	if err != nil {
		log.Panicf("Error getting schema file: %v", err)
	}
	err = db.WithTX(func(tx *TX) error {
		_, err = tx.Exec(string(instruction))
		return err
	})
	if err != nil {
		log.Panicf("Could not set up database: %v", err)
	}
	return nil
}

// SeedData will add some data to the database
func (db *DB) SeedData() error {
	return db.WithTX(func(tx *TX) error {
		var i int64
		for i = 0; i < 10; i++ {
			if _, err := tx.InsertTimestamp(i * 100000000); err != nil {
				return err
			}
		}
		return nil
	})
}

// BeginTX creates a tx
func (db *DB) BeginTX() (*TX, error) {
	tx, err := db.Beginx()
	if err != nil {
		return nil, err
	}
	return &TX{tx}, nil
}

// WithTX executes the callback in a transaction
func (db *DB) WithTX(callback func(*TX) error) error {
	tx, err := db.BeginTX()
	if err != nil {
		return fmt.Errorf("beginning tx: %v", err)
	}
	if err = callback(tx); err != nil {
		if errTx := tx.Rollback(); errTx != nil {
			return fmt.Errorf("rolling back tx: %v", errTx)
		}
		return err
	}
	if err = tx.Commit(); err != nil {
		return fmt.Errorf("committing tx: %v", err)
	}
	return nil
}

// SQL finds sql by file
func (db *DB) SQL(filename string) string {
	return sql(filename)
}

// SQL finds sql by file
func (tx *TX) SQL(filename string) string {
	return sql(filename)
}

func sql(filename string) string {
	file := fmt.Sprintf("sql/%s", filename)
	contents, err := Asset(file)
	if err != nil {
		log.Panicf("Error getting sql file %s: %v", filename, err)
	}
	return string(contents)
}
