package main

import (
	"fmt"
	"math"
)

var pl = fmt.Println
var pf = fmt.Printf

func main() {
	pl("5 + 4 = ", 5+4)
	pl("5 - 4 = ", 5-4)
	pl("5 * 4 = ", 5*4)
	pl("5 / 4 = ", 5/4)
	pl("5 % 4 = ", 5%4)

	n := 1
	pl("Num : ", n)
	n = n+1
	pl("Num : ", n)
	n += 1
	pl("Num : ", n)
	n++
	pl("Num : ", n)

	math.Cbrt()
}