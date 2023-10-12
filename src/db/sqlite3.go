package db

import (
	"database/sql"
	"log"
)

type SQLiteDB struct {
	*sql.DB
}

var SQLCli *SQLiteDB

func init() {
	var err error
	SQLCli, err = NewSQLiteDB()
	if err != nil {
		log.Fatalln(err)
	}
}

func NewSQLiteDB() (*SQLiteDB, error) {
	db, err := sql.Open("sqlite3", "config.db")
	if err != nil {
		return nil, err
	}

	// 执行创建表语句 , 如果不存在的话
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS pair (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			key TEXT UNIQUE,
			value TEXT,
			description TEXT
		)
	`)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS script (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT,
			path TEXT UNIQUE,
			value TEXT,
			description TEXT
		)
	`)
	if err != nil {
		return nil, err
	}

	return &SQLiteDB{db}, nil
}
