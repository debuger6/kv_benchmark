package main

import (
	"context"
	"flag"
	"fmt"
	"kv_benchmark/benchmark"
	_ "kv_benchmark/db/pika"
	_ "kv_benchmark/db/tikv"
	"kv_benchmark/kvdb"
	"os"
	"strings"
)

var (
	dbArg  = new(kvdb.DBArg)
	dbName string
	addrs  string

	threadCount uint
	opCount     uint
	opType      string

	keyLen   uint
	valueLen uint

	batchSize uint

	readRatio float64
)

func initVar() {
	flag.StringVar(&dbName, "db", "tikv", "db's name to test")
	flag.StringVar(&addrs, "addr", "127.0.0.1:2379", "db's addrs to connect")
	flag.UintVar(&dbArg.MaxConn, "max_conn", 128, "max connection to db")
	flag.UintVar(&dbArg.BatchSize, "batch", 1, "batch size of request to db")

	flag.UintVar(&threadCount, "tc", 1, "thread count to test")
	flag.UintVar(&opCount, "oc", 1000, "operation count to test")
	flag.StringVar(&opType, "ot", "w", `operation type ['r' or 'w']`)

	flag.UintVar(&keyLen, "klen", 10, "the length of key")
	flag.UintVar(&valueLen, "vlen", 64, "the length of value")

	flag.Float64Var(&readRatio, "rr", 0.5, "read ratio of all operation")

	flag.Parse()

	if opCount < threadCount {
		fmt.Fprintln(os.Stderr, "operation count must gte thread count")
		os.Exit(-1)
	}

	dbArg.Addrs = strings.Split(addrs, ",")
}

func main() {
	initVar()
	scheduler := benchmark.NewScheduler(dbArg, dbName, threadCount, opCount, keyLen, valueLen, readRatio)
	scheduler.Run(context.Background())
}
