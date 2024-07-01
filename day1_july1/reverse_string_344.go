package main

func reverseString(s []byte) []byte {
	left, right := 0, len(s)-1
	for left < right {
		// Swap characters
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
	return s
}
