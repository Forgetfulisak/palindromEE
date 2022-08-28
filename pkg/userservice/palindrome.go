package userservice

func isPalindrome(s string) bool {
	letters := []rune(s)

	for start, letter := range letters {
		end := len(letters) - start - 1
		if start >= end {
			return true
		}
		if letter != letters[end] {
			return false
		}
	}
	panic("should not reach")
}
