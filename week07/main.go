package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	// var now time.Time = time.Now()
	// //var month time.Month = now.Month()
	// fmt.Printf("Now %dyear %dmonth %d day \n", now.Year(), int(now.Month()), now.Day())
	// fmt.Printf("Now %dhour %dminute %dsecond ", now.Hour(), now.Minute(), now.Second())

	// var army string = "오늘은 $가와 $만에 중서을 다하는 대한민 $육군임니다"
	// armyFixed := strings.NewReplacer("$", "국")
	// fmt.Println(army)
	// fmt.Println(armyFixed.Replace(army))

	// var army string = "오늘은 $가와 $만에 중서을 다하는 대한민 $육군임니다"
	// armyFixed := strings.NewReplacer("$", "국")
	// fmt.Println(army)
	// fmt.Println(armyFixed.Replace(army))

	in := bufio.NewReader(os.Stdin)
	fmt.Println("Input your name:")
	name, err := in.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(name)
	}

}
