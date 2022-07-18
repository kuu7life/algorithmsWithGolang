package main

import "fmt"

func numInList(list []int, num int) bool {
	for i, v := range list {
		if v == num {
			fmt.Printf("%d in index of %d\n", v, i)
			return true
		}
	}
	return false
}

func main() {
	numms := []int{1, 2, 4, 4, 3, 5, 33, 221, 2, 3}
	numInList(numms, 4)
}
