// Modify the echo program to print the index and value of each arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	for index, value := range os.Args[1:] {
		fmt.Println(index, value)
	}
}
