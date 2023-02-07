package main

import (
	"fmt"
	"math"
)

func main1() {
	arr := []int{1, 2, 3}
	arr1 := []int{4, 5, 6}
	fmt.Println(arr, arr1)
	arr3 := append(arr, arr1...)
	fmt.Printf("%#v\n", arr3)
	fmt.Println(math.Pow(2, 10))

	var x []int
	x = append(x, 1)
	x = append(x, 2)
	x = append(x, 3)

	if len(x) > 5 {
		println("Smaller")
	} else {
		println("Bigger")
	}
	for i, i2 := range x {
		fmt.Println(i, i2)
	}
	A, B := myValue(5)
	fmt.Println(A, B)

	var obj = struct {
		x int
		y string
	}{}
	obj.y = "Hello"
	fmt.Println(obj)

	type obj2 struct {
		name string
		age  int
	}
}

func myValue(x int) (z int, y int) {
	z = x
	y = z
	return
}
