package main

// cannot use int(1000) (constant 1000 of type int) as Duration value in assignment
// Por mas que Duration sea un alias de int16, no se puede asignar un int a un Duration (go los reconoce ceomo tipos distintos)
//

type Duration int16

func main() {

	var duration Duration

	duration = int(1000)

	var numberInt int = 100

	int.String()
	var str string = "HelloWorldv"
	str.String()

	// Conversión explícita de int a Duration
	duration = Duration(1000)

	//	duration.

}
