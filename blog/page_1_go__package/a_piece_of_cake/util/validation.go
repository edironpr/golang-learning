package util

import "unicode"

// ValidateUsername 校验用户名
func ValidateUsername(username string) bool {
	if len(username) < 3 || len(username) > 20 {
		return false
	}

	if !unicode.IsLetter(rune(username[0])) {
		return false
	}
	return true
}