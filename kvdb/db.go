package kvdb

import (
	"context"
	"errors"
)

type DB interface {
	Put(ctx context.Context, key, value string) error
	Get(ctx context.Context, key string) (string, error)
	BatchPut(ctx context.Context, keys, values []string) error
	BatchGet(ctx context.Context, keys []string) ([]string, error)
	Scan(ctx context.Context, startKey string, limit int) (map[string]string, error)

	Close() error
}

type DBArg struct {
	Addrs []string

	MaxConn uint
	BatchSize uint
}

type Builder interface {
	Build(arg *DBArg) (DB, error)
}

var dbBuilderMap = make(map[string]Builder)

func RegisterDBBuilder(dbName string, builder Builder) {
	dbBuilderMap[dbName] = builder
}

func BuildDB(name string, arg *DBArg) DB {
	builder, ok := dbBuilderMap[name]
	if !ok {
		err := errors.New(name + " is not support now.")
		panic(err)
	}

	db, err := builder.Build(arg)
	if err != nil {
		panic(err)
	}

	return db
}
