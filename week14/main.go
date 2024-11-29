package main

import (
	"fmt"
)

func main() {
	var student1 struct {
		id   int
		name string
		gpa  float32
	}
	student1.id = 202388007
	student1.name = "Dilnoza"
	student1.gpa = 4.5
	fmt.Println(student1.gpa)

	var student2 struct {
		id   int
		name string
		gpa  float32
	}
	student2.id = 202388007
	student2.name = "Aidana"
	student2.gpa = 4.43
	fmt.Println(student2.id)
}
