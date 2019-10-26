package main

import (
	"fmt"
	"testing"
)

func BenchmarkSprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := fmt.Sprintf("foo %s", "bar")
		// Do something with the result to avoid optimizing it away
		if len(s) > 100*b.N {
			panic("foo")
		}
	}
}

func BenchmarkSprint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := fmt.Sprint("foo ", "bar")
		// Do something with the result to avoid optimizing it away
		if len(s) > 100*b.N {
			panic("foo")
		}
	}
}
