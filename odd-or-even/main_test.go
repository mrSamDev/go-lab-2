package main

import "testing"

func BenchmarkGenerateNumberPreAllocated(b *testing.B) {
	for i := 0; i < b.N; i++ {
		generateNumber(100)
	}
}

func BenchmarkGenerateNumberRuntimeAllocated(b *testing.B) {
	for i := 0; i < b.N; i++ {
		generateNumberWithoutMake(100)
	}
}
