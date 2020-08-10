// code snippet of parsing cli args
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("parse cli args")
	fmt.Printf("\nprogram name: %v\n", os.Args[0])

	// args are returned as a string
	args := os.Args[1:]
	fmt.Printf("\nargs: %v\n", args)

	name := args[0]

	count, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("error converting count to int", err)
	}

	fork, err := strconv.ParseBool(args[2])
	if err != nil {
		fmt.Println("error converting fork to bool", err)
	}

	fmt.Printf("\nname: %s\tcount: %d\tfork: %v\n", name, count, fork)
}
