package capture

import (
	"strings"
)

func move(from Point, next rune) Point {

	switch next {
	case NORTH:
		from.Y++
	case SOUTH:
		from.X--
	case EAST:
		from.X++
	case WEST:
		from.Y--
	}

	return from
}

// AmountOfCapturedPokemons represents the amount of captured pokemons.
func AmountOfCapturedPokemons(movements string) int {

	type Empty struct{}
	var empty Empty

	currentPoint := Point{0, 0}
	directions := strings.ToUpper(movements)

	jorney := make(map[Point]struct{}, len(movements))
	jorney[currentPoint] = empty
	for _, value := range directions {
		currentPoint = move(currentPoint, value)
		jorney[currentPoint] = empty
	}

	return len(jorney)
}
