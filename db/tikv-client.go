package main

import (
	"context"
	"flag"
	"fmt"
	_ "kv_benchmark/db/tikv"
	"kv_benchmark/kvdb"
)

func main() {
	startKey := ""
	limit := 0

	flag.StringVar(&startKey, "skey", "", "start key to scan")
	flag.IntVar(&limit, "l", 10, "limit to scan")
	flag.Parse()

	arg := kvdb.DBArg{
		Addrs:     []string{"127.0.0.1:2379"},
		MaxConn:   128,
		BatchSize: 1,
	}

	db := kvdb.BuildDB("tikv", &arg)
	kv, err := db.Scan(context.Background(), startKey, limit)
	if err != nil {
		return
	}

	for k, v := range kv {
		fmt.Printf("key: %s, value: %s\n", k, v)
	}

	db.Put(context.Background(), "hello", "world")
}
