package main

import "fmt"

type MyInt int

type MyMapa map[MyInt]int

func main() {

	var myMap map[string]int

	fmt.Println(myMap["key"])

	fmt.Println(len(myMap))

	myMap2 := make(map[string]int)

	fmt.Println(myMap2["key"])

	changeKeyKeyfromMap(myMap2)

	fmt.Println(myMap2["key"])

	commits := map[string]int{
		"rsc": 3711,
		"r":   2138,
		"gri": 1908,
		"adg": 912,
	}

	fmt.Println(commits)

	v, ok := commits["rsc"]
	fmt.Println("rsc:", v, ok)

	for key, value := range commits {

		fmt.Println("Key:", key, "Value:", value)
	}

	var myTypeMap MyMapa

	myTypeMap[2] = 10
	// myTypeMap["2"] = 10

	fmt.Println(myTypeMap)

}

func changeKeyKeyfromMap(mymap map[string]int) {
	mymap["key"] = 10
}
