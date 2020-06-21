package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type point struct {
	X int
	Y int
}

const (
	norte = 78
	sul   = 79
	leste = 69
	oeste = 83
)

func move(from point, next rune) point {

	switch next {
	case norte:
		from.Y++
	case sul:
		from.X--
	case leste:
		from.X++
	case oeste:
		from.Y--
	}

	return from
}

// AmountOfCapturedPokemons represents the amount of captured pokemons.
func AmountOfCapturedPokemons(movements string) int {

	last := point{0, 0}
	var points []point
	directions := strings.ToUpper(movements)

	jorney := make(map[point]string, len(movements))
	jorney[last] = "O"

	for _, value := range directions {
		last = move(last, value)
		points = append(points, last)
		jorney[last] = string(value)
	}

	return len(jorney)
}

var pool = "NSEO"

func randomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = pool[rand.Intn(3)]
	}
	return string(bytes)
}

func processar(movements string) {

	start := time.Now()
	amount := AmountOfCapturedPokemons(movements)
	elapsed := time.Since(start)
	fmt.Println("O Ash encontrou ", amount, " pokemon(s).", "Tempo de processamento: ", elapsed.Minutes(), "minutos")

}

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Movimentos válidos: N, S, E ou O sendo (Norte, Sul, Este, Oeste)")
	fmt.Print("Digite a jornada do Ash: ")
	movimentos, _ := reader.ReadString('\n')
	processar(movimentos)
	processar(randomString(99000000))
}
