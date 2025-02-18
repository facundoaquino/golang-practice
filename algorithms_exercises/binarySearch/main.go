package main

import "fmt"

const sizeArr = 100000000

func sortedArrGeneration(size int) []int {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = i
	}

	return arr
}

func main() {

	arrToSearch := sortedArrGeneration(sizeArr)

	searched := binarySearch(arrToSearch, 199)

	fmt.Println("Serached: ", searched)

	searchedRecursive := binarySearchRecursive(arrToSearch, 500, 0, len(arrToSearch)-1)

	fmt.Println("Searched Recursive: ", searchedRecursive)
}

func binarySearch(arr []int, x int) bool {

	leftIndx := 0

	rightIndx := len(arr) - 1

	middleindx := (leftIndx + rightIndx) / 2

	for leftIndx <= rightIndx {
		if arr[middleindx] == x {
			return true
		}

		if x < arr[middleindx] {
			rightIndx = middleindx - 1

		}

		if x > arr[middleindx] {

			leftIndx = middleindx + 1
		}

		middleindx = (leftIndx + rightIndx) / 2

	}

	return false
}

func binarySearchRecursive(arr []int, x int, leftIndx int, rightIndx int) bool {
	if leftIndx > rightIndx {
		return false
	}

	middleInx := (leftIndx + rightIndx) / 2

	if arr[middleInx] == x {
		return true
	}

	if x < arr[middleInx] {
		return binarySearchRecursive(arr, x, leftIndx, middleInx-1)
	}

	return binarySearchRecursive(arr, x, middleInx+1, rightIndx)

}
