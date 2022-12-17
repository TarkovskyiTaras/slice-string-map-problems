package main

import "fmt"

/*
Given two strings, return true if either of the strings appears at the very end of the other string,
ignoring upper/lower case differences (in other words, the computation should not be "case sensitive").
Note: str.toLowerCase() returns the lowercase version of a string.

endOther("Hiabc", "abc") → true
endOther("AbC", "HiaBc") → true
endOther("abc", "abXabc") → true
*/

func stringToLowerCase(s string) string {
	resultString := ""
	for i := 0; i < len(s); i++ {
		if 'A' <= s[i] && s[i] <= 'Z' {
			resultString += string(s[i] + 'a' - 'A')
		} else {
			resultString += string(s[i])
		}
	}
	return resultString
}

func endOther(s1, s2 string) bool {
	if s1 == s2 {
		return true
	}
	if len(s1) < len(s2) {
		s1, s2 = s2, s1
	}

	s1 = stringToLowerCase(s1)
	s2 = stringToLowerCase(s2)

	t := len(s1) - len(s2)
	if s1[t:] == s2 {
		return true
	}
	return false
}

func main() {
	fmt.Println(endOther("TaRaS", "Ass"))
}
