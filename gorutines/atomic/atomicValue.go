package main

import "sync/atomic"

/**
EL COMPILADOR DE GO DETECTA SI EN LA EJECUCION DE NUESTRO CODIGO HAY CONDICIONES DE CARRERA SI HAY ARRAJO UN RESULTADO COMO ESTE:

Previous write at 0x00c00008a00c by goroutine 7:
  main.main.func1()
      C:/Users/Redbee/Desktop/dev/golang/gorutines/atomic/atomicValue.go:10 +0x44
  main.main.gowrap1()
      C:/Users/Redbee/Desktop/dev/golang/gorutines/atomic/atomicValue.go:11 +0x41

Goroutine 9 (running) created at:
  main.main()
      C:/Users/Redbee/Desktop/dev/golang/gorutines/atomic/atomicValue.go:9 +0x64

Goroutine 7 (finished) created at:
  main.main()
      C:/Users/Redbee/Desktop/dev/golang/gorutines/atomic/atomicValue.go:9 +0x64
==================
Found 2 data race(s)
exit status 66



$ go run -race atomicValue.go
*/
func main() {

	var anInt int32

	/*for i := 0; i < 100; i++ {

		go func(i int) {
			anInt += int32(i)
		}(i)
	}
	*/

	for i := 0; i < 10; i++ {
		go func(i int) {

			atomic.AddInt32(&anInt, int32(i))
		}(i)
	}

}
