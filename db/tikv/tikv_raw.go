package tikv

import (
	"context"
	"errors"
	"fmt"
	"github.com/tikv/client-go/rawkv"
	"os"
)

type rawDB struct {
	c      *rawkv.Client
}

func (db *rawDB) Put(ctx context.Context, key, value string) error {
	return db.c.Put(ctx, []byte(key), []byte(value))
}

func (db *rawDB) Get(ctx context.Context, key string) (string, error) {
	row, err := db.c.Get(ctx, []byte(key))
	return string(row), err
}

func (db *rawDB) BatchPut(ctx context.Context, keys, values []string) error {
	keysLen := len(keys)
	valuesLen := len(values)
	if keysLen != valuesLen {
		return errors.New("the len of keys and values must be equal")
	}

	keysBuf := make([][]byte, 0, keysLen)
	valuesBuf := make([][]byte, 0, valuesLen)
	for i:=0; i<keysLen; i++ {
		keysBuf = append(keysBuf, []byte(keys[i]))
		valuesBuf = append(valuesBuf, []byte(values[i]))
	}

	return db.c.BatchPut(ctx, keysBuf, valuesBuf)
}

func (db *rawDB) BatchGet(ctx context.Context, keys []string) ([]string, error) {
	keysBuf := make([][]byte, 0, len(keys))
	for _, key := range keys {
		keysBuf = append(keysBuf, []byte(key))
	}

	rows, err := db.c.BatchGet(ctx, keysBuf)

	values := make([]string, 0, len(rows))
	for _, row := range rows {
		values = append(values, string(row))
	}

	return values, err
}

func (db *rawDB) Scan(ctx context.Context, startKey string, limit int) (map[string]string, error) {
	keys, values, err := db.c.Scan(ctx, []byte(startKey), nil, limit)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return nil, err
	}

	kv := make(map[string]string)
	for i := 0; i < len(keys); i++ {
		kv[string(keys[i])] = string(values[i])
	}

	return kv, nil
}

func (db *rawDB) Close() error {
	return db.c.Close()
}



