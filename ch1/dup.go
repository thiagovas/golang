package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
  counts := make(map[string]int)
  input := bufio.NewScanner(os.Stdin)
  for input.Scan() {
    counts[input.Text()]++
  }

  // One thing to notice is that, unlike in C++, a loop over a map
  // doesn't occur in lexicographical order on the key.
  for line, n := range counts {
    if n > 1 {
      fmt.Printf("%d\t%s\n", n, line)
    }
  }
}
