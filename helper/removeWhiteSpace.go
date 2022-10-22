package helper

import "strings"

func RemoveWhiteSpace(text string) string {
	return strings.ReplaceAll(text, " ", "")
}
