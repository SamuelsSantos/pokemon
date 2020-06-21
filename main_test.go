package main

import (
	"testing"
)

func TestMain(t *testing.T) {

	t.Run("E=1", func(t *testing.T) {
		var expected = 2
		var result = AmountOfCapturedPokemons("E")
		if result != expected {
			t.Error(
				"expected", expected,
				"got", result,
			)
		}
	})
	t.Run("NESO=4", func(t *testing.T) {
		var expected = 4
		var result = AmountOfCapturedPokemons("NESO")
		if result != expected {
			t.Error(
				"expected", expected,
				"got", result,
			)
		}
	})
	t.Run("NSNSNSNSNS=2", func(t *testing.T) {
		var expected = 2
		var result = AmountOfCapturedPokemons("NSNSNSNSNS")
		if result != expected {
			t.Error(
				"expected", expected,
				"got", result,
			)
		}
	})
	t.Run("99000000!=0", func(t *testing.T) {
		var expected = 0
		value := randomString(99000000)
		var result = AmountOfCapturedPokemons(value)
		if result == expected {
			t.Error(
				"expected", expected,
				"got", result,
			)
		}

	})
}

var result int

func benchmarkAmountOfCapturedPokemons(i int, b *testing.B) {
	var r int
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		jorney := randomString(i)
		b.StartTimer()
		r = AmountOfCapturedPokemons(jorney)
	}
	result = r
}

func Benchmark1000(b *testing.B) {
	benchmarkAmountOfCapturedPokemons(1000, b)
}
func Benchmark50000(b *testing.B) {
	benchmarkAmountOfCapturedPokemons(5000, b)
}
func Benchmark10000(b *testing.B) {
	benchmarkAmountOfCapturedPokemons(100000, b)
}
func Benchmark500000(b *testing.B) {
	benchmarkAmountOfCapturedPokemons(500000, b)
}
func Benchmark1000000(b *testing.B) {
	benchmarkAmountOfCapturedPokemons(1000000, b)
}
func Benchmark50000000(b *testing.B) {
	benchmarkAmountOfCapturedPokemons(5000000, b)
}
