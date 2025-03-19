package main

import (
	"fmt"
	"greet"
	"os"
)

func main() {
	fmt.Println(greet.Greet(os.Args[1:]))
}
