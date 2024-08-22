package mathlib

func Add(a, b int) int {
	return a + b
}

func Subtract(a, b int) int {
	return a - b
}

func Multiply(a, b int) int {
	return a * b
}

func Divide(a, b int) (int, error) {
	if b == 0 {
		panic("Cannot divide by zero")
	}
	return a / b, nil
}

func Factorial(a int) int {
	if a == 0 {
		return 1
	}
	return a * Factorial(a-1)
}
