package main

import (
	"encoding/json"
	"fmt"
	"math"
	"math/big"
	"sort"
	"strconv"
	"strings"
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
	case "structLang":
		// 6. struct
		structLang()
	case "bigLang":
		// 7. big number
		bigLang()
	case "stringLang":
		// 8. string
		stringLang()
	case "arrayLang":
		// 9. array
		arrayLang()
	case "mapLang":
		// 10. map
		mapLang()
	case "combineLang":
		// 12. combine
		combineLang()
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

type wgs84 struct {
	//latitude  float64
	//longitude float64
	latitude, longitude float64
}

type Person struct {
	Name     string   `json:"name"`
	Age      int      `json:"age"`
	Birthday datetime `json:"-"`
}

func structLang() {
	//var point struct {
	//	latitude  float64
	//	longitude float64
	//}

	var w1 wgs84
	w1.latitude = 30.4
	w1.longitude = 114.4
	//w2 := wgs84{30.4, 114.4}
	w2 := wgs84{latitude: 30.4, longitude: 114.4}
	w3 := w2
	w3.longitude += 0.1
	fmt.Printf("%v\n", w1)
	fmt.Printf("%+v\n", w2)
	fmt.Printf("%+v\n", w3)

	p1 := Person{"Aaric", 18, 1668744000000}
	bts, err := json.Marshal(p1)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bts))

	var p2 Person
	dat := []byte(`{"name":"Aaric","age":18,"birthday":1668744000000}`)
	if err := json.Unmarshal(dat, &p2); err != nil {
		panic(err)
	}
	fmt.Printf("%+v", p2)
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

func arrayLang() {
	p1 := [...]string{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune"}
	fmt.Println(p1)
	fmt.Println(p1[0:4], p1[4:6], p1[6:8])
	fmt.Println(p1[:4], p1[4:6], p1[6:])

	for i := 0; i < len(p1); i++ {
		fmt.Println(p1[i])
	}

	p2 := []string{"    Mercury", "Venus    ", "    Earth    "}
	for i, e := range p2 {
		fmt.Printf("%v %v\n", i, strings.TrimSpace(e))
	}
	fmt.Println(strings.Join(p2, ","))

	//p3 := sort.StringSlice{"Mercury", "Venus", "Earth"}
	p3 := sort.StringSlice(p1[:])
	p3.Sort()
	fmt.Println(p3)

	p4 := p1[3:6:8]
	fmt.Printf("%v -> length=%v, capacity=%v\n", p4, len(p4), cap(p4))

	p5 := make([]string, 0, 10)
	p5 = append(p5, "Mars", "Jupiter", "Saturn")
	fmt.Printf("%v -> length=%v, capacity=%v\n", p5, len(p5), cap(p5))
}

func mapLang() {
	virus := map[string]string{
		"Alpha": "B.1.1.7",
		"Beta":  "B.1.351",
		"Gamma": "P.1",
		"Delta": "B.1.617.2",
	}
	v1 := virus["Alpha"]
	fmt.Println(v1)

	virus["Alpha"] = "B.1.1.8"
	virus["Omicron"] = "B.1.1529"
	fmt.Println(virus)

	zeta := virus["Zeta"]
	fmt.Println(zeta)

	if zeta, ok := virus["Zeta"]; ok {
		fmt.Printf("Zeta is %v\n", zeta)
	} else {
		fmt.Printf("Zeta is nil\n")
	}

	covid19 := virus
	covid19["Alpha"] = "B.1.1.9"
	fmt.Println(virus)
	fmt.Println(covid19)

	delete(virus, "Alpha")
	fmt.Println(virus)
	fmt.Println(covid19)

	temps := []float64{
		-5.0, 18.0, 15.0, -5.0, 15.0, 10.0, -5.0,
	}

	counters := make(map[float64]int, 8)
	for _, e := range temps {
		counters[e]++
	}
	for k, v := range counters {
		fmt.Printf("%+.2f has %v times\n", k, v)
	}

	groups := make(map[float64][]float64)
	for _, e := range temps {
		g := math.Trunc(e/10) * 10
		groups[g] = append(groups[g], e)
	}
	for k, v := range groups {
		fmt.Printf("%v contians %v\n", k, v)
	}

	flags := make(map[float64]bool)
	for _, e := range temps {
		flags[e] = true
	}
	if flags[-5.0] {
		fmt.Println("it has -5.0 item")
	}
}

type temp struct {
	high, low celsius
}

func (t temp) avg() celsius {
	return (t.high + t.low) / 2
}

type loc struct {
	lat, long float64
}

//type report struct {
//	sol  int
//	temp temp
//	loc  loc
//}

// combine
type report struct {
	sol int
	temp
	loc
}

func combineLang() {
	mLoc := loc{30.4, 114.4}
	myTemp := temp{high: 18.0, low: -4.0}
	myReport := report{sol: 5, loc: mLoc, temp: myTemp}
	//fmt.Println("avg temp:", myReport.temp.avg())
	fmt.Printf("avg temp: %v", myReport.avg())
}
