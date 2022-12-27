package main

import (
	"encoding/json"
	"fmt"
	"math/big"
	"time"
)

// Go Snippets
// https://goproxy.io/zh/
func main1(fun string) {
	switch fun {
	case "printLang":
		// 1. print
		printLang()
	case "typeLang":
		// 2. types
		typeLang()
	case "constLang":
		// 3. consts
		constLang()
	case "functionLang":
		// 4. functions
		sum, muli := functionLang(1, 2)
		fmt.Println(sum, muli)
	case "jsonLang":
		// 5. json
		jsonLang()
	case "bigLang":
		bigLang()
	default:
		fmt.Println("not match")
	}
}

func printLang() {
	time.Sleep(3 * time.Second)
	fmt.Println("hello go")
}

func typeLang() {
	var a int
	var b1, b2 int = 10, 20
	var (
		c1 = 30
		c2 = true
	)
	d := 40
	fmt.Printf("%d %d %d %d %t %T\n", a, b1, b2, c1, c2, d)
}

func constLang() {
	const PI = 3.14
	fmt.Println(PI)

	const (
		C  = iota
		GO = iota + 2
		JAVA
	)
	fmt.Println(C, GO, JAVA)
}

func functionLang(a int, b int) (int, int) {
	return a + b, a * b
}

type datetime = int64

type Person struct {
	Name     string   `json:"name"`
	Age      int      `json:"age"`
	Birthday datetime `json:"-"`
}

func jsonLang() {
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

func bigLang() {
	m := big.NewInt(3e15)
	n := new(big.Int)
	n.SetString("2000000000000000", 10)

	result := new(big.Int)
	result.Add(m, n)
	fmt.Println(result)
}
