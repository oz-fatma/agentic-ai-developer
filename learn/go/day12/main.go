package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	nums = append(nums, 4, 5)
	fmt.Println("slice:", nums, "len:", len(nums), "cap:", cap(nums))

	sub := nums[1:4]
	fmt.Println("sub:", sub)

	counts := map[string]int{
		"go":  3,
		"git": 2,
	}
	counts["http"] = 1

	fmt.Println("map:")
	for key, value := range counts {
		fmt.Printf("  %s: %d\n", key, value)
	}

	shared := make([]int, 2, 4)
	copySlice := shared
	copySlice = append(copySlice, 99)
	fmt.Println("shared:", shared)
	fmt.Println("copySlice:", copySlice)
}
