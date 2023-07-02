package main

import (
	"os"
	"testing"
)

func BenchmarkMain(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Write("/Users/artem/go/src", os.Stdout, ',')
	}
	b.ReportAllocs()
}
