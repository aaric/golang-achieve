package main

import (
	"time"
	"fmt"
)

// Go Snippets
// https://goproxy.io/zh/
func main() {
  // 1. print
  //print()

  // 2. types
  //types()

  // 3. consts
  //consts()

  // 4. functions
  sum, muli := functions(1, 2)
  fmt.Println(sum, muli)
}

func print() {
	time.Sleep(3 * time.Second)
  fmt.Println("hello go")
}

func types() {
  var a int
  var b1, b2 int = 10, 20
  var (
    c1 = 30
    c2 = true
  )
  d := 40
  fmt.Printf("%d %d %d %d %t %T\n", a, b1, b2, c1, c2, d)
}

func consts() {
  const PI =  3.14
  fmt.Println(PI)

  const (
    C = iota
    GO = iota + 2
    JAVA
  )
  fmt.Println(C, GO, JAVA)
}

func functions(a int, b int) (int, int) {
  return a + b, a * b;
}
