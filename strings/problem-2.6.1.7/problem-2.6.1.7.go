package main

import "fmt"

/*
Return a version of the given string, where for every star (*) in the string the star and the chars immediately
to its left and right are gone. So "ab*cd" yields "ad" and "ab**cd" also yields "ad".

starOut("ab*cd") → "ad"
starOut("ab**cd") → "ad"
starOut("sm*eilly") → "silly"
*/

/*
-----------------------------
"ta*ast*ras"  -->  "tsas"
"ta*as"  -->  "ts"
"*aras"  -->  "ras"
"tara*"  -->  "tar"
*/

func starOut1(givenString string) string {
	resultString := ""
	for i := 0; i < len(givenString); i++ {
		if i == 0 && givenString[i] == '*' {
			i++
			continue
		}
		if givenString[i] == '*' {
			resultString = resultString[:len(resultString)-1]
			i++
			continue
		}
		resultString += string(givenString[i])
	}
	return resultString
}

func main() {
	fmt.Println(starOut1("ta*ast*ras")) // tsas
	fmt.Println(starOut1("*aras"))
	fmt.Println(starOut1("tara*"))

	fmt.Println(starOut1("*ara*taras*aras*ara*")) //"rararar" == "rararar"
}
