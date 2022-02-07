package main

import "fmt"

func main() {
	str1 := "abba"
	str2 := "cdd"
	str3 := "dvdf"
	str4 := "bbbbbbb"
	fmt.Println(lengthOfLongestSubstring(str1))
	fmt.Println(lengthOfLongestSubstring(str2))
	fmt.Println(lengthOfLongestSubstring(str3))
	fmt.Println(lengthOfLongestSubstring(str4))
}

func lengthOfLongestSubstring(s string) int {
	lastOccurred := make(map[byte]int)
	start := 0
	maxLength := 0
	for i, ch := range []byte(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		lengthNow := i - start + 1
		if lengthNow > maxLength {
			maxLength = lengthNow
		}
		lastOccurred[ch] = i
	}
	return maxLength
}
