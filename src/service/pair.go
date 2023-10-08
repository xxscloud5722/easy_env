package service

import (
	"github.com/nuwa/server.v3/bean"
	"github.com/nuwa/server.v3/db"
	"github.com/samber/lo"
	"log"
)

type Pair struct {
}

var pairDB *db.PairDB

func init() {
	var err error
	pairDB, err = db.NewPairDB()
	if err != nil {
		log.Fatalln(err)
	}
}

func (p *Pair) KeyValue(key string) (string, error) {
	value, err := pairDB.Get(key)
	return lo.IfF(value != nil, func() string { return *value }).Else(""), err
}

func (p *Pair) Save(pair bean.Pair) (any, error) {
	err := pairDB.Put(pair.Key, pair.Value, pair.Description)
	return nil, err
}

func (p *Pair) RemoveByKey(key string) (any, error) {
	err := pairDB.Remove(key)
	return nil, err
}

func (p *Pair) List(prefix string) ([]bean.Pair, error) {
	return pairDB.List(prefix)
}
