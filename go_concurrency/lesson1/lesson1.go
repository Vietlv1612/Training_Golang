package main

import (
	"fmt"
	"time"
)

func printMessage(message string) {
	fmt.Println(message)
}

func main() {
	go printMessage("Hello from Goroutine!")
	fmt.Println("Hello from main function")

	// Đợi một khoảng thời gian để goroutine có thời gian chạy trước khi chương trình kết thúc
	time.Sleep(time.Second)
}
