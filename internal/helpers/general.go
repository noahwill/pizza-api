package helpers

import (
	"regexp"
)

// IsOnlyAlphabetic : validates that word only contains alphabetic characters
func IsOnlyAlphabetic(word string) bool {
	isOnlyAlphabetic := regexp.MustCompile(`^[a-zA-Z" "]+$`).MatchString
	valid := isOnlyAlphabetic(word)
	return valid
}
