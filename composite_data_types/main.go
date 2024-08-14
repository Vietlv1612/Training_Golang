package main

import "fmt"

func main() {
	b := []int{1, 2, 3, 4, 5}

	for i := 0; i < len(b); i++ {
		fmt.Println(b[i])
	}

	s := b[1:3]

	// append item to array b
	b = append(b, 6)

	// Define struct type

	type Person struct {
		Name string
		Age  int
	}

	type Employee struct {
		Person
		Position string
	}

	p := Person{Name: "Viet", Age: 30}

	v := Employee{
		Person: Person{
			Name: "Viet",
			Age:  30,
		},
		Position: "Developer",
	}

	fmt.Println(s)

	fmt.Println(b)

	fmt.Println(p)
	fmt.Println(p.Name)

	fmt.Println(v)
}
