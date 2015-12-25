// Try fetchall with longer argument lists, such as
// samples from the top million web sites available
// at alexa.com. How does the program behave if a 
// web site just doesn't respond?
package main

import (
  "bufio"
  "fmt"
  "io"
  "net/http"
  "os"
  "strings"
  "time"
)

func main() {
  urls, err := readLines(os.Args[1])
  if err != nil {
    fmt.Printf("reading url list from %s: %v", os.Args[1], err)
    os.Exit(1)
  }

  start := time.Now()
  ch := make(chan string)
  for _, url := range urls {
    go fetech(url, ch) // start a gorountine
  }
  for range urls {
    fmt.Println(<-ch) // receive from channel ch
  }
  fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func readLines(filepath string) ([]string, error) {
  file, err := os.Open(filepath)
  if err != nil {
    return nil, err
  }
  defer file.Close()

  var lines []string
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }
  return lines, scanner.Err()
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
