package random


func IsLower(s string) bool {
	i := 0
	for _, v := range s {
		if rune(v) >= 'a' && rune(v) <= 'z' {
			i++
		} else {
			return false
		}
	}
	return true
}