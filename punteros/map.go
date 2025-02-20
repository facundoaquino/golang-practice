package main

import "fmt"

type MyMap map[string]string

func (m MyMap) Add(key, value string) {
	m[key] = value
}

func (m MyMap) Cotains(key string) bool {
	_, ok := m[key]
	return ok
}

func main() {

	var myMap MyMap = make(MyMap)

	myMap.Add("key1", "value1")

	fmt.Println(myMap["key1"])

	changeValueMap(myMap)

	fmt.Println(myMap["key1"])
}

// los mapas pasan por referencia afectando el valor original

func changeValueMap(m MyMap) {
	m["key1"] = "value2"
}
