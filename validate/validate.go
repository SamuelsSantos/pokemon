package validate

import (
	"regexp"
	"strings"
)

const validCharacteres = "nseoNSEO"

// IsValidCharacteres is a verification the text contains just valid characters
func IsValidCharacteres(text string) bool {

	if isEmpty(text) {
		return false
	}

	rule := "^[" + validCharacteres + "]*$"
	match, _ := regexp.MatchString(rule, text)

	return match
}

// isEmpty is a verification the text is empty
func isEmpty(text string) bool {
	return strings.TrimSpace(text) == ""
}
