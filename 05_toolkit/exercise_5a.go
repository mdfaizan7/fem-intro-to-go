package main

import (
	"fem-intro-to-go/05_toolkit/utils/packages"
	"fmt"
)

func calculateImportantData() int {
	totalValue := packages.Add(1, 2, 3, 4, 5)
	return totalValue
}

func main() {
	fmt.Println("Packages!")
	total := calculateImportantData()
	fmt.Println(total)
	fmt.Println(packages.MakeExcited("hello nice to meet you!"))
}
