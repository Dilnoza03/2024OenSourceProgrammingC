// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {

// 	// var now time.Time = time.Now()
// 	// //var month time.Month = now.Month()
// 	// fmt.Printf("Now %dyear %dmonth %d day \n", now.Year(), int(now.Month()), now.Day())
// 	// fmt.Printf("Now %dhour %dminute %dsecond ", now.Hour(), now.Minute(), now.Second())

// 	var army string = "오늘은 $가와 $만에 중서을 다하는 대한민 $육군임니다"
// 	armyFixed := strings.NewReplacer("$", "국")
// 	fmt.Println(army)
// 	fmt.Println(armyFixed.Replace(army))

// }

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	in := bufio.NewReader(os.Stdin)
	fmt.Println("Input your score:")
	i, err := in.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	i = strings.TrimSpace(i)
	score, _ := strconv.ParseInt(i, 10, 32)
	if score >= 60 {
		fmt.Println("A")
		fmt.Printf("%d\n", score)
	} else {
		fmt.Println("BCDF")
		fmt.Printf("%d\n", score)
	}

}

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var int int = 99
// 	a := 1
// 	fmt. Println(int + a)
// }
