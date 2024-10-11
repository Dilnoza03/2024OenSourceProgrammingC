package main

import (
	"fmt"
	"time"
)

func main() {
	var now time.Time = time.Now()
	//var month time.Month = now.Month()
	fmt.Printf("Now %dyear %dmonth %d day \n", now.Year(), int(now.Month()), now.Day())
	fmt.Printf("Now %dhour %dminute %dsecond ", now.Hour(), now.Minute(), now.Second())

}
