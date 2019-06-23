package utils

import "strings"

func CapToLowerCase(s string) string {
	return strings.ReplaceAll(strings.ToLower(s), " ", "_")
}
