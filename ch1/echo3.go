package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("os.Args[0]: ")
	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1:])
	fmt.Println(strings.Join(os.Args[1:], " "))
}
