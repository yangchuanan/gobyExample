package main

import "fmt"

type person struct {
	name string
	age  int
}

func structDemo() {
	fmt.Println(person{"Bob", 20})

	fmt.Println(person{name: "Alice", age: 24})

	fmt.Println(person{name: "Fred"})

	fmt.Println(&person{name: "Ann", age: 25})

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 29
	fmt.Println(sp.age)
}

type rect struct {
	width, height int
}

func (r *rect) area() int {
	r.width = 20
	return r.width * r.height
}

func (r rect) perim() int {
	r.width = 20
	return 2*r.width + 2*r.height
}

func methodDemo() {
	r := rect{width: 10, height: 5}
	fmt.Println("area:", r.area())
	fmt.Println("perim:", r.perim())
	fmt.Println("r.width:", r.width)

	rp := &r
	fmt.Println("area:", rp.area())
	fmt.Println("perim:", rp.perim())
	fmt.Println("r.width:", r.width)

}
