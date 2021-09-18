package main

import "fmt"

func Worker(id int, jobs <-chan int, result chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker with id %d started fib with %d\n", id, job)

		fib := Fibionacci(job)
		fmt.Printf("Worker with id %d, job %d and fib %d\n", id, job, fib)

		result <- fib
	}
}

func Fibionacci(n int) int {
	if n <= 1 {
		return n
	}

	return Fibionacci(n-1) + Fibionacci(n-2)
}

func main() {
	tasks := []int{2, 3, 4, 5, 7, 10, 12, 40, 50}
	nWorker := 3
	jobs := make(chan int, len(tasks))
	results := make(chan int, len(tasks))

	for i := 0; i < nWorker; i++ {
		go Worker(i, jobs, results)
	}

	for _, value := range tasks {
		jobs <- value
	}

	close(jobs)

	for r := 0; r < len(tasks); r++ {
		<-results
	}
}
