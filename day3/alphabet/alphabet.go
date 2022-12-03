package alphabet

import "unicode"

func Num(s rune) int {
	switch {
	case unicode.IsLower(s):
		return int(unicode.ToUpper(s)) - 64
	case unicode.IsUpper(s):
		return int(unicode.ToLower(s) - 70)
	}

	return 0
}
