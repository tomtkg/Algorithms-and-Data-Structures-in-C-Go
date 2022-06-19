package main

import (
	"fmt"
)

func main() {
	hashing(10, 10, []int{11, 3, 32, 4, 50, 34, 19, 39})
	hashing(11, 10, []int{11, 3, 32, 4, 50, 16, 36})
	hashing(8, 10, []int{11, 3, 32, 4, 50, 16, 36})
}
func hashing(n, v int, nums []int) {
	a := make([]int, n)
	for _, num := range nums {
		i := num % v
		for ; a[i] != 0; i = (i + 1) % v {
		}
		a[i] = num
	}
	fmt.Println(a)
}
