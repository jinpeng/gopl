// Experiment to measure the difference in running time between
// our potentially inefficient version and the version
// uses strings.Join.
package main

import (
	"fmt"
	"os"
  "strings"
)

func echoSlow(args []string) string {
  var s, sep string
  for i:=1; i<len(args); i++ {
    s += sep + args[i]
    sep = " "
  }
  return s
}

func echoFast(args []string) string {
  return strings.Join(args[1:], " ")
}

func main() {
  s := echoSlow(os.Args)
  fmt.Println(s)
  
  s = echoFast(os.Args)
  fmt.Println(s)
}

