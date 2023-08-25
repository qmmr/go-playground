package data

import (
	"fmt"

	"github.com/google/uuid"
)

type Duration float32 // in hours

type Course struct {
	Id         uuid.UUID
	Name       string
	Duration   Duration
	Instructor Instructor
}

// The same method as stated in the Interface
func (c Course) Signup() bool {
	return true
}

func (c Course) String() string {
	return fmt.Sprintf("Id: %v, Name: %v by (%v %v), length: %v hours",
		c.Id, c.Name, c.Instructor.FirstName, c.Instructor.LastName, c.Duration)
}
