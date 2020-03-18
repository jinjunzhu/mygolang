package main

import (
	"demo2/lib"
	"flag"
	"os"
)

var name string

var cmdLine = flag.NewFlagSet("question", flag.ExitOnError)
func init(){
	cmdLine.StringVar(&name, "name", "everyone", "The greeting object.")
}
func main() {
	cmdLine.Parse(os.Args[1:])
	lib.PrintHello(name)
}


