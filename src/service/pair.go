package service

import (
	"github.com/nuwa/server.v3/bean"
	"github.com/nuwa/server.v3/db"
	"github.com/samber/lo"
)

type Pair struct {
}

func (p *Pair) KeyValue(key string) (string, error) {
	value, err := db.SQLCli.Get(key)
	return lo.IfF(value != nil, func() string { return *value }).Else(""), err
}

func (p *Pair) Save(pair bean.Pair) (any, error) {
	err := db.SQLCli.Put(pair.Key, pair.Value, pair.Description)
	return nil, err
}

func (p *Pair) RemoveByKey(key string) (any, error) {
	err := db.SQLCli.Remove(key)
	return nil, err
}

func (p *Pair) List(prefix string) ([]bean.Pair, error) {
	return db.SQLCli.List(prefix)
}
