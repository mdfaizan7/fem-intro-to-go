package main

import "fmt"

// task 1
func avgScores(scores [5]float64) float64 {
	total := 0.0
	for _, score := range scores {
		total += score
	}
	return total / 5
}

// task 2
var pets map[string]string = map[string]string{
	"Fido":  "Dog",
	"Jutsu": "Cat",
}

func petNames(name string) bool {
	if _, ok := pets[name]; ok {
		return true
	}

	return false
}

// task 3
var initialGroceries = []string{"lemon", "flour", "fruit", "chicken"}

func groceries(newGroceries ...string) {
	foods := initialGroceries
	for _, grocery := range newGroceries {
		foods = append(foods, grocery)
	}

	fmt.Println(foods)
}

// func main() {
// 	fmt.Println(avgScores([5]float64{4, 5, 6, 7, 8}))
// 	fmt.Println(petNames("Fido"))
// 	groceries("rice", "wheat")
// }
