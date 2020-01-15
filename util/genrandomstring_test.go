package util

import "testing"

func BenchmarkBytesMaskImprSrc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenRandomKey(16)
	}
}
