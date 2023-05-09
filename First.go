package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
	"unicode/utf8"
)

const s string = "constant"

func main() {

	// begin()
	// fmt.Println("add function", add(1, 2))
	// var a, b string = "hello", "world"
	// // var b = "world"
	// println(a, " ", b)
	// a, b = swap(a, b)
	// fmt.Println("after swap : ", a, b)

	// constant()

	// forloop()

	// ifelse()

	// switchL()

	// arrayL()

	// sliceL()
	// mapsL()

	// rangeL()

	// funcL()

	// variadicFunctions()
	// lClosures()

	// lRecursion()

	// lPointers()

	// lStringsAndRunes()

	// lStructs()

	// lMethods()

	lInterfaces()
}

//	func add(x int, y int) int {
//		return x + y
//	}
//
// or you can do like this
func begin() {
	fmt.Println("My faverite number is ", rand.Intn(10))
	fmt.Printf("Now you can have %g problems. \n", math.Sqrt(2))
	fmt.Println(math.Pi)
}
func add(x, y int) int {
	return x + y
}

func constant() {
	fmt.Println(s)

	const n = 500000000
	const d = 3e20 / n
	fmt.Println(d)

}

func swap(x, y string) (string, string) {
	return y, x
}

func forloop() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for i := 7; i <= 9; i++ {
		fmt.Println(i)
	}

	for {
		fmt.Println("Loop")
		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}

func ifelse() {
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisble by 4")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		println("has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}

func switchL() {
	i := 3
	fmt.Println("write ", i, " as ")

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Print("It's a weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm a integer")
		default:
			fmt.Printf("Dont know type %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")

}

func arrayL() {
	var a [5]int
	fmt.Println("emp : ", a)

	a[4] = 100
	fmt.Println("set ; ", a)
	fmt.Println("get : ", a[4])

	fmt.Println("len : ", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl: ", b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d : ", twoD)
}

func sliceL() {
	s := make([]string, 3)
	fmt.Println("emp : ", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set : ", s)
	fmt.Println("get : ", s[2])
	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd : ", s)

	c := make([]string, len(s))

	copy(c, s)

	fmt.Println("cpy : ", c)

	l := s[2:5]
	fmt.Println("sl1: ", l)

	l = s[:5]
	fmt.Println("sl2 : ", l)

	l = s[3:]
	fmt.Println("sl3 : ", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl : ", t)
	t = append(t, "k")
	fmt.Println("dcl : ", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d ; ", twoD)
}

func mapsL() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map : ", m)

	v1 := m["k1"]
	fmt.Println("v1 : ", v1)

	v3 := m["k3"]
	fmt.Println("v3: ", v3)

	fmt.Println("len : ", len(m))

	delete(m, "k2")
	fmt.Println("map : ", m)

	_, prs := m["k1"]
	fmt.Println("prs: ", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map", n)
}

func rangeL() {
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		fmt.Println(num)
		sum += num
	}

	fmt.Println("sub : ", sum)
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index : ", i)
			fmt.Println("index : ", num)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}

	for k, v := range kvs {
		fmt.Printf("key --> %s : value --> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key:", k)
	}

	for _, v := range kvs {
		fmt.Println("value:", v)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}
}

func funcL() {
	res := plus(1, 2)
	fmt.Println("1+2 : ", res)
	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3", res)
	a, b := vals()
	fmt.Println("a : ", a)
	fmt.Println("b : ", b)
	_, c := vals()
	fmt.Println("print only c ", c)
}
func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return 1 + b + c
}

func vals() (int, int) {
	return 3, 7
}

func variadicFunctions() {
	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	sum(nums...)
}

func sum(nums ...int) {
	fmt.Println(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println("total : ", total)
}

func lClosures() {
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	nextInts := intSeq()
	fmt.Println(nextInts())
	fmt.Println(nextInts())

}

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func lRecursion() {
	fmt.Println("fact 7 : ", fact(7))
	fmt.Println("fib 7 : ", fib(7))

}

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func lPointers() {
	i := 1
	fmt.Println("initial : ", i)

	zeroval(i)
	fmt.Println("zeroval : ", i)

	zeroptr(&i)
	fmt.Println("zeroptr : ", i)

	fmt.Println("pointer:", &i)

}

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0

	// it change in memory
}

func lStringsAndRunes() {
	const s = "สวัสดี"

	fmt.Println("Length : ", len(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()
	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	fmt.Println("\nUsing DecodeRuneInString")

	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width
		examineRune(runeValue)
	}

}
func examineRune(r rune) {
	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("Found so sua")

	}
}

type person struct {
	name string
	age  int
}

func lStructs() {
	fmt.Println(person{name: "Bob", age: 20})
	fmt.Println(person{name: "Alice", age: 30})
	fmt.Println(person{name: "Fred"})
	fmt.Println(person{name: "Ann", age: 40})
	fmt.Println(newPerson("Jon"))
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)
	sp := &s
	fmt.Println(sp.age)
	sp.age = 51
	fmt.Println(sp.age)
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

type react struct {
	width, heigh int
}

func (r *react) area() int {
	return r.width * r.heigh
}

func (r react) perim() int {
	return 2*r.width + 2*r.heigh
}

func lMethods() {
	r := react{width: 10, heigh: 5}
	fmt.Println("area : ", r.area())
	fmt.Println("perim : ", r.perim())

	rp := &r
	fmt.Println("area : ", rp.area())
	fmt.Println("perim : ", rp.perim())

}

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return 2 * math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}
func lInterfaces() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}
	measure(r)
	measure(c)
}
