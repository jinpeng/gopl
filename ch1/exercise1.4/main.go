// Modify dup2 to print the names of all files
// in which each duplicated line occurs.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
  line2file := make(map[string]map[string]bool)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, "STDIN", counts, line2file)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "exercise 1.4: %v\n", err)
				continue
			}
			countLines(f, arg, counts, line2file)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%s\n", n, line, getMapKeysString(line2file[line]))
		}
	}
}

func countLines(f *os.File, filename string, counts map[string]int, line2file map[string]map[string]bool) {
	input := bufio.NewScanner(f)
	for input.Scan() {
    line := input.Text()
		counts[line]++
    if line2file[line] == nil {
      line2file[line] = make(map[string]bool)
    }
    line2file[line][filename] = true
	}
}

func getMapKeysString(amap map[string]bool) string {
  var s string
  var sep string
  for key, _ := range amap {
    s += sep + key
    sep = ", "
  }
  return s
}

