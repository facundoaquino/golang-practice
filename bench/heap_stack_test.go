package main

import (
	"fmt"

	"testing"
)

// https://blog.boot.dev/golang/pointers-faster-than-values/

/**

go test -bench=. -benchmem

**/

type data struct {
	a, b, c, d, e, f, g, h, i, j int64
}

var globalPtr *data

var globalValue data

func newDataPtr(i int) *data {

	data := &data{int64(i), int64(i + 1), int64(i + 2), int64(i + 3), int64(i + 4), int64(i + 5), int64(i + 6), int64(i + 7), int64(i + 8), int64(i + 9)}

	return data

}

func newData(i int) data {

	data := data{int64(i), int64(i + 1), int64(i + 2), int64(i + 3), int64(i + 4), int64(i + 5), int64(i + 6), int64(i + 7), int64(i + 8), int64(i + 9)}

	return data

}

func BenchmarkProcessValue(b *testing.B) {

	for i := 0; i < b.N; i++ {

		globalValue = newData(i)

	}

	// use it to avoid compiler optimization

	fmt.Println(globalValue.a)

}

func BenchmarkProcessPointer(b *testing.B) {

	for i := 0; i < b.N; i++ {

		globalPtr = newDataPtr(i)

	}

	// use it to avoid compiler optimization

	fmt.Println(globalPtr.a)

}
