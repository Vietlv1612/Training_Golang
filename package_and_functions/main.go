package main

import (
	"package_and_functions/mathlib"
	"package_and_functions/product"
)

func main() {
	var a int = 10
	var b int = 5

	// Add
	println("Addition: ", mathlib.Add(a, b))
	//Subtract
	println("Subtraction: ", mathlib.Subtract(a, b))
	//Multiply
	println("Multiplication: ", mathlib.Multiply(a, b))
	//Divide
	result, err := mathlib.Divide(a, b)
	if err != nil {
		println(err)
	} else {
		println("Division: ", result)
	}
	//Factorial
	println("Factorial: ", mathlib.Factorial(a))

	// Thêm sản phẩm hợp lệ
	p1 := product.Product{Name: "Laptop", Price: 1500.00}
	product.HandleAddProduct(p1)

	// Thêm sản phẩm không hợp lệ (giá âm)
	p2 := product.Product{Name: "Smartphone", Price: -500.00}
	product.HandleAddProduct(p2)
}
