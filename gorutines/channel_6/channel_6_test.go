package main

import (
	"fmt"
	"testing"
)

// si ejecutamos go test --race  podemos ver las condiciones de carrera porque no usamos canales

/**


Goroutine 10 (running) created at:
  mi_proyecto/gorutines/channel_5.TestAddUser()
      C:/Users/Redbee/Desktop/dev/golang/gorutines/channel_5/channel_5_test.go:12 +0x84
  testing.tRunner()
      C:/Program Files/Go/src/testing/testing.go:1792 +0x1dc
  testing.(*T).Run.gowrap1()
      C:/Program Files/Go/src/testing/testing.go:1851 +0x44

Goroutine 8 (finished) created at:
  mi_proyecto/gorutines/channel_5.TestAddUser()
      C:/Users/Redbee/Desktop/dev/golang/gorutines/channel_5/channel_5_test.go:12 +0x84
  testing.tRunner()
      C:/Program Files/Go/src/testing/testing.go:1792 +0x1dc
  testing.(*T).Run.gowrap1()
      C:/Program Files/Go/src/testing/testing.go:1851 +0x44
==================
--- FAIL: TestAddUser (0.00s)
    testing.go:1490: race detected during execution of test


**/

func TestAddUser(t *testing.T) {
	server := NewServer()
	server.Start()
	// for i := 0; i < 10; i++ {
	// 	go server.AddUser(fmt.Sprintf("user_%d", i))
	// }

	for i := 0; i < 10; i++ {

		server.userChannel <- fmt.Sprintf("user_%d", i)

	}

	fmt.Println("finish loop")

}
