package main

import (
	"fmt"
	_ "sort"
)

func main() {
	//	findPrimeNumber(2)

	is, cIf, cFor := isPrime(3)
	fmt.Println(is, cIf, cFor)
}

func findPrimeNumber(n int32) {
	var intIf int32 = 0
	var intFor int32 = 0
	var primes []int32
	var i int32
	for i = 1; i < n; i++ {
		is, cIf, cFor := isPrime(i)
		intIf = intIf + cIf
		intFor = intFor + cFor
		if is == true {
			primes = append(primes, i)
		}
	}

	//	sort.Ints([]int(primes))
	num := int32(len(primes) - 1)
	fmt.Println("最大素数为:", primes[num])
	fmt.Println("if执行次数:", intIf)
	fmt.Println("for执行次数:", intFor)
	//	for _,v := range primes {
	//
	//	}

}

func isPrime(value int32) (bool, int32, int32) {
	var is bool
	var int32If int32
	var int32For int32
	if value <= 3 {
		is = (value >= 2)
		int32If++
	}

	if value%2 == 0 || value%3 == 0 {
		is = false
		int32If++
	}

	var i int32
	for i = 5; i*i <= value; i += 6 {
		int32For++
		if value%i == 0 || value%(i+2) == 0 {
			is = false
			int32If++
		}
	}

	return is, int32If, int32For
}

func isPrimeB(value int) bool {
	if value <= 3 {
		return value >= 2
	}
	if value%2 == 0 || value%3 == 0 {
		return false
	}
	for i := 5; i*i <= value; i += 6 {
		if value%i == 0 || value%(i+2) == 0 {
			return false
		}
	}
	return true
}
