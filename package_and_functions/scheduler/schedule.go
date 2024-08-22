package main

import (
	"fmt"
	"time"
)

type Task func()

type Scheduler struct {
	tasks []Task
}

func NewScheduler() *Scheduler {
	return &Scheduler{tasks: []Task{}}
}

func (s *Scheduler) Schedule(t Task) {
	s.tasks = append(s.tasks, t)
}

func (s *Scheduler) Run() {
	for _, task := range s.tasks {
		task()
	}
}

func main() {
	scheduler := NewScheduler()

	// Lên lịch một công việc in ra thông báo sau 3 giây
	scheduler.Schedule(func() {
		defer fmt.Println("Task 1 Completed")
		time.Sleep(3 * time.Second)
		fmt.Println("Task 1 Running")
	})

	// Lên lịch một công việc in ra thông báo sau 2 giây
	scheduler.Schedule(func() {
		defer fmt.Println("Task 2 Completed")
		time.Sleep(2 * time.Second)
		fmt.Println("Task 2 Running")
	})

	// Chạy tất cả các công việc đã lên lịch
	scheduler.Run()
}
