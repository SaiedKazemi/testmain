package main

import (
	"flag"
	"fmt"
	"log"
)

var (
	panicOnError = false
	fatalError   = ""

	arg = flag.String("arg", "", "Some string argument.")
)

func main() {
	defer func() {
		recover() // error already printed in fatal()
	}()
	flag.Parse()
	if *arg != "valid" {
		fatal("invalid argument: %s", *arg)
	}
}

func fatal(format string, v ...interface{}) {
	fatalError = fmt.Sprintf(format, v...)
	if panicOnError {
		log.Panic(fatalError)
	} else {
		log.Fatal(fatalError)
	}
}
