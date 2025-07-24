package main

import "fmt"

func Product(a, b int) int {
	return a * b
}
func main() {
	result := Product(4, 5)
	fmt.Println("Product: ",result)
}