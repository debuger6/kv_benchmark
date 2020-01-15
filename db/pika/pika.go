package pika

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"os"
)

type pikaDB struct {
	c *redis.Client
}

func (db *pikaDB) Put(ctx context.Context, key, value string) error {
	return db.c.Set(key, value, 0).Err()
}

func (db *pikaDB) Get(ctx context.Context, key string) (string, error) {
	return db.c.Get(key).Result()
}

func (db *pikaDB) BatchGet(ctx context.Context, keys []string) ([]string, error) {
	pipe := db.c.Pipeline()
	values := make([]string, 0, len(keys))
	for _, key := range keys {
		pipe.Get(key)
	}
	result, err := pipe.Exec()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return nil, err
	}
	for _, ret := range result {
		m, err := ret.(*redis.StringStringMapCmd).Result()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		for _, v := range m {
			values = append(values, v)
		}
	}

	return values, nil
}

func (db *pikaDB) BatchPut(ctx context.Context, keys, values []string) error {
	return nil
}

func (db *pikaDB) Scan(ctx context.Context, startKey string, limit int) (map[string]string, error) {
	return nil, nil
}

func (db *pikaDB) Close() error {
	return db.c.Close()
}

