package main

func findWordsContaining(words []string, x byte) []int {
	indexArr := []int{}
	for i, word := range words {
		for _, char := range word {
			if byte(char) == x {
				indexArr = append(indexArr, i)
				break
			}
		}
	}

	return indexArr
}
