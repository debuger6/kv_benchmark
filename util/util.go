package util

import (
	"math/rand"
	"strconv"
	"sync/atomic"
	"time"

	cmap "github.com/orcaman/concurrent-map"
)

const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits

	letterBytes = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

var (
	src = rand.NewSource(time.Now().UnixNano())

	valueMap = make(map[uint]string)
	keyMap   = cmap.New()
	keyCount int64
)

func init() {
	valueMap[1] = randStringBytesMaskImplSrc(1)
	for i := 0; i < 12; i++ {
		n := 2 << i
		valueMap[uint(n)] = randStringBytesMaskImplSrc(n)
	}
}

func randStringBytesMaskImplSrc(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}

func GenRandomKey(n uint) string {
	key := randStringBytesMaskImplSrc(int(n))
	v := atomic.AddInt64(&keyCount, 1)
	keyMap.Set(strconv.Itoa(int(v)), key)
	return key
}

func GetValue(n uint) string {
	value, ok := valueMap[n]
	if !ok {
		return "helloworld"
	}
	return value
}

func GetKey() string {
	if keyCount == 0 {
		return "hello"
	}
	r := rand.Int63n(keyCount)
	key, ok := keyMap.Get(strconv.Itoa(int(r)))
	if ok {
		return key.(string)
	}
	return "hello"
}
