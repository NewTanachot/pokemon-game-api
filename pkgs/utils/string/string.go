package stringutils

import "strings"

const (
	Empty string = ""
)

func IsNilOrEmpty(str *string) bool {
	if str == nil || strings.TrimSpace(*str) == Empty {
		return true
	}
	return false
}
