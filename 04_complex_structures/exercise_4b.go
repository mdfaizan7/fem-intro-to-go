package main

func average2(nums ...float64) float64 {
	var total float64
	for _, num := range nums {
		total += +num
	}

	return total / 3
}

// func main() {
// 	var avg float64 = average(3, 10, 6)
// 	fmt.Println(avg)
// }
