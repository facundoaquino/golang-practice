package main

// (x, y, z int) is the return type of the function

func returnNakedValues(a int, b int, c int) (x, y, z int) {

	x = a + b
	y = b + c
	z = c + a

	return
}

func main() {
	a, b, c := returnNakedValues(1, 2, 3)

	println(a, b, c)
}
