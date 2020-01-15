package pika

import (
	"fmt"
	"github.com/go-redis/redis"
	"kv_benchmark/kvdb"
)

type pikaBuilder struct {

}

func (b *pikaBuilder) Build(arg *kvdb.DBArg) (kvdb.DB, error) {
	c := redis.NewClient(&redis.Options{
		Addr:     arg.Addrs[0],
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := c.Ping().Result()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &pikaDB{c:c}, nil
}

func init() {
	kvdb.RegisterDBBuilder("pika", &pikaBuilder{})
}
