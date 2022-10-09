package main

import "fmt"

func main() {
	arr := []string{"hello", "mr", "snoop", "and", "pharell"}

	println("This has a different address in each iteration")
	for _, w := range arr {
		w := w
		println(fmt.Sprintf("address: %s\n", &w))
		println(fmt.Sprintf("value: %s\n", w))
	}

	println("This one has the same address in every iteration")
	for _, w := range arr {
		println(fmt.Sprintf("address: %s\n", &w))
		println(fmt.Sprintf("value: %s\n", w))
	}
}
