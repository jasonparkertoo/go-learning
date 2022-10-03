package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var pl = fmt.Println
var pf = fmt.Printf

func main() {
	for i := 1; i <= 5; i++ {
		pl(i)
	}
	
	for j := 5; j >= 1; j-- {
		pl(j)
	}

	seed := time.Now().Unix()
	rand.Seed(seed)
	num := rand.Intn(50) + 1

	count := 0
	for true {
		count++
		pf("Guess #%v: ", count)

		reader := bufio.NewReader(os.Stdin)
		guess, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		guess = strings.TrimSpace(guess)
		guessNum, err := strconv.Atoi(guess)
		if err != nil {
			log.Fatal(err)
		}

		if guessNum < num {
			pl("Guess is to low guess again...")
		} else if guessNum > num {
			pl("Guess is to high guess again ...")
		} else {
			pl("Congradulations, you guess the number!")
			break
		}
	}
}