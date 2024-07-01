package main

func numJewelsInStones(jewels string, stones string) int {
	jewelSet := make(map[rune]bool)
	count := 0

	for _, jewel := range jewels {
		jewelSet[jewel] = true
	}

	for _, stone := range stones {
		if jewelSet[stone] {
			count++
		}
	}
	return count
}
