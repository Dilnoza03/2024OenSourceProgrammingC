package main

import (
	"fmt"
	"reflect"
)

func main() {
	var gpa [5]float64 = [5]float64{3.5, 4.2, 4.5, 3.9, 4.23}
	gpa_slice := gpa[1:4]
	fmt.Println(gpa, reflect.TypeOf(gpa))
	fmt.Println(gpa_slice, reflect.TypeOf(gpa_slice))
}
