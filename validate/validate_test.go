package validate

import (
	"testing"
)

func TestValidate(t *testing.T) {

	t.Run("When_Text_Is_NSEO_Then_Is_Valid", func(t *testing.T) {
		var expected = true
		var result = IsValidCharacteres("NSEO")
		if result != expected {
			t.Error(
				"expected", expected,
				"got", result,
			)
		}
	})

	t.Run("When_Text_Is_nseo_Then_Is_Valid", func(t *testing.T) {
		var expected = true
		var result = IsValidCharacteres("nseo")
		if result != expected {
			t.Error(
				"expected", expected,
				"got", result,
			)
		}
	})

	t.Run("When_Text_Is_NSEOnseo0_Then_Is_InValid", func(t *testing.T) {
		var expected = false
		var result = IsValidCharacteres("NSEOnseo0")
		if result != expected {
			t.Error(
				"expected", expected,
				"got", result,
			)
		}
	})

	t.Run("When_Text_Is_Empty_Then_Is_InValid", func(t *testing.T) {
		var expected = false
		var result = IsValidCharacteres("")
		if result != expected {
			t.Error(
				"expected", expected,
				"got", result,
			)
		}
	})

	t.Run("When_Text_Is_NESO_Then_Is_Valid", func(t *testing.T) {
		var expected = false
		var result = isEmpty("NESO")
		if result != expected {
			t.Error(
				"expected", expected,
				"got", result,
			)
		}
	})

	t.Run("When_Text_Is_Empty_Then_Result_True", func(t *testing.T) {
		var expected = true
		var result = isEmpty("")
		if result != expected {
			t.Error(
				"expected", expected,
				"got", result,
			)
		}
	})

	t.Run("When_Text_Is_With_Space_Then_Is_False", func(t *testing.T) {
		var expected = true
		var result = isEmpty("  ")
		if result != expected {
			t.Error(
				"expected", expected,
				"got", result,
			)
		}
	})
}
