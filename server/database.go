package server

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

const schema = `
	CREATE TABLE IF NOT EXISTS scheduler (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		date DATE,
		title TEXT,
		comment TEXT,
		repeat TEXT
	);
	CREATE INDEX IF NOT EXISTS idx_date ON scheduler(date);
	`

func InitDB(filename string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	_, err = db.Exec(schema)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func CheckDB(filename string) error {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return fmt.Errorf("database file %s does not exist", filename)
	}
	return nil
}
