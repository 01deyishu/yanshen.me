package main

import (
	"fmt"
)

func main() {
	fmt.Println("Start a algorithm code ...")
	l := strStr("aaaaa", "bba")
	fmt.Println(l)
}

func strStr(haystack string, needle string) int {
	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}
