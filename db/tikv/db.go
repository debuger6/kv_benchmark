package tikv

import (
	"context"
	"github.com/tikv/client-go/config"
	"github.com/tikv/client-go/rawkv"
	"kv_benchmark/kvdb"
)

type tikvBuilder struct {

}

func (b *tikvBuilder) Build(arg *kvdb.DBArg) (kvdb.DB, error) {
	conf := config.Default()
	conf.RPC.MaxConnectionCount = arg.MaxConn
	conf.RPC.Batch.MaxBatchSize = arg.BatchSize

	c, err := rawkv.NewClient(context.Background(), arg.Addrs, conf)
	if err != nil {
		return nil, err
	}

	return &rawDB{c:c}, nil
}

func init() {
	kvdb.RegisterDBBuilder("tikv", &tikvBuilder{})
}
