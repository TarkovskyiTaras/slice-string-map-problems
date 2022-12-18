package main

import "fmt"

/*
We'll say that a String is xy-balanced if for all the 'x' chars in the string, there exists a 'y' char somewhere later in the string.
So "xxy" is balanced, but "xyx" is not. One 'y' can balance multiple 'x's. Return true if the given string is xy-balanced.

xyBalance("aaxbby") → true
xyBalance("aaxbb") → false
xyBalance("yaaxbb") → false
*/
func xyBalance(s string) bool {
	isXRight := false
	isY := false
	for j := len(s) - 1; j >= 0; j-- {
		if s[j] == 'x' {
			isXRight = true
		}
		if s[j] == 'y' && !isXRight {
			isY = true
		}
	}
	if isXRight && isY {
		return true
	}
	return false
}

func main() {
	fmt.Println(xyBalance("tayxasyt"))
}
