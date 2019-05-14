package main

import (
	"fmt"
	"testing"
)

func BenchmarkF(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := fmt.Sprintf("foo %s", "bar")
		if len(s) > 100*b.N {
			panic("foo")
		}
	}
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := fmt.Sprint("foo ", "bar")
		if len(s) > 100*b.N {
			panic("foo")
		}
	}
}
