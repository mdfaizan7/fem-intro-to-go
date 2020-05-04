package main

import "fmt"

func main() {
	var sentence string = "This is a sentence"

	for index, letter := range sentence {
		if index%2 != 0 {
			fmt.Println("index:", index, " letter:", string(letter))
		}
	}
}
