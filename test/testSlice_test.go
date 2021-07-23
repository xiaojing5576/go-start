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

func modifySlice2(innerSlice *[]string) {
	*innerSlice = append(*innerSlice, "a")
	(*innerSlice)[0] = "b"
	(*innerSlice)[1] = "b"
	fmt.Println(innerSlice)
	fmt.Printf("%p\n", innerSlice)
}

func BenchmarkSliceAndSlicePointer2(b *testing.B) {
	outerSlice := make([]string, 0, 3)
	outerSlice = append(outerSlice, "a", "a")
	fmt.Printf("%p\n", &outerSlice)
	modifySlice2(&outerSlice)
	outerSlice = append(outerSlice, "c")
	//fmt.Println(outerSlice,outerSlice)
	//fmt.Println(innerSlice,innerSlice)
	fmt.Printf("%p\n", &outerSlice)
}

func BenchmarkArrayAndSlice2(b *testing.B) {
	arrayA := [2]int{100, 200}
	//var arrayB [2]int

	fmt.Printf("arrayA : %p , %v\n", &arrayA, arrayA)
	fmt.Println("112")
	arrayB := arrayA

	fmt.Printf("arrayB : %p , %v\n", &arrayB, arrayB)

	testArray(arrayA)
}

func testArray(x [2]int) {
	fmt.Printf("func Array : %p , %v\n", &x, x)
}
