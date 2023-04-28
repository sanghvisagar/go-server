package main

import (
	"fmt"
	"reflect"
)

func main() {
	arrayfunc()
}

func switchFunc() {
	const a = 11
	if a%3 == 0 {
		fmt.Println("inside if")
	} else if a%2 == 0 {
		fmt.Println("inside else")
	} else {
		fmt.Println("Inside final else")
	}

	i := 20
	switch i {
	case 1:
		fmt.Print("i is equal to 1\n")
	case 2:
		fmt.Print("i is equal to 2\n")
	default:
		fmt.Print("i is not equal to 1or2\n")
	}

	switchOnType(i)

}

func switchOnType(i interface{}) {
	switch t := i.(type) {
	case bool:
		fmt.Println("I'm bool")
	case int:
		fmt.Println("I'm int")
	default:
		fmt.Printf("Dont know type of %T\n", t)
	}
}

func arrayfunc() {
	// var a []int
	a := []int{1, 2, 3, 4}
	b := make([]int, 5)

	fmt.Println("first element of array is ", a[0])
	fmt.Println("type of a is ", reflect.TypeOf(a))
	fmt.Println("type of b is ", reflect.TypeOf(b))

	var c []int
	c = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("capacity of c", cap(c))

	c = c[3:5]
	fmt.Println("capacity of c", cap(c))

	c = c[:7]
	fmt.Println("capacity of c", cap(c))

	c = append(c, 10, 11, 12, 13, 14, 15, 16)
	fmt.Println("capacity of c", cap(c))

	d := []int{17, 18, 19, 20, 21, 22, 23}
	c = append(c, d...)
	// c = append(c, []int{17, 18, 19, 20, 21, 22, 23}...)

	// fmt.Println("length of c", len(c))
	fmt.Println("capacity of c", cap(c))

	fmt.Println("c ", c)
}
