package main

import (
    "fmt"
    "os"
)

func main() {
  for _, v := range os.Args[1:] {
    fmt.Print(v + " ")
  }
  fmt.Print("\n")
}
