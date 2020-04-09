package main

import "fmt"

func main() {
	array1 := [3]string{"a", "b", "c"}
	fmt.Printf("The array: %v\n", array1)
	array2 := modifyArray(array1)
	fmt.Printf("The modified array: %v\n", array2)
	fmt.Printf("The old array: %v\n", array1)

	complexArray1 := [3][]string{
		[]string{"d", "e", "f"},
		[]string{"g", "h", "i"},
		[]string{"j", "k", "l"},
	}

	complexArray2 := modifyArray1(complexArray1)
	fmt.Printf("The modified complexArray1: %v\n", complexArray1)
	fmt.Printf("The old complexArray1: %v\n", complexArray2)
}

func modifyArray(a [3]string) [3]string {
	a[1] = "x"
	return a
}

func modifyArray1(a [3][]string) [3][]string {
	a[1][1] = "x"
	return a
}
