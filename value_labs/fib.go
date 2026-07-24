//create a fibonnaci series
// take an 0,1,1,2,3  input - 5 output - 3

package main

import "fmt"

func main() {

	// fibRes := map[int]int{}

	// input := 20
	// for i := 0; i < input; i++ {
	// 	res := fib(i, fibRes)
	// 	fibRes[i] = res
	// 	fmt.Println(res)
	// }
	fmt.Println(lengthOfLongestSubstring("abcabcdabc"))
}

func fib(n int, fibRes map[int]int) int {
	if n <= 1 {
		return n
	}

	res1, ok := fibRes[n-1]
	if !ok {
		res1 = fib(n-1, fibRes)
	}

	res2, ok := fibRes[n-2]
	if !ok {
		res2 = fib(n-2, fibRes)
	}

	return res1 + res2
}

// fib(5) fib(4), fib(3)
// fib(6) fib(5), fib(4), fib(3)

func fibonnaciPrint(n int) {

}

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	longestLength := 1
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			isUniqueSubstr := isUnique(s[i : j+1])
			if isUniqueSubstr && longestLength < j-i+1 {
				longestLength = j - i + 1
			}
		}
	}

	return longestLength
}

func isUnique(s string) bool {
	m := make(map[rune]int)
	for _, char := range s {
		if _, ok := m[char]; ok {
			return false
		} else {
			m[char] = 1
		}
	}
	return true
}

func isDuplicate(a string, b string) bool {
	m := make(map[rune]int)
	for _, char := range t {
		if _, ok := m[char]; ok {
			return false
		} else {
			m[char] = 1
		}
	}
	return true
}
