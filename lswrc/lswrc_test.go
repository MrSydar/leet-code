package lswrc

import (
	"testing"
)

func TestExample1(t *testing.T) {
	s := "abcabcabc"
	expected := 3
	if actual := LengthOfLongestSubstrings(s); actual != expected {
		t.Fatalf("expected %v, but actual is %v", expected, actual)
	}
}

func TestExample2(t *testing.T) {
	s := "bbbbb"
	expected := 1
	if actual := LengthOfLongestSubstrings(s); actual != expected {
		t.Fatalf("expected %v, but actual is %v", expected, actual)
	}
}

func TestExample3(t *testing.T) {
	s := "pwwkew"
	expected := 3
	if actual := LengthOfLongestSubstrings(s); actual != expected {
		t.Fatalf("expected %v, but actual is %v", expected, actual)
	}
}

func TestOneChar(t *testing.T) {
	s := " "
	expected := 1
	if actual := LengthOfLongestSubstrings(s); actual != expected {
		t.Fatalf("expected %v, but actual is %v", expected, actual)
	}
}
