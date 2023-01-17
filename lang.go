package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"math/big"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
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
	case "structCombineLang":
		// 12. struct combine
		structCombineLang()
	case "structInterfaceLang":
		// 13. struct interface
		structInterfaceLang()
	case "pointerLang":
		// 14. pointer
		pointerLang()
	case "nilLang":
		// 15. nil
		nilLang()
	case "errorLang":
		// 16. error
		errorLang()
	case "deferLang":
		// 17. defer
		deferLang()
	case "threadLang":
		// 18. thread
		threadLang()
	case "chanLang":
		// 19. chan
		chanLang()
	case "lockLang":
		// 20. lock
		lockLang()
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

func structCombineLang() {
	mLoc := loc{30.4, 114.4}
	myTemp := temp{high: 18.0, low: -4.0}
	myReport := report{sol: 5, loc: mLoc, temp: myTemp}
	//fmt.Println("avg temp:", myReport.temp.avg())
	fmt.Printf("avg temp: %v", myReport.avg())
}

type talker interface {
	talk() string
}

type cat struct {
}

func (c cat) talk() string {
	return "meow"
}

type dog int

func (d dog) talk() string {
	return strings.Repeat("bark", int(d))
}

func speek(t talker) {
	fmt.Println(t.talk())
}

type pet struct {
	talker
}

func structInterfaceLang() {
	//var talker interface {
	//	talk() string
	//}
	//talker = cat{}
	//fmt.Println(talker.talk())
	//talker = dog(3)
	//fmt.Println(talker.talk())

	speek(cat{})
	speek(dog(3))

	p := pet{cat{}}
	speek(p)
}

//func birthday(p *Person) {
func (p *Person) birthday() {
	p.Age++
}

func pointerLang() {
	age := 18
	var agePtr *int
	agePtr = &age
	fmt.Printf("age ptr is %T, the value is %v\n", agePtr, *agePtr)

	str := "aaa"
	strPtr := &str
	str = "bbb"
	fmt.Printf("str = %v, *strPtr = %v\n", str, *strPtr)

	p := &Person{
		Name: "zhangsan",
		Age:  18,
	}
	(*p).Name = "lisi"
	p.Name = "wangwu"
	//birthday(p)
	p.birthday()
	fmt.Printf("%+v\n", p)

	virusPtr := &[3]string{"Alpha", "Beta", "Gamma"}
	fmt.Println(virusPtr[0])
	fmt.Println(virusPtr[1:2])
}

func tail(strings []string) []string {
	return append(strings, "Alpha", "Beta", "Gamma")
}

func nilLang() {
	var numPtr *int
	if numPtr != nil {
		fmt.Println(*numPtr)
	}
	fmt.Println(numPtr)
	//fmt.Println(*numPtr)

	var fn func(a, b int) int
	fmt.Println(fn == nil)

	fmt.Println(tail(nil))

	var virus map[string]int
	fmt.Println(virus == nil)

	virus = make(map[string]int)
	virus["Alpha"] = 2019
	virus["Beta"] = 2020
	virus["Gamma"] = 2021
	v, ok := virus["Delta"]
	if ok {
		fmt.Println(v)
	}
	for k, v := range virus {
		fmt.Println(k, v)
	}
}

func errorLang() error {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, file := range files {
		fmt.Println(file.Name())
	}

	file, err := os.Create("test.log")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	_, err = fmt.Fprintln(file, "hello world")
	if err != nil {
		fmt.Println(err)
		file.Close()
		os.Exit(1)
	}

	var ErrCustom = errors.New("custom error")
	if file != nil {
		fmt.Println("custom error")
		return ErrCustom
	}

	return nil
}

func deferLang() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println(e)
		}
		fmt.Println("defer block")
	}()

	panic("let's go")
}

func sleep1s(i int, c chan int) {
	time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
	fmt.Printf("sleep-%v 1s ok\n", i)
	c <- i
}

func threadLang() {
	c := make(chan int)
	for i := 0; i < 5; i++ {
		go sleep1s(i, c)
	}

	timeout := time.After(2 * time.Second)
	for i := 0; i < 5; i++ {
		//tId := <-c
		//fmt.Printf("sleep-%v finished\n", tId)

		select {
		case tId := <-c:
			fmt.Printf("sleep-%v finished\n", tId)
		case <-timeout:
			fmt.Printf("task timeout\n")
			return
		}
	}
	fmt.Println("task running...")
}

func filter1(c chan string) {
	for _, v := range []string{"Alpha", "Beta", "Gamma", "Delta", "Omicron", "Zeta"} {
		c <- v
	}
	//c <- ""
	close(c)
}

func filter2(c1, c2 chan string) {
	for {
		//v := <-c1
		v, ok := <-c1
		//if v == "" {
		if !ok {
			//c2 <- ""
			close(c2)
			return
		}
		if !strings.Contains(v, "ta") {
			c2 <- v
		}
	}
}

func filter3(c chan string) {
	//for {
	//	v := <-c
	//	if v == "" {
	//		return
	//	}
	//	fmt.Println(v)
	//}
	for v := range c {
		fmt.Println(v)
	}
}

func chanLang() {
	c1 := make(chan string)
	c2 := make(chan string)
	go filter1(c1)
	go filter2(c1, c2)
	filter3(c2)
}

type Visited struct {
	key     sync.Mutex
	visited map[string]int
}

func (v Visited) visit(url string) int {
	v.key.Lock()
	defer v.key.Unlock()

	count := v.visited[url]
	count++
	v.visited[url] = count
	return count
}

func letVisit(i int, v Visited, url string) {
	count := v.visit(url)
	fmt.Printf("v-%v visit %v %v times\n", i, url, count)
}

func lockLang() {
	v := Visited{visited: make(map[string]int)}
	for i := 0; i < 100; i++ {
		go letVisit(i, v, fmt.Sprintf("http://let/%v", rand.Intn(10)))
	}
}
