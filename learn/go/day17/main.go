package main

import "fmt"

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		results <- job * job
		fmt.Println("worker", id, "processed", job)
	}
}

func main() {
	jobs := make(chan int, 5)
	results := make(chan int, 5)

	for w := 1; w <= 2; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	for i := 0; i < 5; i++ {
		fmt.Println("result:", <-results)
	}
}
