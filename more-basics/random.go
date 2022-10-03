package main

import (
	"fmt"
	"math/rand"
	"time"
)

var pl = fmt.Println
var pf = fmt.Printf

func main() {
	seed := time.Now().Unix()
	rand.Seed(seed)
	randNum := rand.Intn(50) + 1
	pl("Random : ", randNum);
}