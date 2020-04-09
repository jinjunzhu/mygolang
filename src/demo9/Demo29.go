package main

import (
	"errors"
	"fmt"
)

type operator func(x, y int) int

func calculate(x int, y int, op operator) (int, error) {
	if op == nil {
		return 0, errors.New("invalid operation")
	}
	return op(x, y), nil
}

type calculateFunc func(x int, y int) (int, error)

func genCalculator(op operator) calculateFunc {
	return func(x int, y int) (int, error) {
		if op == nil {
			return 0, errors.New("invalid operation")
		}
		return op(x, y), nil
	}
}

func main() {
	add := func(x, y int) int {
		return x + y
	}

	x, y := 12, 13
	result, err := calculate(x, y, add)
	fmt.Printf("The result: %d (error: %v)\n", result, err)
	result, err = calculate(x, y, nil)
	fmt.Printf("The result: %d (error: %v)\n", result, err)

	x, y = 25, 26
	add1 := genCalculator(add)
	result, err = add1(x, y)
	fmt.Printf("The result: %d (error: %v)\n", result, err)

}
