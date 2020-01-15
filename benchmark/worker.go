package benchmark

import (
	"context"
	"fmt"
	"kv_benchmark/kvdb"
	"kv_benchmark/util"
	"os"
	"time"
)

type worker struct {
	db    kvdb.DB

	opCount uint
	opType  string

	keyLen   uint
	valueLen uint

	batchSize uint
}

func newWorker(db kvdb.DB, opType string, opCount, keyLen, valueLen, batchSize uint) *worker {
	return &worker{
		db:        db,
		opCount:   opCount,
		opType:    opType,
		keyLen:    keyLen,
		valueLen:  valueLen,
		batchSize: batchSize,
	}
}

func (w *worker) run(ctx context.Context) {
	var opDone uint
	for opDone < w.opCount {
		startTime := time.Now()

		w.op(ctx)

		opDone += w.batchSize
		Stat(startTime)
	}
}

func (w *worker) op(ctx context.Context) {
	switch w.opType {
	case "r":
		w.read(ctx)
	case "w":
		w.write(ctx)
	}
}

func (w *worker) read(ctx context.Context) {
	if w.batchSize == 1 {
		w.Get(ctx)
	} else {
		w.BatchGet(ctx)
	}
}

func (w *worker) write(ctx context.Context) {
	if w.batchSize == 1 {
		w.Put(ctx)
	} else {
		w.BatchPut(ctx)
	}
}

func (w *worker) Get(ctx context.Context) {
	w.db.Get(ctx, util.GetKey())
}

func (w *worker) BatchGet(ctx context.Context) {
	keys := make([]string, 0, w.batchSize)
	for i:=0; i<int(w.batchSize); i++ {
		keys = append(keys, util.GetKey())
	}
	w.db.BatchGet(ctx, keys)
}

func (w *worker) Put(ctx context.Context) {
	key := util.GenRandomKey(w.keyLen)
	value := util.GetValue(w.valueLen)

	if err := w.db.Put(ctx, key, value); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func (w *worker) BatchPut(ctx context.Context) {
	keys := make([]string, 0, w.batchSize)
	values := make([]string, 0, w.batchSize)

	for i := 0; i<int(w.batchSize); i++ {
		keys = append(keys, util.GenRandomKey(w.keyLen))
		values = append(values, util.GetValue(w.valueLen))
	}

	if err := w.db.BatchPut(ctx, keys, values); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
