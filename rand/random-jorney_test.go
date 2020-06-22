package rand

import (
	"pokemon/validate"
	"testing"
)

func TestRandomJorney(t *testing.T) {

	t.Run("When_Call_GetRandomJorney({100})_Then_Return_A_ValidJorney_With_100_Steps", func(t *testing.T) {
		var expected = true
		var jorney = GetRandomJorney(100)
		var validSize = len(jorney) == 100
		var validJorney = validate.IsValidCharacteres(jorney)
		var result = validSize && validJorney

		if result != expected {
			t.Error(
				"expected", expected,
				"got", result,
			)
		}
	})
}
