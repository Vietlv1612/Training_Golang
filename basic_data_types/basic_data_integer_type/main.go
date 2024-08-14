package main

import "fmt"

func main() {
	// Integer
	var a int = 19

	// Float
	var d float32 = 16.12
	var e float64 = 16.1212

	// Không thể sử dụng như sau
	// var e float = 14.12 (Error cần phải rõ ràng kiểu dữ liệu)

	// Cú pháp tường minh rút gọn
	// Cú pháp ngắn gọn với kiểu dữ liệu tự suy luận
	number := 14
	pi := 3.14
	text := "Hello, Go!"
	isGoFun := true

	const Pi = 3.14
	const Greeting = "Hello, constant world!"

	fmt.Println("Giá trị của Pi:", Pi)
	fmt.Println(Greeting)

	var i int = 10
	var f float64 = float64(i)
	var u uint = uint(f)

	var complex complex128 = complex(5, 7)

	fmt.Println("Giá trị của complex (complex128):", complex)

	fmt.Println("Giá trị của i (int):", i)
	fmt.Println("Giá trị của f (float64):", f)
	fmt.Println("Giá trị của u (uint):", u)

	fmt.Println(a)
	fmt.Println(d)
	fmt.Println(e)

	fmt.Println(number)
	fmt.Println(pi)
	fmt.Println(text)
	fmt.Println(isGoFun)

	// Check kiểu dữ liệu
	fmt.Printf("Kiểu dữ liệu của number: %T\n", number)
	fmt.Printf("Kiểu dữ liệu của pi: %T\n", pi)
	fmt.Printf("Kiểu dữ liệu của text: %T\n", text)
	fmt.Printf("Kiểu dữ liệu của isGoFun: %T\n", isGoFun)

}
