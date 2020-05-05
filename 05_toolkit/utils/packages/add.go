package packages

import "fmt"

// Add numbers together
func Add(nums ...int) int {
	total := 0
	for _, v := range nums {
		fmt.Println("Number:", v)
		total += v
	}
	return total
}
