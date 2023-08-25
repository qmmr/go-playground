package main

import (
	"fmt"

	"github.com/google/uuid"
	"qmmr.xyz/go/structs/data"
)

func main() {
	var courses [3]data.Signable

	joe := data.Instructor{
		FirstName: "Joe",
		LastName:  "Doe",
		Age:       33,
		Score:     66,
	}
	jane := data.NewInstructor("Jane", "Doe", 22, 99)

	// joe.Print()

	goCourse := data.Course{Id: uuid.New(), Name: "Go fundamentals", Instructor: joe, Duration: 1.2}
	jsCourse := data.Course{Id: uuid.New(), Name: "JavaScript fundamentals", Instructor: jane, Duration: 2.5}
	tsWorkshop := data.NewWorkshop("TypeScript Workshop", joe)

	courses[0] = goCourse
	courses[1] = jsCourse
	courses[2] = tsWorkshop

	for _, course := range courses {
		fmt.Println(course)
	}
}
