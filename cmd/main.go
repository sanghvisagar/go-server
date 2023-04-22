package main

import (
	"fmt"
	"strconv"
)

var testMap map[string]string

func main() {
	fmt.Println("Hello World")

	testMap = make(map[string]string)

	testMap["companyName"] = "Infy"

	val, ok := testMap["compnayLocation"]
	if !ok {
		fmt.Println("compnayLocation is not defined")
	} else {
		fmt.Println("compnayLocation ", val)
	}

	val1, ok := testMap["companyName"]
	if !ok {
		fmt.Println("companyName is not defined")
	} else {
		fmt.Println("companyName ", val1)
	}

	fmt.Println(testMap)
	delete(testMap, "companyName")
	fmt.Println(testMap)

	//c := make(chan string)

	//	//	1st goroutine
	//	//var routine1 *string
	//	go func() string {fmt.Println("Goroutine 1st"); return "1st Done";}()
	//	//c <- go func() string {fmt.Println("Goroutine 1st"); return "1st Done";}()
	//	c <-
	//
	////  2nd goroutine
	//routine2 := go func() string {fmt.Println("Goroutine 2nd"); return "2nd Done";}()
	//
	////channel
	//
	//c <- routine1

	ch := make(chan string)
	for i := 1; i <= 10; i++ {
		go func(i int) {
			for j := 0; j < 10; j++ {
				ch <- "Goroutine : " + strconv.Itoa(i)
			}
		}(i)
	}

	for j := 1; j <= 100; j++ {
		fmt.Println(j, <-ch)
	}

}

//func goroutine1() string {
//	fmt.Println()
//}
//
//func goroutine2() string {
//	fmt.Println()
//}

//func add ()

//func (funcMap testMap) temp () int {
//
//}
