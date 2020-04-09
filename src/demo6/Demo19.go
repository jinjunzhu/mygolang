package main

import "fmt"

func main() {
	var m map[string]int

	key := "two"
	elem, ok := m["two"]
	fmt.Printf("The element paired with key %q in nil map: %d (%v)\n", key, elem, ok)

	fmt.Printf("The length of nil map: %d\n", len(m))

	fmt.Printf("Delete the key-element pair by key %q...\n", key)
	delete(m, key)

	elem = 2
	//fmt.Println("Add a key-element pair to a nil map...")
	//m["two"] = elem // 这里会引发panic。

	aMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	k := "two"
	v, ok := aMap[k]
	if ok {
		fmt.Printf("The element of key %q: %d\n", k, v)
	} else {
		fmt.Println("Not found!")
	}
	aMap["four"] = 6
	fmt.Printf("the new element four,%d", aMap["four"])
}
