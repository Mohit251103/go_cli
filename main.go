package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	greet := flag.Bool("greet", false, "Greet the user")
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("No args")
		os.Exit(1)
	}

	if *greet {
		fmt.Printf("Hello ,%s\n", args[0])
		os.Exit(0)
	} else {
		fmt.Printf("Name : %s\n", args[0])
		os.Exit(0)
	}
}
