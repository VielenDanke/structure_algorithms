package main

import (
	"fmt"
	"regexp"
	"unicode"
)

/*
"hello"
"my phone equals 123"
"Hello hi"
"" - what to return (for all invalid inputs)

all numbers all letters in lower case
*/

func charCount(str string) map[string]int {
	// create holder for char-int
	holder := make(map[string]int)
	r, _ := regexp.Compile("[a-z_0-9]")
	// loop over the string
	for _, v := range str {
		lv := string(unicode.ToLower(v))
		// check if regexp matches
		if !r.Match([]byte(lv)) {
			continue
		}
		n, ok := holder[lv]
		if !ok {
			// if no letter/number found - add to holder and set value to 1
			holder[lv] = 1
			continue
		}
		// if the char is letter/number AND already in holder - increment the count value
		holder[lv] = n + 1
	}
	// return holder
	return holder
}

func charCountPerformance(str string) map[rune]int {
	holder := make(map[rune]int)

	for _, v := range str {
		if !(v > 47 && v < 58) && !(v > 64 && v < 91) && !(v > 96 && v < 123) {
			continue
		}
		ch, ok := holder[v]
		if !ok {
			holder[v] = 1
		} else {
			holder[v] = ch + 1
		}
	}
	return holder
}

func main() {
	str := "asdanwdubwkjd nstwa yv !xdwa"
	res := charCount(str)
	resSec := charCountPerformance(str)
	fmt.Println(res)
	fmt.Print(resSec)
}
