package test

import (
	"fmt"
	"testing"
)

func array() [1024]int {
	var x [1024]int
	for i := 0; i < len(x); i++ {
		x[i] = i
	}
	return x
}

func slice() []int {
	x := make([]int, 1024)
	for i := 0; i < len(x); i++ {
		x[i] = i
	}
	return x
}

func BenchmarkArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		array()
	}
}

func BenchmarkSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slice()
	}
}

func BenchmarkSliceToArray(b *testing.B) {
	array := [8]int{}
	slice := array[:]
	fmt.Println(&array, array)
	array[3] = 5
	fmt.Println(&slice, slice[3])
}

func modifySlice(innerSlice []string) {
	innerSlice = append(innerSlice, "a")
	innerSlice[0] = "b"
	innerSlice[1] = "b"
	fmt.Println(innerSlice)
}

func BenchmarkSliceAndSlicePointer(b *testing.B) {
	outerSlice := make([]string, 0, 3)
	outerSlice = append(outerSlice, "a", "a")
	modifySlice(outerSlice)
	fmt.Println(outerSlice)
}

func modifySlice2(innerSlice []string) []string {
	innerSlice = append(innerSlice, "a")
	innerSlice[0] = "b"
	innerSlice[1] = "b"
	fmt.Println(innerSlice)
	return innerSlice
}

func BenchmarkSliceAndSlicePointer2(b *testing.B) {
	outerSlice := make([]string, 0, 3)
	outerSlice = append(outerSlice, "a", "a")
	innerSlice := modifySlice2(outerSlice)
	fmt.Println(outerSlice)
	outerSlice = append(outerSlice, "c")
	fmt.Println(outerSlice)
	fmt.Println(innerSlice)
}
