package data

import "fmt"

type Instructor struct {
	Id        int
	FirstName string
	LastName  string
	Age       int
	Score     int
}

func (i Instructor) Print() {
	fmt.Printf("%v, %v (score: %v)\n", i.LastName, i.FirstName, i.Score)
}

// Factory method
func NewInstructor(firstName string, lastName string, age int, score int) Instructor {
	return Instructor{FirstName: firstName, LastName: lastName, Age: age, Score: score}
}
