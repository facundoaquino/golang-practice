package main

import "fmt"

func main() {
	var n, n1, n2, n3 int
	n = 0
	n1 = 1
	n2 = 3
	n3 = 4

	sumIntArgs(n1, n2, n3, n)

	//ahother way to send parameters

	nums2 := []int{1, 2, 3}
	sumIntArgs(nums2...)

}

func sumIntArgs(nums ...int) {
	fmt.Println(nums)

	totalSum := 0

	for _, num := range nums {
		totalSum += num
	}

	fmt.Println(totalSum)
}
