package main

import (
	"fmt"
	"time"
)

func main() {
	var now time.Time = time.Now()
	var month int = now.Year()
	fmt.Println(month)
	fmt.Println(now)

}
