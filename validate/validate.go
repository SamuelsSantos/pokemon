package validate

import (
	"errors"
	"regexp"
	"strings"
)

const validCharacteres = "nseoNSEO"

// IsValidCharacteres is a verification the text contains just valid characters
func IsValidCharacteres(text string) (bool, error) {

	if isEmpty(text) {
		return false, errors.New("A jornada não foi informada")
	}

	rule := "^[" + validCharacteres + "]*$"
	matched, err := regexp.MatchString(rule, text)
	if err != nil || !matched {
		return false, errors.New("A jornada possui characteres inválidos")
	}

	return matched, nil
}

// isEmpty is a verification the text is empty
func isEmpty(text string) bool {
	return strings.TrimSpace(text) == ""
}
