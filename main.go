package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"pokemon/capture"
	"pokemon/validate"
)

func processar(movements string) (int, error) {

	if _, err := validate.IsValidCharacteres(movements); err != nil {
		return 0, err
	}

	return capture.AmountOfCapturedPokemons(movements), nil
}

func doProcessar(stdIn io.Reader) (int, error) {
	scanner := bufio.NewScanner(stdIn)
	var movimentos string
	fmt.Println("Movimentos válidos: N, S, E ou O sendo (Norte, Sul, Este, Oeste)")
	fmt.Print("Digite a jornada do Ash: ")
	if scanner.Scan() {
		movimentos = scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}

	amount, err := processar(movimentos)
	if err != nil {
		return 0, err
	}

	return amount, nil
}

func main() {
	if amount, err := doProcessar(os.Stdin); err != nil {
		fmt.Printf("Falha: %+v\n", err)
	} else {
		fmt.Println(amount)
	}
}
