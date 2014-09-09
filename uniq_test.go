package uniq

import (
	"testing"
)

func BenchmarkMutex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		uniqFn()
	}
}

func BenchmarkChannel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		<-uniq
	}
}

func TestFoo(t *testing.T) {}
