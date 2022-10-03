package main

import (
	"fmt"
	"reflect"
)

var pl = fmt.Println

func main() {
	pl(reflect.TypeOf(25))
	pl(reflect.TypeOf(3.4))
	pl(reflect.TypeOf(true))
	pl(reflect.TypeOf("Hi"))
	pl(reflect.TypeOf(1))
}