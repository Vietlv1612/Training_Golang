package main

import (
	"fmt"
	"math"
)

// Bài tập 1: Viết hàm Max sử dụng Generics để tìm giá trị lớn nhất trong hai số

// Hàm Max sử dụng Generics với constraint 'comparable' để có thể so sánh các giá trị
// Tạo constraint tùy chỉnh cho các kiểu dữ liệu có thể sử dụng các toán tử >, <, >=, <=
type Ordered interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64 | string
}

// Hàm Max sử dụng Generics với constraint 'Ordered' để có thể so sánh các giá trị
func Max[T Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// Bài tập 2: Viết hàm Filter sử dụng Generics để lọc các phần tử trong slice

// Hàm Filter nhận một slice và một hàm kiểm tra điều kiện, trả về một slice mới chỉ chứa các phần tử thỏa mãn điều kiện
func Filter[T any](s []T, f func(T) bool) []T {
	var result []T
	for _, v := range s {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

// Bài tập 3: Viết cấu trúc dữ liệu Stack sử dụng Generics

// Định nghĩa Stack sử dụng Generics
type Stack[T any] struct {
	elements []T
}

// Phương thức Push để thêm phần tử vào Stack
func (s *Stack[T]) Push(element T) {
	s.elements = append(s.elements, element)
}

// Phương thức Pop để lấy phần tử ra khỏi Stack
func (s *Stack[T]) Pop() T {
	if len(s.elements) == 0 {
		var zero T
		return zero
	}
	lastIndex := len(s.elements) - 1
	element := s.elements[lastIndex]
	s.elements = s.elements[:lastIndex]
	return element
}

// Bài tập 4: Viết hàm Generics để tính tổng của các phần tử trong slice

// Hàm Sum tính tổng các phần tử trong slice
func Sum[T int | float64](s []T) T {
	var sum T
	for _, v := range s {
		sum += v
	}
	return sum
}

// Bài tập 5: Viết hàm Generics để tìm giá trị tuyệt đối của một số

// Hàm Abs tính giá trị tuyệt đối của một số
func Abs[T int | float64](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

// Bài tập 6: Viết hàm Generics để tìm khoảng cách giữa hai điểm trong không gian 2D

// Định nghĩa struct Point2D sử dụng Generics
type Point2D[T float64] struct {
	X, Y T
}

// Hàm Distance tính khoảng cách giữa hai điểm trong không gian 2D
func Distance[T float64](p1, p2 Point2D[T]) T {
	return T(math.Sqrt(math.Pow(float64(p2.X-p1.X), 2) + math.Pow(float64(p2.Y-p1.Y), 2)))
}

func main() {
	// Bài tập 1: Sử dụng hàm Max
	fmt.Println("Max of 3 and 5:", Max(3, 5))
	fmt.Println("Max of 3.14 and 2.71:", Max(3.14, 2.71))

	// Bài tập 2: Sử dụng hàm Filter
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	evenNumbers := Filter(numbers, func(n int) bool { return n%2 == 0 })
	fmt.Println("Even numbers:", evenNumbers)

	// Bài tập 3: Sử dụng Stack
	var stack Stack[int]
	stack.Push(10)
	stack.Push(20)
	fmt.Println("Stack pop:", stack.Pop()) // Output: 20
	fmt.Println("Stack pop:", stack.Pop()) // Output: 10

	// Bài tập 4: Sử dụng hàm Sum
	fmt.Println("Sum of int slice:", Sum([]int{1, 2, 3, 4, 5}))
	fmt.Println("Sum of float64 slice:", Sum([]float64{1.1, 2.2, 3.3}))

	// Bài tập 5: Sử dụng hàm Abs
	fmt.Println("Abs of -10:", Abs(-10))
	fmt.Println("Abs of -3.14:", Abs(-3.14))

	// Bài tập 6: Sử dụng hàm Distance
	p1 := Point2D[float64]{X: 1, Y: 2}
	p2 := Point2D[float64]{X: 4, Y: 6}
	fmt.Println("Distance between p1 and p2:", Distance(p1, p2))
}
