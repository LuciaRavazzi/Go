package main

import "fmt"

type Student struct {
	name   string
	grades []int
	age    int
}

// Change the value of the Student class, si the pointer is needed.
func (s *Student) setAge(age int) {
	(*s).age = age
}

func (s Student) getAge() int {
	return s.age
}

func main() {
	s1 := Student{"Me", []int{70, 90, 80, 85}, 23}
	fmt.Println(s1)

	s1.setAge(24)
	fmt.Println(s1)

}
