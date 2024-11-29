package main

import (
	"fmt"
)

type student struct {
	id   int
	name string
	gpa  float32
}

func main() {

	var student1 student
	student1.id = 202388007
	student1.name = "Dilnoza"
	student1.gpa = 4.5
	fmt.Println(student1.gpa)

	var student2 student
	student2.id = 202444990
	student2.name = "Aidana"
	student2.gpa = 4.43
	fmt.Println(student2.id)
}
