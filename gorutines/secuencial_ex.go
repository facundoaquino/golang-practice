package main

import (
	"mi_proyecto/utils"
)

func main() {

	var sizeElements int = 200000000
	arr1 := utils.GenerateRandomInts(sizeElements)
	arr2 := utils.GenerateRandomInts(sizeElements)

	for i := 0; i < len(arr1); i++ {
		arr1[i] = arr1[i] * 2
	}

	for i := 0; i < len(arr2); i++ {
		arr2[i] = arr2[i] * 2
	}

}
