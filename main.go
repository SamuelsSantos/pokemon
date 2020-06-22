package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func processar(movements string) {

	if isValidCharacteres(movements) {

		start := time.Now()
		amount := AmountOfCapturedPokemons(movements)
		elapsed := time.Since(start)
		fmt.Println("O Ash encontrou ", amount, " pokemon(s).", "Tempo de processamento: ", elapsed.Minutes(), "minutos")

	} else {
		fmt.Println("Jornada inválida!")
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Movimentos válidos: N, S, E ou O sendo (Norte, Sul, Este, Oeste)")
	fmt.Print("Digite a jornada do Ash: ")
	movimentos, _ := reader.ReadString('\n')
	processar(movimentos[:len(movimentos)-1])
}
