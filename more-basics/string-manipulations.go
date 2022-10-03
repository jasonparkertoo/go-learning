package main

import (
	"fmt"
	"strings"
)

var pl = fmt.Println

func main() {

	dataStr := "This is a test of the emergency broadcast system. This is only a test"
	replacer := strings.NewReplacer("test", "cow")

	newString := replacer.Replace(dataStr)
	pl(newString)
	pl("Length: ", len(newString))
	pl("Contains cow : ", strings.Contains(newString, "cow"))
	pl("o index : ", strings.Index(dataStr, "o"))
	pl("Replace : ", strings.Replace(dataStr, "o", "0", -1))
	
	anotherNewString := "\nSome Words\n"
	anotherNewString = strings.TrimSpace(anotherNewString)
	
	pl("Split : ", strings.Split("a-b-c-d", "-"))
	pl("Lower : ", strings.ToLower(dataStr))
	pl("Lower : ", strings.ToUpper(dataStr))
	pl("Prefix : ", strings.HasPrefix(dataStr, "This"))
	pl("Suffix : ", strings.HasSuffix(dataStr, "test"))
}