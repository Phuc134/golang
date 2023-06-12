package main

import "fmt"

type Person struct {
	name   string
	age    int
	job    string
	salary int
}

func PrintPerson(person Person) {
	fmt.Println("Name: ", person.name)
	fmt.Println("Age: ", person.age)
	fmt.Println("Job: ", person.job)
	fmt.Println("Salary: ", person.salary)
}

func main() {
	var person1 Person
	var person2 Person
	// Pers1 specification
	person1.name = "Hege"
	person1.age = 45
	person1.job = "Teacher"
	person1.salary = 6000

	// Pers2 specification
	person2.name = "Cecilie"
	person2.age = 24
	person2.job = "Marketing"
	person2.salary = 4500
	PrintPerson(person1)
	PrintPerson(person2)
}
