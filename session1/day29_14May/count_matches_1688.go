package main

func numberOfMatches(n int) int {
	var matches, teams int
	teams = n
	for teams > 1 {
		if teams%2 == 0 {
			matches += teams / 2
			teams /= 2
		} else {
			matches += (teams - 1) / 2
			teams = ((teams - 1) / 2) + 1
		}
	}
	return matches
}

func main() {

}
