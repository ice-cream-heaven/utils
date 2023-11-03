package stringx

import "unicode/utf16"

func Utf16Len(str string) int {
	if str == "" {
		return 0
	}
	return len(utf16.Encode([]rune(str)))
}
