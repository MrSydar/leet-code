package main

import (
	"bytes"
	"fmt"
	"mrsydar/leetcode/lswrc"
	"os"
)

func main() {
	var buf bytes.Buffer
	buf.ReadFrom(os.Stdin)
	longest := lswrc.LengthOfLongestSubstrings(buf.String())
	fmt.Print(longest)
}
