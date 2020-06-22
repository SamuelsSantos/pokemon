package main

import (
	"testing"
)

func TestCapture(t *testing.T) {

	t.Run("When_Jornada_Is_{'E'}_Then_Found_{1}_Pokemons", func(t *testing.T) {
		var expected = 2
		var result = AmountOfCapturedPokemons("E")
		if result != expected {
			t.Error(
				"expected", expected,
				"got", result,
			)
		}
	})
	t.Run("When_Jornada_Is_{'NESO'}_Then_Found_4_Pokemons", func(t *testing.T) {
		var expected = 4
		var result = AmountOfCapturedPokemons("NESO")
		if result != expected {
			t.Error(
				"expected", expected,
				"got", result,
			)
		}
	})
	t.Run("When_Jornada_Is_{'NSNSNSNSNS'}_Then_Found_{2}_Pokemons", func(t *testing.T) {
		var expected = 2
		var result = AmountOfCapturedPokemons("NSNSNSNSNS")
		if result != expected {
			t.Error(
				"expected", expected,
				"got", result,
			)
		}
	})
	t.Run("1000000!=0", func(t *testing.T) {
		var expected = 0
		value := getJorney(100000)
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
		jorney := getJorney(i)
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
