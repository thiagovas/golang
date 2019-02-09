package main

import (
  "fmt"
  "io"
  "io/ioutil"
  "net/http"
  "os"
  "time"
)

func main() {
  start := time.Now()
  ch := make(chan string)

  for i, _ := range os.Args[1:] {
    addHttpPrefix(&os.Args[i+1])
  }

  for _, url := range os.Args[1:] {
    go fetch(url, ch)
  }

  for range os.Args[1:] {
    fmt.Println(<-ch)
  }
  fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

// fetch gets the url and outputs stats on the channel.
func fetch(url string, ch chan<- string) {
  start := time.Now()
  resp, err := http.Get(url)

  if err != nil {
    ch <- fmt.Sprint(err)
    return
  }

  nbytes, err := io.Copy(ioutil.Discard, resp.Body)
  resp.Body.Close()
  if err != nil {
    ch <- fmt.Sprintf("while reading %s: %v", url, err)
    return
  }

  latency := time.Since(start).Seconds()
  ch <- fmt.Sprintf("%.2fs \t  %7d bytes \t %s", latency, nbytes, url)
}

// addHttpPrefix checks whether the url has a http[s] prefix, and adds
// a http prefix in case it doesn't have.
func addHttpPrefix(url *string) {
  if len(*url) < 8 {
    *url = "http://" + *url
  }
  if (*url)[:4] != "http" {
    *url = "http://" + *url
  }
}
