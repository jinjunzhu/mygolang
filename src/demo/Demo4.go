package main

import (
	"flag"
	"fmt"
	"os"
)

var name string

func init() {
	flag.CommandLine = flag.NewFlagSet("", flag.PanicOnError)
	flag.CommandLine.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
		flag.PrintDefaults()
	}
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}
func main() {
	flag.Parse()
	fmt.Printf("hollo, %s!\n", name)
}