package main

import "fmt"

type Human struct {
	name string
	age int
	phone string
}

type Student struct {
	Human
	school string
	loan float64
}

type Employee struct {
	Human
	company string
	money float64
}

func (h Human) SayHi()  {
	fmt.Printf("Hi, name: %s, phone: %s", h.name, h.phone)
}

func (h Human) Sing(lyrics string)  {
	fmt.Println("La la la...", lyrics)
}

func (e Employee) SayHi() {
	fmt.Printf("Hi name: %s, company: %s, phone: %s",
		e.name, e.company, e.phone)
}

type Men interface{
	SayHi()
	Sing(lyrics string)
}

func main()  {
	mike := Student{Human{"Mike", 25, "222-222-XXX"}, "MIT", 0.00}
	paul := Student{Human{"Paul", 26, "111-222-XXX"}, "Harvard", 100}
	sam := Employee{Human{"Sam", 36, "444-222-XXX"}, "Golang Inc.", 1000}
	Tom := Employee{Human{"Sam", 36, "444-222-XXX"}, "Things Ltd.", 5000}

	var i Men
	i = mike
	fmt.Println("This is Mike, a Student:")
	i.SayHi()
	i.Sing("laa")

	i = Tom
	fmt.Println("This is Tom, an Employee:")
	i.SayHi()
	i.Sing("Born to be wild")

	fmt.Println("let's use a slice of Men and what happens")
	x := make([]Men, 3)
	x[0], x[1], x[2] = paul, sam, mike
	for _, value := range x{
		value.SayHi()
	}
}