package main

import "fmt"

func main() {
	score := 85

	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else {
		fmt.Println("Grade: C or below")
	}

	fmt.Println("Counting 1 to 5:")
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	subjects := []string{"Go", "Git", "HTTP"}
	fmt.Println("Subjects:")
	for _, subject := range subjects {
		fmt.Println("-", subject)
	}

	day := 2
	switch day {
	case 1:
		fmt.Println("Variables and types")
	case 2:
		fmt.Println("Control flow")
	default:
		fmt.Println("Another topic")
	}

	fmt.Println("Odd numbers under 10:")
	for n := 1; n < 10; n++ {
		if n%2 == 0 {
			continue
		}
		if n > 7 {
			break
		}
		fmt.Println(n)
	}
}
