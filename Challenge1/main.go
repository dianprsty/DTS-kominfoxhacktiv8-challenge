package main

import (
	"fmt"
	"strconv"
)

func main() {
	i := 21
	j := true
	fmt.Println(i)

	fmt.Printf("%T\n", i)
	fmt.Printf("%%\n")

	fmt.Printf("%t\n", j)
	
	fmt.Printf("%c\n", i)
	fmt.Printf("%b\n", i)
	var ya = 1071
	fmt.Printf("%c\n", ya)
	fmt.Printf("%d\n", i)
	fmt.Printf("%o\n", i)
	
	var x = fmt.Sprintf("%x", i)
	var s,_ = strconv.Atoi(x)
	fmt.Printf("%x\n", s)
	fmt.Printf("%X %x\n", s,i-2)

	fmt.Printf("%U\n", '\u402f')

	var k float64 = 123.456
	fmt.Printf("%f\n", k)
	fmt.Printf("%e\n", k)

}