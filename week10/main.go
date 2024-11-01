package main

import (
	"bufio"
	"fmt" // Импортируем пакет для форматированного ввода-вывода
	"log" // Импортируем пакет для логирования ошибок
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	// fmt.Printf("%f\n", math.Sqrt(19.0))
	fmt.Println("Input number:")
	in := bufio.NewReader(os.Stdin)
	i, err := in.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	i = strings.TrimSpace(i)
	n, err := strconv.Atoi(i)
	if err != nil {
		log.Fatal(err)
	}
	var isPrime bool = true
	//bug fix
	if n <= 1 {
		isPrime = false
	} else {
		j := 2
		for j <= int(math.Sqrt(float64(n))) {
			if n%j == 0 {
				isPrime = false
				break
			}
			fmt.Printf("%d ", j) // Check j loop
			j++
		}
	}

	if isPrime {
		fmt.Printf("%d is prime number.", n)
	} else {
		fmt.Printf("%d is NOT prime number.", n)
	}
}
