package lswrc

func LengthOfLongestSubstrings(s string) int {
	runePos := make(map[byte]int)

	longestWindowLength := 0

	currWindowStart := 0
	for currWindowEnd := 0; currWindowEnd < len(s); currWindowEnd++ {
		if _, visited := runePos[s[currWindowEnd]]; visited && runePos[s[currWindowEnd]] >= currWindowStart {
			currWindowStart = runePos[s[currWindowEnd]] + 1
		}

		runePos[s[currWindowEnd]] = currWindowEnd

		currWindowLength := currWindowEnd - currWindowStart + 1

		if currWindowLength > longestWindowLength {
			longestWindowLength = currWindowLength
		}
	}

	return longestWindowLength
}
