package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	fmt.Println("Go UNIX Interaction Demo")

	// 1. Chạy lệnh UNIX `ls -la`
	fmt.Println("\n1. Listing files with `ls -la`:")
	cmd := exec.Command("ls", "-la")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error running `ls -la`:", err)
		return
	}
	fmt.Println(string(output))

	// 2. Tạo một tệp mới và ghi dữ liệu vào tệp đó
	fmt.Println("\n2. Creating and writing to a file:")
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	file.WriteString("Hello, GoLang and UNIX!\n")
	fmt.Println("File created and written successfully")

	// 3. Quản lý tín hiệu hệ thống
	fmt.Println("\n3. Handling system signals:")
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println("Received signal:", sig)
		os.Exit(0)
	}()

	// Giả lập một quá trình chạy để chờ tín hiệu
	fmt.Println("Waiting for signal (Ctrl+C to interrupt)...")
	for {
		time.Sleep(1 * time.Second)
	}
}
