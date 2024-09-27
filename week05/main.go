package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

func main() {
	// var f float32 = 3.99

	i := 55
	f := 3.99
	fmt.Println(reflect.TypeOf(i), reflect.TypeOf(f))

	fmt.Println(f, math.Ceil(3.49))
	fmt.Println(strings.Title("dika"))
	fmt.Println("i %s\n", i)
	fmt.Println("i", i, "\n")
	fmt.Println("i ", i)

}
