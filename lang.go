package main

import (
	"encoding/json"
	"fmt"
	"time"
)

// Go Snippets
// https://goproxy.io/zh/
func main1(fun string) {
	switch fun {
	case "print":
		// 1. print
		print()
	case "types":
		// 2. types
		types()
	case "consts":
		// 3. consts
		consts()
	case "functions":
		// 4. functions
		sum, muli := functions(1, 2)
		fmt.Println(sum, muli)
	case "json":
		// 5. json
		json2str()
	default:
		fmt.Println("not match")
	}
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
	const PI = 3.14
	fmt.Println(PI)

	const (
		C  = iota
		GO = iota + 2
		JAVA
	)
	fmt.Println(C, GO, JAVA)
}

func functions(a int, b int) (int, int) {
	return a + b, a * b
}

type datetime = int64

type Person struct {
	Name     string   `json:"name"`
	Age      int      `json:"age"`
	Birthday datetime `json:"-"`
}

func json2str() {
	p1 := Person{"Aaric", 18, 1668744000000}
	bts, err := json.Marshal(p1)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bts))

	var p2 Person
	dat := []byte(`{"Name":"Aaric","Age":18,"Birthday":1668744000000}`)
	if err := json.Unmarshal(dat, &p2); err != nil {
		panic(err)
	}
	fmt.Println(p2.Name, p2.Age, p2.Birthday)
}
