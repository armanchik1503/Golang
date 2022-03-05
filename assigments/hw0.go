package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(Fizzbuzz(102))
	fmt.Println(IsPrime(3))
	fmt.Println(IsPalindrome("kazak"))
}

func Fizzbuzz(n int) string {
	if n%3 == 0 {
		return "Fizz"
	} else if n%5 == 0 {
		return "Buzz"
	} else {
		return "FizzBuzz"
	}

	return ""
}

func IsPrime(n int) bool {
	for i := 2; i <= int(math.Floor(float64(n)/2)); i++ { //6
		fmt.Println(i)
		if n%i == 0 {
			return false
		}
	}
	return true
}

func IsPalindrome(str string) bool {
	result := []byte{}
	for i := len(str) - 1; i >= 0; i-- { //5   5 > 0 0 > 4
		fmt.Println(i)
		result = append(result, str[i])
	}

	if str == string(result) {
		return true
	}
	return false
}
