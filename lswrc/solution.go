package lswrc

func LengthOfLongestSubstrings(s string) int {
	if len(s) == 0 {
		return 0
	}

	visited := make(map[rune]bool)
	for _, b := range s {
		if visited[b] {
			break
		}
		visited[b] = true
	}

	currBest := len(visited)
	bestOfTheRest := LengthOfLongestSubstrings(s[1:])

	if currBest > bestOfTheRest {
		return currBest
	} else {
		return bestOfTheRest
	}
}
