package main

import "strings"

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

func toString(p point) string {
	return string(p.X) + string(p.Y)
}

// AmountOfCapturedPokemons represents the amount of captured pokemons.
func AmountOfCapturedPokemons(movements string) int {

	currentPoint := point{0, 0}
	directions := strings.ToUpper(movements)

	jorney := make(map[point]string, len(movements))
	jorney[currentPoint] = toString(currentPoint)

	for _, value := range directions {
		currentPoint = move(currentPoint, value)
		jorney[currentPoint] = toString(currentPoint)
	}

	return len(jorney)
}
