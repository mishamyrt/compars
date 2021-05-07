package substring

import "strings"

// Trim given comment
func Trim(text string) string {
	return strings.Trim(text, "\t *")
}

// GetSubsequent text
func GetSubsequent(text string, delimeter string) string {
	offset := strings.Index(text, delimeter)
	if offset < 0 {
		return ""
	}
	return text[offset+len(delimeter):]
}

// GetPrevious text
func GetPrevious(text string, delimeter string) string {
	offset := strings.Index(text, delimeter)
	if offset < 0 {
		return ""
	}
	return text[:offset]
}
