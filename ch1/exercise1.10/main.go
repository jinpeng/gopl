// Find a web site that produces a large amount
// of data. Investigate caching by running fetchall
// twice in succession to see whether the reported
// time changes much. Do you get the same content
// each time? Modify fetechall to print its output
// to a file so it can be examined.
package main

import (
  "fmt"
  "io"
  "net/http"
  "os"
  "strings"
  "time"
)

func main() {
  start := time.Now()
  ch := make(chan string)
  for _, url := range os.Args[1:] {
    go fetech(url, ch) // start a gorountine
  }
  for range os.Args[1:] {
    fmt.Println(<-ch) // receive from channel ch
  }
  fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetech(url string, ch chan<-string) {
  start := time.Now()
  resp, err := http.Get(url)
  if err != nil {
    ch <- fmt.Sprint(err) // send to channel ch
    return
  }

  filename := generateFilename(url)
  f, err := os.Create(filename)
  if err != nil {
    ch <- fmt.Sprintf("while creating file %s: %v", filename, err)
    return
  }

  nbytes, err := io.Copy(f, resp.Body)
  resp.Body.Close()
  f.Close()
  if err != nil {
    ch <- fmt.Sprintf("while reading %s: %v", url, err)
    return
  }
  secs := time.Since(start).Seconds()
  ch <- fmt.Sprintf("%.2fs    %7d    %s", secs, nbytes, url)
}

func generateFilename(url string) string {
  filename := strings.Replace(url, "http://", "", -1)
  filename = strings.Replace(filename, ".", "_", -1)
  filename = strings.Replace(filename, "/", "_", -1)
  return filename
}
