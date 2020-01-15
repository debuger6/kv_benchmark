package benchmark

import (
	"context"
	"kv_benchmark/kvdb"
	"math/rand"
	"sync"
	"time"
)

type Scheduler struct {
	dbName string
	dbArg  *kvdb.DBArg

	threadCount uint
	opCount     uint

	keyLen   uint
	valueLen uint

	readRatio  float64
}

func NewScheduler(arg *kvdb.DBArg, dbName string,
	threadCount, opCount, keyLen, valueLen uint, readRatio float64) *Scheduler {
	return &Scheduler{
		dbArg:       arg,
		dbName:     dbName,
		threadCount: threadCount,
		opCount:     opCount,
		keyLen:      keyLen,
		valueLen:    valueLen,
		readRatio:   readRatio,
	}
}

func (s *Scheduler) Run(ctx context.Context) {
	defer Summary()
	var wg sync.WaitGroup

	db := kvdb.BuildDB(s.dbName, s.dbArg)
	db.Put(ctx, "hello", "world")
	workers := make([]*worker, 0, s.threadCount)

	GlobalStartTime = time.Now()

	for i := 0; i < int(s.threadCount); i++ {
		// generate op tye
		opType := "r"
		f := rand.Float64()
		if f > s.readRatio {
			opType = "w"
		}
		w := newWorker(db, opType, s.opCount / s.threadCount, s.keyLen, s.valueLen, s.dbArg.BatchSize)
		workers = append(workers, w)
	}

	for _, w := range workers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			w.run(ctx)
		}()
	}

	go func() {
		t := time.NewTicker(time.Duration(3) * time.Second)
		defer t.Stop()

		for {
			select {
			case <-t.C:
				Summary()
			default:
			}
		}
	}()

	wg.Wait()
}
