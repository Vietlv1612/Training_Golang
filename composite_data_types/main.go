package main

import (
	"errors"
	"fmt"
)

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

	// Conditionals
	if p.Age > 18 {
		fmt.Println("Adult")
	} else {
		fmt.Println("Child")
	}

	// Conditionals with switch
	switch p.Age {
	case 18:
		fmt.Println("Adult")
	case 30:
		fmt.Println("Middle age")
	default:
		fmt.Println("Child")
	}

	// Conditionals with logic operators
	if p.Age > 18 && p.Age < 30 {
		fmt.Println("Young adult")
	}

	fmt.Println(s)

	fmt.Println(b)

	fmt.Println(p)
	fmt.Println(p.Name)

	fmt.Println(v)

	// Call func
	fmt.Println(add(1, 2))

	// Call func with multiple return values
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	// Call func with variadic parameters
	fmt.Println(sum(1, 2, 3, 4, 5))

	// Define a func as a variable
	var f func(int, int) int
	f = add
	fmt.Println(f(1, 2))

	// Call a func as a parameter
	fmt.Println(apply(add, 1, 2))

	// Define a func that returns a func
	multiplyBy2 := multiplier(2)
	fmt.Println(multiplyBy2(2))

	// Call a recursive func
	fmt.Println(factorial(5))

	// Khởi tạo slice
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("Initial Slice:", slice)

	// Thêm phần tử vào slice
	slice = append(slice, 6)
	fmt.Println("Slice after append:", slice)

	// Trích xuất một phần của slice
	subSlice := slice[2:5]
	fmt.Println("Sub Slice:", subSlice)

	// Thay đổi giá trị trong slice
	slice[0] = 10
	fmt.Println("Slice after modification:", slice)

	// Độ dài và sức chứa của slice
	fmt.Println("Length of Slice:", len(slice))
	fmt.Println("Capacity of Slice:", cap(slice))

	// Tạo một slice mới bằng make
	newSlice := make([]int, 3, 5)
	fmt.Println("New Slice created with make:", newSlice)

	// Thêm phần tử vào slice mới
	newSlice = append(newSlice, 1, 2)
	fmt.Println("New Slice after append:", newSlice)

	// Khởi tạo một map với các giá trị ban đầu
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	fmt.Println("Initial Map:", m)

	// Truy cập một giá trị trong map
	value := m["two"]
	fmt.Println("Value for 'two':", value)

	// Kiểm tra sự tồn tại của một key
	value, ok := m["three"]
	if ok {
		fmt.Println("Value for 'three':", value)
	} else {
		fmt.Println("'three' not found")
	}

	// Thêm một cặp key-value mới
	m["four"] = 4
	fmt.Println("Map after adding 'four':", m)

	// Cập nhật giá trị cho một key
	m["two"] = 22
	fmt.Println("Map after updating 'two':", m)

	// Xóa một cặp key-value khỏi map
	delete(m, "one")
	fmt.Println("Map after deleting 'one':", m)

	// Duyệt qua tất cả các cặp key-value trong map
	for key, value := range m {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}
	// Sử dụng make để tạo một slice
	slice1 := make([]int, 5)
	fmt.Println("Slice:", slice1) // Output: [0 0 0 0 0]

	// Sử dụng make để tạo một map
	m1 := make(map[string]int)
	m1["key1"] = 100
	m1["key2"] = 200
	fmt.Println("Map:", m1) // Output: map[key1:100 key2:200]

	// Sử dụng make để tạo một channel
	ch := make(chan int, 2)
	ch <- 10
	ch <- 20
	fmt.Println("Channel:", <-ch, <-ch) // Output: 10 20
}

// Define a function
func add(a int, b int) int {
	return a + b
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Can not divide by 0")
	} else {
		return a / b, nil
	}
}

func sum(numbers ...int) int {
	result := 0
	for _, number := range numbers {
		result += number
	}
	return result
}

func apply(f func(int, int) int, a, b int) int {
	return f(a, b)
}

func multiplier(factor int) func(int) int {
	return func(a int) int {
		return a * factor
	}
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
