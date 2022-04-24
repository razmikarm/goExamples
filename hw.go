package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("Hello, world.")

	// An int is a positive or negative number without decimals
	// Versions
	// uint8 : unsigned  8-bit integers (0 to 255)
	// uint16 : unsigned 16-bit integers (0 to 65535)
	// uint32 : unsigned 32-bit integers (0 to 4294967295)
	// uint64 : unsigned 64-bit integers (0 to 18446744073709551615)
	// int8 : signed  8-bit integers (-128 to 127)
	// int16 : signed 16-bit integers (-32768 to 32767)
	// int32 : signed 32-bit integers (-2147483648 to 2147483647)
	// int64 : signed 64-bit integers (-9223372036854775808 to
	// 9223372036854775807)

	var age int = 40

	const pi float64 = 3.1415926535897932 // 3846264338327950

	var num0ne = 1.00

	var num99 = .999

	randNum := 26

	fmt.Println(randNum, pi, num0ne-num99)

	fmt.Printf("Pi is %f \n", pi)
	fmt.Printf("Pi is %.10f \n", pi)

	fmt.Printf("Type of 'age' is %T \n", age)

	fmt.Printf("%d \n", 100)

	// %b prints in binary

	fmt.Printf("%b \n", 100)

	// %c prints the character associated with a keycode

	fmt.Printf("%c \n", 65)

	// %x prints in hexcode

	fmt.Printf("%x \n", 465)

	// %e prints in scientific notation

	fmt.Printf("%e \n", pi)

	var (
		name    string = "John"
		surname        = `Doe`
	)

	fmt.Println(name, surname)

	var isTrue = true

	fmt.Printf("The value is %t \n", isTrue)

	next()
	withSlice()
	withMaps()

	list := []float64{0.9, 1.2, 3.4, 5.6, 7.8}
	res1 := addThemUp(list)
	res2 := addThemUp(list[1:])
	fmt.Println(res1, res2)

	num1 := 89
	num2 := 123
	num1, num2 = switch2Vals(num1, num2)
	fmt.Println(num1, num2)

	defer printTwo() // runs after main finished
	fmt.Println(factorial(5))
	printOne()

	res := safeDiv(3, 0)
	fmt.Println(safeDiv(3, 2))
	fmt.Println(res)

	val := 6
	updateVal(&val)
	updateVal(&val)
	fmt.Println(val)

	newPtr := *new(int)
	var oldPtr int
	fmt.Println(newPtr == oldPtr) // true

	rect1 := Rectangle{0, 50, 10, 10}
	fmt.Println(rect1.rightPoints())
	// fmt.Println(getArea(rect1)) // Rectangle does not implement Shape (missing area method)
	circle := Circle{1}
	triangle := Triangle{5.5, 9}
	fmt.Println(getArea(circle), triangle.area())
	fmt.Println(Less(circle, triangle))

	sampString := "Hello World"
	fmt.Println(strings.Contains(sampString, "lo"))
	fmt.Println(strings.Index(sampString, "lo"))
	fmt.Println(strings.Count(sampString, "o"))
	fmt.Println(strings.Replace(sampString, "l", "x", 2))
	fmt.Println(sampString)

	csvStr := "1,2,3,4,56,7,89"
	vals := strings.Split(csvStr, ",")
	vals[0] += "1"
	fmt.Println(vals)

	newVals := []string{"H", "e", "l", "o", "l"}
	sort.Strings(newVals)
	fmt.Println("Ordered: ", newVals)
	fmt.Println("Joined: ", strings.Join(newVals, ""))

	file, err := os.Create("samp.txt")
	if err != nil {
		log.Fatal(err)
	}
	file.WriteString("Rand text")
	file.Close()

	stream, err := ioutil.ReadFile("samp.txt")
	if err != nil {
		log.Fatal(err)
	}
	readStr := string(stream)
	fmt.Println(readStr)

	randInt := 5
	randFloat := 10.9
	randString := "1297053"
	randStringBinary := "0b100111100101010011101"
	randStringFloat := "250.5"
	fmt.Println(float64(randInt))
	fmt.Println(int(randFloat))
	newInt10base, _ := strconv.ParseInt(randString, 0, 64)
	fmt.Println(newInt10base)
	newInt2base, _ := strconv.ParseInt(randStringBinary, 0, 64)
	fmt.Println(newInt2base)
	newFloat, _ := strconv.ParseFloat(randStringFloat, 64)
	fmt.Println(newFloat)

	http.HandleFunc("/", homePage)
	http.HandleFunc("/earth", earthPage)
	http.ListenAndServe(":8080", nil)

	for i := 0; i < 10; i++ {
		go count(i)
	}
	time.Sleep(time.Millisecond * 10000)

	stringChan := make(chan string)
	for i := 0; i < 3; i++ {

		go makeDough(stringChan)
		go addToppings(stringChan)
		go addSauce(stringChan)
		time.Sleep(time.Millisecond * 5000)
	}

}
