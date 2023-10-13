package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/nuwa/server.v3/bean"
)

func (db *SQLiteDB) Get(key string) (*string, error) {
	rows, err := db.Query("SELECT `value` FROM `pair` WHERE `key` = ? LIMIT 1", key)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)

	for rows.Next() {
		var value string
		err = rows.Scan(&value)
		if err != nil {
			return nil, err
		}
		return &value, nil
	}
	return nil, nil
}

func (db *SQLiteDB) Put(key, value, description string) error {
	_, err := db.Exec("INSERT OR REPLACE INTO `pair` (`key`, `value`, `description`) VALUES (?, ?, ?)", key, value, description)
	if err != nil {
		return err
	}
	return nil
}

func (db *SQLiteDB) Remove(key string) error {
	_, err := db.Exec("DELETE FROM `pair` WHERE `key` = ?", key)
	if err != nil {
		return err
	}
	return nil
}

func (db *SQLiteDB) List(prefix string) ([]bean.Pair, error) {
	var rows *sql.Rows
	var err error
	if prefix == "" {
		rows, err = db.Query("SELECT `id`, `key`, IFNULL(`value`, ''), IFNULL(`description`, '') FROM `pair` ORDER BY `id` DESC")
	} else {
		rows, err = db.Query("SELECT `id`, `key`, IFNULL(`value`, ''), IFNULL(`description`, '') FROM `pair` WHERE `key` LIKE ? || '%' ORDER BY `id` DESC", prefix)
	}
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)

	var pairs []bean.Pair
	for rows.Next() {
		var id int
		var key string
		var value string
		var description string
		err = rows.Scan(&id, &key, &value, &description)
		if err != nil {
			return nil, err
		}
		pairs = append(pairs, bean.Pair{
			Id:          id,
			Key:         key,
			Value:       value,
			Description: description,
		})
	}
	return pairs, nil
}
