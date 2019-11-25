package threesum

import (
	"math/rand"
	"sort"
	"testing"
)

type test struct {
	input    []int
	expected bool
}

var (
	int100  []int
	int1000 []int
	int2000 []int

	testData = []test{
		{[]int{-1, 0, 1}, true},
		{[]int{-1, 0, 0}, false},
	}
)

func init() {
	int100 = intSlice(100)
	int1000 = intSlice(1000)
	int2000 = intSlice(2000)
}

func intSlice(n int) []int {
	ints := make([]int, n)
	for i := range ints {
		ints[i] = rand.Int()
	}
	sort.Ints(ints)

	return ints
}

func runTest(f func([]int) bool, t *testing.T) {
	for _, d := range testData {
		if f(d.input) != d.expected {
			t.Errorf("invalid result for %v", d.input)
		}
	}
}

func TestNaive(t *testing.T) {
	runTest(naive, t)
}

func TestHash(t *testing.T) {
	runTest(hash, t)
}

func TestTwoPointer(t *testing.T) {
	runTest(twopointer, t)
}

func TestBinarySearch(t *testing.T) {
	runTest(binarysearch, t)
}

func runBenchmark(evaluator func([]int) bool, input []int, b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		evaluator(input)
	}
}

func BenchmarkNaive100(b *testing.B) {
	runBenchmark(naive, int100, b)
}

func BenchmarkNaive1000(b *testing.B) {
	runBenchmark(naive, int1000, b)
}

func BenchmarkNaive2000(b *testing.B) {
	runBenchmark(naive, int2000, b)
}

func BenchmarkHash100(b *testing.B) {
	runBenchmark(hash, int100, b)
}

func BenchmarkHash1000(b *testing.B) {
	runBenchmark(hash, int1000, b)
}

func BenchmarkHash2000(b *testing.B) {
	runBenchmark(hash, int2000, b)
}

func BenchmarkTwoPointer100(b *testing.B) {
	runBenchmark(twopointer, int100, b)
}

func BenchmarkTwoPointer1000(b *testing.B) {
	runBenchmark(twopointer, int1000, b)
}

func BenchmarkTwoPointer2000(b *testing.B) {
	runBenchmark(twopointer, int2000, b)
}

func BenchmarkBinarySearch100(b *testing.B) {
	runBenchmark(binarysearch, int100, b)
}

func BenchmarkBinarySearch1000(b *testing.B) {
	runBenchmark(binarysearch, int1000, b)
}

func BenchmarkBinarySearch2000(b *testing.B) {
	runBenchmark(binarysearch, int2000, b)
}
