// Modify the echo program to also print os.Args[0].
package main

import (
	"fmt"
	"os"
  "strings"
)

func main() {
	fmt.Println(strings.Join(os.Args, " "))
}
