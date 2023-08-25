package data

import (
	"time"

	"github.com/google/uuid"
)

type Workshop struct {
	Course
	Instructor
	Date time.Time
}

// The same method as stated in the Interface
func (w Workshop) Signup() bool {
	return true
}

func NewWorkshop(name string, instructor Instructor) Workshop {
	w := Workshop{}
	w.Name = name // takes Name from Course
	w.Instructor = instructor
	// It's not easy to pass in properties from Embedded structs :|
	w.Course.Id = uuid.New()         // easy, created on the fly
	w.Course.Instructor = instructor // easy, same as Workshop
	w.Course.Duration = 2.2          // ?!? needs to be passed as argument

	return w
}
