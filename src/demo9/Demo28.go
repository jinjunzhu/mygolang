package main

import "fmt"

type Printer func(content string) (n int, err error)

func printToStd(content string) (bytesNum int, err error) {
	return fmt.Printf(content)
}

func main() {
	var p Printer
	p = printToStd
	p("something")
}
