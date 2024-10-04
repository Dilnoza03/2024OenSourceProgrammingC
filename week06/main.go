// package main

// import (
// 	"fmt"
// 	"reflect"
// )

// func main() {
// 	i := 55              //var i int = 55
// 	var f float64 = 12.9 //f := 12.9
// 	c1 := 'Z'   // 90
// 	c2 := "김" //

// 	// fmt.Printf("value i :%d,value f:%f\n", i, f)
// 	// //fmt.Printf("%d*%f=%f/n", i, f, i*f)
// 	// fmt.Printf("%d*%f=%f\n", i, f, i*int(f))
// 	// fmt.Println(reflect.TypeOf(i), reflect.TypeOf(f))

// 	fmt.Println(c1, c2)
// 	fmt.Println(reflect.TypeOf(i), reflect.TypeOf(f), reflect.TypeOf(c1), reflect.TypeOf(c2))

// }

package main

import (
	"fmt"
)

func main() {
	var f float64
	var i int
	var b bool
	var s string

	fmt.Println(f, b, s, i)
	fmt.Printf("%f%t%s%d\n", f, b, s, i)

	f = 2.7
	i = 3
	fmt.Printf("\n\n", (f <= float64(i)), "\n") //comparison (true/false)

}
