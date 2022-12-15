package main

import "fmt"

/*
Returns true if for every '*' (star) in the string, if there are chars both immediately before and after the star, they are the same.

sameStarChar("xy*yzz") → true
sameStarChar("xy*zzz") → false
sameStarChar("*xa*az") → true
*/

func sameStarChar(givenString string) bool {
	sameCharAstr := false
	for i := 1; i < len(givenString)-1; i++ {
		if string(givenString[i]) == "*" {
			if givenString[i-1] != givenString[i+1] {
				return false
			} else {
				sameCharAstr = true
			}
		}
	}
	return sameCharAstr
}

func main() {
	fmt.Println(sameStarChar("tas*s"))
}
