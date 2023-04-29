package main

import (
	"fmt"
	"reflect"
)

func main() {
	// fmt.Println(fact(10))
	fmt.Println(anoyRecFunc(5))
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
	// in below case , capacity gets doubled to keep future req. in mind
	c = append(c, d...)
	// c = append(c, []int{17, 18, 19, 20, 21, 22, 23}...)

	// fmt.Println("length of c", len(c))
	fmt.Println("capacity of c", cap(c))

	fmt.Println("c ", c)
}

func mapfunc() {
	var a map[int]string
	a = map[int]string{1: "sagar"}
	a[1] = "sagar"
	// a[2] = "sanghvi"
	fmt.Println(a)

	b := make(map[string]string)
	b["name"] = "sagar"
	b["surname"] = "sanghvi"

	delete(b, "name")

	fmt.Println(b)

	value, ok := b["name"]
	fmt.Println(value, ok)

	value, ok = b["surname"]
	fmt.Println(value, ok)

}

func rangefunc() {
	a := []int{1, 2, 3, 4, 5}
	for i, v := range a {
		fmt.Println(i, v)
	}

	b := map[string]string{"name": "sagar", "lastname": "sanghvi"}
	for key := range b {
		fmt.Println("key is", key)
	}

	for key, value := range b {
		fmt.Println("key value", key, value)
	}
}

func funcfunc() {
	a := func() string {
		return "abcd"
	}

	fmt.Println(a())
}

func variadicfunc() {
	a := []int{1, 2, 3, 4}
	sum(a...)
}

func sum(nums ...int) {
	for i := range nums {
		fmt.Println(i)
	}
}

func closurefunc() {
	a := nextInt()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())

	b := nextInt()
	fmt.Println(b())
}

func nextInt() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func fact(n int) int {
	if n == 1 {
		return 1
	}
	return n * fact(n-1)
}

func anoyRecFunc(n int) int {
	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}

	return fib(n)
}
