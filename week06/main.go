package main

import (
	"fmt"
	"reflect"
)

func main() {
	i := 55              //var i int = 55
	var f float64 = 12.9 //f := 12.9
	// fmt.Printf("value i :%d,value f:%f\n", i, f)
	//fmt.Printf("%d*%f=%f/n", i, f, i*f)
	fmt.Printf("%d*%f=%f\n", i, f, i*int(f))
	fmt.Println(reflect.TypeOf(i), reflect.TypeOf(f))

}
