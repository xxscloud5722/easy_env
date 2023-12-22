package service

import (
	"encoding/json"
	"errors"
	"github.com/dgraph-io/badger/v4"
	"github.com/xxscloud5722/easy_env/server/src/bean"
	"log"
)

var (
	db *badger.DB
)

func init() {
	var err error
	db, err = badger.Open(badger.DefaultOptions("config/pair"))
	if err != nil {
		log.Fatal(err)
	}
}

type Pair struct{}

func (p *Pair) KeyValue(key string) (string, error) {
	var result *string
	err := db.View(func(txn *badger.Txn) error {
		value, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}
		snapshot, err := value.ValueCopy(nil)
		if err != nil {
			return err
		}
		var pair bean.Pair
		err = json.Unmarshal(snapshot, &pair)
		if err != nil {
			return err
		}
		result = &pair.Value
		return nil
	})
	if err != nil {
		if errors.Is(err, badger.ErrKeyNotFound) {
			return "", nil
		}
		return "", err
	}
	return *result, err
}

func (p *Pair) Save(pair bean.Pair) (any, error) {
	err := db.Update(func(txn *badger.Txn) error {
		byteArray, err := json.Marshal(pair)
		if err != nil {
			return err
		}
		return txn.Set([]byte(pair.Key), byteArray)
	})
	if err != nil {
		return false, err
	}
	return true, err
}

func (p *Pair) RemoveByKey(key string) (any, error) {
	err := db.Update(func(txn *badger.Txn) error {
		return txn.Delete([]byte(key))
	})
	if err != nil {
		return false, err
	}
	return true, err
}

func (p *Pair) List(prefix string) ([]bean.Pair, error) {
	var pairs []bean.Pair
	var prefixBytes = []byte(prefix)
	err := db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.Prefix = prefixBytes

		it := txn.NewIterator(opts)
		defer it.Close()

		for it.Seek(prefixBytes); it.ValidForPrefix(prefixBytes); it.Next() {
			item := it.Item()
			value, err := item.ValueCopy(nil)
			if err != nil {
				return err
			}

			var pair bean.Pair
			err = json.Unmarshal(value, &pair)
			if err != nil {
				return err
			}
			pairs = append(pairs, pair)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return pairs, nil
}
