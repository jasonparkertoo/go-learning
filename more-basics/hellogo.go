package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var pl = fmt.Println
var pf = fmt.Printf

func main() {
	pf("What is your name? ")
	reader := bufio.NewReader(os.Stdin)
	name, err :=reader.ReadString('\n')
	if err == nil {
		pf("Hello, %v", name)
	} else {
		log.Fatal(err)
	}
}