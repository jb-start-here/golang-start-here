package main

import (
	"fmt"
	"time"
)

func main() {
	tasks := []func(){
		func() { time.Sleep(3 * time.Second); fmt.Println("task 1...") },
		func() { time.Sleep(3 * time.Second); fmt.Println("task 2...") },
		func() { time.Sleep(3 * time.Second); fmt.Println("task 3...") },
		func() { time.Sleep(3 * time.Second); fmt.Println("task 4...") },
		func() { time.Sleep(3 * time.Second); fmt.Println("task 5...") },
		func() { time.Sleep(3 * time.Second); fmt.Println("task 6...") },
		func() { time.Sleep(3 * time.Second); fmt.Println("task 7...") },
		func() { time.Sleep(3 * time.Second); fmt.Println("task 8...") },
		func() { time.Sleep(3 * time.Second); fmt.Println("task 9...") },
		func() { time.Sleep(3 * time.Second); fmt.Println("task 10...") },
		func() { time.Sleep(3 * time.Second); fmt.Println("task 11...") },
		func() { time.Sleep(3 * time.Second); fmt.Println("task 12...") },
		func() { time.Sleep(3 * time.Second); fmt.Println("task 13...") },
		func() { time.Sleep(3 * time.Second); fmt.Println("task 14...") },
		func() { time.Sleep(3 * time.Second); fmt.Println("task 15...") },
	}
	workerCount := 4

	<-executeAllFuncs(tasks, workerCount)
}
