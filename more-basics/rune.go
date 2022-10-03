package main

import (
	"fmt"
	"unicode/utf8"
)

var pl = fmt.Println
var pf = fmt.Printf

func main() {
	rStr := "abcdefg"
	pl("Rune Count :", utf8.RuneCountInString(rStr))
	for i, runeVal := range rStr {
		fmt.Printf("%d : %#U : %c\n", i, runeVal, runeVal)
	}
}