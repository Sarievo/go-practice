package main

import (
  "fmt"
  "github.com/fatih/color"
  "example/hello"
)

func main() {
  color.Blue("I'm blue")
  fmt.Println(hello.Welcome)
  fmt.Println(hello.HelloWorld())
}
