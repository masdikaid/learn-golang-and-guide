package basic

import (
	"fmt"
	"time"
)

// struct is one of data type that can be a template for multi data type collection
// struct also can be work like object in oop consep
// struct can tied with func that make struct can have method

// struct as multi data type collection
// declaring struct
type Student struct {
	firstname, lastname, classroom string
	age                            int
	bornDate                       time.Time
	active                         bool
}

// delaring method
func (s *Student) sayHello() {
	fmt.Println("hello my name is", s.firstname)
}

func Struct() {
	// implement struct
	a := Student{firstname: "dika", lastname: "setiawan", classroom: "IX", age: 19, bornDate: time.Now(), active: true}
	// or
	x := Student{"nanang", "subarjo", "XII", 2, time.Now(), true} // follow properties or field order declared
	// or
	b := new(Student)
	b.firstname = "ahmad"
	b.lastname = "cayadi"
	// so on
	// or
	var c Student
	c.firstname = "Eko"
	c.lastname = "Juanda"

	// access struct
	fmt.Println(a, b, c, x)
	fmt.Println(a.firstname, "kelas", a.classroom) //access field of struct

	// using method
	a.sayHello()

}
