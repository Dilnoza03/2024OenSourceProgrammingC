package greeting

import "fmt"

// Функция для печати приветствия "Hi"
func Hi(name string) {
	fmt.Printf("Hi %s!\n", name)
}

// Функция для печати приветствия "Hello"
func Hello(name string) {
	fmt.Printf("Hello %s~\n", name)
}

// Функция для вызова приветствий на английском
func EnglishGreetings(name string) {
	Hello(name)
	Hi(name)
}
