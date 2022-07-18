package main

import "fmt"

//implementing recursion
func sum(numbers []int) int {
	fmt.Println(numbers)
	if len(numbers) == 0 {
		return 0
	}
	return numbers[0] + sum(numbers[1:])
}

func main() {
	numbers := []int{1, 2, 23, 4, 5, 3, 4}
	answer := sum(numbers)
	fmt.Println(answer)
}
