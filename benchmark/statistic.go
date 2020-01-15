package benchmark

import (
	"fmt"
	"math"
	"sync/atomic"
	"time"
)

var (
	TotalOpCount int64

	TotalLatency int64
	MinLatency   int64 = math.MaxInt64
	MaxLatency   int64 = -1
	Latency90    int64
	Latency99    int64
	Latency999   int64

	Ops float64

	GlobalStartTime time.Time
)

// statistic latency and total operation after every op
// this function must be high performance to reduce the affect of ops
func Stat(start time.Time) {
	lat := int64(time.Now().Sub(start) / time.Microsecond)
	atomic.AddInt64(&TotalLatency, lat)
	atomic.AddInt64(&TotalOpCount, 1)

	updateMax(lat)
	updateMin(lat)

}

func updateMin(v int64) {
	for {
		old := atomic.LoadInt64(&MinLatency)
		if v >= old {
			break
		}

		if atomic.CompareAndSwapInt64(&MinLatency, old, v) {
			break
		}
	}
}


func updateMax(v int64) {
	for {
		old := atomic.LoadInt64(&MaxLatency)
		if v <= old {
			break
		}

		if atomic.CompareAndSwapInt64(&MaxLatency, old, v) {
			break
		}
	}
}

// Summary
func Summary() {
	min := atomic.LoadInt64(&MinLatency)
	max := atomic.LoadInt64(&MaxLatency)
	totalOp := atomic.LoadInt64(&TotalOpCount)
	totalLat := atomic.LoadInt64(&TotalLatency)
	avg := int64(float64(totalLat) / float64(totalOp))

	takes := time.Now().Sub(GlobalStartTime).Seconds()
	ops := float64(totalOp) / takes

	fmt.Printf("Takes(s): %.2f, Total_op: %d, OPS: %.2f, Min_lat(us): %d, Avg_lat(us): %d, max_lat(us): %d\n",
		takes, totalOp, ops, min, avg, max)
}