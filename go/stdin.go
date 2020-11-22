package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("input: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("[error]: could not read input")
	}

	fmt.Println("input was: " + input)
}
