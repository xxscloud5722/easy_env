package db

import (
	"database/sql"
	"github.com/nuwa/server.v3/bean"
	"github.com/samber/lo"
	"strings"
)

func (db *SQLiteDB) GetScript(id string) (*bean.Script, error) {
	rows, err := db.Query("SELECT `id`, `name`, `path`, `value`, `description` FROM `script` WHERE `id` = ? LIMIT 1", id)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)
	result, err := getQueryResult(rows)
	if err != nil {
		return nil, err
	}
	return lo.IfF(len(result) > 0, func() *bean.Script {
		return &result[0]
	}).Else(nil), nil
}

func (db *SQLiteDB) GetScriptByPath(path string) (*bean.Script, error) {
	rows, err := db.Query("SELECT `id`, `name`, `path`, `value`, `description` FROM `script` WHERE `path` = ? LIMIT 1", path)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)
	result, err := getQueryResult(rows)
	if err != nil {
		return nil, err
	}
	return lo.IfF(len(result) > 0, func() *bean.Script {
		return &result[0]
	}).Else(nil), nil
}

func getQueryResult(rows *sql.Rows) ([]bean.Script, error) {
	var result []bean.Script
	for rows.Next() {
		var id int
		var name string
		var path string
		var value string
		var description string
		err := rows.Scan(&id, &name, &path, &value, &description)
		if err != nil {
			return nil, err
		}
		result = append(result, bean.Script{
			Id:          id,
			Name:        name,
			Path:        path,
			Value:       value,
			Description: description,
		})
	}
	return result, nil
}

func (db *SQLiteDB) PutScript(id *int, name, path, value, description string) error {
	_, err := db.Exec("INSERT OR REPLACE INTO `script` (`id`, `name`, `path`, `value`, `description`) VALUES (?, ?, ?, ?, ?)", id, name, path, value, description)
	if err != nil {
		return err
	}
	return nil
}

func (db *SQLiteDB) RemoveScript(id int) error {
	_, err := db.Exec("DELETE FROM `script` WHERE `id` = ?", id)
	if err != nil {
		return err
	}
	return nil
}

func (db *SQLiteDB) RemoveScriptByPath(path string) error {
	_, err := db.Exec("DELETE FROM `script` WHERE `path` = ?", path)
	if err != nil {
		return err
	}
	return nil
}

func (db *SQLiteDB) ListScript(name, path string) ([]bean.Script, error) {
	var rows *sql.Rows
	var err error
	var statement = "SELECT `id`, `name`, `path`, `value`, `description` FROM `script` "
	var params []string
	if name != "" {
		statement += lo.If(strings.HasSuffix(statement, " "), "WHERE").Else("AND") + " `name` LIKE '%' || ? || '%' "
		params = append(params, name)
	}
	if path != "" {
		statement += lo.If(strings.HasSuffix(statement, " "), "WHERE").Else("AND") + " `path` LIKE '%' || ? || '%' "
		params = append(params, path)
	}
	if len(params) > 0 {
		rows, err = db.Query(statement, params)
	} else {
		rows, err = db.Query(statement)
	}
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)

	return getQueryResult(rows)
}
