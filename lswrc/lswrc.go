package lswrc

func lengthOfLongestSubstringsTR(prevLongest int, b []byte) int {
	if len(b) == 0 {
		return prevLongest
	}

	visited := make(map[byte]bool)
	for _, _b := range b {
		if visited[_b] {
			break
		}
		visited[_b] = true
	}

	currLongest := len(visited)

	if currLongest < prevLongest {
		currLongest = prevLongest
	}

	return lengthOfLongestSubstringsTR(currLongest, b[1:])
}

func LengthOfLongestSubstrings(s string) int {
	b := []byte(s)
	return lengthOfLongestSubstringsTR(0, b)
}
