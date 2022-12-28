package main

import (
	"encoding/json"
	"fmt"
	"math/big"
	"strconv"
	"time"
	"unicode/utf8"
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
	case "methodLang":
		// 5. method
		methodLang()
	case "jsonLang":
		// 6. json
		jsonLang()
	case "bigLang":
		// 7. big number
		bigLang()
	case "stringLang":
		// 8. string
		stringLang()
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

type celsius float64
type kelvin float64

func kelvinToCelsius(k kelvin) celsius {
	return celsius(k - 273.15)
}

func (k kelvin) celsius() celsius {
	return celsius(k - 273.15)
}

var k2c = func(k kelvin) celsius {
	return celsius(k - 273.15)
}

func methodLang() {
	var k kelvin = 100
	fmt.Println(kelvinToCelsius(k))
	fmt.Println(k.celsius())
	fmt.Println(k2c(k))
	func() {
		fmt.Println(celsius(k - 273.15))
	}()
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

func stringLang() {
	fmt.Println("hello \t go")
	fmt.Println(`hello \t go`)
	fmt.Println(`
Dear all:
  Happy New Year!`)

	var pi rune = 960
	var alpha rune = 940
	var omega rune = 963
	var bang rune = 33
	fmt.Printf("%v %v %v %v\n", pi, alpha, omega, bang)
	fmt.Printf("%c %c %c %c\n", pi, alpha, omega, bang)
	fmt.Println(string(pi), string(alpha), string(omega), string(bang))

	msg := "πάσ is chars!"
	fmt.Println(len(msg), "bytes")
	fmt.Println(utf8.RuneCountInString(msg), "runes")
	c, size := utf8.DecodeRuneInString(msg)
	fmt.Printf("%c is %v bytes\n", c, size)
	for i, c := range msg {
		fmt.Printf("%v %c\n", i, c)
	}

	countdown := 10
	str := "Using " + strconv.Itoa(countdown) + " seconds, status is " + fmt.Sprintf("%v", true)
	fmt.Println(str)
}
