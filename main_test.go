package main

import (
	"bytes"
	"testing"
)

func TestMain(t *testing.T) {

	t.Run("When_Input_NSEO_Then_Return_4", func(t *testing.T) {

		var stdin bytes.Buffer
		stdin.Write([]byte("NESO\n"))
		var amount, err = doProcessar(&stdin)
		var expected = 4
		if amount != expected {
			t.Error(
				"expected", expected,
				"got", amount,
				"msg", err,
			)
		}
	})

	t.Run("When_Input_TEXTOINVALIDO_Then_Return_Error", func(t *testing.T) {
		var stdin bytes.Buffer
		stdin.Write([]byte("TEXTOINVALIDO"))
		var expected = 0
		var result int
		if amount, err := doProcessar(&stdin); err != nil {
			result = 0
		} else {
			result = amount
		}
		if result != expected {
			t.Error(
				"expected", expected,
				"got", result,
			)
		}
	})

}
